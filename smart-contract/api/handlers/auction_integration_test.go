package handlers_test

import (
	"api/auction"
	"api/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// setupTestRouter initializes the dependencies and Gin router for our integration tests.
func setupTestRouter(t *testing.T) *gin.Engine {
	// These variables are expected to be set for the integration test to run.
	contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
	if contractAddressHex == "" {
		t.Skip("Skipping integration test: CONTRACT_ADDRESS environment variable is not set")
	}

	nodeURL := os.Getenv("ETH_NODE_URL")
	if nodeURL == "" {
		nodeURL = "http://127.0.0.1:8545" // Default Anvil RPC
	}

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to eth client: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	contractWrapper, err := contract.NewContract(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to create contract wrapper: %v", err)
	}

	h := handlers.NewHandler(contractWrapper)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	
	api := router.Group("/api/auction")
	{
		api.POST("/start", h.StartAuction)
		api.POST("/bid", h.CreateBid)
		api.POST("/end", h.EndAuction)
		api.POST("/withdraw", h.Withdraw)
	}

	return router
}

func TestAuctionIntegration(t *testing.T) {
	router := setupTestRouter(t)

	// Anvil default accounts (Do not use in production!)
	deployerKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	bidder1Key := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	bidder2Key := "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
	chainID := "31337"

	// 1. Test Starting the Auction
	t.Run("Start Auction", func(t *testing.T) {
		payload := map[string]interface{}{
			"openingBid": 100,
			"duration":   3600,
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest(http.MethodPost, "/api/auction/start", bytes.NewBuffer(body))
		req.Header.Set("X-Private-Key", deployerKey)
		req.Header.Set("X-Chain-ID", chainID)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %d. Response: %s", w.Code, w.Body.String())
		}
	})

	// Wait briefly for transaction to be mined
	time.Sleep(1 * time.Second)

	// 2. Test Placing a Bid (User 1)
	t.Run("Place Bid User 1", func(t *testing.T) {
		payload := map[string]interface{}{
			"amount": 200,
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest(http.MethodPost, "/api/auction/bid", bytes.NewBuffer(body))
		req.Header.Set("X-Private-Key", bidder1Key)
		req.Header.Set("X-Chain-ID", chainID)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %d. Response: %s", w.Code, w.Body.String())
		}
	})

	time.Sleep(1 * time.Second)

	// 3. Test Placing a Higher Bid (User 2)
	t.Run("Place Bid User 2", func(t *testing.T) {
		payload := map[string]interface{}{
			"amount": 300,
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest(http.MethodPost, "/api/auction/bid", bytes.NewBuffer(body))
		req.Header.Set("X-Private-Key", bidder2Key)
		req.Header.Set("X-Chain-ID", chainID)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %d. Response: %s", w.Code, w.Body.String())
		}
	})

	time.Sleep(1 * time.Second)

	// 4. Test Withdrawing outbid funds (User 1)
	t.Run("Withdraw Funds User 1", func(t *testing.T) {
		payload := map[string]interface{}{} // empty body
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest(http.MethodPost, "/api/auction/withdraw", bytes.NewBuffer(body))
		req.Header.Set("X-Private-Key", bidder1Key)
		req.Header.Set("X-Chain-ID", chainID)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			// This might fail if the smart contract state wasn't successfully updated in previous tests.
			t.Fatalf("Expected status 200, got %d. Response: %s", w.Code, w.Body.String())
		}
	})
}
