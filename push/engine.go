package push

type MsgId interface{}

type PushOptions struct {
	ttl       int
	sendNo    int64
	iosBadge  string
	iosExtras interface{}
	iosSound  interface{}

	disableIOS     bool // 禁止推向ios设备
	disableAndroid bool // 禁止推向android设备
}

const DefaultTtl int = -1
const DefaultSendNo int64 = 1

var InvalidMsgId MsgId = nil

func Options() *PushOptions {
	return &PushOptions{
		ttl:      DefaultTtl,
		sendNo:   DefaultSendNo,
		iosSound: "default",
	}
}

func (this *PushOptions) SendNo(sendNo int64) *PushOptions {
	this.sendNo = sendNo
	return this
}

func (this *PushOptions) Ttl(ttl int) *PushOptions {
	this.ttl = ttl
	return this
}

func (this *PushOptions) IosBadge(iosBadge string) *PushOptions {
	this.iosBadge = iosBadge
	return this
}

func (this *PushOptions) IosExtras(iosExtras interface{}) *PushOptions {
	this.iosExtras = iosExtras
	return this
}

func (this *PushOptions) DisableIOS() {
	this.disableIOS = true
}
func (this *PushOptions) EnableIOS() {
	this.disableIOS = false
}
func (this *PushOptions) DisableAndroid() {
	this.disableAndroid = true
}
func (this *PushOptions) EnableAndroid() {
	this.disableAndroid = false
}
func (this *PushOptions) IosSound(sound interface{}) *PushOptions {
	this.iosSound = sound
	return this
}
func (this *PushOptions) DisableIosSound() {
	this.iosSound = nil
}

type PushEngine interface {
	PushAll(title string, body interface{}, options *PushOptions) (MsgId, error)
	PushTags(title string, msgBody interface{}, options *PushOptions, tags ...string) (MsgId, error)
	PushAlias(title string, msgBody interface{}, options *PushOptions, alias []string) (MsgId, error)
}
