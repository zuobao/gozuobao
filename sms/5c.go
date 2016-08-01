package sms

import (
	"net/url"
	"log"
	"net/http"
	"errors"
)


/**


 */

// 美联短信发送（可发国际短信）： http://www.5c.com.cn/
type MLSms struct {
	Username, Password string
	ApiKey, Url        string
	Suffix             string
}

func (me *MLSms) Send(mobile, msg string) error {

	if len(me.Suffix) > 0 {
		msg = me.Suffix + msg
	}

	params := url.Values{}
	params.Set("mobile", mobile)
	params.Set("username", me.Username)
	params.Set("password", me.Password)
	params.Set("content", msg)
	params.Set("apikey", me.ApiKey)


	resp, err := http.PostForm(me.Url, params)
	if err != nil {
		log.Println(err)
		return errors.New("网络错误: " + err.Error())
	}

	smsResult := YpResult{}

	defer resp.Body.Close()
//	buf, _ := ioutil.ReadAll(resp.Body)// 此时buf已经是字节类型

	// 转换为 string
//	content := string(buf)

	if smsResult.Code == 0 {

		return nil

	}else{

		log.Println("错误信息：", &smsResult)

		return &smsResult

	}


	return nil
}
