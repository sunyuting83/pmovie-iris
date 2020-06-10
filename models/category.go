package models

import (
	orm "pornplay/database"
)

// CategoryList Category List
type CategoryList struct {
	ID     int64  `json:"id" gorm:"primary_key, column:id"`
	CID    int64  `json:"cid" gorm:"column:cid"`
	Title  string `json:"title" gorm:"column:title"`
	Cover  string `json:"cover" gorm:"column:cover"`
	More   string `json:"more" gorm:"column:more"`
	Region string `json:"region" gorm:"column:region"`
}

// TableName change table name
func (CategoryList) TableName() string {
	return "movie"
}

// GetCategory 列表
func (category *CategoryList) GetCategory(id, page int64) (categorys []CategoryList, err error) {
	p := makePage(page)
	if err = orm.Eloquent.
		Select("id, cid, title, cover, more, region").
		Order("id desc").
		Limit(12).Offset(p).
		Find(&categorys, "cid = ?", id).Error; err != nil {
		return
	}
	return
}

// makePage make page
func makePage(p int64) int64 {
	p = p - 1
	if p <= 0 {
		p = 0
	}
	page := p * 12
	return page
}
