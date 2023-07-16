package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/json-iterator/go"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	// 设置创建时间
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil

}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	// 设置修改时间
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil

}

// GetTags  新增标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return tags
}

// GetTagTotal 获取总数
func GetTagTotal(maps interface{}) (count int) {
	// 通过where条件获取总数
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count
}

// ExistTagByName 判断标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// ExistTagById 判断标签是否存在
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// AddTag 新增标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// DeleteTag  修改标签
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

// EditTag 修改标签
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}
