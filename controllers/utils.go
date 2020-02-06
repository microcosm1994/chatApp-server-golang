package controllers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"github.com/astaxie/beego/config"
)

/*UtilsController 工具类 */
type UtilsController struct {
	MainController
}

/*GenValidateCode 生成随机验证码 */
func (c *UtilsController) GenValidateCode(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var code strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&code, "%d", numeric[rand.Intn(r)])
	}
	return code.String()
}

/*SendMessage 给用户发送短信*/
func (c *UtilsController) SendMessage(phone string) bool {
	iniconf, err := config.NewConfig("ini", "conf/aliyun.conf")
	if err != nil {
		return false
	}
	code := c.GenValidateCode(4)
	// 短信天极流量控制
	c.SetSession("code", code)
	fmt.Println("code", code)
	var (
		gatewayURL      = "http://dysmsapi.aliyuncs.com/"
		accessKeyID     = iniconf.String("AccessKeyID")
		accessKeySecret = iniconf.String("AccessKeySecret")
		phoneNumbers    = phone
		signName        = "chatApp"
		templateCode    = iniconf.String("templateCode")
		templateParam   = "{\"code\":\"" + code + "\"}"
	)
	smsClient := aliyunsmsclient.New(gatewayURL)
	result, _ := smsClient.Execute(accessKeyID, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	if result.IsSuccessful() {
		// 保存短信验证码
		c.SetSession("code", code)
		return true
	}
	return false
}
