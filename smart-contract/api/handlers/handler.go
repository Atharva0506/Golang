package handlers

import (
	"api/auction"
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// Handler serves as the receiver for our API endpoints.
// It holds the initialized smart contract wrapper so that 
// all endpoint functions can access and interact with the contract.
type Handler struct {
	Contract *contract.Contract
}

// NewHandler initializes a new Handler with the provided contract wrapper.
func NewHandler(c *contract.Contract) *Handler {
	return &Handler{
		Contract: c,
	}
}

// getTransactOpts is a helper function to construct transaction options
// required to send a transaction to the Ethereum blockchain.
// In a production environment, reading private keys from headers is highly discouraged.
// Users should typically sign transactions via a wallet (like MetaMask) on the frontend.
func getTransactOpts(c *gin.Context) (*bind.TransactOpts, error) {
	// 1. Extract the private key from the request header.
	privKeyHex := c.GetHeader("X-Private-Key")
	if privKeyHex == "" {
		return nil, errors.New("missing X-Private-Key header")
	}

	// 2. Extract the chain ID from the request header.
	chainIDStr := c.GetHeader("X-Chain-ID")
	if chainIDStr == "" {
		return nil, errors.New("missing X-Chain-ID header")
	}

	chainIDInt, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid chain ID format")
	}
	chainID := big.NewInt(chainIDInt)

	// 3. Parse the private key from hex into an ECDSA struct.
	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, errors.New("invalid private key")
	}

	// 4. Create the keyed transactor, which signs the transaction for this network.
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
