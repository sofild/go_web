package main

import (
	//"fmt"
	"strconv"
	"time"
)

func getArticles(condition string, order string, limit string) []map[string]string {
	//获取文章
	table.Name = "articles"
	table.SetField("id", "cateid", "title", "description", "pic", "content", "addtime")
	table.Limit = limit
	table.Order = order
	table.Condition = condition
	blog_data := table.Select()
	cateids := make([]string, len(blog_data))
	for i, data := range blog_data {
		cateids[i] = data["cateid"]
	}
	cateNames := getCateNames(cateids)
	for k, datas := range blog_data {
		cateid := datas["cateid"]
		blog_data[k]["catename"] = cateNames[cateid]
		addtime, err := strconv.ParseInt(datas["addtime"], 10, 64)
		checkErr(err, 25)
		addtime_str := time.Unix(addtime, 0).Format("2006-01-02 15:04:05")
		blog_data[k]["addtime"] = addtime_str
	}
	return blog_data
}
