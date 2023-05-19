package main

import (
	"fmt"
	"github.com/pkg6/go-paginate"
	"github.com/pkg6/go-paginate/adapter"
)

var source = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 22, 33, 44, 55, 66, 77, 88, 99, 199,
	111, 222, 333, 444, 555, 666, 777, 888, 999, 199,
	1111, 2222,
}
var adapt = adapter.SliceAdapter(source)

func Simple() {
	myPage := paginate.SimplePaginate(adapt, 10, 5)
	var dest []int
	myPage.Get(&dest)
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

func total() {
	myPage := paginate.TotalPaginate(adapt, 10, 1, int64(len(source)))
	var dest []int
	myPage.Get(&dest)
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
	total()
	Simple()
}
