package main

import (
	"fmt"
	"strconv"
	"strings"
	//"time"
)

func getCate(id int64) map[string]string {
	table.Name = "cate"
	table.SetField("id", "parent_id", "name", "addtime")
	table.Limit = ""
	table.Order = "id desc"
	table.Condition = fmt.Sprintf("id=%s", strconv.FormatInt(id, 10))
	data := table.Find()
	return data
}

func getCateNames(ids []string) map[string]string {
	table.Name = "cate"
	table.SetField("id", "name")
	table.Order = "id desc"
	table.Condition = fmt.Sprintf("id in (%s)", strings.Join(ids, ","))
	cates := table.Select()
	data := make(map[string]string)
	for _, cate := range cates {
		cate_id := cate["id"]
		data[cate_id] = cate["name"]
	}
	return data
}

func getCateList() []map[string]string {
	table.Name = "cate"
	table.SetField("id", "parent_id", "name", "addtime")
	table.Order = "id desc"
	table.Condition = "parent_id=0"
	cates := table.Select()
	return cates
}
