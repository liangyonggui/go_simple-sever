package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/index.html")
	now := time.Now()
	timeStr := now.Format("20060102 15:04:05")

	data := map[string]string{
		"name":    "chiral",
		"someStr": "这是一个开始",
		"str1": fmt.Sprintf("当前时间%s", timeStr),
		"str2": "天气转冷，注意保暖哦~",
	}

	t.Execute(w, data)

	// fmt.Fprintln(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb)

	//指定相对路径./static 为文件服务路径
	staticHandle := http.FileServer(http.Dir("./static/"))
	//将/js/路径下的请求匹配到 ./static/js/下
	http.Handle("/js/", staticHandle)
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static"))))

	fmt.Println("服务器即将开启，访问地址 http://localhost:8080")

	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
