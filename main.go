package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func myWeb(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/index.html")
	// 1. 定义目标日期：2025年9月13日（时区使用本地时区，也可指定UTC）
	targetDate := time.Date(
		2025,           // 年
		time.September, // 月（使用time包的月份常量避免数字错误）
		13,             // 日
		0,              // 时
		0,              // 分
		0,              // 秒
		0,              // 纳秒
		time.Local,     // 时区（Local为本地时区，UTC为世界协调时间）
	)

	// 2. 获取当前时间（本地时区）
	now := time.Now()

	// 3. 截断当前时间到日期级别（忽略时分秒，只保留年月日）
	// 避免因当天时间不同导致的天数计算误差（比如当前是23点，目标是0点，直接相减会少算一天）
	nowTruncated := now.Truncate(24 * time.Hour)
	targetTruncated := targetDate.Truncate(24 * time.Hour)

	// 4. 计算时间差（纳秒），转换为天数
	duration := nowTruncated.Sub(targetTruncated)
	days := int(duration.Hours() / 24)

	data := map[string]string{
		"name":    "李雪纯",
		"someStr": "小傻蛋",
		"str1": fmt.Sprintf("我们已经在一起%d天啦~", days),
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
