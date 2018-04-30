package controllers

import (
	"encoding/json"
	"errorReportTrackingServer/models"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"fmt"
	"time"
)

// TTrackingrecordController operations for TTrackingrecord
type TTrackingrecordController struct {
	beego.Controller
}

// URLMapping ...
func (c *TTrackingrecordController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetbyTime",c.GetbyTime)
}

// Post ...
// @Title Post
// @Description create TTrackingrecord
// @Param	body		body 	models.TTrackingrecord	true		"body for TTrackingrecord content"
// @Success 201 {int} models.TTrackingrecord
// @Failure 403 body is empty
// @router / [post]
func (c *TTrackingrecordController) Post() {
	var v models.TTrackingrecord
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTTrackingrecord(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get TTrackingrecord by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TTrackingrecord
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TTrackingrecordController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTTrackingrecordById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TTrackingrecord
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TTrackingrecord
// @Failure 403
// @router / [get]
func (c *TTrackingrecordController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	l, err := models.GetAllTTrackingrecord(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
		/*for _,v := range l{
			value:=v.(models.TTrackingrecord)
			value.TRTimestamp
		}*/
		/*fmt.Println("fields: ", fields)
		fmt.Println("query: ", query)
		fmt.Println("limit: ",limit)
		fmt.Println("offset: ", offset)
		fmt.Println("sortby: ", sortby)
		fmt.Println("order: ",order)*/
	}
	c.ServeJSON()
}
func (c *TTrackingrecordController) GetbyTime() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	var T1,T2 time.Time
	if v := c.GetString("time"); v != "" {
			Time := strings.SplitN(v, "|", 2)
			if len(Time) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}

			Tstart,_ := time.Parse("2006-01-02 03:04:05 PM",Time[0])
			Tend,_ := time.Parse("2006-01-02 03:04:05 PM",Time[1])
			fmt.Println(Tstart)
			fmt.Println(Tend)
			T1=Tstart
			if Time[1]=="" {
				T20:=time.Now()
				T21 := T20.Format("2006-01-02 03:04:05 PM")
				T2,_ =time.Parse("2006-01-02 03:04:05 PM",T21)
			}else {T2=Tend}
			fmt.Println(T1)
			fmt.Println(T2)
	}
	l, err := models.GetAllTTrackingrecord(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var vlist []models.TTrackingrecord
		for _,v := range l{
			value:=v.(models.TTrackingrecord)
			vts := value.TRTimestamp.Format("2006-01-02 03:04:05 PM")
			vtt,_ :=  time.Parse("2006-01-02 03:04:05 PM",vts)
			if vtt.After(T1) && vtt.Before(T2) {
				vlist = append(vlist,value)
			}
		}
		c.Data["json"] = vlist

		/*for _,v := range l{
			value:=v.(models.TTrackingrecord)
			value.TRTimestamp
		}*/
		/*fmt.Println("fields: ", fields)
		fmt.Println("query: ", query)
		fmt.Println("limit: ",limit)
		fmt.Println("offset: ", offset)
		fmt.Println("sortby: ", sortby)
		fmt.Println("order: ",order)*/
	}
	c.ServeJSON()
}
// Put ...
// @Title Put
// @Description update the TTrackingrecord
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TTrackingrecord	true		"body for TTrackingrecord content"
// @Success 200 {object} models.TTrackingrecord
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TTrackingrecordController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.TTrackingrecord{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTTrackingrecordById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the TTrackingrecord
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TTrackingrecordController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTTrackingrecord(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
