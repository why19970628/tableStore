package main

import (
	"fmt"
	"net/http"
	"tableStore/api"
	"tableStore/lib"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func init() {
	lib.LoadConfig() // 初始化配置文件
	lib.TsClient = tablestore.NewClient(lib.Conf.Endpoint, lib.Conf.InstanceName, lib.Conf.AccessKeyID, lib.Conf.AccessKeySecret)
	// models.InitFilmTable()
}

func main() {
	http.HandleFunc("/dicom", api.DICOMHandler)

	// 启动
	fmt.Println("starting server on port ", lib.Conf.Port, "...")
	http.ListenAndServe(":"+lib.Conf.Port, nil)
}
