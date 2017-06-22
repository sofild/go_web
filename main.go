package main

import (
	"fmt"
	"io"
	"net/http"
	//"strconv"
	"text/template"
)

func main() {
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/fonts/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/photo", photo)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("html/index.html", "html/header.html", "html/footer.html")
	if err != nil {
		fmt.Println("Template parse failed:", err.Error())
	}
	datas := map[string]string{"datas": "Hello World!"}
	html.Execute(w, datas)
}

func login(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("html/login.html", "html/header.html", "html/footer.html")
	if err != nil {
		fmt.Println("Template parse failed:", err.Error())
	}
	datas := map[string]string{"datas": "Login"}
	html.Execute(w, datas)
}

func blog(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("html/blog.html", "html/header.html", "html/footer.html")
	if err != nil {
		fmt.Println("Template parse failed:", err.Error())
	}
	datas := map[string]string{"datas": "blog"}
	html.Execute(w, datas)
}

func photo(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("html/photo.html", "html/header.html", "html/footer.html")
	if err != nil {
		fmt.Println("Template parse failed:", err.Error())
	}
	datas := map[string]string{"datas": "photo"}
	html.Execute(w, datas)
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

	var table Table
	table.Config.Db = "gotest"
	table.Config.Host = "192.168.37.170"
	table.Config.Port = "3306"
	table.Config.User = "root"
	table.Config.Pass = "root"

	table.Conn()
	table.Name = "cate"
	table.Field = append(table.Field, "id", "parent_id", "name", "addtime")
	table.Limit = ""
	table.Order = "id desc"
	table.Condition = "parent_id=0"

	find_data := table.Find()
    fmt.Println(find_data)
    select_data := table.Select()
    fmt.Println(select_data)
}
