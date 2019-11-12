package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type OrderStatus int32

const (
	UnPaid OrderStatus = 1 //待付款
	Cancel OrderStatus = 2 //已取消
	Paid   OrderStatus = 3 //已付款
	Closed OrderStatus = 4 //已关闭
)

type paramSqlPrepare struct {
	i    int
	sql  string
	args []string
	lens []int
}

func (p *paramSqlPrepare) replace(str string) string {
	var dstr []string
	l := 0
	for l < p.lens[p.i] {
		dstr = append(dstr, "?")
		l++
	}

	p.i++

	return strings.Join(dstr, ",")

}

func (p *paramSqlPrepare) prepare() (sql string, args []string) {

	for _, arg := range p.args {

		v_arr := strings.Split(arg, ",")
		l := len(v_arr)
		p.lens = append(p.lens, l)
		if l > 1 {
			for _, v := range v_arr {
				args = append(args, v)

			}
		} else {
			args = append(args, arg)
		}

	}

	rep, _ := regexp.Compile("\\?")
	sql = rep.ReplaceAllStringFunc(p.sql, p.replace)

	return sql, args

}
func main() {
	db, err := sql.Open("mysql", "root:sferwlwe@(host.docker.internal:3306)/user_db?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		panic(err)
	}
	sql := "select order_sn from tb_orders where order_status in(?)"

	psp := paramSqlPrepare{
		i:    0,
		sql:  sql,
		args: []string{"1,2"},
	}
	sql, args := psp.prepare()

	var iargs []interface{}
	for _, v := range args {
		iargs = append(iargs, v)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(iargs...) //查询待付款，已取消，已关闭
	defer rows.Close()
	for rows.Next() {
		var order_sn string
		if ers := rows.Scan(&order_sn); ers == nil {
			fmt.Println(order_sn)
		}
	}
}
