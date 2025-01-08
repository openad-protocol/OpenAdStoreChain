package ton_chain

import (
	"AdServerCollector/conf"
	"AdServerCollector/logger"
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/tvm/cell"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

type TonNetWorkClient struct {
	client *liteclient.ConnectionPool
	w      *wallet.Wallet
	api    *ton.APIClient
}

func NewTonNetWorkClient() *TonNetWorkClient {
	var url *string
	var walletSeed *string
	var client *liteclient.ConnectionPool
	var err error
	var ctx context.Context
	var w *wallet.Wallet
	ctx = context.Background()
	client = liteclient.NewConnectionPool()
	if conf.ISTEST {
		url = &conf.Config.TonNetwork.Testnet.RPCUrl
		walletSeed = &conf.Config.TonNetwork.Testnet.Seed
	} else {
		url = &conf.Config.TonNetwork.Mainnet.RPCUrl
		walletSeed = &conf.Config.TonNetwork.Mainnet.Seed
	}
	if err = client.AddConnectionsFromConfigUrl(ctx, *url); err != nil {
	}

	api := ton.NewAPIClient(client)
	// 取得钱包
	wWord := strings.Split(*walletSeed, " ")
	if w, err = wallet.FromSeed(api, wWord, wallet.V3); err != nil {
		panic(err)
	}
	return &TonNetWorkClient{client: client, w: w, api: api}
}

// SendHashString 发送hash字符串
func (t *TonNetWorkClient) SendHashString(hash string) (string, error) {
	bigInt := new(big.Int)
	// Set the hex string into the big.Int, base 16 for hexadecimal
	bigInt.SetString(hash, 16)

	msgBody := cell.BeginCell().
		MustStoreUInt(0x7e8764ef, 32).
		MustStoreUInt(0, 64).
		MustStoreBigUInt(bigInt, 256).
		EndCell()

	// 目标合约
	addr := address.MustParseAddr("kQDy5OOP1zrMyut4J6-omWFQsJ5E6TD-mG1LYXwH_5KmyGnV")

	mint := wallet.SimpleMessage(addr, tlb.MustFromTON("0.02"), msgBody)

	msg, err := t.w.SendManyWaitTxHash(context.Background(), []*wallet.Message{mint})
	if err != nil {
		return "", err
	}
	txHash := hex.EncodeToString(msg)
	logger.Infof("send root hash: %s, tx hash: %s", hash, txHash)
	return txHash, nil
}
