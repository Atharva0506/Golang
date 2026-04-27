package handlers

import (
	"api/models"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartAuction handles the API request to start a new auction.
// It requires a transaction option (auth), opening bid, and duration.
func (h *Handler) StartAuction(c *gin.Context) {
	// 1. Parse the incoming JSON request into our model.
	var req models.StartAuctionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	// 2. Obtain transaction options by reading headers (private key, chain ID).
	auth, err := getTransactOpts(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to authenticate transaction", "details": err.Error()})
		return
	}

	// 3. Convert integer values to big.Int as required by the Ethereum bindings.
	openingBid := big.NewInt(req.OpeningBid)
	duration := big.NewInt(req.Duration)

	// 4. Call the Start method on the smart contract wrapper.
	// This broadcasts a transaction to the Ethereum network.
	tx, err := h.Contract.Start(auth, openingBid, duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start auction", "details": err.Error()})
		return
	}

	// 5. Return success response with the transaction hash.
	c.JSON(http.StatusOK, gin.H{
		"message": "Auction started successfully",
		"txHash":  tx.Hash().Hex(),
	})
}

// CreateBid handles the API request to place a bid on the active auction.
// Since bidding requires sending tokens, this is a payable transaction.
func (h *Handler) CreateBid(c *gin.Context) {
	// 1. Parse the incoming JSON request containing the bid amount.
	var req models.BidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	// 2. Obtain transaction options.
	auth, err := getTransactOpts(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to authenticate transaction", "details": err.Error()})
		return
	}

	// 3. IMPORTANT: For a payable function, we must specify the amount of tokens
	// we want to send within the Value field of the transaction options.
	auth.Value = big.NewInt(req.Amount)

	// 4. Call the Bid method on the smart contract wrapper.
	tx, err := h.Contract.Bid(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to place bid", "details": err.Error()})
		return
	}

	// 5. Return success response.
	c.JSON(http.StatusOK, gin.H{
		"message": "Bid placed successfully",
		"txHash":  tx.Hash().Hex(),
	})
}

// EndAuction handles the API request to end an active auction.
// It can only be called once the auction duration has passed.
func (h *Handler) EndAuction(c *gin.Context) {
	// 1. Obtain transaction options.
	auth, err := getTransactOpts(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to authenticate transaction", "details": err.Error()})
		return
	}

	// 2. Call the End method on the contract. No arguments are needed besides auth.
	tx, err := h.Contract.End(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to end auction", "details": err.Error()})
		return
	}

	// 3. Return success response.
	c.JSON(http.StatusOK, gin.H{
		"message": "Auction ended successfully",
		"txHash":  tx.Hash().Hex(),
	})
}

// Withdraw handles the API request for users to withdraw their funds.
// Bidders whose bids were outbid can withdraw their tokens.
func (h *Handler) Withdraw(c *gin.Context) {
	// 1. Obtain transaction options.
	auth, err := getTransactOpts(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to authenticate transaction", "details": err.Error()})
		return
	}

	// 2. Call the Withdraw method.
	tx, err := h.Contract.Withdraw(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to withdraw funds", "details": err.Error()})
		return
	}

	// 3. Return success response.
	c.JSON(http.StatusOK, gin.H{
		"message": "Withdrawal successful",
		"txHash":  tx.Hash().Hex(),
	})
}
