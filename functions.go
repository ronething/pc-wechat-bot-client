package wechat_bot

import (
	"time"
)

func getId() string {
	return time.Now().Format("20060102150405")
}

type GetContactList struct {
	Content []*Friend `json:"content"`
	Type    int64     `json:"type"`
}

type GetMemberList struct {
	Content []*Member `json:"content"`
	Type    int64     `json:"type"`
}

type Friend struct {
	Name string `json:"name"`
	WxId string `json:"wxid"`
}

type Member struct {
	NickName string `json:"nickname"`
	RoomId   string `json:"roomid"`
	WxId     string `json:"wxid"`
}

type Bot interface {
	//群聊发送 @member 消息
	SendAtMsg(content, roomId, atWxId, nick string) error
	//发送文本消息
	SendTxtMsg(content, wxId string) error
	//发送图片
	SendPic(wxId, filePath string) error
	//发送附件
	SendAttach(wxId, filePath string) error
	//获取群成员昵称
	GetMemberNick(roomId string) (res []*Member, err error)
	//获取群成员 id
	GetMemberId() (resp interface{}, err error)
	//获取微信通讯录好友和群列表
	GetContactList() (res []*Friend, err error)
	//刷新好友列表
	RefreshMemberList() error
	//动态卸载 ddl
	SendDestroy() error
}
