package routers

import (
    "github.com/simman/Image_Beego_Qiniu/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/my", &controllers.HomeController{}, "*:My")
    beego.Router("/hot", &controllers.HomeController{}, "*:Hot")
    beego.Router("/Upload", &controllers.HomeController{}, "post:Upload")
}
