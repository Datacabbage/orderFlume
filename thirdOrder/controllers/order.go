package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"order_flume/lib/beelog"
	"order_flume/models"
	"order_flume/thirdOrder/conf"
	"order_flume/thirdOrder/service"
	"order_flume/thirdOrder/sign"
	"order_flume/utils"

	"github.com/astaxie/beego"
)

type ThirdController struct {
	beego.Controller
}

func (c *ThirdController) Meituan() {
	var t = make(map[string]interface{}, 0)
	var value = make(map[string]interface{}, 0)
	info := make(map[string]string, 0)
	info["app_id"] = c.GetString("app_id")
	info["timestamp"] = c.GetString("timestamp")
	signMsg := c.GetString("sign")
	signs := sign.GetMeituanSign(info, conf.Conf.MeituanSecretSign, "http://yema.journeyui.com/order/status/sync")
	//	statusType := c.GetString("status_type")
	//	order_source := c.GetString("order_source")
	//	third_user_id := c.GetString("third_user_id")
	//	status := c.GetString("status")
	//	status_update_time := c.GetString("status_update_time")
	wm_order_view_id := c.GetString("wm_order_view_id")
	wm_user_phone := c.GetString("wm_user_phone")
	body := GetMeituanWaimai(wm_user_phone, wm_order_view_id)
	json.Unmarshal(body, &value)
	beelog.Debug(value)
	if !strings.EqualFold(signs, signMsg) {
		t["code"] = 201
		t["msg"] = "sign not same"
		t["data"] = nil
	} else {
		t["code"] = 200
		t["msg"] = ""
		t["data"] = nil
	}
	ret := service.SendToKafka(value, "美团外卖")
	beelog.Debug(ret)
	c.Data["json"] = t
	c.ServeJSON()
}

func (c *ThirdController) Post() {
	//request := string(o.Ctx.Input.RequestBody)
	info := make(map[string]string, 0)
	info["orderid"] = c.GetString("orderid")
	info["uid"] = c.GetString("uid")
	info["sid"] = c.GetString("sid")
	info["total"] = c.GetString("total")
	info["direct"] = c.GetString("direct")
	info["quantity"] = c.GetString("quantity")
	info["dealid"] = c.GetString("dealid")
	info["smstitle"] = c.GetString("smstitle")
	info["paytime"] = c.GetString("paytime")
	info["modtime"] = c.GetString("modtime")
	info["type"] = c.GetString("type")
	beelog.Debug(info)
	ret := service.SendToKafka(info, "美团")
	c.Data["json"] = ret
	c.ServeJSON()
}

func (c *ThirdController) Maoyan() {
	var t = make(map[string]interface{}, 0)
	var value map[string]interface{}
	info := make(map[string]string, 0)
	info["order"] = c.GetString("order")
	beelog.Debug(info["order"])
	info["timestamp"] = c.GetString("timestamp")
	signMsg := c.GetString("signMsg")
	signs := sign.GetSign(info, conf.Conf.MaoyanKey)

	if !strings.EqualFold(signs, signMsg) {
		t["code"] = "FAILURE"
		t["msg"] = "sign not same"
	} else {
		t["code"] = "SUCCESS"
		t["msg"] = ""
	}
	json.Unmarshal([]byte(info["order"]), &value)
	beelog.Debug(value)
	service.SendToKafka(value, "猫眼")
	c.Data["json"] = t
	c.ServeJSON()
}

