package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"tcgasstation-backend/utils/eth/contract/ethbridge"
	"tcgasstation-backend/utils/eth/contract/generative_nft_contract"
	"tcgasstation-backend/utils/eth/contract/tcbridge"
	"tcgasstation-backend/utils/eth/safe"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"

	"tcgasstation-backend/utils/rpccaller"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Client struct {
	eth *ethclient.Client
}

func NewClient(eth *ethclient.Client) *Client {
	return &Client{eth}
}

func (c *Client) GetClient() *ethclient.Client {
	return c.eth
}

func GenerateAddress() (privKey, pubKey, address string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		err = errors.Wrap(err, "crypto.GenerateKey")
		return
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("failed to cast public key to ECDSA")
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

func (c *Client) GenerateAddressFromPrivKey(privKey string) (pubKey, address string, err error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		err = errors.Wrap(err, "crypto.HexToECDSA")
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("failed to cast public key to ECDSA")
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

func (c *Client) GenPubPriKeyFromIncPriKey(incPrivateKey []byte) (ecdsa.PrivateKey, ecdsa.PublicKey) {
	priKey := new(ecdsa.PrivateKey)
	priKey.Curve = crypto.S256()
	priKey.D = c.b2ImN(incPrivateKey)
	priKey.PublicKey.X, priKey.PublicKey.Y = priKey.Curve.ScalarBaseMult(priKey.D.Bytes())
	return *priKey, priKey.PublicKey
}

func (c *Client) b2ImN(bytes []byte) *big.Int {
	x := big.NewInt(0)
	x.SetBytes(crypto.Keccak256Hash(bytes).Bytes())
	for x.Cmp(crypto.S256().Params().N) != -1 {
		x.SetBytes(crypto.Keccak256Hash(x.Bytes()).Bytes())
	}
	return x
}

func (c *Client) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := c.eth.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, errors.Wrap(err, "c.eth.BalanceAt")
	}
	return balance, nil
}

func (c *Client) GetTransaction(ctx context.Context, txAddress string) (uint64, error) {
	hash := common.HexToHash(txAddress)
	receipt, err := c.eth.TransactionReceipt(ctx, hash)
	if err != nil {
		return 0, errors.Wrap(err, "c.eth.GetTransaction")
	}
	return receipt.Status, nil
}

func (c *Client) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.eth.PendingNonceAt(ctx, address)
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.eth.SuggestGasPrice(ctx)
}

func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.eth.NetworkID(ctx)
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.eth.SendTransaction(ctx, tx)
}

func (c *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.eth.TransactionReceipt(ctx, txHash)
}

func (c *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return c.eth.BlockByHash(ctx, hash)
}

func (c *Client) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return c.eth.TransactionByHash(ctx, hash)
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
func (c *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return c.eth.BlockByNumber(ctx, number)
}

// TransactionsByBlockNumber returns all transactions from the current canonical chain. If number is nil, the
// latest known block is returned.
func (c *Client) TransactionsByBlockNumber(ctx context.Context, number *big.Int) (types.Transactions, error) {
	block, err := c.eth.BlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return block.Transactions(), nil
}

const ADDRESS_0 = "0x0000000000000000000000000000000000000000"

func (c *Client) getEthHeader(
	blockHash common.Hash,
) (*types.Header, error) {
	blockByHash, err := c.BlockByHash(context.Background(), blockHash)
	if err != nil {
		return nil, errors.Wrap(err, "c.BlockByHash")
	}

	blockByNumber, err := c.BlockByNumber(context.Background(), blockByHash.Number())
	if err != nil {
		return nil, errors.Wrap(err, "c.BlockByHash")
	}

	if blockByNumber.Hash().String() != blockByHash.Hash().String() {
		return nil, errors.New("the requested eth BlockHash is being on fork branch, rejected")
	}

	return blockByHash.Header(), nil
}

func (c *Client) encodeForDerive(list types.DerivableList, i int, buf *bytes.Buffer) []byte {
	buf.Reset()
	list.EncodeIndex(i, buf)
	// It's really unfortunate that we need to do perform this copy.
	// StackTrie holds onto the values until Hash is called, so the values
	// written to it must not alias.
	return common.CopyBytes(buf.Bytes())
}

