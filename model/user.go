package model

import (
	"log"

	apierrors "github.com/forbearing/go-blog/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(500);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 一定要实现指针方法而不是结构体方法
func (u *User) BeforeSave(_ *gorm.DB) error {
	u.Password = ScyptPw(u.Password)
	return nil
}

// 查询用户是否存在
func CheckUser(username string) int {
	user := &User{}
	db.Select("id").Where("username = ?", username).First(user)
	if user.ID > 0 {
		return apierrors.ErrUsernameUsed // 1001
	}

	return apierrors.Success
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	//data.Password = ScyptPw(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return apierrors.Failed // 500
	}
	return apierrors.Success
}

// GetUser 查询单个用户
func GetUser(id int) (*User, int) {
	user := &User{}
	err := db.Limit(1).Where("ID = ?", id).Find(user).Error
	if err != nil {
		return user, apierrors.Failed
	}
	return user, apierrors.Success
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if len(username) > 0 {
		db.Select("id,username,role,created_at").Where("username LIKE ?", username+"%").
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where("username LIKE ?", username+"%").Count(&total)
		return users, total
	}
	err := db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	db.Model(&users).Count(&total)
	if err != nil {
		return users, 0
	}
	return users, total
}

// EditUser 编辑用户, 如果要修改密码需要额外的一个方法
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return apierrors.Failed
	}
	return apierrors.Success
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ? ", id).Delete(&user).Error
	logrus.Info(err)
	if err != nil {
		return apierrors.Failed
	}
	return apierrors.Success
}

// ScyptPw 密码加密
func ScyptPw(passwd string) string {
	const cost = 10
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPasswd)
}
