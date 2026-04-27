# 🚀 English Auction - Go API Service

This is the Go backend API for interacting with the **English Auction** smart contract. It provides a RESTful API built with **Gin** and uses `go-ethereum` (abigen bindings) to communicate with the deployed smart contract on the blockchain.

## 📁 Project Structure

```
├── api/
│   ├── auction/          # Auto-generated Go bindings (from abigen)
│   ├── handlers/         # API Endpoint logic & Integration Tests
│   ├── models/           # JSON request payload structures
│   ├── main.go           # Entry point (Server, Router, EthClient setup)
│   ├── go.mod            # Go modules
│   └── .gitignore        # Git ignore for compiled binaries
```

## 🛠️ Prerequisites

- **Go 1.21+** installed
- A running Ethereum node (e.g., **Anvil**)
- The smart contract deployed to the network (You need its `CONTRACT_ADDRESS`)

---

## 🚀 Quick Setup & Run

### 1. Start the Local Blockchain & Deploy Contract
(This is done in the Solidity repository: [go-solidity-contracts](https://github.com/Atharva0506/go-solidity-contracts))
```bash
# In your solidity repo terminal:
anvil

# In another terminal in the solidity repo:
forge script script/EnglishAuction.s.sol:EnglishAuctionScript --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
```
*Note the `EnglishAuction` deployed address from the output!*

### 2. Configure Environment & Run Server
Navigate to the `api` directory in this Go repository:
```bash
cd api/
```

Set the environment variables. Replace the contract address with the one from your deployment step:
```powershell
# For PowerShell
$env:CONTRACT_ADDRESS="0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
$env:ETH_NODE_URL="http://127.0.0.1:8545"
$env:CGO_ENABLED="0"

# Run the server
go run main.go
```
*The server will start on `http://localhost:8080`.*

---

## 🧪 Testing the API

### Option 1: Automated Integration Tests
Make sure the server is NOT running, but Anvil IS running and the contract is freshly deployed (no active auction yet).
```powershell
cd api/
$env:CONTRACT_ADDRESS="0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
$env:CGO_ENABLED="0"

# Run the tests
go test ./handlers/... -v
```

### Option 2: Manual cURL Tests (End-to-End)

Once the server is running (`go run main.go`), open a new terminal and run these commands. 
*(Note: We use Anvil's default test private keys).*

**1. Start Auction** (Deployer: Account 0)
```powershell
curl -X POST http://localhost:8080/api/auction/start `
  -H "X-Private-Key: ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" `
  -H "X-Chain-ID: 31337" `
  -H "Content-Type: application/json" `
  -d '{"openingBid": 100, "duration": 3600}'
```

**2. Place a Bid** (Bidder 1: Account 1)
```powershell
curl -X POST http://localhost:8080/api/auction/bid `
  -H "X-Private-Key: 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" `
  -H "X-Chain-ID: 31337" `
  -H "Content-Type: application/json" `
  -d '{"amount": 200}'
```

**3. Place a Higher Bid** (Bidder 2: Account 2)
```powershell
curl -X POST http://localhost:8080/api/auction/bid `
  -H "X-Private-Key: 5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a" `
  -H "X-Chain-ID: 31337" `
  -H "Content-Type: application/json" `
  -d '{"amount": 300}'
```

**4. Withdraw Outbid Funds** (Bidder 1 gets their 200 Wei back)
```powershell
curl -X POST http://localhost:8080/api/auction/withdraw `
  -H "X-Private-Key: 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" `
  -H "X-Chain-ID: 31337" `
  -H "Content-Type: application/json" `
  -d '{}'
```

---

## 📖 API Endpoints Summary

| Method | Endpoint | Payload | Description |
|--------|----------|---------|-------------|
| POST | `/api/auction/start` | `{"openingBid": 100, "duration": 3600}` | Starts the auction |
| POST | `/api/auction/bid` | `{"amount": 200}` | Places a bid (payable) |
| POST | `/api/auction/end` | `{}` | Ends the auction (after duration) |
| POST | `/api/auction/withdraw`| `{}` | Withdraws outbid funds |

*All endpoints require `X-Private-Key` and `X-Chain-ID` headers to authorize the transaction on the blockchain.*
