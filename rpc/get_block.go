package rpc

import "context"

type GetBlockResponse JsonRpcResponse[GetBlock]

type GetBlock struct {
	Blockhash         string                `json:"blockhash"`
	BlockTime         *int64                `json:"blockTime"`
	BlockHeight       *int64                `json:"blockHeight"`
	PreviousBlockhash string                `json:"previousBlockhash"`
	ParentSlot        uint64                `json:"parentSlot"`
	Transactions      []GetBlockTransaction `json:"transactions"`
	Signatures        []string              `json:"signatures"`
	Rewards           []Reward              `json:"rewards"`
}

type GetBlockTransaction struct {
	Transaction any              `json:"transaction"`
	Meta        *TransactionMeta `json:"meta"`
	Version     any              `json:"version"`
}

type GetBlockTransactionDetail struct {
	Message    any      `json:"message"`
	Signatures []string `json:"signatures"`
}

type GetBlockTransactionMessage struct {
	AccountKeys         []AccountKey `json:"accountKeys"`
	AddressTableLookups []any        `json:"addressTableLookups"`
	Instructions        []any        `json:"instructions"`
	RecentBlockhash     string       `json:"recentBlockhash"`
}

type AccountKey struct {
	Pubkey   string `json:"pubkey"`
	Signer   bool   `json:"signer"`
	Source   string `json:"source"`
	Writable bool   `json:"writable"`
}

type GetBlockConfig struct {
	Encoding                       GetBlockConfigEncoding           `json:"encoding,omitempty"`                       // default: "json"
	TransactionDetails             GetBlockConfigTransactionDetails `json:"transactionDetails,omitempty"`             // default: "full", either "full", "signatures", "none"
	Rewards                        *bool                            `json:"rewards,omitempty"`                        // default: true
	Commitment                     Commitment                       `json:"commitment,omitempty"`                     // "processed" is not supported
	MaxSupportedTransactionVersion *uint8                           `json:"maxSupportedTransactionVersion,omitempty"` // default: nil legacy only
}

type GetBlockConfigEncoding string

const (
	GetBlockConfigEncodingJson       GetBlockConfigEncoding = "json"
	GetBlockConfigEncodingJsonParsed GetBlockConfigEncoding = "jsonParsed"
	GetBlockConfigEncodingBase58     GetBlockConfigEncoding = "base58"
	GetBlockConfigEncodingBase64     GetBlockConfigEncoding = "base64"
)

type GetBlockConfigTransactionDetails string

const (
	GetBlockConfigTransactionDetailsFull       GetBlockConfigTransactionDetails = "full"
	GetBlockConfigTransactionDetailsSignatures GetBlockConfigTransactionDetails = "signatures"
	GetBlockConfigTransactionDetailsNone       GetBlockConfigTransactionDetails = "none"
)

// GetBlock returns identity and transaction information about a confirmed block in the ledger
func (c *RpcClient) GetBlock(ctx context.Context, slot uint64) (JsonRpcResponse[*GetBlock], error) {
	return call[JsonRpcResponse[*GetBlock]](c, ctx, "getBlock", slot)
}

// GetBlockWithConfig returns identity and transaction information about a confirmed block in the ledger
func (c *RpcClient) GetBlockWithConfig(ctx context.Context, slot uint64, cfg GetBlockConfig) (JsonRpcResponse[*GetBlock], error) {
	return call[JsonRpcResponse[*GetBlock]](c, ctx, "getBlock", slot, cfg)
}
