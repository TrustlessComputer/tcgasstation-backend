package usecase

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"tcgasstation-backend/utils/config"
	"tcgasstation-backend/utils/eth"
	"testing"
	"time"

	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/ethclient"
	"gotest.tools/assert"
)

func TestGetListUTXOFromBlockStream(t *testing.T) {
	u := Usecase{
		Config: &config.Config{
			BitcoinParams: config.BitcoinParamsRegtest,
			BlockStream:   "https://blockstream.regtest.trustless.computer/regtest/",
		},
	}
	// address := "bcrt1pj2t2szx6rqzcyv63t3xepgdnhuj2zd3kfggrqmd9qwlg3vsx37fq7lj7tn"

	u.getBlockStreamConfig()

	// u.getUTXOFromBlockStream(address, 6, 1000000)
}
func ParseTx(data string) (*wire.MsgTx, error) {
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	tx := &wire.MsgTx{}
	err = tx.Deserialize(strings.NewReader(string(dataBytes)))
	if err != nil {
		return nil, err
	}
	return tx, nil
}
func TestGetListTCPrivateKey(t *testing.T) {
	res := getListValidatorPrivKey()
	fmt.Printf("res: %v - %v\n", res, len(res))
}

func TestParseTx(t *testing.T) {
	hexTx := "010000000001019cb5c1f8ae0b0a59fc85fb6bb2f23c7800bcedffaf006e00e0d7badd1f895ec500000000005019000002f082000000000000225120d5254f2c52e2672daea941a86c99232693149fd0423ef523fe4e0dcb12a68d53004601000000000022002045e9b957847aa39e386a1961dc92233bffd150c9217e729e65a1b557b5156db60200f155210374d80fd52aba67dfc80d4d5186b67104c4a613786ff3ae0b80773245581e6d612103e2f2a5515bd55fc8921e3baa503f29766023db3f24bcdd89f3784b28024807622103bceeb1fc74e5d207dea737fce7e3d05f587a5cfa475594722203682b6f9dd1162102eb5a2102594dc0f25e80ebf592722c64a058d54a1f9cb2f81f8b7dd12f3f8021210227b01c631ba52abb1a3495e8ef338cb27eee596825c324a94bcbbc81119faf2e210232a24e2f1e070d64a5e3b0872b58a0cba4cdd7d632374b012fe19dfa4a195f02210349c484fac836f2c33c88a60c136a4dd29da39a081d3b55586f40e65ac452071b57ae00000000"
	tx, err := ParseTx(hexTx)
	fmt.Printf("Err: %v\n", err)
	fmt.Printf("Tx: %+v\n", tx)

	fmt.Printf("TxIn: %+v\n", tx.TxIn[0])
	fmt.Printf("TxOut: %+v\n", tx.TxOut[0])
	fmt.Printf("TxOut: %+v\n", tx.TxOut[1])

	hex2 := "020000000001017ffeeefefacda8936d0088323a211e17d74b537e864c16765f3386ff11e9854a0000000000fdffffff0232c8943e00000000225120f6478c2c4fb888c5495d53606de5be63b39035bfe00d133c3200921674a6e7f9e56c0000000000002251200ac938db67d78e55934b0d64634e241d34e6109cb30923ce3379851f66bfdf510140f4afdad00a92b35753d11cd216950b3857019268014346836b5959a8daea93b83179b7226a0edb5289ecc9cb1d5e240b89eb26aa967de28af2fd16aac52e792b00000000"

	tx2, err := ParseTx(hex2)
	fmt.Printf("Err: %v\n", err)
	fmt.Printf("Tx2: %+v\n", tx)

	fmt.Printf("TxIn 2: %+v\n", tx2.TxIn[0])
	fmt.Printf("TxOut 2: %+v\n", tx2.TxOut[0])
	fmt.Printf("TxOut 2: %+v\n", tx2.TxOut[1])
}

func TestGetAmountFormatStr(t *testing.T) {
	fmt.Println(getAmountFormatStr("2223516", 1e6))
	assert.Equal(t, false, true)
}

func TestIsForkedChain(t *testing.T) {

	tcClientWrap, _ := ethclient.Dial("https://tc-node.trustless.computer")
	tcClient := eth.NewClient(tcClientWrap)
	u := Usecase{
		Config: &config.Config{
			BitcoinParams: config.BitcoinParamsMaintest,
			BlockStream:   "https://blockstream.regtest.trustless.computer/regtest/",
		},
		TcClient: tcClient,
	}

	// address := "bcrt1pj2t2szx6rqzcyv63t3xepgdnhuj2zd3kfggrqmd9qwlg3vsx37fq7lj7tn"
	tcTxID := "0xbf06fb0909f97267e3f7acfb91c5cda16bd343ae52396cae4f4f093746780850"

	res, err := u.isMainChain(tcTxID)
	fmt.Printf("RES: %v - %v\n", res, err)

	// u.getUTXOFromBlockStream(address, 6, 1000000)
}
func TestRandomAmount(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	min := 30.0
	max := 50.0
	amount := min + rand.Float64()*(max-min)
	fmt.Printf("Random amount between $%.2f and $%.2f: $%.2f\n", min, max, amount)

	ethAmount := amount / 1826.15
	fmt.Printf("$%.2f is currently worth %.8f ETH\n", amount, ethAmount)

	ethBigInt := new(big.Int)
	ethBigInt.SetString(fmt.Sprintf("%.0f", ethAmount*math.Pow10(18)), 10)
	fmt.Printf("$%.2f is currently worth %v ETH\n", amount, ethBigInt)

	assert.Equal(t, false, true)

}

func TestBLock(t *testing.T) {

	eth := 1.23456789 // The amount of ETH

	gwei := new(big.Int)
	ethInWei := new(big.Float).Mul(big.NewFloat(eth), big.NewFloat(1e18))
	ethInWei.Int(gwei)

	fmt.Println("ETH:", eth)
	fmt.Println("Gwei (BigInt):", gwei)

	var blockFrom uint64 = 10000
	var blockTo uint64 = 10000

	var block uint64 = 0
	var maxBlock uint64 = 1000
	for block = blockFrom; block <= blockTo; block++ {

		max := blockTo - block

		if max > (maxBlock) {
			max = (maxBlock)
		}
		fromBlockTemp := block
		toBlockTemp := fromBlockTemp + max

		block = toBlockTemp

		log.Println("Filter tc burn checking block from, to", fromBlockTemp, ":", toBlockTemp)
	}
	assert.Equal(t, false, true)
}
