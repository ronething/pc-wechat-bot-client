package wechat_bot

//var wsBot *websocketBot
//var httpBot *httpBot
//
//func TestMain(m *testing.M) {
//	url := "ws://127.0.0.1:5555"
//	wsBot = NewWebsocketBot(url)
//	wsBot.Connect()
//	prefix := "http://127.0.0.1:5555"
//	httpBot = NewHTTPBot(prefix)
//	defer wsBot.Close()
//	m.Run()
//}
//
//func TestWebsocketBot_SendTxtMsg(t *testing.T) {
//	wxid := "wxid_xxxxx"
//	err := wsBot.SendTxtMsg("今天是个好日子", wxid)
//	t.Logf("err: %v\n", err)
//}
