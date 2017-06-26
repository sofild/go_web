package main

import (
	"crypto/md5"
	"fmt"
	"time"
    "strconv"
)

func getUser(username string, password string) map[string]string {
	pass := []byte(password)
	pwd := md5.Sum(pass)
	pwd_str := fmt.Sprintf("%x", pwd)
	table.Name = "user"
	table.SetField("id", "username", "password", "avatar", "logintime")
	table.Order = "id desc"
	table.Condition = fmt.Sprintf("username='%s' and password='%s'", username, pwd_str)
	data := table.Find()
	logintime,err := strconv.ParseInt(data["logintime"], 10, 64)
    checkErr(err, 19)
    if logintime > 0 {
		data["logintime"] = time.Unix(logintime, 0).Format("2006-01-02 15:04:05")
	}
	return data
}

func updateLogin(id string) int64 {
	table.Name = "user"
	table.SetField("logintime")
	table.SetValue(time.Now().Unix())
	return table.Update()
}
