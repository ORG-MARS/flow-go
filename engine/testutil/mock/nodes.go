package mock

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/rs/zerolog"

	collectioningest "github.com/dapperlabs/flow-go/engine/collection/ingest"
	"github.com/dapperlabs/flow-go/engine/collection/provider"
	consensusingest "github.com/dapperlabs/flow-go/engine/consensus/ingestion"
	"github.com/dapperlabs/flow-go/engine/consensus/propagation"
	"github.com/dapperlabs/flow-go/engine/execution/blocks"
	"github.com/dapperlabs/flow-go/engine/execution/execution"
	"github.com/dapperlabs/flow-go/engine/execution/execution/state"
	"github.com/dapperlabs/flow-go/engine/execution/execution/virtualmachine"
	"github.com/dapperlabs/flow-go/engine/execution/receipts"
	"github.com/dapperlabs/flow-go/engine/verification/verifier"
	"github.com/dapperlabs/flow-go/module"
	"github.com/dapperlabs/flow-go/module/mempool"
	"github.com/dapperlabs/flow-go/network/stub"
	"github.com/dapperlabs/flow-go/protocol"
	"github.com/dapperlabs/flow-go/storage"
	"github.com/dapperlabs/flow-go/storage/ledger/databases/leveldb"
)

// GenericNode implements a generic in-process node for tests.
type GenericNode struct {
	Log   zerolog.Logger
	DB    *badger.DB
	State protocol.State
	Me    module.Local
	Net   *stub.Network
}

// CollectionNode implements an in-process collection node for tests.
type CollectionNode struct {
	GenericNode
	Pool            mempool.Transactions
	Collections     storage.Collections
	IngestionEngine *collectioningest.Engine
	ProviderEngine  *provider.Engine
}

// ConsensusNode implements an in-process consensus node for tests.
type ConsensusNode struct {
	GenericNode
	Pool              mempool.Guarantees
	IngestionEngine   *consensusingest.Engine
	PropagationEngine *propagation.Engine
}

// ExecutionNode implements a mocked execution node for tests.
type ExecutionNode struct {
	GenericNode
	BlocksEngine    *blocks.Engine
	ExecutionEngine *execution.Engine
	ReceiptsEngine  *receipts.Engine
	BadgerDB        *badger.DB
	LevelDB         *leveldb.LevelDB
	VM              virtualmachine.VirtualMachine
	State           state.ExecutionState
}

// VerificationNode implements an in-process verification node for tests.
type VerificationNode struct {
	GenericNode
	Receipts       mempool.Receipts
	Blocks         mempool.Blocks
	Collections    mempool.Collections
	VerifierEngine *verifier.Engine
}
