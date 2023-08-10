package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pkg6/go-paginate"
	"github.com/pkg6/go-paginate/gormp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Post struct {
	ID     uint `gormp:"primarykey" json:"id"`
	Number int  `json:"number"`
}
type PostIndex struct {
	ID     uint `gormp:"primarykey" json:"id"`
	Number int  `json:"number"`
	Index  int  `json:"index"`
}

var db, _ = gorm.Open(sqlite.Open("gormp.db?cache=shared"), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
})

func init() {
	_ = db.AutoMigrate(&Post{})
	for i := 1; i <= 100; i++ {
		p := Post{
			Number: i,
		}
		db.Save(&p)
	}
}
func simple() {
	q := db.Model(Post{})
	var dest []Post
	var adapt = gormp.Adapter(q)
	myPage := paginate.SimplePaginate(adapt, 10, 1)
	_ = myPage.Get(&dest)
	//获取最后页码
	page, err := myPage.GetLastPage()
	fmt.Println(fmt.Sprintf("获取最后页码:%v", page))
	fmt.Println(fmt.Sprintf("获取最后页码错误信息:%v", err))
	//获取总数
	total, err := myPage.GetTotal()
	fmt.Println(fmt.Sprintf("获取总数:%v", total))
	fmt.Println(fmt.Sprintf("获取总数错误信息:%v", err))
	fmt.Println(fmt.Sprintf("当前页码:%v", myPage.GetCurrentPage()))
	fmt.Println(fmt.Sprintf("每页显示多少条数:%v", myPage.GetListRows()))
	fmt.Println(fmt.Sprintf("是否还可以进行分页:%v", myPage.HasPages()))
	fmt.Println(dest)
}

func Total() {
	q := db.Model(Post{}).Where([]int64{20, 21, 22}).Order("id desc")
	var dest []Post
	var adapt = gormp.Adapter(q)
	t, _ := adapt.Length()
	myPage := paginate.TotalPaginate(adapt, 10, 1, t)
	_ = myPage.Get(&dest)
	var destIndex []PostIndex
	for i, post := range dest {
		destIndex = append(destIndex, PostIndex{ID: post.ID, Number: post.Number, Index: i})
	}
	// 先执行当get方法此时为nil data数据为dest
	// 不执行get方法 此时传入dest，data就是dest
	// 执行get方法 此时在传入新的结构体 data 就是data
	// 不执行get 此时传入nil 就返回nil
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
	total, err := myPage.GetTotal()
	fmt.Println(fmt.Sprintf("获取总数:%v", total))
	fmt.Println(fmt.Sprintf("获取总数错误信息:%v", err))
	fmt.Println(fmt.Sprintf("当前页码:%v", myPage.GetCurrentPage()))
	fmt.Println(fmt.Sprintf("每页显示多少条数:%v", myPage.GetListRows()))
	fmt.Println(fmt.Sprintf("是否还可以进行分页:%v", myPage.HasPages()))
	fmt.Println(dest)
}

func main() {
	Total()
	simple()
}
