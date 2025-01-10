package user_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/models"
	"BlogServer/utlis/jwts"
	"BlogServer/utlis/pwd"
	"github.com/gin-gonic/gin"
)

type PwdLoginRequest struct {
	Val      string `json:"val" binding:"required"` // 这里有可能是用户名，也可能是邮箱
	Password string `json:"password" binding:"required"`
}

func (UserApi) PwdLoginApi(c *gin.Context) {
	var cr PwdLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if !global.Config.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("站点未启动密码登录", c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "(username = ? or email = ?) and password <> ''",
		cr.Val, cr.Val).Error

	if err != nil {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	// 颁发token
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})

	res.OkWithData(token, c)

}