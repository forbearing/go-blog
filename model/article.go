package model

import (
	apierrors "github.com/forbearing/go-blog/pkg/errors"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return apierrors.Failed // 500
	}
	return apierrors.Success
}

// GetCategoryArticle 查询分类下的所有文章
func GetCategoryArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64

	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, apierrors.ErrCategoryNotExist, 0
	}
	return cateArtList, apierrors.Failed, total
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var art Article
	err := db.Where("id = ?", id).Preload("Category").First(&art).Error
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, apierrors.ErrArticleNotExist
	}
	return art, apierrors.Success
}

// GetArticle 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	err = db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, apierrors.Failed, 0
	}
	return articleList, apierrors.Success, total
}

// SearchArticle 搜索文章标题
func SearchArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	//单独计数
	db.Model(&articleList).Where("title LIKE ?",
		title+"%",
	).Count(&total)

	if err != nil {
		return nil, apierrors.Failed, 0
	}
	return articleList, apierrors.Success, total
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return apierrors.Failed
	}
	return apierrors.Success
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var art Article
	err := db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return apierrors.Failed
	}
	return apierrors.Success
}
