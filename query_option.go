package bunmodel

import (
	"github.com/uptrace/bun"
)

type QueryOption func(q bun.Query)

func Relations(relations ...string) QueryOption {
	return func(q bun.Query) {
		if selectQuery, ok := q.(*bun.SelectQuery); ok {
			for _, relation := range relations {
				selectQuery.Relation(relation)
			}
		} else {
			panic("Relations only works with SelectQuery")
		}
	}
}

func selectFor(_for string) QueryOption {
	return func(q bun.Query) {
		if selectQuery, ok := q.(*bun.SelectQuery); ok {
			selectQuery.For(_for + " OF ?TableAlias")
		} else {
			panic("SelectForXXX only works with SelectQuery")
		}
	}
}

func SelectForUpdate() QueryOption { return selectFor("UPDATE") }

func WhereDeleted() QueryOption {
	return func(q bun.Query) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.WhereDeleted()
		case *bun.UpdateQuery:
			q.WhereDeleted()
		case *bun.DeleteQuery:
			q.WhereDeleted()
		default:
			panic("WhereDeleted only works with SelectQuery, UpdateQuery, DeleteQuery")
		}
	}
}

func WhereAllWithDeleted() QueryOption {
	return func(q bun.Query) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.WhereAllWithDeleted()
		case *bun.UpdateQuery:
			q.WhereAllWithDeleted()
		case *bun.DeleteQuery:
			q.WhereAllWithDeleted()
		default:
			panic("WhereAllWithDeleted only works with SelectQuery, UpdateQuery, DeleteQuery")
		}
	}
}

func Limit(limit int) QueryOption {
	return func(q bun.Query) {
		if selectQuery, ok := q.(*bun.SelectQuery); ok {
			selectQuery.Limit(limit)
		} else {
			panic("Limit only works with SelectQuery")
		}
	}
}

func Returning(ret string) QueryOption {
	return func(q bun.Query) {
		switch q := q.(type) {
		case *bun.InsertQuery:
			q.Returning(ret)
		case *bun.UpdateQuery:
			q.Returning(ret)
		case *bun.DeleteQuery:
			q.Returning(ret)
		default:
			panic("Returning only works with InsertQuery, UpdateQuery, DeleteQuery")
		}
	}
}

func ReturningAll() QueryOption { return Returning(RetAll) }

func QueryOptions[Q bun.Query](q Q, options ...QueryOption) Q {
	for _, opt := range options {
		opt(q)
	}

	return q
}
