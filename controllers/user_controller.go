package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/pppercyWang/iris-gorm-demo/models"
	"github.com/pppercyWang/iris-gorm-demo/service"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: service.NewUserServices()}
}

// @Tags User
// @Summary Login 用户登录
// @Description login by username and password
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Result
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /user/login [post]
func (g *UserController) PostLogin(userName, passwd string) models.Result {
	var m map[string]string
	//err := g.Ctx.ReadJSON(&m)
	//if err != nil {
	//	log.Println("ReadJSON Error:", err)
	//}
	m = make(map[string]string, 2)
	m[userName] = userName
	m[passwd] = passwd
	result := g.Service.Login(m)
	return result
}

// @Tags Save
// @Summary Save 用户注册
// @Description save user
// @Param user body models.User true "用户信息" zhangsan
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Result
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /user/save [post]
func (g *UserController) PostSave(user models.User) (result models.Result) {
	//var user models.User
	//if err := g.Ctx.ReadJSON(&user); err != nil {
	//	log.Println(err)
	//	result.Msg = "数据错误"
	//	return
	//}
	//user.Name=username
	//user.Password=password

	return g.Service.Save(user)
}
