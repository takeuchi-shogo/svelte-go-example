package controller

import (
	"strconv"
	"svelte-go-sample/models/db"
	"svelte-go-sample/models/entity"

	"github.com/gin-gonic/gin"
)

//FetchAllTasks は全ての商品情報を取得する
func FetchAllTasks(c *gin.Context) {
	resultTasks := db.FindAllTasks()

	//URLのアクセスに対してJSONを返す
	c.JSON(200, resultTasks)
}

//AddTask は商品情報をDBに登録する
func AddTask(c *gin.Context) {
	taskName := c.PostForm("taskname")

	var task = entity.Task{
		Name: taskName,
	}

	db.InsertTask(&task)
}

//DeleteTask は商品情報をDBから削除する
func DeleteTask(c *gin.Context) {
	taskIDStr := c.PostForm("taskID")

	taskID, _ := strconv.Atoi(taskIDStr)

	db.DeleteTask(taskID)
}
