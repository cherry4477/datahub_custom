package models

import "github.com/astaxie/beego/orm"

type Requirement struct {
	ReqId               int    `orm:"auto";"pk" json:"omitempty"`
	Name                string `orm:"size(64)" json:"name"`
	Phone               string `orm:"size(20) json:"phone"`
	Email               string `orm:"size(64) json:"email"`
	Company             string `orm:"size(64) json:"company"`
	Requirement_content string `orm:"type(text)" json:"requirementContent"`

	Remark              string `orm:"type(text)" json:"remark, omitempty"`
	Status              string `orm:"size(2)" json:"status, omitempty"`
}

func init() {
	orm.RegisterModel(new(Requirement))
}
