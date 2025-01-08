package attestation

import (
	"AdServerCollector/logger"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
)

type SaveVersionDB struct {
	db     *leveldb.DB
	dbPath *string
}

func NewSaveVersionDB(dbPath string) *SaveVersionDB {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}
	return &SaveVersionDB{
		db:     db,
		dbPath: &dbPath,
	}
}
func (s *SaveVersionDB) GetDb() *leveldb.DB {
	return s.db
}

// SetLocalStateVersion sets the local state version
func (s *SaveVersionDB) SetLocalStateVersion(t int, localStateVersion uint64) error {
	_strVersion := fmt.Sprintf("%d", localStateVersion)
	_localStateVersion := fmt.Sprintf("%d:localStateVersion", t)
	return s.db.Put([]byte(_localStateVersion), []byte(_strVersion), nil)
}

// GetLocalStateVersion gets the local state version
func (s *SaveVersionDB) GetLocalStateVersion(t int) (version uint64) {
	_localStateVersion := fmt.Sprintf("%d:localStateVersion", t)

	data, err := s.db.Get([]byte(_localStateVersion), nil)
	if err != nil {
		logger.Errorf("GetLocalStateVersion err %s", err)
		return 0
	}
	if version, err = strconv.ParseUint(string(data), 10, 64); err != nil {
		logger.Errorf("GetLocalStateVersion err %s", err)
		return 0
	}
	return version
}

// SetChainStateVersion 	sets the chain version
func (s *SaveVersionDB) SetChainStateVersion(t int, chainVersion uint64) error {
	_strVersion := fmt.Sprintf("%d", chainVersion)
	_chainVersion := fmt.Sprintf("%d:chainVersion", t)
	return s.db.Put([]byte(_chainVersion), []byte(_strVersion), nil)
}

// GetChainStateVersion gets the chain version
func (s *SaveVersionDB) GetChainStateVersion(t int) (chainVersion uint64) {
	//_strVersion := fmt.Sprintf("%d", chainVersion)
	_chainVersion := fmt.Sprintf("%d:chainVersion", t)
	data, err := s.db.Get([]byte(_chainVersion), nil)
	if err != nil {
		return 0
	}
	_strVersion := string(data)
	if chainVersion, err = strconv.ParseUint(_strVersion, 10, 64); err != nil {
		logger.Errorf("GetLocalStateVersion err %s", err)
		return 0
	}
	return chainVersion
}

// SetChainHash 设置链上Hash
func (s *SaveVersionDB) SetChainHash(t int, version uint64, chainHash string) error {
	_version := fmt.Sprintf("%d:%d", t, version) //第一个是类型，第二个是版本
	return s.db.Put([]byte(_version), []byte(chainHash), nil)
}

func (s *SaveVersionDB) GetChainHash(t int, version uint64) (result string) {
	_version := fmt.Sprintf("%d:%d", t, version) //第一个是类型，第二个是表本
	result = "0000000000000000000000000000000000000000000000000000000000000000"
	_hByte, err := s.db.Get([]byte(_version), nil)
	if err != nil {
		logger.Errorf("GetChainHash err %s", err)
		return result
	}
	result = string(_hByte)
	return result
}

func (s *SaveVersionDB) Close() error {
	return s.db.Close()
}
