package filters

import(
	"fmt"
	"github.com/astaxie/beego/context"
)

type Sys struct {}

/*TokenFilter 验证token*/
func TokenFilter(ctx *context.Context) {
	fmt.Println(ctx.Input.Header("t"))
	fmt.Println(ctx.Request.RequestURI)
}
