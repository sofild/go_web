package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	//"strings"
	"reflect"
)

type DbConf struct {
	Host string
	Port string
	Db   string
	User string
	Pass string
}

type Table struct {
	Config    DbConf
	Name      string
	Condition string
	Field     []interface{}
	Value     []interface{}
	Order     string
	Limit     string
}

var db *sql.DB

func (table *Table) Conn() *sql.DB {
	dbConf := table.Config
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Db)
	fmt.Println(conn)
	var err error
	db, err = sql.Open("mysql", conn)
	checkErr(err)
	return db
}

func (table *Table) Find() {
	if table.Condition == "" {
		table.Condition = "1=1"
	}
	if table.Order == "" {
		table.Order = "id desc"
	}
	sql := fmt.Sprintf("select * from %s where %s order by %s limit 1", table.Name, table.Condition, table.Order)
	rows, err1 := db.Query(sql)
	checkErr(err1)
	defer rows.Close()
    fields := make([]interface{}, len(table.Field))
    for idx,_ := range fields {
        var field interface{}
        fields[idx] = &field
    }
	for rows.Next() {
		data := make(map[string]string)
        err := rows.Scan(fields...)
		checkErr(err)
		for i, key := range table.Field {
			val := reflect.Indirect(reflect.ValueOf(fields[i])).Interface()
            var value string
            switch val.(type){
                case []byte:
                    if v, ok := val.([]byte); ok {
                        value = string(v)
                    }
                case string:
                    if v, ok := val.(string); ok {
                        value = v
                    }
                case int:
                    if v, ok := val.(int); ok {
                        value = strconv.Itoa(v)
                    }
                case int64:
                    if v, ok := val.(int64); ok {
                        value = strconv.FormatInt(v, 10)
                    }
            }
            data[key] = value
            fmt.Println(key, value)
		}

	}
	return data
}

/*
func (table *Table) Select() []map[string]string {
	if table.Condition == "" {
		table.Condition = "1=1"
	}
	if table.Order == "" {
		table.Order = "id desc"
	}
	sql := fmt.Sprintf("select * from %s where %s order by %s limit %s", table.Name, table.Condition, table.Order, table.Limit)
	fmt.Println(sql)
	rows, err1 := db.Query(sql)
	checkErr(err1)
	defer rows.Close()
	var datas []map[string]string
	for rows.Next() {
		data := make(map[string]string)
		err := rows.Scan(table.Field...)
		checkErr(err)
		for k, v := range table.Value {
			key := table.Field[k]
			switch vType := v.(type) {
			case string:
				data[key] = v
			case int:
				data[key] = strconv.Itoa(v)
			case int64:
				data[key] = strconv.FormatInt(v, 10)
			default:
				data[key] = v
			}
		}
		datas = append(datas, data)
	}
	return datas
}

//新增
func (table *Table) Add() int64 {
	sql := fmt.Sprintf("insert into %s (%s) values (?,?,?)", table.Name, strings.Join(table.Field, ","))
	fmt.Println(sql)
	stmt, err1 := db.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(table.Value...)
	checkErr(err2)
	id, err3 := res.LastInsertId()
	checkErr(err3)
	return id
}

//更新
func (table *Table) Update() int64 {
	var st []string
	for k, v := range table.Field {
		param := fmt.Sprintf("%s='%s'", k, table.Value[k])
		st = append(st, param)
	}
	params := strings.Join(st, ",")
	sql := fmt.Sprintf("update %s set %s where %s", table.Name, params, table.Condition)
	fmt.Println(sql)
	res, err := db.Exec(sql)
	checkErr(err)
	num, err1 := res.RowsAffected()
	checkErr(err1)
	return num
}

//删除
func (table *Table) Del() int64 {
	sql := fmt.Sprintf("Delete from %s where %s", table.Name, table.Condition)
	fmt.Println(sql)
	res, err1 := stmt.Exec(sql)
	checkErr(err1)
	num, err2 := res.RowsAffected()
	checkErr(err2)
	return num
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
*/
