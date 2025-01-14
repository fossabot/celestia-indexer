package rollback

import (
	"context"

	"github.com/dipdup-io/celestia-indexer/internal/storage/postgres"
	"github.com/dipdup-io/celestia-indexer/pkg/types"
)

func (module *Module) rollbackTransactions(ctx context.Context, tx postgres.Transaction, height types.Level) error {
	txs, err := tx.RollbackTxs(ctx, height)
	if err != nil {
		return nil
	}

	if len(txs) == 0 {
		return nil
	}

	ids := make([]uint64, len(txs))
	for i := range txs {
		ids[i] = txs[i].Id
	}

	if err := tx.RollbackSigners(ctx, ids); err != nil {
		return err
	}

	return nil
}
