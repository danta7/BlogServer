package site_api

import (
	"BlogServer/models/enum"
	"BlogServer/service/log_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	fmt.Println("1")
	log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	log_service.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "danta", "1234")
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
	return
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowRequestHeader()
	log.ShowResponseHeader()
	log.ShowResponse()
	log.SetTitle("更新站点")
	log.SetItemInfo("请求时间", time.Now())
	log.SetImage("/xxx/xxx")
	log.SetLink("gin链接", "https://gin-gonic.com/zh-cn/")
	c.Header("xxx", "xxxxe")

	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.SetError("参数绑定失败", err)
	}

	log.SetItemInfo("结构体", cr)
	log.SetItemInfo("切片", []string{"a", "b"})
	log.SetItemInfo("字符串", "你好")
	log.SetItemInfo("数字", 123)

	id := log.Save()
	fmt.Println(1, id)
	id = log.Save()
	fmt.Println(2, id)
	// id := log.Save()
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息", "data": 1})
	return
}