package wechat_bot

import (
	"encoding/json"
	"time"

	"github.com/sacOO7/gowebsocket"
)

type websocketBot struct {
	socket gowebsocket.Socket
}

func NewWebsocketBot(url string) *websocketBot {
	return &websocketBot{
		socket: gowebsocket.New(url), // TODO: bind func and connect
	}
}

type WrapperMsg struct {
	Id      string `json:"id"`
	MsgType int64  `json:"type"`
	Content string `json:"content"`
	WxId    string `json:"wxid"`
}

type WrapperAtMsg struct {
	WrapperMsg
	RoomId   string `json:"roomid"`
	NickName string `json:"nickname"`
}

func NewWrapperAtMsg(wrapperMsg WrapperMsg, roomId string, nickName string) *WrapperAtMsg {
	return &WrapperAtMsg{WrapperMsg: wrapperMsg, RoomId: roomId, NickName: nickName}
}

func NewWrapperMsg(msgType int64, content string, wxId string) *WrapperMsg {
	return &WrapperMsg{Id: getId(), MsgType: msgType, Content: content, WxId: wxId}
}

func (w *websocketBot) Close() {
	w.socket.Close()
}

func (w *websocketBot) BindOnTextMessage(bindFunc func(message string, socket gowebsocket.Socket)) {
	w.socket.OnTextMessage = bindFunc
}

func (w *websocketBot) Connect() {
	w.socket.Connect()
}

//Send 最终会调用这个方法进行 websocket 的消息发送
func (w *websocketBot) Send(msg interface{}) error {
	time.Sleep(time.Second) // 统一限制一秒频率
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	w.socket.SendBinary(b)
	return nil
}

func (w *websocketBot) SendAtMsg(content, roomId, atWxId, nick string) error {
	msg := NewWrapperAtMsg(*NewWrapperMsg(AtMsg, content, atWxId), roomId, nick)
	return w.Send(msg)
}

func (w *websocketBot) SendTxtMsg(content, wxId string) error {
	return w.Send(NewWrapperMsg(TxtMsg, content, wxId))
}

func (w *websocketBot) SendPic(wxId, filePath string) error {
	return w.Send(NewWrapperMsg(PicMsg, filePath, wxId))
}

func (w *websocketBot) SendAttach(wxId, filePath string) error {
	panic("implement me") // 这个使用 http 的方法去请求
}

func (w *websocketBot) GetMemberNick(roomId string) (res []*Member, err error) {
	panic("implement me") // 这个使用 http 的方法去请求
}

func (w *websocketBot) GetMemberId() (resp interface{}, err error) {
	panic("implement me")
}

func (w *websocketBot) GetContactList() (res []*Friend, err error) {
	panic("implement me") // 这个使用 http 的方法去请求
}

func (w *websocketBot) RefreshMemberList() error {
	panic("implement me") // 这个使用 http 的方法去请求
}

func (w *websocketBot) SendDestroy() error {
	panic("implement me")
}

func (w *websocketBot) GetPersonalInfo() error {
	op := "op:personal info"
	msg := NewWrapperMsg(PersonalInfo, op, "ROOT")
	return w.Send(msg)
}
