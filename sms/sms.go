package sms

type SmsEngine interface {
	Send(mobile string, msg string) error
}

type MobileSetting struct {
	Prefix      string
	Country     string
	PhoneLength int
}

var mobileSettings = []*MobileSetting {
	&MobileSetting{"86", "中国", 11},
	&MobileSetting{"886", "中国香港", 9},
	&MobileSetting{"886", "中国澳门", 9},
	&MobileSetting{"886", "台湾", 9},

	&MobileSetting{"886", "美国", 9},

	&MobileSetting{"886", "英国", 9},
	&MobileSetting{"886", "法国", 9},
	&MobileSetting{"886", "德国", 9},
	&MobileSetting{"886", "意大利", 9},
	&MobileSetting{"886", "俄罗斯", 9},

	&MobileSetting{"886", "澳大利亚", 9},

	&MobileSetting{"886", "巴西", 9},

	&MobileSetting{"81", "日本", 9},
	&MobileSetting{"886", "泰国", 9},
	&MobileSetting{"82", "韩国", 9},

	&MobileSetting{"886", "加拿大", 9},
	&MobileSetting{"886", "新加坡", 9},
	&MobileSetting{"886", "马来西亚", 9},
}

