package attestation

import (
	"AdServerCollector/core/ton_chain"
	"AdServerCollector/libs/common"
	"AdServerCollector/logger"
	"AdServerCollector/utils"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tendermint/iavl"
	tmdb "github.com/tendermint/tm-db"
	"os"
	"time"
)

const (
	defaultIAVLCacheSize = 10000
)

var (
	StateTreeDb        = "state_avl"
	SaveVersionLevelDB *SaveVersionDB
	GetAdTree          *AttestationStoreImp
	LogTree            *AttestationStoreImp
	ClickTree          *AttestationStoreImp
	CallBackTree       *AttestationStoreImp
)

func init() {
	SaveVersionLevelDB = NewSaveVersionDB("save_version.db") //保存localStateVersion,chainVersion,chainHash
	LogTree = NewAttestationStore("log_tree.db", 0)
	ClickTree = NewAttestationStore("click_tree.db", 1)
	GetAdTree = NewAttestationStore("getad_tree.db", 2)
	CallBackTree = NewAttestationStore("callback_tree.db", 3)
}

type AttestationStoreImp struct {
	stateTree                  *iavl.MutableTree
	localStateVersion          uint64 // 本地内存version
	chainVersion               uint64
	txHash                     string
	statVersionFile            *os.File //保存链上Version
	savingAttestationSemaphore chan bool
	closing                    bool
	chainSave                  chan string
	chainClient                *ton_chain.TonNetWorkClient
	dbDir                      string
	saveType                   int
	lastSaveTime               *int64
	lastSaveChainTime          *int64
}

func NewAttestationStore(dataDir string, saveType int) *AttestationStoreImp {
	//todo: 链URL和钱包seed
	chainClient := ton_chain.NewTonNetWorkClient()
	attestationStore := &AttestationStoreImp{
		savingAttestationSemaphore: make(chan bool, 1),
		chainClient:                chainClient,
		dbDir:                      dataDir,
		saveType:                   saveType,
	}

	stateDb, err := tmdb.NewGoLevelDB(StateTreeDb, dataDir)
	if err != nil {
		panic(fmt.Errorf("NewGoLevelDB err %s", err))
		//return nil, fmt.Errorf("NewGoLevelDB err %s", err)
	}
	stateTree, err := iavl.NewMutableTree(stateDb, defaultIAVLCacheSize)
	if err != nil {
		panic(fmt.Errorf("NewMutableTree err %s", err))
		//return nil, fmt.Errorf("NewMutableTree err %s", err)
	}
	attestationStore.stateTree = stateTree
	err = attestationStore.Init()
	if err != nil {
		panic(fmt.Errorf("attestationStore.Init err %s", err))
	}
	logger.Infof("init attestation store success leveldb name:%s, root hash: %s, local version: %d, chain version: %d", dataDir, hex.EncodeToString(attestationStore.stateTree.Hash()), attestationStore.localStateVersion, attestationStore.chainVersion)
	logger.Infof("my address %v", &attestationStore)
	return attestationStore
}

func (this *AttestationStoreImp) Init() error {
	if err := this.ReadVersionStateFile(); err != nil {
		return err
	}
	var hashes []common.Uint256
	err := this.recoverStore(hashes)
	if err != nil {
		return fmt.Errorf("recoverStore error %s", err)
	}
	//todo: 实始化订阅nats
	go this.saveToChan()
	return nil
}

func (this *AttestationStoreImp) recoverStore(hashes []common.Uint256) error {
	this.stateTree.LoadVersion(int64(this.localStateVersion))
	for _, hash := range hashes {
		err := this.submitAttestation(hash)
		if err != nil {
			return fmt.Errorf("submit attestation: error %s", err)
		}
	}
	return nil
}

func (this *AttestationStoreImp) GetStoreProof(key []byte, version int64) ([]byte, []byte, []byte, uint32, error) {
	latestTree, err := this.stateTree.GetImmutable(version)
	if err != nil {
		return nil, nil, nil, 0, err
	}
	value, proof, err := latestTree.GetWithProof(key)
	if err != nil {
		return nil, nil, nil, 0, err
	}
	data := common.NewZeroCopySink(nil)
	storeProof := common.StoreProof(*proof)
	storeProof.Serialization(data) //序列化存证数据
	rootHash := latestTree.Hash()
	return value, data.Bytes(), rootHash, uint32(version), err
}

func (this *AttestationStoreImp) GetStoreProofJson(key []byte, version int64) ([]byte, *common.StoreProof, error) {
	latestTree, err := this.stateTree.GetImmutable(version)
	if err != nil {
		return nil, nil, err
	}
	value, proof, err := latestTree.GetWithProof(key)
	if err != nil {
		return nil, nil, err
	}
	storeProof := common.StoreProof(*proof)
	return value, &storeProof, nil
}

func (this *AttestationStoreImp) SubmitAttestation(hashes common.Uint256) error {
	this.getSavingAttestationLock()
	defer this.releaseSavingAttestationLock()
	if this.closing {
		return errors.New("save block error: ledger is closing")
	}

	err := this.submitAttestation(hashes)
	if err != nil {
		return fmt.Errorf("saveBlock error %s", err)
	}
	return nil
}

func (this *AttestationStoreImp) tryGetSavingAttestationLock() (hasLocked bool) {
	select {
	case this.savingAttestationSemaphore <- true:
		return false
	default:
		return true
	}
}

func (this *AttestationStoreImp) getSavingAttestationLock() {
	this.savingAttestationSemaphore <- true
}

