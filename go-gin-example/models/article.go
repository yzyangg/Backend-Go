package models

// 类的定义，类的属性，类的一些基本方法，类的一些特殊方法
import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	//当一个字段被标记为 gorm:"index" 时，GORM会在数据库表中为该字段创建一个索引
	//以便在查询中快速定位相关数据。这可以提高查询效率，特别是对于频繁进行查询的字段。
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate 在创建之前
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	// 设置创建时间
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// BeforeUpdate 在更新之前
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	// 设置修改时间
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// ExistArticleByName 判断文章是否存在
func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// GetArticleTotal 获取文章总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return count
}

// GetArticles 获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	//Preload("Tag") 是一个方法调用，它用于预加载与模型关联的 Tag 模型数据。
	//这意味着在执行查询时，会将关联的 Tag 数据一起加载到结果中，以避免后续的延迟加载。
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return articles
}

// GetArticle 获取单个文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	//这意味着根据 article 中定义的关联关系，将加载与之关联的 Tag 数据到 article.Tag 字段中。
	db.Model(&article).Related(&article.Tag)
	return article
}

// EditArticle 编辑文章
func EditArticle(id int, data interface{}) bool {

	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

// AddArticle 新增文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

// AddArticleByStruct 新增文章
func AddArticleByStruct(data Article) bool {
	db.Create(&data)
	return true
}

// DeleteArticle 删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}
