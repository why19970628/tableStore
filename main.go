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

	// 启动
	fmt.Println("starting server on port ", config.Conf.Port, "...")
	http.ListenAndServe(":"+config.Conf.Port, nil)
}
