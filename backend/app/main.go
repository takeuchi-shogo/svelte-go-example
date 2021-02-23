package main

import (
	"log"

	"svelte-go-sample/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("svelte-go-sample/frontend/public/index.html")

	//cors対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5000"}
	router.Use(cors.New(config))

	//全てのタスク情報をJSONで返す
	router.GET("/tasks", controller.FetchAllTasks)

	//一つのタスク情報をJSONで返す
	router.POST("/tasks", controller.AddTask)

	//タスク情報を削除する
	router.POST("/tasks/:id", controller.DeleteTask)

	if err := router.Run(); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
