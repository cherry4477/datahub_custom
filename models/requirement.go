package models

import (
	"github.com/asiainfoLDP/datahub_custom/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	mysqlUser     string
	mysqlPassword string
	mysqlDatabase string
	mysqlUrl      string
)

func AddOne(requirement *Requirement) (int64, error) {
	beego.Info("begin create a requirement model.")

	o := orm.NewOrm()
	o.Using("datahub")

	requirementId, err := o.Insert(requirement)
	if err != nil {
		return 0, err
	}
	beego.Debug("requirementId:", requirementId)

	history := new(History)
	history.Remark = requirement.Remark
	history.Status = requirement.Status
	history.Available = "Y"
	history.Requirement = requirement
	beego.Debug(history)

	_, err = o.Insert(history)
	if err != nil {
		return requirementId, err
	}

	beego.Info("end create a requirement model.")
	return requirementId, err
}

type GetByParamsFilterUserResult struct {
	Id                  int
	Name                string
	Phone               string
	Email               string
	Company             string
	Requirement_content string
	Create_time         time.Time
	Remark              string
	Status              string
}

func GetByParamsFilterUser(params map[string]string, offset int64, size int) ([]GetByParamsFilterUserResult, error) {
	beego.Info("begin get requirement for datahub by param model.")

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
	count, err := qs.All(&requirements)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return []GetByParamsFilterUserResult{}, nil
	}
	beego.Debug("count:", count)

	validateOffsetAndLimit(count, &offset, &size)

	for _, requirement := range requirements {
		history := new(History)
		_, err := o.QueryTable("dh_rm_history").Filter("requirement_id", requirement.Id).Filter("available__exact", "Y").All(&history)
		if err != nil {
			return nil, err
		}
		requirement.History = append(requirement.History, history)
	}
	beego.Debug(requirements)

	results := []GetByParamsFilterUserResult{}
	result := GetByParamsFilterUserResult{}

	for key, value := range requirements {
		result.Id = value.Id
		result.Name = value.Name
		result.Phone = value.Phone
		result.Email = value.Email
		result.Company = value.Email
		result.Requirement_content = value.Requirement_content
		result.Create_time = value.Create_time
		result.Remark = value.History[key].Remark
		result.Status = value.History[key].Status

		results = append(results, result)
	}

	beego.Info("end get requirement for datahub by param model.")
	return results, err
}

func GetByParams(params map[string]string) ([]*Requirement, error) {
	beego.Info("begin get requirement by param model.")

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

	beego.Debug("requirements:", requirements)

	for _, requirement := range requirements {
		history := new(History)
		_, err := o.QueryTable("dh_rm_history").Filter("requirement_id", requirement.Id).Filter("available__exact", "Y").All(history)
		if err != nil {
			return nil, err
		}
		requirement.History = append(requirement.History, history)
	}

	beego.Debug(requirements)

	beego.Info("end get requirement by param model.")
	return requirements, err
}
func GetById(reqId int) (*Requirement, error) {
	beego.Info("begin get a requirement without history model.")

	o := orm.NewOrm()
	o.Using("datahub")

	requirement := Requirement{Id: reqId}

	err := o.Read(&requirement)
	if err != nil {
		return nil, err
	}

	beego.Info("end get a requirement without history model.")
	return &requirement, err
}

func GetNewHistory(reqId int) (*History, error) {
	beego.Info("begin get a requirement with history model.")

	o := orm.NewOrm()
	o.Using("datahub")

	history := new(History)

	qs := o.QueryTable("dh_rm_history")
	qs = qs.Filter("requirement_id", reqId).Filter("available__exact", "Y")

	_, err := qs.All(history)
	if err != nil {
		return nil, err
	}
	beego.Debug("history:", history)

	_, err = o.LoadRelated(history, "requirement")
	if err != nil {
		return nil, err
	}

	beego.Info("end get a requirement with history model.")
	return history, err
}

func GetAll(offset int64, size int) ([]*History, error) {
	beego.Info("begin get all requirements model.")

	o := orm.NewOrm()
	o.Using("datahub")

	count, err := o.QueryTable("dh_rm_history").Filter("available__exact", "Y").Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return []*History{}, nil
	}
	beego.Debug("count:", count)

	validateOffsetAndLimit(count, &offset, &size)

	var historys []*History
	_, err = o.QueryTable("dh_rm_history").Filter("available__exact", "Y").Limit(size, offset).All(&historys)
	if err != nil {
		return nil, err
	}
	beego.Debug("historys:", historys)

	for _, history := range historys {
		_, err = o.LoadRelated(history, "requirement")
		if err != nil {
			return nil, err
		}
	}

	beego.Info("end get all requirements model.")
	return historys, err
}
func validateOffsetAndLimit(count int64, offset *int64, limit *int) {
	if *limit < 1 {
		*limit = 1
	}
	if *offset >= count {
		*offset = count - int64(*limit)
	}
	if *offset < 0 {
		*offset = 0
	}
	if *offset+int64(*limit) > count {
		*limit = int(count - *offset)
	}
}

func Update(req *Requirement) (int64, error) {
	beego.Info("begin update requirement model.")

	o := orm.NewOrm()
	o.Using("datahub")

	rows, err := o.Update(req)
	if err != nil {
		return 0, err
	}

	beego.Debug("Update reqId:", req.Id)
	var idList orm.ParamsList
	_, err = o.Raw("SELECT id FROM dh_rm_history WHERE available = 'Y' AND requirement_id = ?", req.Id).ValuesFlat(&idList)
	if err != nil {
		return 0, err
	}
	beego.Debug(idList)

	rs := o.Raw("UPDATE dh_rm_history SET available = ? WHERE id = ?", "N", idList[len(idList)-1])
	_, err = rs.Exec()
	if err != nil {
		return 0, err
	}

	history := new(History)
	history.Remark = req.Remark
	history.Status = req.Status
	history.Review_user = req.Review_user
	history.Available = "Y"
	history.Requirement = req
	_, err = o.Insert(history)
	if err != nil {
		return 0, err
	}

	beego.Info("end update requirement model.")
	return rows, err
}

func Delete(id int) error {
	beego.Info("begin delete requirement model.")

	o := orm.NewOrm()
	o.Using("datahub")

	rs := o.Raw("UPDATE dh_requirement SET status = ? WHERE id = ?", "需求取消", id)
	_, err := rs.Exec()
	if err != nil {
		return err
	}

	beego.Info("end delete requirement model.")
	return err
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

	//orm.DefaultTimeLoc = time.UTC
}
