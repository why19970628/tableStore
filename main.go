package main

import (
	"fmt"
	"net/http"
	"tableStore/api"
	"tableStore/config"
	"tableStore/lib"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func init() {
	config.Load() // 初始化配置文件
	lib.TsClient = tablestore.NewClient(config.Conf.Endpoint, config.Conf.InstanceName, config.Conf.AccessKeyID, config.Conf.AccessKeySecret)
	// models.InitFilmTable()
}

func main() {
	http.HandleFunc("/dicom", api.DICOMHandler)

	// 添加
	// f := models.Film{
	// 	Code:      "wangzhicheng",
	// 	OtherCode: "哦也",
	// 	ImageURL:  "我是地址"}
	// cc := models.AddOneFilm(f)
	// fmt.Println("code =>", cc)

	// 读取
	// film, _ := models.GetOneFilm("wangzhicheng2")
	// fmt.Println(film)
	// fmt.Println(film.OtherCode)

	// 启动
	fmt.Println("starting server on port ", config.Conf.Port, "...")
	http.ListenAndServe(":"+config.Conf.Port, nil)
}
