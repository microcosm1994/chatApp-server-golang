package filters

import (
	"fmt"
	// _ "unsafe"
	"chatAppServer/models"
	"chatAppServer/utils"
	"strings"

	"github.com/astaxie/beego/context"
)

/*Sys 用户权限验证*/
type Sys struct {
	*utils.Redis
}

/*ErrorResult 错误返回*/
type ErrorResult struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

/*TokenFilter 验证token*/
func TokenFilter(ctx *context.Context) {
	fmt.Println(ctx.Request.RequestURI)
	url := ctx.Request.RequestURI
	if strings.Contains(url, "/login") || strings.Contains(url, "/checkCode") {
		fmt.Println("-------登陆相关---------")
	} else {
		fmt.Println("-------其他接口---------")
		// 取出session中的用户信息
		u := ctx.Input.Session("userInfo")
		if u == nil {
			data := ErrorResult{
				Status: 401,
				Msg:    "用户登陆已失效",
			}
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(data, true, true)
		} else {
			phone := u.(models.SysUser).Phone
			// 获取header中的token
			t := ctx.Input.Header("t")
			// 取出token值（token在redis中是key）
			val := Sys.GetRedis(Sys{}, t)
			if val != phone {
				fmt.Println("--------token验证错误---------")
				data := ErrorResult{
					Status: 401,
					Msg:    "用户登陆已失效",
				}
				ctx.Output.SetStatus(401)
				ctx.Output.JSON(data, true, true)
			}
		}
	}
}
