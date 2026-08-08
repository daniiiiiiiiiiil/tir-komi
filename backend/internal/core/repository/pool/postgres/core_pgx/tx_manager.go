package core_pgx

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

type PgxTxManager struct {
	pool postgres.Pool
}

func NewPgxTxManager(pool postgres.Pool) *PgxTxManager {
	return &PgxTxManager{pool: pool}
}

func (m *PgxTxManager) WithinTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return m.pool.BeginFunc(ctx, func(tx postgres.Tx) error {
		txCtx := WithTx(ctx, tx)
		return fn(txCtx)
	})
}
func (m *PgxTxManager) WithinTxIsolated(
	ctx context.Context,
	level postgres.IsolationLevel,
	fn func(ctx context.Context) error,
) error {
	opts := &postgres.TxOptions{IsolationLevel: level}

	return m.pool.BeginTxFunc(ctx, opts, func(tx postgres.Tx) error {
		txCtx := WithTx(ctx, tx)
		return fn(txCtx)
	})
}
