package models

import (
	"github.com/asiainfoLDP/datahub_custom/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlUser     string
	mysqlPassword string
	mysqlDatabase string
	mysqlUrl      string
)

func AddOne(requirement Requirement) (int64, error) {
	o := orm.NewOrm()
	o.Using("datahub")
	requirement.Status = "A"

	id, err := o.Insert(&requirement)
	if err != nil {
		return 0, err
	}
	beego.Debug(id)

	return id, err
}

func Get(params map[string]string) []*Requirement {
	o := orm.NewOrm()
	o.Using("datahub")

	var requirements []*Requirement
	qs := o.QueryTable("requirement")

	if name, _ := params["name"]; name != "" {
		beego.Debug(name)
		qs = qs.Filter("name__iexact", name)
	}
	if phone, _ := params["phone"]; phone != "" {
		qs = qs.Filter("phone__contains", phone)
	}
	if email, _ := params["email"]; email != "" {
		qs = qs.Filter("email__contains", email)
	}
	if company, _ := params["company"]; company != "" {
		qs = qs.Filter("company__iexact", company)
	}
	if content, _ := params["content"]; content != "" {
		qs = qs.Filter("requirement_content__icontains", content)
	}
	qs.All(&requirements)

	beego.Debug(requirements)

	return requirements
}

func GetById(id int) Requirement {
	o := orm.NewOrm()
	o.Using("datahub")

	requirement := Requirement{Id: id}

	o.Read(&requirement)

	return requirement
}

func GetAll() []*Requirement {
	o := orm.NewOrm()
	o.Using("datahub")

	var requirements []*Requirement
	_, err := o.QueryTable("requirement").All(&requirements)
	checkErr(err)

	beego.Debug(requirements)

	return requirements

}

func Update(req Requirement) {
	o := orm.NewOrm()
	o.Using("datahub")

	_, err := o.Update(&req)
	checkErr(err)
}

func Delete(ObjectId string) {
}

func checkErr(err error) {
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func init() {
	mysqlUser = conf.GetMysqlUser()
	mysqlPassword = conf.GetMysqlPassword()
	mysqlDatabase = conf.GetMysqlDatabase()
	mysqlUrl = conf.GetMysqlUrl()

	beego.Debug(mysqlUser, mysqlPassword, mysqlDatabase)

	connstr := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlUrl + ")/" + mysqlDatabase + "?charset=utf8&loc=Asia%2FShanghai"

	orm.RegisterDataBase("default", "mysql", connstr, 30)

	orm.RunSyncdb("default", false, true)

	//
}
