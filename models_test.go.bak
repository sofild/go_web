package main

import (
	"fmt"
	"testing"
)

func initConn() *Table{
    var table *Table
	table.Config.Db = "gotest"
	table.Config.Host = "192.168.37.170"
	table.Config.Port = "3306"
	table.Config.User = "root"
	table.Config.Pass = "root"
	table.Conn()
    return table
}

func TestConn(t *testing.T){
    var table *Table
    table.Config.Db = "gotest"
    table.Config.Host = "192.168.37.170"
    table.Config.Port = "3306"
    table.Config.User = "root"
    table.Config.Pass = "root"
    table.Conn()
    if table==nil{
        t.Fatalf("Test failed.")
    }
    fmt.Println(table)
}

func TestFind(t *testing.T){
    table := initConn()
    table.Name = "cate"
    table.Field = append(table.Field, "id", "parent_id", "name", "addtime")
    table.Order = "id desc"
    table.Condition = "parent_id=0"

    find_data := table.Find()
    fmt.Println(find_data)
    if find_data==nil {
        t.Fatalf("Test failed.")
    }
}

func TestSelect(t *testing.T){
    table := initConn()
    table.Name = "cate"
    table.Field = append(table.Field, "id", "parent_id", "name", "addtime")
    table.Order = "id desc"
    table.Condition = "parent_id=0"
    select_data := table.Select()
    fmt.Println(select_data)
    if select_data == nil {
        t.Fatalf("Test failed.")
    }
}

func TestAdd(t *testing.T){
    table := initConn()
    table.Name = "articles"
    table.Field = append(table.Field, "cateid", "title", "description", "pic", "content", "addtime")
     table.Value = append(table.Value,4, "单元测试哈哈哈哈哈", "单元测试测试一下描述", "http://s2.juancdn.com/bao/160921/d/9/57e26ab2151ad118338b45a9_800x800.jpg", "单元测试需要爱啊一下内容", 1497704998)
     res := table.Add()
     fmt.Println(res)
     if res < 1{
        t.Fatalf("Test failed.")
     }
}
