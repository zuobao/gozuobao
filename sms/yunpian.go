package sms

import (
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"errors"
	"encoding/json"
)

type YpSms struct {
	Apikey, Url string
	Suffix string
}

type YpResult struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
	Result struct {
		Count int `json:"count"`
		Fee int `json:"fee"`
		Sid int `json:"sid"`
	}
	Detail string `json:"detail"`

}

func (me *YpResult) Error() string {
	return me.Msg
}


func (me *YpSms) Send(mobile, msg string) error {

	if len(me.Suffix) > 0 {
		msg = me.Suffix + msg
	}

	params := url.Values{"apikey": {me.Apikey}, "mobile": {mobile}, "text": {msg}}
	resp, err := http.PostForm(me.Url, params)
	if err != nil {
		log.Println(err)
		return errors.New("network error: " + err.Error())
	}

	smsResult := YpResult{}

	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)// 此时buf已经是字节类型

	err = json.Unmarshal(buf,&smsResult)// json.Unmarshal 第一个参数为字节

	if smsResult.Code == 0 {

		return nil

	}else{

		log.Println("错误信息：", smsResult.Code, smsResult.Msg)

		return &smsResult

	}

}












