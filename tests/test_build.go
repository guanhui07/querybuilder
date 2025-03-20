package main

import (
	"fmt"
	"github.com/goal-web/contracts"
	builder "github.com/goal-web/querybuilder"
	"github.com/gookit/goutil/dump"
)

// 数组
func main() {
	dealFunc12()
}

func dealFunc12() {
	TestInsert()
}

// where having group by order by limit
func TestInsert() {
	sql, bindings := builder.New[string]("users").CreateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})

	fmt.Println(sql)
	fmt.Println(bindings)

	dump.P(sql)

}

func TestUpdate() {
	sql, bindings := builder.New[string]("users").
		Where("id", ">", 1).
		UpdateSql(contracts.Fields{
			"name": "qbhy", "age": 18, "money": 100000000000,
		})
	fmt.Println(sql)
	fmt.Println(bindings)

	dump.P(sql)

}

func TestDelete() {
	sql, bindings := builder.New[string]("users").Where("id", ">", 1).DeleteSql()
	fmt.Println(sql)
	fmt.Println(bindings)

	dump.P(sql)

}

// where having group by order by limit
func TestSimpleQueryBuilder() {
	query := builder.New[string]("users").
		Where("name", "qbhy").
		Where("age", ">", 18).
		Where("gender", "!=", 0).
		OrWhere("amount", ">=", 100).
		OrderByDesc("age").
		OrderBy("id").
		Limit(10).
		WhereIsNull("avatar")

	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())

	query = builder.New[string]("users").
		Where("name", "qbhy").
		Where("age", ">", 18).
		Where("gender", "!=", 0).
		OrWhere("amount", ">=", 100).
		GroupBy("name").
		OrderByDesc("age").
		OrderBy("id").
		Limit(10).
		WhereIsNull("avatar")

	fmt.Println(query.ToSql())

	fmt.Println(query.GetBindings())
	// select * from users where name = ? and age > ? and gender != ? and avatar is null or amount >= ?
	// [qbhy 18 0 100]

	dump.P(query.ToSql())

}
