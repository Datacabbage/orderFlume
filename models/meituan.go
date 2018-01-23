package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TMeituan struct {
	Id         int       `orm:"column(id);pk"`
	Orderid    string    `orm:"column(orderid);null"`
	UID        string    `orm:"column(uid);size(50);null"`
	Sid        string    `orm:"column(sid);size(50);null"`
	Total      string    `orm:"column(total);size(50);null"`
	Direct     string    `orm:"column(direct);size(50);null"`
	Type       string    `orm:"column(type);size(50);null"`
	Quantity   string    `orm:"column(quantity);size(50);null"`
	Dealid     string    `orm:"column(dealid);size(50);null"`
	Smstitle   string    `orm:"column(smstitle);null"`
	Paytime    string    `orm:"column(paytime);null"`
	Modtime    string    `orm:"column(modtime);null"`
	CreateTime time.Time `orm:"column(createtime);type(datetime)"`
}

func (t *TMeituan) TableName() string {
	return "t_meituan"
}

func init() {
	orm.RegisterModel(new(TMeituan))
}

// AddTPcactive insert a new TPcactive into database and returns
// last inserted Id on success.
func AddTMeituan(m *TMeituan) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
func FindUsedTMeituan() (count []TMeituan, err error) {
	o := orm.NewOrm()
	var ch []TMeituan
	num, err := o.QueryTable(new(TMeituan)).Filter("useState", 0).All(&ch)
	if err == nil && num > 0 {
		return ch, nil
	}
	return nil, err

}
func FindTMeituan() (count []TMeituan, err error) {
	o := orm.NewOrm()
	var ch []TMeituan
	num, err := o.QueryTable(new(TMeituan)).All(&ch)
	if err == nil && num > 0 {
		return ch, nil
	}
	return nil, err

}
func FindTMeituanWithCondition(state int, packageName string) (count int, err error) {
	o := orm.NewOrm()
	var ch []TMeituan
	num, err := o.QueryTable(new(TMeituan)).Filter("useState", state).Filter("packageName__contains", packageName).All(&ch)
	if err == nil && num > 0 {
		return int(num), nil
	}
	return 0, err

}

// GetTPcactiveById retrieves TPcactive by Id. Returns error if
// Id doesn't exist
func GetTMeituanById(id int) (v *TMeituan, err error) {
	o := orm.NewOrm()
	v = &TMeituan{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func FindTMeituanDesc() (date []TMeituan, err error) {
	var ch []TMeituan
	o := orm.NewOrm()
	num, err := o.QueryTable(new(TMeituan)).Limit(1).OrderBy("-id").All(&ch)
	if err == nil && num > 0 {
		return ch, err
	}
	return nil, err
}
func FindTMeituanCond(isuse int) (date []TMeituan, err error) {
	var ch []TMeituan
	o := orm.NewOrm()
	num, err := o.QueryTable(new(TMeituan)).Filter("isUse", isuse).All(&ch)
	if err == nil && num > 0 {
		return ch, err
	}
	return nil, err
}

// UpdateTPcactive updates TPcactive by Id and returns error if
// the record to be updated doesn't exist
func UpdateTMeituanById(m *TMeituan) (err error) {
	o := orm.NewOrm()
	v := TMeituan{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			beego.Debug("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTPcactive deletes TPcactive by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTMeituan(id int) (err error) {
	o := orm.NewOrm()
	v := TMeituan{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TMeituan{Id: id}); err == nil {
			beego.Debug("Number of records deleted in database:", num)
		}
	}
	return
}
