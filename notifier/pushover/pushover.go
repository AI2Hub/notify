package pushover

import (
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	rTypes "github.com/sari3l/requests/types"
)

const DefaultWebhook = "https://api.pushover.net/1/messages.json"

type Option struct {
	types.BaseOption `yaml:",inline"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Token   string `yaml:"token" dict:"token"`
	User    string `yaml:"user" dict:"user"`
	Message string `yaml:"message" dict:"message"`
	//Attachment string `yaml:"attachment" dict:"attachment"` // 预留，待解决 requests multipart/form-data 后恢复
	Device   *string `yaml:"device,omitempty" dict:"device,omitempty"`
	Html     *int    `yaml:"html,omitempty" dict:"html,omitempty"`
	Priority *int    `yaml:"priority,omitempty" dict:"priority,omitempty"`
	Sound    *string `yaml:"sound,omitempty" dict:"sound,omitempty"`
	Title    *string `yaml:"title,omitempty" dict:"title,omitempty"`
	Url      *string `yaml:"url,omitempty" dict:"url,omitempty"`
	UrlTitle *string `yaml:"urlTitle,omitempty" dict:"urlTitle,omitempty"`
}

type Notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *Notifier {
	noticer := &Notifier{}
	noticer.Option = opt
	return noticer
}

func (n *Notifier) format(messages []string) (string, rTypes.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	data := utils.FormatAnyWithMap(n.MessageParams, formatMap)
	dict := utils.StructToDict(data)
	return DefaultWebhook, ext.Form(dict)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("PushOver", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
