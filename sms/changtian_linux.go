package sms

import (
	"net/http"
	"net/url"
	iconv "github.com/djimenez/iconv-go"
	"log"
	"io/ioutil"
	"errors"
)



type CtySms struct {
	Username, Password, Url string
	Suffix string
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

	if len(me.Suffix) > 0 {
		msg = msg + me.Suffix
	}

	converted_msg, _ := iconv.ConvertString(msg , "utf-8", "gbk")

	params := url.Values{"un": {me.Username}, "pwd": {me.Password}, "mobile": {mobile}, "msg": {converted_msg}}
	resp, err := http.PostForm(me.Url, params)
	if err != nil {
		log.Println(err)
		return errors.New("网络错误: " + err.Error())
	}

	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	log.Println(resp.TransferEncoding)
	results , err := url.ParseQuery(string(buf))
	result := results.Get("result")
	resultObj, _ := ctyResults[result]

	if !resultObj.succeed {
		return &resultObj
	} else {
		return nil
	}
}












