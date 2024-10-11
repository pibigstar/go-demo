package main

import (
	"fmt"

	"github.com/xwb1989/sqlparser"
)

func main() {
	sql := `select
  shop_id,
  count(*) as __value__
from
  app_cse_dehc_voc_gather_di t1 left join app_cse t2 on t1.shop_id = t2.shop_id
where
  t1.date = '${date}'
  and t1.label = '赠品问题'
  and t2.date = '${date}'
group by
  shop_id
order by __value__ desc limit 2000`
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		panic(err)
	}

	// 判断语句类型是否为 SELECT
	sel, ok := stmt.(*sqlparser.Select)
	if !ok {
		panic("Not a SELECT statement")
	}

	var tableNames []string
	// 获取所有表名
	for _, f := range sel.From {
		switch expr := f.(type) {
		case *sqlparser.AliasedTableExpr:
			if t := getTable(expr); t != "" {
				tableNames = append(tableNames, t)
			}
		case *sqlparser.JoinTableExpr:
			if at, ok := expr.LeftExpr.(*sqlparser.AliasedTableExpr); ok {
				if t := getTable(at); t != "" {
					tableNames = append(tableNames, t)
				}
			}
			if at, ok := expr.RightExpr.(*sqlparser.AliasedTableExpr); ok {
				if t := getTable(at); t != "" {
					tableNames = append(tableNames, t)
				}
			}
		}
	}
	fmt.Println(tableNames)

	var columns []string
	// 获取 SELECT 子句中的列名
	for _, expr := range sel.SelectExprs {
		switch expr := expr.(type) {
		case *sqlparser.AliasedExpr:
			// 如果有别名，则获取别名
			if asName := expr.As.String(); asName != "" {
				columns = append(columns, asName)
				continue
			}
			if col, ok := expr.Expr.(*sqlparser.ColName); ok {
				columns = append(columns, col.Name.String())
			}
		}
	}
	fmt.Println(columns)

	// 获取 SELECT 子句中的列名
	where := sqlparser.String(sel.Where)
	fmt.Println(where)

	var groupby []string
	// 获取 group by字段
	for _, expr := range sel.GroupBy {
		switch expr := expr.(type) {
		case *sqlparser.ColName:
			groupby = append(groupby, expr.Name.String())
		}
	}
	fmt.Println(groupby)

	// 新增where 条件
	additionalCondition := "t2.name = '1111' and start_time > '111'"
	sql = AddConditionToWhere(sql, additionalCondition)
	fmt.Println(sql)
}

func getTable(expr *sqlparser.AliasedTableExpr) string {
	if t, ok := expr.Expr.(sqlparser.TableName); ok {
		return t.Name.String()
	}
	return ""
}

func AddConditionToWhere(sql string, additionalCondition string) string {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return ""
	}

	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		newStm, err := sqlparser.Parse(fmt.Sprintf("select * from t Where %s", additionalCondition))
		if err != nil {
			return sql
		}
		sel, _ := newStm.(*sqlparser.Select)
		stmt.AddWhere(sel.Where.Expr)

		sql = sqlparser.String(stmt)
	}
	return sql
}
