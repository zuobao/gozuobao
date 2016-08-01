package email

import (
	"net/smtp"
	"encoding/base64"
	"bytes"

	"fmt"
	"git.oschina.net/zuobao/gozuobao/logger"
)


type SmtpEngine interface {
	Send(to string, title, content string) error
}


type Smtps []*Smtp

type Smtp struct {
	Host string
	Port string
	Username string
	Password string
	TLS bool
	From string
}

var b64 = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func (me Smtps) Send(to string, title, content string) error {
	for _, engine := range me {
		err := engine.Send(to, title, content)
		if err == nil {
			break
		}
	}
	return nil
}

func (me *Smtp) String() string {
	return fmt.Sprintf("%#v", *me)
}

func (me *Smtp) Send(to string, title, content string) error {

	auth := smtp.PlainAuth("", me.Username, me.Password, me.Host)

	header := make(map[string]string)
	header["From"] = me.From
//	header["Sender"] = "nslaile<" + me.DefaultFrom + ">"
	header["To"] = to

//	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(title)))
	header["Subject"] = title
//	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
//	header["Content-Transfer-Encoding"] = "base64"

	message := bytes.NewBufferString("")
	for k, v := range header {
		message.WriteString(k)
		message.WriteString(": ")
		message.WriteString(v)
		message.WriteString("\r\n")
	}
	message.WriteString("\r\n")
	message.WriteString(content)
//	message.WriteString(b64.EncodeToString([]byte(content)))

//	logger.Debugln(me.Host + ":" + me.Port)

	var err error

	if me.TLS {
		err = SendMailUsingTLS(me.Host + ":" + me.Port, auth, me.Username, []string{to}, []byte(message.String()))
	} else {
		err = smtp.SendMail(me.Host + ":" + me.Port, auth, me.Username, []string{to}, []byte(message.String()))
	}
	if err != nil {
		logger.Errorln(err, to)
	}

	return err
}