func (c *Client) GetNonce(txs []string, ctx context.Context) (*big.Int, error) {
	if len(txs) == 0 {
		return nil, nil // No tx => not retry => nonce = nil to auto estimate
	}

	for _, tx := range txs {
		t, _, err := c.TransactionByHash(ctx, common.HexToHash(tx))
		if err != nil {
			continue
		}
		return big.NewInt(int64(t.Nonce())), nil
	}
	return nil, fmt.Errorf("failed getting nonce %v", txs)
}

func (c *Client) GetNonceByTx(tx string, ctx context.Context) (*big.Int, error) {
	if len(tx) == 0 {
		return nil, nil // No tx => not retry => nonce = nil to auto estimate
	}

	t, _, err := c.TransactionByHash(ctx, common.HexToHash(tx))
	if err != nil {
		return nil, err
	}

	return big.NewInt(int64(t.Nonce())), nil
}

func (c *Client) GetMaxGasPrice(txs []string) (*big.Int, error) {
	if len(txs) == 0 {
		return big.NewInt(0), nil
	}

	maxGasPrice := big.NewInt(0)
	for _, tx := range txs {
		t, _, err := c.TransactionByHash(context.Background(), common.HexToHash(tx))
		if err != nil {
			continue
		}
		p := t.GasPrice()
		if p.Cmp(maxGasPrice) > 0 {
			maxGasPrice = p
		}
	}
	return maxGasPrice, nil
}

func ValidateAddress(address string) bool {
	return common.IsHexAddress(address)
}

func ConvertWeiFromEther(val float64) *big.Int {
	return new(big.Int).Mul(big.NewInt(int64(val*1e3)), big.NewInt(1e15))
}

// transfer:
func (c *Client) Transfer(senderPrivKey, receiverAddress string, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(senderPrivKey)
	if err != nil {
		return "", errors.Wrap(err, "crypto.HexToECDSA")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", errors.Wrap(err, "s.ethClient.PendingNonceAt")
	}

	gasLimit := uint64(21000)
	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "s.ethClient.SuggestGasPrice")
	}

	fee := new(big.Int)
	fee.Mul(big.NewInt(int64(gasLimit)), gasPrice)

	fmt.Println("fee: ", fee)

	toAddress := common.HexToAddress(receiverAddress)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "c.NetworkID")
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", errors.Wrap(err, "types.SignTx")
	}
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.Wrap(err, "c.SendTransaction")
	}
	return signedTx.Hash().Hex(), nil
}

// Transfer max
func (c *Client) TransferMax(privateKeyStr, receiveAddress string) (string, string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		fmt.Println(err)
		return "", "0", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the balance of the account
	balance, err := c.eth.BalanceAt(context.Background(), fromAddress, nil)

	if err != nil {
		fmt.Println(err)
		return "", "0", err
	}

	// Create the transaction
	gasPrice, err := c.eth.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return "", "0", err
	}

	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", "0", errors.Wrap(err, "s.ethClient.PendingNonceAt")
	}

	gasLimit := uint64(21000) // limit for sending ETH
	value := new(big.Int).Sub(balance, new(big.Int).Mul(new(big.Int).SetUint64(gasPrice.Uint64()), new(big.Int).SetUint64(gasLimit)))

	fmt.Println("amount ETH to send: ", value)

	toAddress := common.HexToAddress(receiveAddress)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Sign the transaction
	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", "0", errors.Wrap(err, "c.NetworkID")
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
		return "", "0", err
	}

	// Send the transaction
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return "", "0", err
	}

	fmt.Printf("Sent %s ETH from %s to %s\n", value.String(), fromAddress.Hex(), toAddress.Hex())

	return signedTx.Hash().Hex(), value.String(), nil
}

