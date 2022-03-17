package main

import (
	"fmt"
	"gocamp/ch5/restrict"
	"net/http"
)

// 原理：创建一个队列，初始化这个队列大小，每次请求计算当前时间和最开始的时间差是否在阀值内
func main() {
	sw := restrict.New(3)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if sw.Add() {
			fmt.Fprintf(w,"请求成功")
		} else {
			fmt.Fprintf(w,"请求失败")
		}
	})
	http.ListenAndServe(":8000", nil)
}
