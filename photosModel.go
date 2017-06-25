package main

import (
	//"fmt"
	"strconv"
	//"strings"
	"time"
)

func getPhotoList(condition string, order string, limit string) []map[string]string {
	//获取图片
	table.Name = "photos"
	table.SetField("id", "cateid", "image", "title", "addtime")
	table.Limit = limit
	table.Order = order
	table.Condition = condition
	photo_data := table.Select()
	cateids := make([]string, len(photo_data))
	for i, data := range photo_data {
		cateids[i] = data["cateid"]
	}
	cateNames := getCateNames(cateids)
	for k, datas := range photo_data {
		cateid := datas["cateid"]
		photo_data[k]["catename"] = cateNames[cateid]
		addtime, err := strconv.ParseInt(datas["addtime"], 10, 64)
		checkErr(err, 25)
		addtime_str := time.Unix(addtime, 0).Format("2006-01-02 15:04:05")
		photo_data[k]["addtime"] = addtime_str
	}
	return photo_data
}
