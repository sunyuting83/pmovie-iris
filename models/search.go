package models

import (
	orm "porn_movie/database"
)

// SearchKey 列表
func (search *CategoryList) SearchKey(key string) (searchs []CategoryList, err error) {
	if err = orm.Eloquent.
		Where("title LIKE ?", "%"+key+"%").
		Order("id desc").
		Limit(10).
		Find(&searchs).Error; err != nil {
		return
	}
	return
}

// Search 列表
func (search *CategoryList) Search(key string, page int64) (searchs []CategoryList, err error) {
	p := makePage(page)
	if err = orm.Eloquent.
		Where("title LIKE ?", "%"+key+"%").
		Order("id desc").
		Limit(12).
		Offset(p).
		Find(&searchs).Error; err != nil {
		return
	}
	return
}
