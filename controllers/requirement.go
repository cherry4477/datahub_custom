package controllers

import (
	_ "github.com/asiainfoLDP/datahub_custom/logs"
	"github.com/asiainfoLDP/datahub_custom/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
)

// Operations about object
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
	beego.Informational(this.Ctx.Request.URL, "admin add a requirement.")

	//this.auth()

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
	beego.Informational(this.Ctx.Request.URL, "get requirements by params.")

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
	beego.Informational(this.Ctx.Request.URL, "get all requirement.")

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
	beego.Informational(this.Ctx.Request.URL, "Update a requirement.")

	reqId := this.Ctx.Input.Param(":reqId")

	var object models.Requirement
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &object)
	beego.Debug(string(this.Ctx.Input.RequestBody))
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(object)

	object.ReqId, _ = strconv.Atoi(reqId)
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

func (this *ARequirementController) auth() {
	loginName := this.Ctx.Request.Header.Get("User")
	loginName = strings.Split(loginName, "+")[0]
	if loginName == "datahub" || loginName == "" {
		beego.Notice("not authorized.")
		this.Abort("401")
	}
}
