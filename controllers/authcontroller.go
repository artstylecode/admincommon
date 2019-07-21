package controllers

import (
	"github.com/astaxie/beego"
)

//AuthController 权限控制器基类
type AuthController struct {
	BaseController
}

//验证用户登录
func checkLogin(this *AuthController) bool {

	return this.GetUID() != -1
}

//跳转至登录页面
func (ac *AuthController) redirectToLogin() {
	loginURL := beego.AppConfig.String("login_url")
	ac.Redirect(loginURL, 302)
}

//验证用户权限
func checkAuth(this *AuthController) bool {
	return true
}

//GetUID 获取当前登陆用户id，未登陆返回-1
func (ac *AuthController) GetUID() int {
	loginKey := beego.AppConfig.String("admin_session_key")
	sessionUID := ac.GetSessionInt(loginKey)
	return sessionUID
}

// Prepare 添加登录与角色验证
func (ac *AuthController) Prepare() {
	if checkLogin(ac) {
		if checkAuth(ac) {

		} else {

		}
	} else {
		ac.redirectToLogin()
		return
	}
}
