package models

// StartAuctionRequest represents the payload required to start an auction.
// It includes the initial opening bid and the duration of the auction.
type StartAuctionRequest struct {
	// OpeningBid is the starting amount for the auction, typically in wei.
	OpeningBid int64 `json:"openingBid" binding:"required"`
	// Duration is the length of the auction in seconds.
	Duration int64 `json:"duration" binding:"required"`
}

// BidRequest represents the payload for placing a new bid.
type BidRequest struct {
	// Amount is the number of tokens (in wei) the user wants to bid.
	// This corresponds to the 'value' field in an Ethereum payable transaction.
	Amount int64 `json:"amount" binding:"required"`
}

// Note: EndAuction and Withdraw do not require specific body payloads, 
// as they only need the transaction options (private key, chain ID) 
// to authorize the interaction with the smart contract.
