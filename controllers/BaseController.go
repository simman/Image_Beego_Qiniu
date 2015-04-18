package controllers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego/utils/pagination"
    "github.com/astaxie/beego/config"
)

var Logger *logs.BeeLogger
var Configer config.ConfigContainer

type BaseController struct {
	beego.Controller
}

func init() {

    Configer = beego.AppConfig
    Logger = logs.NewLogger(10000)      // 10000 表示缓存的大小
    Logger.SetLogger("console", "")     // 包括：console、file、conn、smtp
    Logger.EnableFuncCallDepth(true)    // 输出文件名和行号
}

func (this *BaseController) SendJson(arr []map[string]string) {
	this.Data["json"] = arr
	this.ServeJson()
}

// @Title SetPage
// @Description 设置分页
// @param per int 每页数量
// @param counts int64 总数量
func (this *BaseController) SetPage(per int, counts int64) *pagination.Paginator {
    p := pagination.NewPaginator(this.Ctx.Request, per, counts)
    this.Data["paginator"] = p
    return p
}
