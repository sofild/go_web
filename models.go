package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strconv"
	"strings"
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
	Field     []string
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

func FormatValue(val interface{}) string {
	var value string
	switch val.(type) {
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
	return value
}

func (table *Table) Find() map[string]string {
	if table.Condition == "" {
		table.Condition = "1=1"
	}
	if table.Order == "" {
		table.Order = "id desc"
	}
	field_str := strings.Join(table.Field, ",")
	sql := fmt.Sprintf("select %s from %s where %s order by %s limit 1", field_str, table.Name, table.Condition, table.Order)
	rows, err1 := db.Query(sql)
	checkErr(err1)
	defer rows.Close()
	fields := make([]interface{}, len(table.Field))
	for idx, _ := range fields {
		var field interface{}
		fields[idx] = &field
	}
	data := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(fields...)
		checkErr(err)
		for i, key := range table.Field {
			val := reflect.Indirect(reflect.ValueOf(fields[i])).Interface()
			value := FormatValue(val)
			data[key] = value
		}
	}
	return data
}

func (table *Table) Select() []map[string]string {
	if table.Condition == "" {
		table.Condition = "1=1"
	}
	if table.Order == "" {
		table.Order = "id desc"
	}
	field_str := strings.Join(table.Field, ",")
	sql := fmt.Sprintf("select %s from %s where %s order by %s", field_str, table.Name, table.Condition, table.Order)
	if table.Limit != "" {
		sql = fmt.Sprintf("%s limit %s", sql, table.Limit)
	}
	fmt.Println(sql)
	rows, err1 := db.Query(sql)
	checkErr(err1)
	defer rows.Close()
	fields := make([]interface{}, len(table.Field))
	for idx, _ := range fields {
		var field interface{}
		fields[idx] = &field
	}
	var datas []map[string]string
	for rows.Next() {
		data := make(map[string]string)
		err := rows.Scan(fields...)
		checkErr(err)
		for i, key := range table.Field {
			val := reflect.Indirect(reflect.ValueOf(fields[i])).Interface()
			value := FormatValue(val)
			data[key] = value
		}
		datas = append(datas, data)
	}
	return datas
}

//新增
func (table *Table) Add() int64 {
	var values []string
	fieldLen := len(table.Field)
	for fieldLen > 0 {
		fieldLen--
		values = append(values, "?")
	}
	sql := fmt.Sprintf("insert into %s (%s) values (%s)", table.Name, strings.Join(table.Field, ","), strings.Join(values, ","))
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
	for _, k := range table.Field {
		param := fmt.Sprintf("%s=?", k)
		st = append(st, param)
	}
	params := strings.Join(st, ",")
	sql := fmt.Sprintf("update %s set %s where %s", table.Name, params, table.Condition)
	fmt.Println(sql)
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err2 := stmt.Exec(table.Value...)
	checkErr(err2)
	num, err1 := res.RowsAffected()
	checkErr(err1)
	return num
}

//删除
func (table *Table) Del() int64 {
	sql := fmt.Sprintf("Delete from %s where %s", table.Name, table.Condition)
	fmt.Println(sql)
	res, err1 := db.Exec(sql)
	checkErr(err1)
	num, err2 := res.RowsAffected()
	checkErr(err2)
	return num
}
