package bunmodel

import (
	"context"

	"github.com/uptrace/bun"
)

const (
	SepAND = " AND "
	SepOR  = " OR "
)

func UpdateColumns(ctx context.Context, db bun.IDB, model any, columns []string, options ...QueryOption) error {
	q := db.NewUpdate().Model(model).WherePK()

	if len(columns) == 0 { // just touch
		q.Set("id = id")
	} else {
		q.Column(columns...)
	}

	_, err := QueryOptions(q, options...).Exec(ctx)
	return err
}
