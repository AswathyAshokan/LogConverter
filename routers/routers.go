package routers

import (
	"github.com/astaxie/beego"
	"LogConverter/Controllers"
)

func init(){

	beego.Router("/logFile",&controllers.FileLogController{},"*:FileLog")

}

