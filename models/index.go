package models

import (
	orm "pornplay/database"
)

// Category struct
type Category struct {
	ID       int64  `json:"id" gorm:"primary_key, column:id"`
	Category string `json:"category" gorm:"column:category"`
	Sort     int64  `json:"sort" gorm:"column:sort"`
	Cover    string `json:"cover" gorm:"column:cover"`
}

// GetIndexs 列表
func (category *Category) GetIndexs() (categorys []Category, err error) {
	if err = orm.Eloquent.Find(&categorys).Error; err != nil {
		return
	}
	return
}
