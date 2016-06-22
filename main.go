package main

import (
	_ "blog/routers"

	_ "blog/initial"
	"fmt"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
)

func main() {
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")

	data := make(map[string]interface{})
	data["content"] = "page not found"
	fmt.Println(data)
	t.Execute(rw, data)
}
