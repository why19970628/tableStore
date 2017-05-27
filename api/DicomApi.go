package api

import (
	"io"
	"net/http"
	"reflect"
	"tableStore/lib"
	"tableStore/models"
)

// DICOMHandler DICOMHandler
func DICOMHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Get(w, r)
	} else if r.Method == "POST" {
		Add(w, r)
	} else if r.Method == "PUT" {
		Update(w, r)
	} else if r.Method == "DELETE" {
		Delete(w, r)
	} else {
		io.WriteString(w, "notice your request method")
	}
}

// Get 读取数据
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	queryString := r.URL.Query()["code"]
	var result *lib.Result
	if len(queryString) == 0 { // 未输入参数
		result = &lib.Result{Success: false, Data: nil, Msg: "未输入参数"}
	} else {
		film, _ := models.GetOneFilm(queryString[0])
		result = &lib.Result{Success: true, Data: film}
	}
	io.WriteString(w, result.ToJSON())
}

// Add 新增数据
func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	film := models.Film{}
	obj := reflect.ValueOf(&film)
	elem := obj.Elem()
	t := reflect.TypeOf(film)
	v := reflect.ValueOf(film)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			value := r.PostFormValue(t.Field(i).Name)
			elem.FieldByName(t.Field(i).Name).SetString(value) //反射注入
		}
	}

	code := models.AddOneFilm(film) //
	result := &lib.Result{Success: true, Msg: code}
	io.WriteString(w, result.ToJSON())
}

// Update 更新数据
func Update(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "coding")
}

// Delete 删除数据
func Delete(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "coding")
}
