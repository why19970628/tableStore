package main

import (
	"tableStore/lib"
	"tableStore/models"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/kataras/iris/context"

	"github.com/kataras/iris"
)

func init() {
	lib.LoadConfig() // 初始化配置文件
	lib.TsClient = tablestore.NewClient(lib.Conf.Endpoint, lib.Conf.InstanceName, lib.Conf.AccessKeyID, lib.Conf.AccessKeySecret)
	// models.InitFilmTable()
}

func main() {
	app := iris.New()

	// Get 读取数据
	app.Get("/dicom/{code}", func(ctx context.Context) {
		code := ctx.Params().Get("code")
		var result *lib.Result
		film, _ := models.GetOneFilm(code)
		result = &lib.Result{Success: true, Data: film, Code: code}
		ctx.JSON(&result)
	})

	// Add 新增数据
	app.Post("/dicom", func(ctx context.Context) {
		var result *lib.Result
		film := models.Film{}
		ctx.ReadForm(&film)
		//
		if film.Code == "" {
			result = &lib.Result{Success: false, Msg: "缺少code参数"}
		} else {
			code := models.AddOneFilm(film) //
			result = &lib.Result{Success: true, Code: code, Msg: "保存成功"}
		}
		ctx.JSON(&result)
	})

	// Delete 删除数据
	app.Delete("/dicom/{code}", func(ctx context.Context) {
		code := ctx.Params().Get("code")
		var result *lib.Result
		err := models.DeleteOneFilm(code)
		if err != nil {
			result = &lib.Result{Success: false, Code: code, Msg: "删除错误: " + err.Error()}
		} else {
			result = &lib.Result{Success: true, Code: code, Msg: "删除成功"}
		}
		ctx.JSON(&result)
	})

	// 启动
	app.Run(iris.Addr(":" + lib.Conf.Port))
}
