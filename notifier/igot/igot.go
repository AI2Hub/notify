package igot

import (
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	rTypes "github.com/sari3l/requests/types"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Content string          `yaml:"content" json:"content"`
	Title   *string         `yaml:"title,omitempty" json:"title,omitempty"`
	Url     *string         `yaml:"url,omitempty" json:"url,omitempty"`
	Detail  *map[string]any `yaml:"detail,omitempty" json:"detail,omitempty"`
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
	data := utils.FormatAnyWithMap(n.Webhook, formatMap)
	json := utils.StructToJson(data)
	return n.Webhook, ext.Json(json)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("iGot", resp, func(request *requests.Response) bool {
		return resp.Ok && resp.Json().Get("ret").Int() == 0
	})
}
