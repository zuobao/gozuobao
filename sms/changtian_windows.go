package sms

import (
	"errors"
)



type CtySms struct {
	Username, Password, Url string
}

type CtyResult struct {
	succeed bool
	code int
	msg string
}

var ctyResults = map[string]CtyResult{
	"1":   {true, 	1, "成功"},
	"-1":  {false, -1, "用户名和密码参数为空或者参数含有非法字符"},
	"-2":  {false, -2, "手机号参数不正确"},
	"-3":  {false, -3, "msg参数为空或长度小于0个字符"},
	"-4":  {false, -4, "msg参数长度超过64个字符"},
	"-6":  {false, -6, "发送号码为黑名单用户"},
	"-8":  {false, -8, "下发内容中含有屏蔽词"},
	"-9":  {false, -9, "下发账户不存在"},
	"-10": {false, -10, "下发账户已经停用"},
	"-11": {false, -11, "下发账户无余额"},
	"-15": {false, -15, "MD5校验错误"},
	"-16": {false, -16, "IP服务器鉴权错误"},
	"-17": {false, -17, "接口类型错误"},
	"-18": {false, -18, "服务类型错误"},
	"-22": {false, -22, "手机号达到当天发送限制"},
	"-23": {false, -23, "同一手机号，相同内容达到当天发送限制"},
	"-24": {false, -24, "模板不存在"},
	"-25": {false, -25, "模板变量超长"},
	"-26": {false, -26, "下发限制，该号码没有上行记录"},
	"-27": {false, -27, "手机号不是白名单用户"},
	"-99": {false, -99, "系统异常"},
}

func (me *CtyResult) Error() string {
	return me.msg
}

func (me *CtySms) Send(mobile, msg string) error {
	return errors.New("windows由于无法导入及编译 github.com/djimenez/iconv-go库，暂不支持发送SMS")
}