func (this *AttestationStoreImp) releaseSavingAttestationLock() {
	select {
	case <-this.savingAttestationSemaphore:
		return
	default:
		panic("can not release in unlocked state")
	}
}

// GetStatVersion 获取当前版本,和上链的根Hash
func (this *AttestationStoreImp) GetStatVersion() (uint64, uint64, string) {
	return this.localStateVersion, this.chainVersion, this.txHash
}

// SaveVersionStateDB 写入树保存版本
func (this *AttestationStoreImp) SaveVersionStateDB(onlyLocal bool) {
	logger.Infof("start save version to file %s, local version:%d,chain version:%d", this.dbDir, this.localStateVersion, this.chainVersion)
	version := this.stateTree.Version()
	this.localStateVersion = uint64(version)

	if onlyLocal {
		if err := SaveVersionLevelDB.SetLocalStateVersion(this.saveType, this.localStateVersion); err != nil {
			logger.Errorf("SaveVersionDB.SetLocalStateVersion error:%s", err)
			return
		}
		logger.Infof("end save version to file, local version:%d ", this.localStateVersion)
	} else if uint64(version) > this.chainVersion {
		// 1小时上一次链
		chainTree, err := this.stateTree.GetImmutable(version)
		if err != nil {
			logger.Errorf("lock onchain version error:%s", err)
			return
		}
		hash, err := this.chainClient.SendHashString(hex.EncodeToString(chainTree.Hash()))
		if err != nil {
			logger.Errorf("send hash string error:%s", err)
			return
		}
		this.chainVersion = uint64(version)
		this.txHash = hash
		if err := SaveVersionLevelDB.SetChainHash(this.saveType, this.chainVersion, this.txHash); err != nil {
			logger.Errorf("SaveVersionDB.SetChainHash error:%s", err)
		}
		if err := SaveVersionLevelDB.SetChainStateVersion(this.saveType, this.chainVersion); err != nil {
			logger.Errorf("SaveVersionDB.SetChainStateVersion error:%s", err)
			return
		}
		logger.Infof("end save version to chain, chain version:%d, txHash:%s", this.chainVersion, this.txHash)
	}
}

func (this *AttestationStoreImp) ReadVersionStateFile() error {
	//var version uint64 = 0
	var localStateVersion uint64
	var chainVersion uint64
	var chainHash string

	localStateVersion = SaveVersionLevelDB.GetLocalStateVersion(this.saveType)
	chainVersion = SaveVersionLevelDB.GetChainStateVersion(this.saveType)
	chainHash = SaveVersionLevelDB.GetChainHash(this.saveType, chainVersion)
	logger.Infof("ReadVersionStateFile localStateVersion:%d, chainVersion:%d, chainHash:%s", localStateVersion, chainVersion, chainHash)
	this.localStateVersion = localStateVersion
	this.chainVersion = chainVersion
	this.txHash = chainHash

	return nil
}

func (this *AttestationStoreImp) submitAttestation(hashes common.Uint256) error {
	this.updateStateToTree(hashes)
	// 最后一次保存时间是否大于60秒
	currentTimestamp := time.Now().Unix()
	if this.lastSaveTime == nil {
		this.lastSaveTime = utils.ValueToPoint(time.Now().Unix())
	} else {
		if currentTimestamp-*this.lastSaveTime < 60 {
			return nil
		}
	}

	rootHash, _, err := this.stateTree.SaveVersion()
	if err != nil {
		return fmt.Errorf("stateTree.SaveVersion height: err %s", err)
	}
	this.SaveVersionStateDB(true)
	logger.Infof("submitAttestation success, commit instance address %v, db folder: %s, root hash: %s, current version: %d, chain version: %d", &this, this.dbDir, hex.EncodeToString(rootHash), this.stateTree.Version(), this.chainVersion)
	this.lastSaveTime = utils.ValueToPoint(currentTimestamp)
	return nil
}

func (this *AttestationStoreImp) updateStateToTree(hashes common.Uint256) {
	this.stateTree.Set(hashes.ToArray(), []byte("v"))

	logger.Infof("update state to tree, key: %s", hex.EncodeToString(hashes.ToArray()))
}

func (this *AttestationStoreImp) saveToChan() {
	for {
		select {
		case msg := <-this.chainSave: //接收广播消息
			{
				if msg == "save" {
					this.SaveVersionStateDB(false)
					logger.Infof("save to chan success,save instance address %v", &this)
				}
			}
		}
	}
}

// Close ledger store.
func (this *AttestationStoreImp) Close() error {
	this.getSavingAttestationLock()
	defer this.releaseSavingAttestationLock()
	this.SaveVersionStateDB(false)
	logger.Info("close ledger store")
	this.closing = true
	// Close state tree
	return nil
}

//func LogFilStatSeekGet() uint64 {
//	offset, err := SaveVersionLevelDB.GetLogFileOffset("seed")
//	if err != nil {
//		logger.Errorf("get seek from file failed: %v", err)
//		return uint64(0)
//	}
//	logger.Infof("get seek from leveldb success: %v", *offset)
//	return *offset
//}
//
//func LogFileStatSeekSet(seek uint64) {
//	err := SaveVersionLevelDB.SetLogFileOffset("seed", seek)
//	if err != nil {
//		logger.Errorf("save seek to file failed: %v", err)
//	}
//	logger.Infof("set seek to leveldb success: %v", seek)
//}
