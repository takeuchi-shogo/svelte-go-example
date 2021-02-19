package main

import (
	"log"

	"svelte-go-sample/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("svelte-go-sample/frontend/public/index.html")

	router.GET("/tasklist", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	//全てのタスク情報をJSONで返す
	router.GET("/alltasks", controller.FetchAllTasks)

	//一つの方品情報をJSONで返す
	router.POST("/addtask", controller.AddTask)

	//タスク情報を削除する
	router.POST("/deletetask", controller.DeleteTask)

	if err := router.Run(); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
