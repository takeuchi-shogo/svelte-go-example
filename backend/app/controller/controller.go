package controller

import (
	"fmt"
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

	var task entity.Task
	//json形式のものをデコードする
	c.BindJSON(&task)
	fmt.Print(task)

	db.InsertTask(&task)
}

//ChangeTask はタスク情報を変更する
func ChangeTask(c *gin.Context) {
	reqTaskID := c.Param("id")
	var reqtaskDone entity.Task

	taskID, _ := strconv.Atoi(reqTaskID)
	c.BindJSON(&reqtaskDone)
	taskDone := reqtaskDone.Done

	if taskDone == false {
		taskDone = true
	} else {
		taskDone = false
	}

	db.UpdateDoneTask(taskID, taskDone)
}

//DeleteTask は商品情報をDBから削除する
func DeleteTask(c *gin.Context) {
	taskIDStr := c.Param("id")
	fmt.Println(taskIDStr)

	taskID, _ := strconv.Atoi(taskIDStr)

	db.DeleteTask(taskID)
}
