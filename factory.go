package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func NewBuilder[T any](table string) *Builder[T] {
	return &Builder[T]{
		table:    table,
		Selects:  []string{fmt.Sprintf("%s.*", table)},
		orderBy:  OrderByFields{},
		bindings: map[bindingType][]any{},
		joins:    Joins{},
		unions:   Unions[T]{},
		groupBy:  GroupBy{},
		wheres: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
		having: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
	}
}

func New[T any](table string) contracts.QueryBuilder[T] {
	return &Builder[T]{
		table:    table,
		Selects:  []string{fmt.Sprintf("%s.*", table)},
		orderBy:  OrderByFields{},
		bindings: map[bindingType][]any{},
		joins:    Joins{},
		unions:   Unions[T]{},
		groupBy:  GroupBy{},
		wheres: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
		having: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
	}
}

func FromSub[T any](callback contracts.QueryProvider[T], as string) contracts.QueryBuilder[T] {
	subQuery := callback()
	builder := NewBuilder[T](as)
	builder.fromSub = &fromSub[T]{as: as, subQuery: subQuery}
	return builder.addBinding(fromBinding, subQuery.GetBindings()...)
}

func FromQuery[T any](subQuery contracts.QueryBuilder[T], as string) contracts.QueryBuilder[T] {
	builder := NewBuilder[T]("")
	builder.fromSub = &fromSub[T]{as: as, subQuery: subQuery}
	return builder.addBinding(fromBinding, subQuery.GetBindings()...)
}
