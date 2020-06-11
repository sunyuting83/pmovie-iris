package models

import (
	orm "pornplay/database"
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
