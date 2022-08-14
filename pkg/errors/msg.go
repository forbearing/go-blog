package errors

const (
	Success = 200
	Failed  = 500

	// code = 1000 用户模块的错误
	ErrUsernameUsed  = 1001
	ErrPasswordWrong = 1002
	ErrUserNotExist  = 1003
	ErrTokenNotExist = 1004 // token 不存在
	ErrTokenTimeout  = 1005 // token 过期
	ErrTokenInvalid  = 1006 // token 无效,
	ErrTokenFormat   = 1007 // token 格式不正确

	// code = 2000 文章模块的错误
	ErrArticleNotExist = 2001

	// code = 3000 分类模块的错误
	ErrCategoryUsed     = 3001
	ErrCategoryNotExist = 3002
)

var codeMsg = map[int]string{
	Success: "SUCCESS",
	Failed:  "FAILED",

	ErrUsernameUsed:  "username already exist",
	ErrPasswordWrong: "username or password wrong",
	ErrUserNotExist:  "usernot exist",
	ErrTokenNotExist: "token not exist",
	ErrTokenTimeout:  "token timeout",
	ErrTokenInvalid:  "token invalid",
	ErrTokenFormat:   "token format error",

	ErrArticleNotExist: "article not exist",

	ErrCategoryUsed:     "category already used",
	ErrCategoryNotExist: "category not exist",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
