// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"order_flume/thirdOrder/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/meituan/data", &controllers.ThirdController{})
	beego.Router("/v1/xiechen/data", &controllers.ThirdController{}, "get:Xiechen")
	beego.Router("/v1/maoyan/data", &controllers.ThirdController{}, "get:Maoyan")
	beego.Router("/order/status/sync", &controllers.ThirdController{}, "get:meituan")
}
