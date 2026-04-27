package main

import (
	"api/auction"
	"api/handlers"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize the Ethereum Client.
	// We need to connect to an Ethereum node (e.g., local node like Anvil, or managed like Infura/Alchemy).
	// We read the RPC URL from an environment variable, falling back to a default localhost endpoint.
	nodeURL := os.Getenv("ETH_NODE_URL")
	if nodeURL == "" {
		nodeURL = "http://127.0.0.1:8545" // Default local blockchain node (e.g., Anvil, Ganache)
	}

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client at %s: %v", nodeURL, err)
	}
	defer client.Close()
	log.Printf("Connected to Ethereum node: %s", nodeURL)

	// 2. Initialize the Smart Contract Wrapper.
	// We need the deployed contract address. In a real scenario, this is obtained after deployment.
	contractAddressHex := os.Getenv("CONTRACT_ADDRESS")
	if contractAddressHex == "" {
		// Provide a placeholder or fail if it's strictly required.
		// For development, we'll just log a warning if it's not set.
		log.Println("WARNING: CONTRACT_ADDRESS environment variable is not set. Contract calls will fail.")
	}
	
	contractAddress := common.HexToAddress(contractAddressHex)

	// NewContract creates a new instance of the contract bound to the specific deployed address.
	contractWrapper, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate the smart contract wrapper: %v", err)
	}
	log.Printf("Smart contract wrapper initialized for address: %s", contractAddressHex)

	// 3. Initialize the Handlers.
	// We inject the contract wrapper into our handler struct.
	// This ensures all our API endpoints have access to the same contract instance.
	h := handlers.NewHandler(contractWrapper)

	// 4. Set up the Gin HTTP router.
	router := gin.Default()

	// 5. Define API Endpoints.
	// We map HTTP methods and paths to their respective handler functions.
	api := router.Group("/api/auction")
	{
		api.POST("/start", h.StartAuction)
		api.POST("/bid", h.CreateBid)
		api.POST("/end", h.EndAuction)
		api.POST("/withdraw", h.Withdraw)
	}

	// 6. Start the API server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting API server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
