package lib

import (
	"encoding/json"
	"fmt"
)

// Result 控制器返回结果
type Result struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
}

// ToJSON 转化对象为JSON字符串
func (result *Result) ToJSON() string {
	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(json)
}
