package v1

import (
	"net/http"
	"strconv"

	"github.com/forbearing/go-blog/model"
	apierrors "github.com/forbearing/go-blog/pkg/errors"
	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUser(data.Username)
	if code == apierrors.Success {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": apierrors.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUserInfo(c *gin.Context) {}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)
	code := apierrors.Success
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": apierrors.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUser(data.Username)
	if code == apierrors.Success {
		model.EditUser(id, &data)
	}
	if code == apierrors.ErrUsernameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": apierrors.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": apierrors.GetErrMsg(code),
	})
}
