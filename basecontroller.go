package admincommon

import (
	"github.com/astaxie/beego"
)

//BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

//GetSessionObject 获取任意类型session值
func (bc *BaseController) GetSessionObject(name string) interface{} {
	sessionVal := bc.GetSession(name)
	return sessionVal
}

//GetSessionString 获取字符串类型session值
func (bc *BaseController) GetSessionString(name string) string {
	sessionVal := bc.GetSessionObject(name)
	if sessionVal != nil {
		return sessionVal.(string)
	}

	return ""
}

//GetSessionBool 获取布尔型session值
func (bc *BaseController) GetSessionBool(name string) bool {
	sessionVal := bc.GetSessionObject(name)
	if sessionVal != nil {
		return sessionVal.(bool)
	}

	return false
}

//GetSessionInt 获取整型session值
func (bc *BaseController) GetSessionInt(name string) int {
	sessionVal := bc.GetSessionObject(name)
	if sessionVal != nil {
		return sessionVal.(int)
	}

	return -1
}

//GetSessionFloat32 获取整型session值
func (bc *BaseController) GetSessionFloat32(name string) float32 {
	sessionVal := bc.GetSessionObject(name)
	if sessionVal != nil {
		return sessionVal.(float32)
	}

	return -1.0
}

//GetSessionFloat64 获取整型session值
func (bc *BaseController) GetSessionFloat64(name string) float64 {
	sessionVal := bc.GetSessionObject(name)
	if sessionVal != nil {
		return sessionVal.(float64)
	}

	return -1.0
}
