package v1

import (
	"net/http"
	"strconv"

	"github.com/forbearing/go-blog/model"
	apierrors "github.com/forbearing/go-blog/pkg/errors"
	"github.com/gin-gonic/gin"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": apierrors.GetErrMsg(code),
	})
}

// GetCategoryArticle 查询分类下的所有文章
func GetCategoryArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.GetCategoryArticle(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": apierrors.GetErrMsg(code),
	})
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": apierrors.GetErrMsg(code),
	})
}

// GetArticle 查询文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	if len(title) == 0 {
		data, code, total := model.GetArticle(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": apierrors.GetErrMsg(code),
		})
		return
	}

	data, code, total := model.SearchArticle(title, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": apierrors.GetErrMsg(code),
	})
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": apierrors.GetErrMsg(code),
	})
}
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": apierrors.GetErrMsg(code),
	})
}
