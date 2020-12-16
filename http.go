package wechat_bot

import (
	"fmt"
	"log"

	"github.com/imroc/req"
)

type httpBot struct {
	client *req.Req
	// TODO: config
	prefix string
}

func NewHTTPBot(prefix string) *httpBot {
	return &httpBot{client: req.New(), prefix: prefix}
}

func (b *httpBot) SendAtMsg(content, roomId, atWxId, nick string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendatmsg", req.QueryParam{
		"id":      getId(),
		"type":    AtMsg,
		"roomid":  roomId,
		"wxid":    atWxId,
		"content": content,
		"nick":    nick,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v\n", resp.String())
	return nil
}

func (b *httpBot) SendTxtMsg(content, wxId string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendtxtmsg", req.QueryParam{
		"id":      getId(),
		"type":    TxtMsg,
		"wxid":    wxId,
		"content": content,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v, status code is %d\n", resp.String(), resp.Response().StatusCode)
	return nil
}

func (b *httpBot) SendPic(wxId, filePath string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendpic", req.QueryParam{
		"id":   getId(),
		"type": PicMsg,
		"wxid": wxId,
		"path": filePath,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v\n", resp.String())
	return nil
}

func (b *httpBot) SendAttach(wxId, filePath string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendattatch", req.QueryParam{
		"id":       getId(),
		"wxid":     wxId,
		"filepath": filePath,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v\n", resp.String())
	return nil
}

func (b *httpBot) GetMemberNick(roomId string) (res []*Member, err error) {
	resp, err := b.client.Get(b.prefix+"/api/getmembernick", req.QueryParam{
		"id":     getId(),
		"type":   ChatroomMemberNick,
		"roomid": roomId,
	})
	if err != nil {
		return nil, err
	}
	var getMemberList GetMemberList
	if err = resp.ToJSON(&getMemberList); err != nil {
		return nil, err
	}
	//fmt.Printf("resp is %v\n", resp)
	return getMemberList.Content, nil
}

func (b *httpBot) GetMemberId() (res interface{}, err error) {
	resp, err := b.client.Get(b.prefix + "/api/getmemberid")
	if err != nil {
		return nil, err
	}
	fmt.Printf("resp is %v\n", resp)
	return resp, nil
}

func (b *httpBot) GetContactList() (res []*Friend, err error) {
	resp, err := b.client.Get(b.prefix + "/api/getcontactlist")
	if err != nil {
		return nil, err
	}
	//fmt.Printf("resp is %v\n", resp.String())
	var getContactList GetContactList
	if err = resp.ToJSON(&getContactList); err != nil {
		return nil, err
	}
	return getContactList.Content, nil
}

func (b *httpBot) RefreshMemberList() (err error) {
	resp, err := b.client.Get(b.prefix+"/api/refresh_chatroom", req.QueryParam{
		"id": getId(),
	})
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp)
	return nil
}

func (b *httpBot) SendDestroy() error {
	resp, err := b.client.Get(b.prefix+"/api/destroy", req.QueryParam{
		"id": getId(),
	})
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp)
	return nil
}
