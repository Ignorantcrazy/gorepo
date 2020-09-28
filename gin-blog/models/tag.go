package models

import (
	"time"

	"gorm.io/gorm"
)

//Tag 标签表
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag *Tag) BeforeCreate(tx *gorm.DB) error {
	tag.CreatedOn = time.Now().Unix()

	return nil
}

func (tag *Tag) BeforeUpdate(tx *gorm.DB) error {
	tag.ModifiedOn = time.Now().Unix()

	return nil
}

func GetTags(pageNum int, PageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(PageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID < 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})
	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}
