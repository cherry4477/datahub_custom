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

func GetByParamsFilterUser(params map[string]string) ([]*Requirement, error) {
	o := orm.NewOrm()
	o.Using("datahub")

	var requirements []*Requirement
	qs := o.QueryTable("dh_requirement")

	if loginUser, ok := params["loginUser"]; ok {
		qs = qs.Filter("create_user__exact", loginUser)
	}
	if name, _ := params["name"]; name != "" {
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
	_, err := qs.All(&requirements)
	if err != nil {
		return nil, err
	}

	beego.Debug(requirements)

	return requirements, err
}

func GetByParams(params map[string]string) ([]*Requirement, error) {
	o := orm.NewOrm()
	o.Using("datahub")

	var requirements []*Requirement
	qs := o.QueryTable("dh_requirement")

	if name, _ := params["name"]; name != "" {
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
	_, err := qs.All(&requirements)
	if err != nil {
		return nil, err
	}

	beego.Debug(requirements)

	return requirements, err
}

func GetById(id int) (*Requirement, error) {
	o := orm.NewOrm()
	o.Using("datahub")

	requirement := Requirement{Id: id}

	err := o.Read(&requirement)
	if err != nil {
		return nil, err
	}

	return &requirement, err
}

func GetAll() ([]*Requirement, error) {
	o := orm.NewOrm()
	o.Using("datahub")

	var requirements []*Requirement
	_, err := o.QueryTable("dh_requirement").All(&requirements)
	if err != nil {
		return nil, err
	}

	beego.Debug(requirements)

	return requirements, err

}

func Update(req *Requirement) (int64, error) {
	o := orm.NewOrm()
	o.Using("datahub")

	rows, err := o.Update(req)
	if err != nil {
		return 0, err
	}

	return rows, err
}

func Delete(id int) error {
	o := orm.NewOrm()
	o.Using("datahub")

	rs := o.Raw("UPDATE dh_requirement SET status = ? WHERE id = ?", "N", id)
	_, err := rs.Exec()
	if err != nil {
		return err
	}

	return err
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

	connstr := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlUrl + ")/" + mysqlDatabase + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", connstr, 30)

	orm.RunSyncdb("default", false, true)

	//orm.DefaultTimeLoc = time.UTC

	//&loc=Asia%2FShanghai
}
