package main

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

/**
*  @Author: leikewei
*  @Date: 2024/4/19
*  @Desc: sql构建
 */

type SQLBuilderReq struct {
	Table   string
	Select  []string
	Where   []string
	GroupBy []string
	OrderBy []string
	Having  []string
	Limit   int
	Offset  int
	Join    *Join
}

type Join struct {
	Table  string
	Option string // left right
	OnExpr []string
}

func SqlBuilder(req *SQLBuilderReq) (string, error) {
	if len(req.Select) == 0 {
		return "", fmt.Errorf("select is empty")
	}
	if len(req.Table) == 0 {
		return "", fmt.Errorf("table is empty")
	}

	sb := sqlbuilder.NewSelectBuilder().From(req.Table)

	sb.Select(req.Select...)

	if len(req.Where) > 0 {
		sb.Where(req.Where...)
	}

	if req.Limit > 0 {
		sb.Limit(req.Limit)
	}

	if req.Offset > 0 {
		sb.Offset(req.Offset)
	}

	if len(req.GroupBy) > 0 {
		sb.GroupBy(req.GroupBy...)
	}

	if len(req.Having) > 0 {
		sb.Having(req.Having...)
	}

	if len(req.OrderBy) > 0 {
		sb.OrderBy(req.OrderBy...)
	}

	if req.Join != nil && len(req.Join.Table) > 0 {
		sb.JoinWithOption(sqlbuilder.JoinOption(req.Join.Option), req.Join.Table, req.Join.OnExpr...)
	}

	return sb.String(), nil
}

func main() {
	req := &SQLBuilderReq{
		Table: "user",
		Select: []string{
			"user_id",
			"count(*) as product_cnt",
			"sum(if(status=1), 1, 0) as unequal_product_cnt",
		},
		Where: []string{
			"date=20240405",
			"name='hhh'",
		},
		Limit:  10,
		Offset: 20,
		GroupBy: []string{
			"user_id",
			"hello",
		},
		OrderBy: []string{
			"user_id desc",
		},
		Having: []string{
			"user_id > 100",
		},
		Join: &Join{
			Table:  "class c",
			Option: "right",
			OnExpr: []string{
				"a.id = c.id",
			},
		},
	}
	sql, err := SqlBuilder(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(sql)
}
