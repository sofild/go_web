package main

import (
	"crypto/md5"
	"fmt"
)

func getUser(username string, password string) map[string]string {
	pass := []byte(password)
	pwd := md5.Sum(pass)
	pwd_str := fmt.Sprintf("%x", pwd)
	table.Name = "user"
	table.SetField("id", "username", "password")
	table.Order = "id desc"
	table.Condition = fmt.Sprintf("username='%s' and password='%s'", username, pwd_str)
	data := table.Find()
	return data
}
