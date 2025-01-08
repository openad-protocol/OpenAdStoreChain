package ton_chain

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func TestCall(t *testing.T) {
	client := liteclient.NewConnectionPool()
	ctx := context.Background()

	// connect to mainnet lite server
	err := client.AddConnectionsFromConfigUrl(ctx, "https://ton.org/testnet-global.config.json")
	if err != nil {
		panic(err)
	}

	// initialize ton api lite connection wrapper
	api := ton.NewAPIClient(client)
	w := getWallet(api)

	walletAddress := w.WalletAddress()
	ws := walletAddress.String()
	log.Println("Deploy wallet:", ws)

	addr := address.MustParseAddr("kQDy5OOP1zrMyut4J6-omWFQsJ5E6TD-mG1LYXwH_5KmyGnV")

	// block, err := api.CurrentMasterchainInfo(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// res, err := api.RunGetMethod(context.Background(), block, addr, "mult", 7, 8)

	hash := "5fba5157df0eb1c19a7ce5b0118f65b563646f296a9300766a4ceb8511a85103"
	bigInt := new(big.Int)
	// Set the hex string into the big.Int, base 16 for hexadecimal
	bigInt.SetString(hash, 16)

	msgBody := cell.BeginCell().
		MustStoreUInt(0x7e8764ef, 32).
		MustStoreUInt(0, 64).
		MustStoreBigUInt(bigInt, 256).
		EndCell()

	mint := wallet.SimpleMessage(addr, tlb.MustFromTON("0.03"), msgBody)

	tx, err := w.SendManyWaitTxHash(ctx, []*wallet.Message{mint})
	if err != nil {
		panic(err)
	}
	txHash := hex.EncodeToString(tx)
	log.Printf("tx hash: %s", txHash)
}

func getWallet(api *ton.APIClient) *wallet.Wallet {
	words := strings.Split(
		"project afford syrup buzz knife chat snack nerve cage jar short balance talent easily august fluid auto version coyote kiwi satisfy crucial journey hurt", " ")
	w, err := wallet.FromSeed(api, words, wallet.V3)
	if err != nil {
		panic(err)
	}
	return w
}

func TestClient(t *testing.T) {
	url := ""
	walletSeed := ""
	c := NewTonNetWorkClient(&url, &walletSeed)
	r, err := c.SendHashString("c2cb3de0d4b384523c1ded9dd55e616caee0e13a2bece9419e7283edeba4eec8")
	if err != nil {
		t.Log(r)
		t.Fatal(err)
	}
}
