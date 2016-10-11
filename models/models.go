package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Requirement struct {
	Id                  int        `orm:"auto;pk" json:"id, omitempty"`
	Name                string     `orm:"size(255)" json:"name"`
	Phone               string     `orm:"size(32)" json:"phone"`
	Email               string     `orm:"size(64)" json:"email"`
	Company             string     `orm:"size(64)" json:"company"`
	Requirement_content string     `orm:"type(text)" json:"requirementContent"`
	Requirement_name    string     `orm:"size(255)" json:"requirementName, omitempty"`
	Attribute           string     `orm:"size(32)" json:"attribute, omitempty"`
	Resourcemap         string     `orm:"size(32)" json:"resourcemap, omitempty"`
	Trade               string     `orm:"size(64)" json:"trade, omitempty"`
	Scope               string     `orm:"size(255)" json:"scope, omitempty"`
	Scene               string     `orm:"size(255)" json:"scene, omitempty"`
	Frequency           string     `orm:"size(64)" json:"frequency, omitempty"`
	Deliver             string     `orm:"size(64)" json:"deliver, omitempty"`
	Priority            string     `orm:"size(8)" json:"priority, omitempty"`
	Create_time         time.Time  `orm:"auto_now_add;type(datetime)"`
	History             []*History `orm:"reverse(many)" json:"history, omitempty"`
	Status              string     `orm:"size(64)" json:"status, omitempty"`
	Create_user         string     `orm:"size(32)" json:"createUser, omitempty"`
	Review_user         string     `orm:"size(32)" json:"review_user, omitempty"`
	Remark              string     `orm:"type(text)" json:"remark, omitempty"`
}

type History struct {
	Id          int          `orm:"auto;pk" json:"id, omitempty"`
	Remark      string       `orm:"type(text)" json:"remark, omitempty"`
	Create_time time.Time    `orm:"auto_now_add;type(datetime)"`
	Status      string       `orm:"size(64)" json:"status, omitempty"`
	Review_user string       `orm:"size(32)" json:"review_user, omitempty"`
	Available   string       `orm:"size(2)" json:"available, omitempty"`
	Requirement *Requirement `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("dh_", new(Requirement))
	orm.RegisterModelWithPrefix("dh_rm_", new(History))
}
