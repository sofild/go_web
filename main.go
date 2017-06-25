package main

import (
	"fmt"
	"io"
	"net/http"
	//"strconv"
	"encoding/json"
	"text/template"
)

var table Table

func main() {
	table.Config.Db = "gotest"
	table.Config.Host = "192.168.80.131"
	table.Config.Port = "3306"
	table.Config.User = "root"
	table.Config.Pass = "root"
	table.Conn()

	http.Handle("/favicon.ico", http.FileServer(http.Dir("myweb/dist")))
	http.Handle("/static/", http.FileServer(http.Dir("myweb/dist")))
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("myweb/dist/index.html")
	if err != nil {
		fmt.Println("Template parse failed:", err.Error())
	}

	blog_data := getArticles("", "id desc", "0,200")
	blog, err2 := json.Marshal(blog_data)
	checkErr(err2, 39)

	photo_data := getPhotoList("", "id desc", "0,270")
	photo, err3 := json.Marshal(photo_data)
	checkErr(err3, 44)

	cate_data := getCateList()
	cate, err4 := json.Marshal(cate_data)
	checkErr(err4, 47)

	datas := map[string]string{"Blog": string(blog), "Photo": string(photo), "Cate": string(cate)}
	html.Execute(w, datas)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) > 0 {
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		user := getUser(username, password)
		var data map[string]string
		if len(user) > 0 {
			uid := user["id"]
			data["uid"] = uid
			fmt.Println(data)
		} else {
			data["err"] = "1"
			data["msg"] = "用户名或密码错误！"
			datas, err := json.Marshal(data)
			checkErr(err, 66)
			fmt.Println(datas)
			io.WriteString(w, string(datas))
		}
	} else {
		html, err := template.ParseFiles("myweb/dist/login.html")
		if err != nil {
			fmt.Println("Template parse failed:", err.Error())
		}
		datas := map[string]string{"datas": "Login"}
		html.Execute(w, datas)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "==================")
	/*
		res := AddCate(0, "数码")
		io.WriteString(w, strconv.FormatInt(res, 10))
		fmt.Println(res)
	*/
	/*
		data := GetCate("parent_id=0", "id desc", "")
		for k, v := range data {
			fmt.Println(strconv.Itoa(k), ":")
			for key, val := range v {
				fmt.Println(key, "=>", val)
			}
			fmt.Println("\n")
		}
	*/
	/*
		data := make(map[string]string)
		data["name"] = "服饰"
		res := UpCate(data, 1)
		fmt.Println(res)
	*/
	/*
		res := DelCate(1)
		fmt.Println(res)
	*/
	/*
		var table Table
		table.Config.Db = "gotest"
		table.Config.Host = "192.168.80.131"
		table.Config.Port = "3306"
		table.Config.User = "root"
		table.Config.Pass = "root"
		table.Conn()
	*/
	/*
		table.Name = "cate"
		table.Field = append(table.Field, "id", "parent_id", "name", "addtime")
		table.Limit = ""
		table.Order = "id desc"
		table.Condition = "parent_id=0"
	*/

	/*
		table.Name = "articles"
		//table.Field = append(table.Field, "id", "cateid", "title", "description", "pic", "content", "addtime")
		table.SetField("id", "cateid", "title", "description", "pic", "content", "addtime")
		table.Limit = ""
		table.Order = "id desc"
		table.Condition = "cateid=4"
		//find_data := table.Find()
		//fmt.Println(find_data)
		select_data := table.Select()
		fmt.Println(select_data)
	*/

	/*
	   type article struct {
	       cateid int64
	       title string
	       description string
	       pic string
	       content string
	       addtime int64
	   }
	*/

	/*
		table.Name = "articles"
		table.SetField("cateid", "title", "description", "pic", "content", "addtime")
		table.SetValue(6, "测试yyy一下标题", "测试yyyy一下描述", "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg", "测试一下内容", 1497704998)
		res := table.Add()
		fmt.Println(res)
	*/

	table.Name = "photos"
	table.SetField("cateid", "image", "title", "addtime")
	table.SetValue(5, "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=d0ba7780b48369418bebf96141eea834&imgtype=0&src=http%3A%2F%2Fimg0.pcauto.com.cn%2Fpcauto%2F1405%2F30%2F4493541_100003604293146_thumb.gif", "好车图", 1497704998)
	res := table.Add()
	fmt.Println(res)
}
