A simple way to implement pagination in Golang.

[![Go Report Card](https://goreportcard.com/badge/github.com/pkg6/go-paginate)](https://goreportcard.com/report/github.com/pkg6/go-paginate)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/pkg6/go-paginate?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/pkg6/go-paginate/-/badge.svg)](https://sourcegraph.com/github.com/pkg6/go-paginate?badge)
[![Release](https://img.shields.io/github/release/pkg6/go-paginate.svg?style=flat-square)](https://github.com/pkg6/go-paginate/releases)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/pkg6/go-paginate/badges/download-count.svg)](https://goproxy.cn)

## Install

~~~
$ go get github.com/pkg6/go-paginate
~~~

## Usage

~~~
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
~~~

## Get & Render method

| Get                         | Render        | data                |
| --------------------------- | ------------- | ------------------- |
| Execute mandatory&data data | nil           | &data               |
| Execute mandatory&data data | Transfer data | Transfer data       |
| Do not execute              | nil           | null                |
| Do not execute              | Transfer data | Get method executed |

## Adapters

An adapter must implement the `Adapter` interface which has 2 methods:
~~~
Length() (int64, error)
Slice(offset, length int64, dest any) error
~~~

### GORM Adapter
~~~
var db, _ = gorm.Open(sqlite.Open("gorm.db?cache=shared"), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
})
q := db.Model(Post{}).Where([]int64{20, 21, 22}).Order("id desc")
var dest []Post
var adapt = gromp.Adapter(q)
t, _ := adapt.Length()
myPage := paginate.TotalPaginate(adapt, 10, 1, t)
~~~

### XORM Adapter
~~~
var engine, _ = xorm.NewEngine("sqlite3", "xorm.db")
session := engine.Table(Post{})
var dest []Post
var adapt = xormp.Adapter(session)
total, _ := adapt.Length()
myPage := paginate.TotalPaginate(adapt, 10, 1, total)
~~~

### Slice adapter
~~~
var source = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 22, 33, 44, 55, 66, 77, 88, 99, 199,
	111, 222, 333, 444, 555, 666, 777, 888, 999, 199,
	1111, 2222,
}
var adapt = slicep.Adapter(source)
myPage := paginate.SimplePaginate(adapt, 10, 5)
~~~

## Join us

If you approve of our open source project and are interested in contributing to the development of go-paginator, we sincerely welcome you to join us in developing and improving it. Whether it's [reporting errors](https://github.com/pkg6/go-paginate/issues) or developing a [Pull Request](https://github.com/pkg6/go-paginate/pulls), even modifying a typo can be a great help.

