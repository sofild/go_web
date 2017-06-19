package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

const (
	host          = "192.168.80.131"
	port          = "3306"
	dbName        = "gotest"
	dbUser        = "root"
	dbPass        = "root"
	cateTable     = "cate"
	articlesTable = "articles"
	photosTable   = "photos"
	userTable     = "user"
)

var myDb *sql.DB

func init() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, host, port, dbName)
	var err error
	myDb, err = sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("mysql connect error:", err.Error())
	}
}

//获取分类
func GetCate(where string, order string, limit string) []map[string]string {
	if where == "" {
		where = "1=1"
	}
	if order == "" {
		order = "id desc"
	}
	sql := fmt.Sprintf("select * from %s where %s order by %s", cateTable, where, order)
	if limit != "" {
		sql = fmt.Sprintf("%s limit %s", sql, limit)
	}
	fmt.Println(sql)
	rows, err1 := myDb.Query(sql)
	checkErr(err1)
	defer rows.Close()
	var datas []map[string]string
	for rows.Next() {
		data := make(map[string]string)
		var id int64
		var parent_id int64
		var name string
		var addtime int64
		err := rows.Scan(&id, &parent_id, &name, &addtime)
		checkErr(err)
		data["id"] = strconv.FormatInt(id, 10)
		data["parent_id"] = strconv.FormatInt(parent_id, 10)
		data["name"] = name
		data["addtime"] = strconv.FormatInt(addtime, 10)
		datas = append(datas, data)
	}
	return datas
}

