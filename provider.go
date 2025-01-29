package siwe

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

// Provider allows making calls to the blockchain for contract verification
type Provider interface {
	// CallContract performs a contract call and returns the raw response
	CallContract(ctx context.Context, to common.Address, data []byte) ([]byte, error)
}

// ERC1271MagicValue is the value that must be returned by a contract's isValidSignature
// function to indicate a valid signature according to EIP-1271
var ERC1271MagicValue = [4]byte{0x16, 0x26, 0xba, 0x7e}

// IsValidSignatureSelector is the function selector for the ERC-1271 isValidSignature function
var IsValidSignatureSelector = [4]byte{0x16, 0x26, 0xba, 0x7e}
