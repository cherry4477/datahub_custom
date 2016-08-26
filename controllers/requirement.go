package controllers

import (
	"encoding/json"
	"github.com/asiainfoLDP/datahub_custom/ds"
	_ "github.com/asiainfoLDP/datahub_custom/logs"
	"github.com/asiainfoLDP/datahub_custom/models"
	"github.com/astaxie/beego"
	"net/http"
	"strconv"
	"strings"
)

type ERequirementController struct {
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
func (this *ERequirementController) Post() {
	beego.Informational(this.Ctx.Request.URL, "Operation create a requirement.")

	this.auth()

	var object models.Requirement
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &object)
	if err != nil {
		beego.Error("Unmarshal err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
	}
	beego.Debug(object)

	_, err = models.AddOne(object)
	if err != nil {
		beego.Error("Model, AddOne err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAddModel, err.Error(), nil)
	}
	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /requirement [get]
func (this *ERequirementController) Get() {
	beego.Informational(this.Ctx.Request.URL, "Operation get requirements by params.")

	this.auth()

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
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", reqs)
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *ERequirementController) GetAll() {
	beego.Informational(this.Ctx.Request.URL, "Operation get all requirement.")

	this.auth()

	reqs, err := models.GetAll()
	if err != nil {
		beego.Error("Model, GetAll err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", reqs)
}

// @Title update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:reqId [put]
func (this *ERequirementController) Put() {
	beego.Informational(this.Ctx.Request.URL, "Operation update a requirement.")

	//this.auth()

	reqId := this.Ctx.Input.Param(":reqId")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
	}

	req, err := models.GetById(id)
	if err != nil {
		beego.Error("Model, GetById err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
	}

	err = json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		beego.Error("Unmarshal err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
	}

	req.Id, err = strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
	}
	beego.Debug(req)

	_, err = models.Update(req)
	if err != nil {
		beego.Error("Model, Update err:", err)
		beego.Error(this.Controller, http.StatusBadRequest, ds.ErrorUpdateModel, err.Error(), nil)
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

// @Title delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:reqId [delete]
func (this *ERequirementController) Delete() {
	reqId := this.Ctx.Input.Param(":reqId")
	id , err := strconv.Atoi(reqId)
	if err != nil {
		beego.Error("Atoi err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAtoi, err.Error(), nil)
	}

	err = models.Delete(id)
	if err != nil {
		beego.Error("Model, Delete err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorDeleteModel, err.Error(), nil)
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

//@router / [post]
func (this *DRequirementController) Post() {
	beego.Informational(this.Ctx.Request.URL, "Datahub create a requirement.")

	loginName := getLoginName(this.Controller)
	var object models.Requirement
	if loginName != "" {
		object.Create_user = loginName
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &object)
	if err != nil {
		beego.Error("Unmarshal err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
	}
	beego.Debug(object)

	_, err = models.AddOne(object)
	if err != nil {
		beego.Error("Model, AddOne err :", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorAddModel, err.Error(), nil)
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", nil)
}

//@router /requirement [get]
func (this *DRequirementController) Get() {
	beego.Informational(this.Ctx.Request.URL, "Datahub get requirements by params.")

	this.auth()
	loginName := getLoginName(this.Controller)

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

	reqs, err := models.GetByParamsFilterUser(params)
	if err != nil {
		beego.Error("Model, GetByParams err:", err)
		sendResult(this.Controller, http.StatusBadRequest, ds.ErrorGetModel, err.Error(), nil)
	}

	sendResult(this.Controller, http.StatusOK, ds.ResultOK, "OK.", reqs)
}

func (this *ERequirementController) auth() {
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

func sendResult(this beego.Controller, statusCode int, code int, msg string, data interface{}) {
	this.Ctx.Output.SetStatus(statusCode)
	result := ds.Result{Code: code, Msg: msg, Data: data}

	this.Data["json"] = &result
	this.ServeJSON()
}

func getLoginName(this beego.Controller) string {
	loginStr := this.Ctx.Request.Header.Get("User")
	if loginStr == "" {
		return ""
	}
	loginName := strings.Split(loginStr, "+")[1]

	return loginName
}
