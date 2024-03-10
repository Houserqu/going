package feishu

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type RobotCardMsgParam struct {
	HookUrl string                    `json:"hook_url"`
	Title   string                    `json:"title"`
	Color   string                    `json:"color"`
	Content string                    `json:"content"`
	Fields  []RobotCardMsgFieldParam  `json:"fields"`
	Remark  string                    `json:"remark"`
	Buttons []RobotCardMsgButtonParam `json:"buttons"`
}

type RobotCardMsgFieldParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RobotCardMsgButtonParam struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

/**
 * 发送飞书机器人卡片消息
 */
func SendRobotCardMsg(params RobotCardMsgParam) (err error) {
	if params.HookUrl == "" {
		err = errors.New("hook url is empty")
		return
	}

	// 请求体
	data := map[string]any{
		"msg_type": "interactive",
	}

	elements := []any{}

	// 处理 fields
	if params.Fields != nil {
		fields := []map[string]any{}
		for _, field := range params.Fields {
			fields = append(fields, map[string]any{
				"is_short": true,
				"text": map[string]any{
					"tag":     "lark_md",
					"content": fmt.Sprintf("**%s: **\n%s", field.Name, field.Value),
				},
			})
		}

		elements = append(elements, map[string]any{
			"tag":    "div",
			"fields": fields,
		})
	}

	// 处理富文本内容
	if params.Content != "" {
		elements = append(elements, map[string]any{
			"tag":     "markdown",
			"content": params.Content,
		})
	}

	// 分割线
	if len(params.Buttons) > 0 || params.Remark != "" {
		elements = append(elements, map[string]any{
			"tag": "hr",
		})
	}

	// 按钮
	if len(params.Buttons) > 0 {
		buttons := []map[string]any{}
		for _, button := range params.Buttons {
			buttons = append(buttons, map[string]any{
				"tag": "button",
				"text": map[string]any{
					"tag":     "plain_text",
					"content": button.Text,
				},
				"type": "primary",
				"multi_url": map[string]string{
					"url":         button.Url,
					"pc_url":      "",
					"android_url": "",
					"ios_url":     "",
				},
			})
		}

		elements = append(elements, map[string]any{
			"tag":     "action",
			"actions": buttons,
		})
	}

	// 备注
	if params.Remark != "" {
		elements = append(elements, map[string]any{
			"tag": "note",
			"elements": []any{
				map[string]string{
					"tag":     "plain_text",
					"content": params.Remark,
				},
			},
		})
	}

	template := "blue"
	if params.Color != "" {
		template = params.Color
	}

	card := map[string]any{
		"header": map[string]any{
			"template": template,
			"title": map[string]any{
				"content": params.Title,
				"tag":     "plain_text",
			},
		},
		"config": map[string]any{
			"wide_screen_mode": true,
		},
	}

	card["elements"] = elements
	data["card"] = card
	str, _ := json.Marshal(data)

	_, err = resty.New().SetDebug(true).R().
		SetBody(string(str)).
		Post(params.HookUrl)

	return
}

/**
 * 发送飞书机器人简单消息
 */
func SendRobotSimpleMsg(hookUrl string, title string, content string) (err error) {
	if hookUrl == "" {
		err = errors.New("hook url is empty")
		return
	}

	params := map[string]any{
		"msg_type": "post",
		"content": map[string]any{
			"post": map[string]any{
				"zh_cn": map[string]any{
					"title": title,
					"content": [][]map[string]any{
						{
							{
								"tag":  "text",
								"text": content,
							},
						},
					},
				},
			},
		},
	}
	str, _ := json.Marshal(params)

	client := resty.New()
	_, err = client.R().SetBody(string(str)).Post(hookUrl)
	return
}
