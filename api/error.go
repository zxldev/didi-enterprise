package api

import "errors"

var ErrorNetWork = errors.New("网络连接错误")
var ErrorDecode = errors.New("数据解析错误")

var ErrorGetAccessToken = errors.New("获取access_token失败")
