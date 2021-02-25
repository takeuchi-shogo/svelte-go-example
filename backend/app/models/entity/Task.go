package entity

//Task はテーブルのモデル
type Task struct {
	ID   int    `gorm:"primary_key;not null" json:"id"`
	Name string `gorm:"type:varchar(200);not null" json:"name"`
	Done bool   `gorm:"not null" json:"done"`
}