func (c *Client) SendMulti(contractAddress, privateKeyStr string, toInfo map[string]*big.Int, gasPrice *big.Int, gasLimit uint64) (string, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	// gasPrice, err := c.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return "", err
	// }

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "crypto.HexToECDSA")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", errors.Wrap(err, "crypto.HexToECDSA")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	// auth.GasLimit = uint64(21000 * len(toInfo)) // in units

	if gasPrice != nil {
		auth.GasPrice = gasPrice
	}

	// auth.GasLimit = gasLimit

	// Create a new instance of the contract with the given address and ABI
	contract, err := NewMultisend(common.HexToAddress(contractAddress), c.GetClient())
	if err != nil {
		return "", errors.Wrap(err, "NewMultisend")
	}

	var listHexAddress []common.Address
	var listAmount []*big.Int

	for k, v := range toInfo {
		listHexAddress = append(listHexAddress, common.HexToAddress(k))
		listAmount = append(listAmount, v)
		auth.Value = auth.Value.Add(auth.Value, v)
	}

	tx, err := contract.MultiTransferOST(auth, listHexAddress, listAmount)

	if err != nil {
		return "", errors.Wrap(err, "contract.MultiTransferOST")
	}

	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

// mint on tc:
func (c *Client) MintTC(contractAddress, privateKeyStr, toAddress string, chunks [][]byte) (string, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		fmt.Println("HexToECDSA err", err)
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	fmt.Println("gasPrice: ", gasPrice)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "crypto.HexToECDSA")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", errors.Wrap(err, "crypto.HexToECDSA")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	// auth.GasLimit = uint64(21000 * len(toInfo)) // in units
	// auth.GasPrice = gasPrice

	// Create a new instance of the contract with the given address and ABI
	contract, err := generative_nft_contract.NewGenerativeNftContract(common.HexToAddress(contractAddress), c.GetClient())
	if err != nil {
		return "", errors.Wrap(err, "NewGenerativeNftContract")
	}

	projectContract, err := contract.Project(nil)
	if err != nil {
		return "", errors.Wrap(err, "contract.Mint")
	}

	if projectContract.Index.Uint64() >= projectContract.MaxSupply.Uint64() {
		err = errors.New("minted_out")
		return "", errors.Wrap(err, "contract.Mint")
	}

	auth.Value = projectContract.MintPrice

	tx, err := contract.Mint(auth, common.HexToAddress(toAddress), chunks)

	if err != nil {
		return "", errors.Wrap(err, "contract.Mint")
	}

	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (c *Client) SafeExecTransaction(fullnode, privateKeyStr, contractBridge, multisigContractAddress string, tcTokens, listValidatorPrivKeyStr, addressesStr []string, amounts []*big.Int, function string) (string, uint64, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", 0, errors.Wrap(err, "crypto.HexToECDSA")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("from address auth: ", fromAddress)

	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", 0, err
	}

	fmt.Println("Nonce: ", nonce)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return "", 0, err
	}

	fmt.Println("gasPrice: ", gasPrice)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", 0, errors.Wrap(err, "crypto.HexToECDSA")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", 0, errors.Wrap(err, "crypto.HexToECDSA")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	// auth.GasLimit = uint64(21000 * len(toInfo)) // in units
	// auth.GasPrice = gasPrice

	// default value:
	var value *big.Int = big.NewInt(0)
	var operation uint8 = 0
	var safeTxGas = big.NewInt(0)
	var baseGas *big.Int = big.NewInt(0)
	var gasPrices *big.Int = big.NewInt(0)
	var gasToken common.Address = common.Address{}
	var refundReceiver common.Address = common.Address{}

	// setup:
	var addresses []common.Address
	for _, addressStr := range addressesStr {
		addresses = append(addresses, common.HexToAddress(addressStr))
	}

	var tokens []common.Address
	for _, token := range tcTokens {
		tokens = append(tokens, common.HexToAddress(token))
	}

	var signatures []byte

	var vaultAbi abi.ABI

	// build data:
	fmt.Println("contractBridge, tcToken, addresses, amounts", contractBridge, tcTokens, addresses, amounts)

	if function == "mint" {
		vaultAbi, err = abi.JSON(strings.NewReader(tcbridge.TcbridgeMetaData.ABI))
	} else { // withdraw
		vaultAbi, err = abi.JSON(strings.NewReader(ethbridge.EthbridgeMetaData.ABI))
	}

	if err != nil {
		return "", 0, errors.Wrap(err, "abi.JSON")
	}

	mintCallData, err := vaultAbi.Pack(function, tokens, addresses, amounts)
	if err != nil {
		return "", 0, errors.Wrap(err, "vaultAbi.Pack")
	}

	fmt.Println("mintCallData: ", mintCallData)

	// get encode tx
	// get nonce
	safeInst, err := safe.NewSafe(common.HexToAddress(multisigContractAddress), c.GetClient())

	if err != nil {
		return "", 0, errors.Wrap(err, "safe.NewSafe")
	}

	multiSigNonce, _ := safeInst.Nonce(nil)
	encodeTx, err := safeInst.EncodeTransactionData(
		nil,
		common.HexToAddress(contractBridge),
		value,
		mintCallData,
		0,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		common.Address{},
		common.Address{},
		multiSigNonce,
	)

	if err != nil {
		fmt.Println("EncodeTransactionData err: ", err)
		return "", 0, errors.Wrap(err, "EncodeTransactionData")
	}

	signData := rawsha3(encodeTx)
	fmt.Println(signData)

	var listValidatorKeyInfo []ValidatorKeyInfo

	for _, validatorPrivKeyStr := range listValidatorPrivKeyStr {

		validatorPrivKey, err := crypto.HexToECDSA(validatorPrivKeyStr)
		if err != nil {
			return "", 0, err
		}
		publicKey := validatorPrivKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

		address := crypto.PubkeyToAddress(*publicKeyECDSA)

		listValidatorKeyInfo = append(listValidatorKeyInfo, ValidatorKeyInfo{
			Address: address.Hex(),
			PrivKey: validatorPrivKey,
			Big:     address.Big(),
		})
	}

	// fmt.Println("before sort:")
	// for i := 0; i < len(listValidatorKeyInfo); i++ {
	// 	fmt.Println(listValidatorKeyInfo[i].Big)
	// }

	// sort:
	sort.Sort(ValidatorKeyInfoSlice(listValidatorKeyInfo))

	// fmt.Println("after sort:")
	for i := 0; i < len(listValidatorKeyInfo); i++ {
		// fmt.Println(listValidatorKeyInfo[i].Big)

		signBytes, err := crypto.Sign(signData, listValidatorKeyInfo[i].PrivKey)
		if err != nil {
			return "", 0, errors.Wrap(err, "crypto.Sign")
		}

		signBytes[64] += 27

		signatures = append(signatures, signBytes...)
	}

	toContract := common.HexToAddress(contractBridge)
	gasLimit, err := EstimateGas(
		ethereum.CallMsg{From: common.HexToAddress(multisigContractAddress), To: &toContract, GasPrice: auth.GasPrice, Value: auth.Value, Data: mintCallData},
		fullnode,
	)
	if err != nil {
		return "", 0, errors.Wrap(err, "contract.ExecTransaction")
	}
	auth.GasLimit = gasLimit * 2

	tx, err := safeInst.ExecTransaction(auth, common.HexToAddress(contractBridge), value, mintCallData, operation,
		safeTxGas, baseGas, gasPrices, gasToken, refundReceiver, signatures)

	if err != nil {
		return "", auth.GasLimit, errors.Wrap(err, "contract.ExecTransaction")
	}

	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	return tx.Hash().Hex(), auth.GasLimit, nil
}

func rawsha3(b []byte) []byte {
	hashF := sha3.NewLegacyKeccak256()
	hashF.Write(b)
	buf := hashF.Sum(nil)
	return buf
}

type EstimateRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

func ToCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func EstimateGas(msg ethereum.CallMsg, fullnode string) (uint64, error) {
	var resp EstimateRes
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{
		ToCallArg(msg),
		"latest",
	}
	err := rpcClient.RPCCall(
		"",
		fullnode,
		"",
		"eth_estimateGas",
		params,
		&resp,
	)
	if err != nil {
		return 0, err
	}

	bb, err := json.Marshal(resp)

	if err != nil {
		return 0, err
	}

	fmt.Println("Estimate gas res: ", string(bb))

	if resp.RPCError != nil {
		return 0, errors.New(resp.RPCError.Message)
	}

	hexString := resp.Result.(string)
	n := new(big.Int)
	n.SetString(hexString, 16)
	if n.Cmp(big.NewInt(0)) != 0 {
		return 0, errors.New("Call estimate gas got error")
	}
	return n.Uint64(), nil
}
