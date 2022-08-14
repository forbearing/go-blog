package v1

import (
	"net/http"
	"strconv"

	"github.com/forbearing/go-blog/model"
	apierrors "github.com/forbearing/go-blog/pkg/errors"
	"github.com/gin-gonic/gin"
)

// 查询分类是否存在

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == apierrors.Success {
		model.CreateCategory(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": apierrors.GetErrMsg(code),
		},
	)
}

// 查询单个分类下的文章
func GetCategoryInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCategoryInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": apierrors.GetErrMsg(code),
		},
	)
}

// 查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetCategory(pageSize, pageNum)
	code := apierrors.Success
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": apierrors.GetErrMsg(code),
		},
	)
}

// 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == apierrors.Success {
		model.EditCategory(id, &data)
	}
	if code == apierrors.ErrCategoryUsed {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": apierrors.GetErrMsg(code),
		},
	)
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCategory(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": apierrors.GetErrMsg(code),
		},
	)
}
