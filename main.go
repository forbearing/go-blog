package main

import (
	"github.com/forbearing/go-blog/model"
	"github.com/forbearing/go-blog/routers"
)

func main() {
	model.InitDB()
	routers.InitRouter()
}
