package models

import (
	"time"

	//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TXiechen struct {
	Id              int       `orm:"column(id);pk"`
	Orderid         string    `orm:"column(orderid);null"`
	Allianceid      string    `orm:"column(allianceid);null"`
	Sid             string    `orm:"column(sid);null"`
	Ouid            string    `orm:"column(ouid);size(50);null"`
	Usestatus       int       `orm:"column(usestatus);null"`
	Orderstatus     string    `orm:"column(orderstatus);null"`
	Orderstatusname string    `orm:"column(orderstatusname);null"`
	Orderamount     string    `orm:"column(orderamount);null"`
	Orderdate       time.Time `orm:"column(orderdate);type(datetime)"`
	Ordertype       string    `orm:"column(ordertype);null"`
	Ordername       string    `orm:"column(ordername);null"`
	Startdatetime   time.Time `orm:"column(startdatetime);type(datetime)"`
	Pushdate        time.Time `orm:"column(pushdate);type(datetime)"`
	Guid            string    `orm:"column(guid);null"`
	CreateTime      time.Time `orm:"column(createtime);type(datetime)"`
}

func (t *TXiechen) TableName() string {
	return "t_xiechen"
}

func init() {
	orm.RegisterModel(new(TXiechen))
}

// AddTPcactive insert a new TPcactive into database and returns
// last inserted Id on success.
func AddTXiechen(m *TXiechen) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
func CheckOrderid(orderID string) bool {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(TXiechen)).Filter("orderid", orderID).Count()
	if err == nil && num > 0 {
		return true
	}
	return false

}