/*allianceid=1&sid=50&orderid=1210669519&ouid=11111&orderstatus=cancelled&orderstatusname
=已取消
&usestatus=0&orderamount=%c2%a5100.00&orderdate=2015-12-15&ordertype=FlightDomestic&ordername=S
earch+Filght+Order+Go2+Soa2.0+OrderDetail&startdatetime=2015-12-15&pushdate=2015-12-15+06%3a38%3a
14&guid=8f584c09eaf94a5e89ed79136a3c11ec*/
func (c *ThirdController) Xiechen() {
	var (
		info    interface{}
		xiechen models.TXiechen
		token   string
		err     error
	)
	beelog.Debug(c.Ctx.Request.RequestURI)
	timeLayout := "2006-01-02 15:04:05"
	dateLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	xiechen.Allianceid = c.GetString("allianceid")
	xiechen.Orderid = c.GetString("orderid")
	xiechen.Ouid = c.GetString("ouid")
	userstatus := c.GetString("usestatus")
	xiechen.Usestatus, _ = strconv.Atoi(userstatus)
	xiechen.Orderstatus = c.GetString("orderstatus")
	xiechen.Orderstatusname = c.GetString("orderstatusname")
	xiechen.Orderamount = c.GetString("orderamount")
	orderdate := c.GetString("orderdate")
	xiechen.Orderdate, _ = time.ParseInLocation(dateLayout, orderdate, loc)
	xiechen.Sid = c.GetString("sid")
	xiechen.Ordertype = c.GetString("ordertype")
	xiechen.Ordername = c.GetString("ordername")
	Startdatetime := c.GetString("startdatetime")
	Pushdate := c.GetString("pushdate")
	xiechen.Startdatetime, _ = time.ParseInLocation(dateLayout, Startdatetime, loc)
	xiechen.Pushdate, _ = time.ParseInLocation(timeLayout, Pushdate, loc)
	xiechen.Guid = c.GetString("guid")
	xiechen.CreateTime = time.Now().UTC()
	test, _ := json.Marshal(xiechen)
	beelog.Debug(string(test))
	if models.CheckOrderid(xiechen.Orderid) {
		c.Ctx.WriteString("success")
		return
	}
	if _, err := models.AddTXiechen(&xiechen); err != nil {
		beelog.Error(err)
	}
	if token, err = utils.Cache.Get("access_token").Result(); err != nil {
		beelog.Error(err)
		token = GetToken()
	}
	info = GetOrderDetail(token, xiechen.Allianceid, xiechen.Sid, xiechen.Orderid, "123456")
	ret := service.SendToKafka(info, "携程酒店")
	if v, ok := ret.(map[string]string); ok {
		if v["errcode"] == "0" {
			c.Ctx.WriteString("success")
			return
		}
	}
	c.Data["json"] = ret
	c.ServeJSON()
}

func GetMeituanWaimai(userPhone, orderId string) []byte {
	meituanurl := conf.Conf.MeituanUrl + "?user_phone" + userPhone + "&order_id" + orderId
	err, body := utils.HttpGet(meituanurl)
	if err != nil {
		beelog.Error(err)
		return nil
	}
	return body
}

func GetToken() string {

	var authtoken string
	tokenurl := conf.Conf.Tokenurl + "?AID=" + conf.Conf.AID + "&SID=" + conf.Conf.Sid + "&KEY=" + conf.Conf.Key

	authtoken = PostToXiechen(tokenurl)
	return authtoken
}
func PostToXiechen(url string) string {
	var authtoken string
	var value map[string]interface{}
	var keepTime time.Duration
	if err, body := utils.HttpPostNo(url); err != nil {
		beelog.Error(err)
		return authtoken
	} else {
		json.Unmarshal(body, &value)
	}
	beelog.Debug(url)
	beelog.Debug(value)
	if token, ok := value["Access_Token"].(string); ok {
		if tm, ok := value["Expires_In"].(int); ok {
			keepTime = time.Duration(tm) * time.Second
		}
		if err := utils.Cache.Set("access_token", token, keepTime).Err(); err != nil {
			beelog.Error(err)
		}
		if refreshToken, ok := value["Refresh_Token"]; ok {
			if err := utils.Cache.Set("refresh_token", refreshToken, 0).Err(); err != nil {
				beelog.Error(err)
			}
		}
		authtoken = token
	}
	return authtoken
}

func RefreshToken() string {
	var authtoken string
	if refreshtoken, err := utils.Cache.Get("refresh_token").Result(); err == nil {
		refreshurl := conf.Conf.Refreshurl + "?AID=" + conf.Conf.AID + "&SID=" + conf.Conf.Sid + "&refresh_token=" + refreshtoken
		authtoken = PostToXiechen(refreshurl)
	} else {
		beelog.Error(err)
		authtoken = GetToken()
	}
	return authtoken
}

func GetOrderDetail(token, aid, sid, orderID, ouid string) interface{} {
	url := conf.Conf.Serviceurl + "?AID=" + aid + "&SID=" + sid + "&ICODE=" + conf.Conf.ICODE + "&UUID=" + utils.CreateId() + "&Token=" + token + "&mode=1&format=json"
	var value = make(map[string]interface{}, 0)
	value["OrderID"] = orderID
	value["AllianceID"], _ = strconv.Atoi(aid)
	value["SID"], _ = strconv.Atoi(sid)
	param, err := json.Marshal(value)
	if err != nil {
		beelog.Error(err)
	}
	beelog.Debug(string(param))
	var header = make(http.Header)
	var result = make(map[string]interface{}, 0)
	header.Set("Content-Type", "application/json")
	body, err := utils.HttpPost(url, string(param), header)
	beego.Debug(string(body))
	json.Unmarshal(body, &result)
	if code, ok := result["ErrCode"].(float64); ok {
		beelog.Debug(code)
		if code == 232 {
			token = GetToken()
			return GetOrderDetail(token, aid, sid, orderID, ouid)
		}
	}
	json.Unmarshal(body, &result)
	result["OUID"] = ouid
	beelog.Debug(result)
	return result
}
