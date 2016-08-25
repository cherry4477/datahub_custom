package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Requirement struct {
	Id                  int       `orm:"auto;pk" json:"id, omitempty"`
	Name                string    `orm:"size(64)" json:"name"`
	Phone               string    `orm:"size(20)" json:"phone"`
	Email               string    `orm:"size(64)" json:"email"`
	Company             string    `orm:"size(64)" json:"company"`
	Requirement_content string    `orm:"type(text)" json:"requirementContent"`
	Create_user         string    `orm:"size(32)" json:"createUser, omitempty"`
	Requirement_name    string    `orm:"size(64)" json:"requirementName, omitempty"`
	Attribute           string    `orm:"size(2)" json:"attribute, omitempty"`
	Resourcemap         string    `orm:"size(2)" json:"resourcemap, omitempty"`
	Trade               string    `orm:"size(64)" json:"trade, omitempty"`
	Scope               string    `orm:"size(64)" json:"scope, omitempty"`
	Frequency           string    `orm:"size(64)" json:"frequency, omitempty"`
	Deliver             string    `orm:"size(64)" json:"deliver, omitempty"`
	Priority            string    `orm:"size(8)" json:"priority, omitempty"`
	Create_time         time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time         time.Time `orm:"auto_now;type(datetime)"`
	Remark              string    `orm:"type(text)" json:"remark, omitempty"`
	Status              string    `orm:"size(2)" json:"status, omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix("dh_", new(Requirement))
}
