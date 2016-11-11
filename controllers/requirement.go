package controllers

import (
	"encoding/json"
	"github.com/asiainfoLDP/datahub_custom/ds"
	_ "github.com/asiainfoLDP/datahub_custom/logs"
	"github.com/asiainfoLDP/datahub_custom/models"
	//"github.com/asiainfoLDP/datahub_custom/pager"
	"github.com/astaxie/beego"
	"net/http"
	"strconv"
	"strings"
)

type ORequirementController struct {
	beego.Controller
}

type DRequirementController struct {
	beego.Controller
}

// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ORequirementController) Create() {
	beego.Informational(this.Ctx.Request.URL, "Operation create a requirement.")

	//this.auth()
	loginName := getLoginName(this.Controller)

	requirement := new(models.Requirement)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, requirement)
	if err != nil {
		beego.Error("Unmarshal err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
		return
	}
	requirement.Create_user = loginName
	beego.Debug(requirement)

	if flag := validateParams(requirement); flag == false {
		beego.Info("输入参数长度有误.")
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorValidateParams, "输入参数长度有误.", nil)
		return
	}

	requirementId, err := models.AddOne(requirement)
	if err != nil {
		beego.Error("Model, AddOne err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAddModel, err.Error(), nil)
		return
	}
	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", requirementId)
}

//@router /requirement/:reqId [get]
func (this *ORequirementController) GetById() {
	beego.Informational(this.Ctx.Request.URL, "Operation get requirements by params.")

	reqId := this.Ctx.Input.Param(":reqId")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
		return
	}

	history, err := models.GetNewHistory(id)
	if err != nil {
		beego.Error("Model, get by id err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModelById, err.Error(), nil)
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", history)
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /requirement [get]
func (this *ORequirementController) Get() {
	beego.Informational(this.Ctx.Request.URL, "Operation get requirements by params.")
	beego.Info("begin get requirement by param handler.")

	//this.auth()

	var name, phone, email, company, content string
	var params = make(map[string]string)

	this.Ctx.Input.Bind(&name, "name")
	this.Ctx.Input.Bind(&phone, "phone")
	this.Ctx.Input.Bind(&email, "email")
	this.Ctx.Input.Bind(&company, "company")
	this.Ctx.Input.Bind(&content, "content")
	params["name"] = name
	params["phone"] = phone
	params["email"] = email
	params["company"] = company
	params["content"] = content
	beego.Debug(params)

	reqs, err := models.GetByParams(params)
	if err != nil {
		beego.Error("Model, GetByParamsFilerUser err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
		return
	}

	beego.Info("end get requirement by param handler.")
	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", reqs)
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *ORequirementController) GetAll() {
	beego.Informational(this.Ctx.Request.URL, "Operation get all requirement.")

	//this.auth()

	offset, size := OptionalOffsetAndSize(this.Ctx.Request, 30, 1, 100)
	beego.Debug("offset, size:", offset, size)

	reqs, err := models.GetAll(offset, size)
	if err != nil {
		beego.Error("Model, GetAll err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
		return
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", NewQueryListResult(int64(len(reqs)), reqs))
}

// @Title update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:reqId [put]
func (this *ORequirementController) Update() {
	beego.Info(this.Ctx.Request.URL, "Operation update a requirement.")
	beego.Info("begin update requirement handler.")
	//this.auth()
	//loginName := getLoginName(this.Controller)

	reqId := this.Ctx.Input.Param(":reqId")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
		return
	}

	req, err := models.GetById(id)
	if err != nil {
		beego.Error("Model, GetById err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
		return
	}

	err = json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		beego.Error("Unmarshal err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
		return
	}

	req.Id, err = strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
		return
	}
	beego.Debug(req)

	_, err = models.Update(req)
	if err != nil {
		beego.Error("Model, Update err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUpdateModel, err.Error(), nil)
		return
	}

	beego.Info("end update requirement handler.")
	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

//@router /review/:reqId [put]
func (this *ORequirementController) Review() {
	beego.Informational(this.Ctx.Request.URL, "Operation review a requirement.")

	//this.auth()
	loginName := getLoginName(this.Controller)

	reqId := this.Ctx.Input.Param(":reqId")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
		return
	}

	req, err := models.GetById(id)
	if err != nil {
		beego.Error("Model, GetById err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
		return
	}

	err = json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		beego.Error("Unmarshal err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
		return
	}

	req.Id, err = strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
		return
	}
	req.Review_user = loginName
	beego.Debug(req)

	_, err = models.Update(req)
	if err != nil {
		beego.Error("Model, Update err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUpdateModel, err.Error(), nil)
		return
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

// @Title delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:reqId [delete]
func (this *ORequirementController) Delete() {
	reqId := this.Ctx.Input.Param(":reqId")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
		return
	}

	err = models.Delete(id)
	if err != nil {
		beego.Error("Model, Delete err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorDeleteModel, err.Error(), nil)
		return
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

//@router / [post]
func (this *DRequirementController) Create() {
	beego.Informational(this.Ctx.Request.URL, "Begin datahub create a requirement.")

	loginName := getLoginName(this.Controller)
	beego.Debug("loginname:", loginName)
	requirement := new(models.Requirement)
	if loginName != "" {
		requirement.Create_user = loginName
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, requirement)
	if err != nil {
		beego.Error("Unmarshal err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
		return
	}
	requirement.Status = "需求提交"
	beego.Debug(requirement)

	if flag := validateParams(requirement); flag == false {
		beego.Info("输入参数长度有误.")
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorValidateParams, "输入参数长度有误.", nil)
		return
	}

	_, err = models.AddOne(requirement)
	if err != nil {
		beego.Error("Model, AddOne err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAddModel, err.Error(), nil)
		return
	}

	beego.Info(this.Ctx.Request.URL, "End datahub create a requirement.")
	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

//@router /requirement [get]
func (this *DRequirementController) Get() {
	beego.Informational(this.Ctx.Request.URL, "Datahub get requirements by params.")

	this.auth()
	loginName := getLoginName(this.Controller)
	//loginName := "wangmeng"

	var name, phone, email, company, content string
	var params = make(map[string]string)

	this.Ctx.Input.Bind(&name, "name")
	this.Ctx.Input.Bind(&phone, "phone")
	this.Ctx.Input.Bind(&email, "email")
	this.Ctx.Input.Bind(&company, "company")
	this.Ctx.Input.Bind(&content, "content")
	params["name"] = name
	params["phone"] = phone
	params["email"] = email
	params["company"] = company
	params["content"] = content
	params["loginUser"] = loginName
	beego.Debug(params)

	offset, size := OptionalOffsetAndSize(this.Ctx.Request, 30, 1, 100)
	beego.Debug("offset, size:", offset, size)

	count, result, err := models.GetByParamsFilterUser(params, offset, size)
	if err != nil {
		beego.Error("Model, GetByParams err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
		return
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", NewQueryListResult(count, result))
}

func (this *ORequirementController) auth() {
	loginStr := this.Ctx.Request.Header.Get("User")
	region := strings.Split(loginStr, "+")[0]
	if region == "" || region == "datahub" {
		beego.Notice("not authorized.")
		sendResult(this.Controller, http.StatusUnauthorized, ds.ErrorUnauthorized, "not authorized.", nil)
	}
}

func (this *DRequirementController) auth() {
	loginStr := this.Ctx.Request.Header.Get("User")
	region := strings.Split(loginStr, "+")[0]
	if region == "" || region != "datahub" {
		beego.Notice("not authorized.")
		sendResult(this.Controller, http.StatusUnauthorized, ds.ErrorUnauthorized, "not authorized.", nil)
	}
}
