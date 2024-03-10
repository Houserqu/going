package feishu

import "testing"

func TestSendRobotCardMsg(t *testing.T) {
	SendRobotCardMsg(RobotCardMsgParam{
		HookUrl: "",
		Title:   "测试卡片消息",
		Color:   "red",
		Content: "这是一条测试卡片消息",
		Fields: []RobotCardMsgFieldParam{
			{
				Name:  "字段1",
				Value: "值1",
			},
			{
				Name:  "字段2",
				Value: "值2",
			},
		},
		Remark: "这是备注",
		Buttons: []RobotCardMsgButtonParam{
			{
				Text: "按钮1",
				Url:  "https://www.feishu.cn",
			},
			{
				Text: "按钮2",
				Url:  "https://www.feishu.cn",
			},
		},
	})
}

func TestSendRobotSimpleMsg(t *testing.T) {
	SendRobotSimpleMsg("", "这是一条测试文本消息", "这是一条测试文本消息")
}
