package models

type Category struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `gorm:"type:varchar(100)" json:"name"`
}

// func (Category) TableName() string {
//     return "kategori"
// }