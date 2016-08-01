package push

import (
	"encoding/json"
	"github.com/zuobao/gozuobao/logger"
	"github.com/wuyongzhi/jpushv3"
)

type JPushV3Engine struct {
	AppKey         string
	MasterSecret   string
	APNSProduction bool
}

func NewJPushV3Engine(appKey, masterSecret string, apns_production bool) PushEngine {

	engine := JPushV3Engine{
		AppKey:         appKey,
		MasterSecret:   masterSecret,
		APNSProduction: apns_production,
	}

	return &engine
}

func (this *JPushV3Engine) PushAll(title string, body interface{}, options *PushOptions) (MsgId, error) {
	var pf jpushv3.Platform
	if options != nil  {
		if options.disableAndroid || options.disableIOS {
			if !options.disableAndroid {
				pf.Add("android")
			}
			if !options.disableIOS {
				pf.Add("ios")
			}
		}
	}
	pf.SetAllIfNil()



	var ad jpushv3.Audience
	ad.All()

	mb := jpushv3.NewMessageAndNoticeBuilder()

	buf, err := json.Marshal(body)
	if err != nil {
		return InvalidMsgId, err
	}

	var msg jpushv3.Message
	msg.Title = title
	msg.Content = string(buf)

	// 如果 title 有内容，则创建通知对象
	if title != "" {
		notify := jpushv3.NewNotification("")
		notify.SetSimpleAlert("")
		notify.SetIOSAlert(title)

		if options != nil {
			if options.iosBadge != "" {
				notify.SetIOSBadge(options.iosBadge)
			}
			if options.iosExtras != nil {
				notify.SetIOSExtras(options.iosExtras)
			}
			if options.iosSound != nil {
				sound, ok := options.iosSound.(string)
				if ok {
					notify.SetIOSSound(sound)
				}
			}
		}

		mb.SetNotification(notify)
	}

	mb.SetMessage(&msg)
	mb.SetAudience(&ad)
	mb.SetPlatform(&pf)
	mb.Options.Apns_production = this.APNSProduction

	client := jpushv3.NewPushClient(this.MasterSecret, this.AppKey)

	msgid, err := client.Send(mb)
	ok := "OK"
	if err != nil {
		logger.Errorf("JPushv3 Push All ERROR: %v. ", ok, err)
	} else {
		logger.Debugf("JPushv3 Push All OK. msgid=%v ", msgid)
	}

	return msgid, err
}


func (this *JPushV3Engine) PushAlias(title string, msgBody interface{}, options *PushOptions, alias []string) (MsgId, error) {

	var ad jpushv3.Audience
	ad.SetAlias(alias)

	return this.push(title, msgBody, options, &ad)
}


func (this *JPushV3Engine) PushTags(title string, msgBody interface{}, options *PushOptions, tags ...string) (MsgId, error) {
	var pf jpushv3.Platform

	// PushOptions 可以控制推送到哪些平台
	if options != nil  {
		if options.disableAndroid || options.disableIOS {
			if !options.disableAndroid {
				pf.Add("android")
			}
			if !options.disableIOS {
				pf.Add("ios")
			}
		}
	}
	pf.SetAllIfNil()

	//	pf.Add(jpushv3.IOS)
	//	pf.Add(jpushv3.ANDROID)

	var ad jpushv3.Audience
	ad.SetTag(tags)

	var msg jpushv3.Message
	buf, err := json.Marshal(msgBody)
	if err != nil {
		return InvalidMsgId, err
	}

	msg.Content = string(buf)
	msg.Title = title
	//	msg.Extras = msgBody
	//	msg.ContentType = "1"

	mb := jpushv3.NewMessageAndNoticeBuilder()
	if title != "" {
		notify := jpushv3.NewNotification("")
		notify.SetSimpleAlert("")
		notify.SetIOSAlert(title)
		if options != nil {
			if options.iosBadge != "" {
				notify.SetIOSBadge(options.iosBadge)
			}
			if options.iosExtras != nil {
				notify.SetIOSExtras(options.iosExtras)
			}
			if options.iosSound != nil {
				sound, ok := options.iosSound.(string)
				if ok {
					notify.SetIOSSound(sound)
				}
			}
		}

		mb.SetNotification(notify)
	}
	mb.SetMessage(&msg)
	mb.SetAudience(&ad)
	mb.SetPlatform(&pf)

	mb.Options.Apns_production = this.APNSProduction

	if options != nil {
		if options.ttl > 300 {
			mb.Options.SetTimelive(options.ttl)
		}
	}

	client := jpushv3.NewPushClient(this.MasterSecret, this.AppKey)

	msgid, err := client.Send(mb)
	if err != nil {
		logger.Errorf("JPushv3 Push tags ERROR, %#v, %v. ", tags, err)
	} else {
		logger.Debugf("JPushv3 Push tags OK, tags=%v,result=%v ", tags, msgid)
	}

	return msgid, err

}


func (this *JPushV3Engine) push(title string, msgBody interface{}, options *PushOptions, audience *jpushv3.Audience) (MsgId, error) {
	var pf jpushv3.Platform

	// PushOptions 可以控制推送到哪些平台
	if options != nil  {
		if options.disableAndroid || options.disableIOS {
			if !options.disableAndroid {
				pf.Add("android")
			}
			if !options.disableIOS {
				pf.Add("ios")
			}
		}
	}

	pf.SetAllIfNil()

	//	pf.Add(jpushv3.IOS)
	//	pf.Add(jpushv3.ANDROID)

	var msg jpushv3.Message
	buf, err := json.Marshal(msgBody)
	if err != nil {
		return InvalidMsgId, err
	}

	msg.Content = string(buf)
	msg.Title = title
	//	msg.Extras = msgBody
	//	msg.ContentType = "1"

	mb := jpushv3.NewMessageAndNoticeBuilder()
	if title != "" {
		notify := jpushv3.NewNotification("")
		notify.SetSimpleAlert("")
		notify.SetIOSAlert(title)
		if options != nil {
			if options.iosBadge != "" {
				notify.SetIOSBadge(options.iosBadge)
			}
			if options.iosExtras != nil {
				notify.SetIOSExtras(options.iosExtras)
			}
			if options.iosSound != nil {
				sound, ok := options.iosSound.(string)
				if ok {
					notify.SetIOSSound(sound)
				}
			}
		}

		mb.SetNotification(notify)
	}
	mb.SetMessage(&msg)
	mb.SetAudience(audience)
	mb.SetPlatform(&pf)

	mb.Options.Apns_production = this.APNSProduction

	if options != nil {
		if options.ttl > 300 {
			mb.Options.SetTimelive(options.ttl)
		}
	}

	client := jpushv3.NewPushClient(this.MasterSecret, this.AppKey)

	msgid, err := client.Send(mb)
	if err != nil {
		logger.Errorf("JPushv3 Push ERROR,  %v. ", err)
	} else {
		logger.Debugf("JPushv3 Push OK, result=%v ", msgid)
	}

	return msgid, err

}