//新增分类
func AddCate(parent_id int, name string) int64 {
	sql := fmt.Sprintf("insert into %s (parent_id, name, addtime) values (?,?,?)", cateTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(parent_id, name, time.Now().Unix())
	checkErr(err2)
	id, err3 := res.LastInsertId()
	checkErr(err3)
	return id
}

//更新分类
func UpCate(data map[string]string, id int) int64 {
	var st []string
	for k, v := range data {
		param := fmt.Sprintf("%s='%s'", k, v)
		st = append(st, param)
	}
	params := strings.Join(st, ",")
	id_str := strconv.Itoa(id)
	sql := fmt.Sprintf("update %s set %s where id=%s", cateTable, params, id_str)
	fmt.Println(sql)
	res, err := myDb.Exec(sql)
	checkErr(err)
	num, err1 := res.RowsAffected()
	checkErr(err1)
	return num
}

//删除分类
func DelCate(id int) int64 {
	sql := fmt.Sprintf("Delete from %s where id=?", cateTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(id)
	checkErr(err2)
	num, err3 := res.RowsAffected()
	checkErr(err3)
	return num
}

//获取文章
func GetArticles(where string, order string, limit string) []map[string]string {
	if where == "" {
		where = "1=1"
	}
	if order == "" {
		order = "id desc"
	}
	sql := fmt.Sprintf("select * from %s where %s order by %s", articlesTable, where, order)
	if limit != "" {
		sql = fmt.Sprintf("%s limit %s", sql, limit)
	}
	fmt.Println(sql)
	rows, err1 := myDb.Query(sql)
	checkErr(err1)
	defer rows.Close()
	var datas []map[string]string
	for rows.Next() {
		data := make(map[string]string)
		var id int64
		var cateid int64
		var title string
		var description string
		var pic string
		var content string
		var addtime int64
		err := rows.Scan(&id, &cateid, &title, &description, &pic, &content, &addtime)
		checkErr(err)
		data["id"] = strconv.FormatInt(id, 10)
		data["cateid"] = strconv.FormatInt(cateid, 10)
		data["title"] = title
		data["description"] = description
		data["pic"] = pic
		data["content"] = content
		data["addtime"] = strconv.FormatInt(addtime, 10)
		datas = append(datas, data)
	}
	return datas
}

//新增文章
func AddArticles(cateid int, title string, description string, pic string, content string) int64 {
	sql := fmt.Sprintf("insert into %s (cateid, title, description, pic, content, addtime) values (?,?,?,?,?,?)", articlesTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(cateid, title, description, pic, content, time.Now().Unix())
	checkErr(err2)
	id, err3 := res.LastInsertId()
	checkErr(err3)
	return id
}

//更新文章
func UpArticles(data map[string]string, id int) int64 {
	var st []string
	for k, v := range data {
		param := fmt.Sprintf("%s='%s'", k, v)
		st = append(st, param)
	}
	params := strings.Join(st, ",")
	id_str := strconv.Itoa(id)
	sql := fmt.Sprintf("update %s set %s where id=%s", articlesTable, params, id_str)
	fmt.Println(sql)
	res, err := myDb.Exec(sql)
	checkErr(err)
	num, err1 := res.RowsAffected()
	checkErr(err1)
	return num
}

//删除文章
func DelArticles(id int) int64 {
	sql := fmt.Sprintf("Delete from %s where id=?", articlesTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(id)
	checkErr(err2)
	num, err3 := res.RowsAffected()
	checkErr(err3)
	return num
}

//获取图片
func GetPotos(where string, order string, limit string) []map[string]string {
	if where == "" {
		where = "1=1"
	}
	if order == "" {
		order = "id desc"
	}
	sql := fmt.Sprintf("select * from %s where %s order by %s", photosTable, where, order)
	if limit != "" {
		sql = fmt.Sprintf("%s limit %s", sql, limit)
	}
	fmt.Println(sql)
	rows, err1 := myDb.Query(sql)
	checkErr(err1)
	defer rows.Close()
	var datas []map[string]string
	for rows.Next() {
		data := make(map[string]string)
		var id int64
		var cateid int64
		var image string
		var title string
		var addtime int64
		err := rows.Scan(&id, &cateid, &image, &title, &addtime)
		checkErr(err)
		data["id"] = strconv.FormatInt(id, 10)
		data["cateid"] = strconv.FormatInt(cateid, 10)
		data["image"] = image
		data["title"] = title
		data["addtime"] = strconv.FormatInt(addtime, 10)
		datas = append(datas, data)
	}
	return datas
}

//新增图片
func AddPhotos(cateid int, image string, title string) int64 {
	sql := fmt.Sprintf("insert into %s (cateid, image, title, addtime) values (?,?,?)", cateTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(cateid, image, title, time.Now().Unix())
	checkErr(err2)
	id, err3 := res.LastInsertId()
	checkErr(err3)
	return id
}

//更新图片
func UpPhotos(data map[string]string, id int) int64 {
	var st []string
	for k, v := range data {
		param := fmt.Sprintf("%s='%s'", k, v)
		st = append(st, param)
	}
	params := strings.Join(st, ",")
	id_str := strconv.Itoa(id)
	sql := fmt.Sprintf("update %s set %s where id=%s", photosTable, params, id_str)
	fmt.Println(sql)
	res, err := myDb.Exec(sql)
	checkErr(err)
	num, err1 := res.RowsAffected()
	checkErr(err1)
	return num
}

//删除图片
func DelPhotos(id int) int64 {
	sql := fmt.Sprintf("Delete from %s where id=?", photosTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(id)
	checkErr(err2)
	num, err3 := res.RowsAffected()
	checkErr(err3)
	return num
}

//获取分类
func GetUser(where string, order string, limit string) []map[string]string {
	if where == "" {
		where = "1=1"
	}
	if order == "" {
		order = "id desc"
	}
	sql := fmt.Sprintf("select * from %s where %s order by %s", userTable, where, order)
	if limit != "" {
		sql = fmt.Sprintf("%s limit %s", sql, limit)
	}
	fmt.Println(sql)
	rows, err1 := myDb.Query(sql)
	checkErr(err1)
	defer rows.Close()
	var datas []map[string]string
	for rows.Next() {
		data := make(map[string]string)
		var id int64
		var username string
		var password string
		var email string
		var avatar string
		var addtime int64
		var logintime int64
		err := rows.Scan(&id, &username, &password, &email, &avatar, &addtime, &logintime)
		checkErr(err)
		data["id"] = strconv.FormatInt(id, 10)
		data["username"] = username
		data["password"] = password
		data["email"] = email
		data["avatar"] = avatar
		data["addtime"] = strconv.FormatInt(addtime, 10)
		data["logintime"] = strconv.FormatInt(logintime, 10)
		datas = append(datas, data)
	}
	return datas
}

//新增分类
func AddUser(username int, password string, email string, avatar string) int64 {
	sql := fmt.Sprintf("insert into %s (username, password, email, avatar, addtime, logintime) values (?,?,?,?,?,?)", userTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(username, password, email, avatar, time.Now().Unix(), time.Now().Unix())
	checkErr(err2)
	id, err3 := res.LastInsertId()
	checkErr(err3)
	return id
}

//更新分类
func UpUser(data map[string]string, id int) int64 {
	var st []string
	for k, v := range data {
		param := fmt.Sprintf("%s='%s'", k, v)
		st = append(st, param)
	}
	params := strings.Join(st, ",")
	id_str := strconv.Itoa(id)
	sql := fmt.Sprintf("update %s set %s where id=%s", userTable, params, id_str)
	fmt.Println(sql)
	res, err := myDb.Exec(sql)
	checkErr(err)
	num, err1 := res.RowsAffected()
	checkErr(err1)
	return num
}

//删除分类
func DelUser(id int) int64 {
	sql := fmt.Sprintf("Delete from %s where id=?", userTable)
	fmt.Println(sql)
	stmt, err1 := myDb.Prepare(sql)
	checkErr(err1)
	res, err2 := stmt.Exec(id)
	checkErr(err2)
	num, err3 := res.RowsAffected()
	checkErr(err3)
	return num
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
