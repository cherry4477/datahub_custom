package controllers

import (
	"github.com/asiainfoLDP/datahub_custom/ds"
	"github.com/asiainfoLDP/datahub_custom/models"
	"github.com/astaxie/beego"
	"net/http"
	"strconv"
	"strings"
)

func sendResult(this beego.Controller, statusCode int, code int, msg string, data interface{}) {
	this.Ctx.Output.SetStatus(statusCode)
	result := ds.Result{Code: code, Msg: msg, Data: data}

	this.Data["json"] = &result
	this.ServeJSON()
}

type QueryListResult struct {
	Total   int64       `json:"total"`
	Results interface{} `json:"results"`
}

func NewQueryListResult(count int64, results interface{}) *QueryListResult {
	return &QueryListResult{Total: count, Results: results}
}
func OptionalOffsetAndSize(r *http.Request, defaultSize int64, minSize int64, maxSize int64) (int64, int) {
	r.ParseForm()
	page := optionalIntParamInQuery(r, "page", 0)

	beego.Debug("page:", page)

	if page < 1 {
		page = 1
	}
	page -= 1

	if minSize < 1 {
		minSize = 1
	}
	if maxSize < 1 {
		maxSize = 1
	}
	if minSize > maxSize {
		minSize, maxSize = maxSize, minSize
	}

	size := optionalIntParamInQuery(r, "size", defaultSize)

	beego.Debug("size:", size)

	if size < minSize {
		size = minSize
	} else if size > maxSize {
		size = maxSize
	}

	return page * size, int(size)
}

func optionalIntParamInQuery(r *http.Request, paramName string, defaultInt int64) int64 {
	return _optionalIntParam(r.Form.Get(paramName), defaultInt)
}

func _optionalIntParam(intStr string, defaultInt int64) int64 {
	if intStr == "" {
		return defaultInt
	}

	beego.Debug("inStr:", intStr)

	i, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return defaultInt
	} else {
		return i
	}
}

func getLoginName(this beego.Controller) string {
	loginStr := this.Ctx.Request.Header.Get("User")
	if loginStr == "" {
		return ""
	}
	loginName := strings.Split(loginStr, "+")[1]

	return loginName
}

func validateParams(object *models.Requirement) bool {

	if len([]rune(object.Requirement_content)) > 200 ||
		len([]rune(object.Requirement_name)) > 100 ||
		len([]rune(object.Trade)) > 52 ||
		len([]rune(object.Scope)) > 200 ||
		len([]rune(object.Scene)) > 200 ||
		len([]rune(object.Frequency)) > 52 ||
		len([]rune(object.Deliver)) > 52 ||
		len([]rune(object.Remark)) > 200 {

		return false
	}

	return true
}
