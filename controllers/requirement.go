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

type ARequirementController struct {
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
func (this *ARequirementController) Post() {
	beego.Informational(this.Ctx.Request.URL, "Admin create a requirement.")

	this.auth()

	var object models.Requirement
	json.Unmarshal(this.Ctx.Input.RequestBody, &object)
	beego.Debug(object)

	models.AddOne(object)
	this.Data["json"] = "Insert success."
	//this.Data["json"] = map[string]int64{"InsertId": insertId}
	this.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /requirement [get]
func (this *ARequirementController) Get() {
	beego.Informational(this.Ctx.Request.URL, "Admin get requirements by params.")

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

	obs := models.Get(params)

	this.Data["json"] = obs
	this.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *ARequirementController) GetAll() {
	beego.Informational(this.Ctx.Request.URL, "Admin get all requirement.")

	this.auth()

	obs := models.GetAll()
	this.Data["json"] = obs
	this.ServeJSON()
}

// @Title update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:reqId [put]
func (this *ARequirementController) Put() {
	beego.Informational(this.Ctx.Request.URL, "Admin update a requirement.")

	reqId := this.Ctx.Input.Param(":reqId")

	var object models.Requirement

	id, _ := strconv.Atoi(reqId)

	object = models.GetById(id)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &object)
	beego.Debug(string(this.Ctx.Input.RequestBody))
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(object)

	object.Id, _ = strconv.Atoi(reqId)
	models.Update(object)
	beego.Debug(object)
	this.ServeJSON()
}

// @Title delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ARequirementController) Delete() {
	//objectId := o.Ctx.Input.Param(":objectId")
	//models.Delete(objectId)
	//o.Data["json"] = "delete success!"
	//o.ServeJSON()
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
		this.SendResult(http.StatusBadRequest, ds.ErrorUnmarshal, err.Error(), nil)
	}
	beego.Debug(object)
	_, err = models.AddOne(object)
	if err != nil {
		beego.Error("models, addone err :", err)
		this.SendResult(http.StatusBadRequest, ds.ErrorAddModel, err.Error(), nil)
	}

	this.SendResult(http.StatusOK, ds.ResultOK, "OK.", nil)
}

func (this *ARequirementController) auth() {
	loginStr := this.Ctx.Request.Header.Get("User")
	region := strings.Split(loginStr, "+")[0]
	if region == ""  || region == "datahub" {
		beego.Notice("not authorized.")
		this.Abort("401")
	}
}

func (this *DRequirementController) auth() {
	loginStr := this.Ctx.Request.Header.Get("User")
	region := strings.Split(loginStr, "+")[0]
	if region == "" || region != "datahub" {
		beego.Notice("not authorized.")
		this.SendResult(http.StatusUnauthorized, ds.ErrorUnauthorized, "not authorized.", nil)
	}
}

func (this *DRequirementController) SendResult(statusCode int, code int, msg string, data interface{}) {
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
