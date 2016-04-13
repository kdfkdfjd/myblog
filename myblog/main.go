package main

import (
	_ "myblog/routers"
	"github.com/astaxie/beego"
)

func main() {
      //changed 20160413
	beego.Run()
}

