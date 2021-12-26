package main

import (
	"fmt"
	"net/http"
)

var userdata map[string]string
//抱歉学长我还没有学数据库 我先看的<图解http>这本书因为我想先了解一些基础知识,理解web框架原理

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userdata[r.Form.Get("name")]=r.Form.Get("passwd")
	fmt.Fprint(w,"感谢注册")
}
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if userdata[r.Form.Get("name")]==r.Form.Get("passwd"){
		fmt.Fprint(w,"欢迎你"+userdata[r.Form.Get("name")])
	} else{
		fmt.Fprint(w,"密码错误")
	}
}
func main() {
	//http://127.0.0.1:9000/login?name=112233&passwd=1122
	//http://127.0.0.1:9000/register?name=112233&passwd=1122
	userdata=make(map[string]string,100)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		return
	}
}