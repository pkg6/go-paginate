package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg6/go-paginate"
	xorm2 "github.com/pkg6/go-paginate/xorm"
	"xorm.io/xorm"
)

type Post struct {
	ID     uint `xorm:"pk autoincr"`
	Number int
}

var engine, _ = xorm.NewEngine("sqlite3", "xorm.db")

func init() {
	err := engine.Sync(new(Post))
	fmt.Println(err)
	for i := 1; i <= 100; i++ {
		post := new(Post)
		post.Number = i
		affected, _ := engine.Insert(post)
		fmt.Println(affected)
	}
}

func Total() {
	session := engine.Table(Post{})
	var dest []Post
	var adapt = xorm2.Adapter(session)
	total, _ := adapt.Length()
	myPage := paginate.TotalPaginate(adapt, 10, 1, total)
	_ = myPage.Get(&dest)
	r := myPage.Render(nil)
	jsonBytes, _ := json.Marshal(r)
	xmlBytes, err := xml.Marshal(r)
	fmt.Println(string(jsonBytes))
	fmt.Println(string(xmlBytes))
	//获取最后页码
	page, err := myPage.GetLastPage()
	fmt.Println(fmt.Sprintf("获取最后页码:%v", page))
	fmt.Println(fmt.Sprintf("获取最后页码错误信息:%v", err))
	//获取总数
	total2, err := myPage.GetTotal()
	fmt.Println(fmt.Sprintf("获取总数:%v", total2))
	fmt.Println(fmt.Sprintf("获取总数错误信息:%v", err))
	fmt.Println(fmt.Sprintf("当前页码:%v", myPage.GetCurrentPage()))
	fmt.Println(fmt.Sprintf("每页显示多少条数:%v", myPage.GetListRows()))
	fmt.Println(fmt.Sprintf("是否还可以进行分页:%v", myPage.HasPages()))
	fmt.Println(dest)
}

func main() {
	//iniDB()
	engine.ShowSQL(true)
	//iniDB()
	Total()
}
