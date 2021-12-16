package ngtypes

import "github.com/ngchain/secp256k1"

type Tx interface {
	GetHash() []byte
	IsSigned() bool
}

type Block interface {
	IsUnsealing() bool
	IsSealed() bool
	IsGenesis() bool
	GetPrevHash() []byte
	GetHash() []byte
	GetHeight() uint64
	GetTx(int) Tx
	GetTimestamp() uint64
}

type Chain interface {
	CheckBlock(Block) error
	GetLatestBlock() Block
	GetLatestBlockHash() []byte
	GetLatestBlockHeight() uint64
	GetBlockByHeight(height uint64) (Block, error)
	GetBlockByHash(hash []byte) (Block, error)
}

// Consensus is an abstract consensus interface.
type Consensus interface {
	GoLoop()
	GetChain() Chain
	ImportBlock(Block) error
	GetBlockTemplate(privateKey *secp256k1.PrivateKey) Block
}
