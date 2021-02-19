package db

import (
	"fmt"
	"time"

	"svelte-go-sample/models/entity"

	"github.com/jinzhu/gorm"
)

func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PROTOCOL := "tcp([127.0.0.1]:3306)"
	DBNAME := "TaskList"
	CONNECT := USER + ":" + "@" + PROTOCOL + "/" + DBNAME

	count := 0

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Println(".")
			time.Sleep(time.Second)
			count++
			if count > 15 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	db.LogMode(true)

	db.AutoMigrate(&entity.Task{})

	return db
}

//FindAllTasks テーブルにレコードを追加
func FindAllTasks() []entity.Task {
	tasks := []entity.Task{}

	db := open()

	db.Order("ID asc").Find(&tasks)

	defer db.Close()

	return tasks
}

//FindTask テーブルを全件取得
func FindTask(taskID int) []entity.Task {
	task := []entity.Task{}

	db := open()

	db.First(&task, taskID)
	defer db.Close()

	return task
}

//InsertTask 商品テーブルにレコードを追加する
func InsertTask(registryTask *entity.Task) {
	db := open()

	db.Create(&registryTask)
	defer db.Close()
}

//DeleteTask 商品テーブルのレコードを削除
func DeleteTask(taskID int) {
	task := []entity.Task{}

	db := open()

	db.Delete(&task, taskID)

	defer db.Close()
}
