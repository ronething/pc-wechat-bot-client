<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [wechat-bot-go](#wechat-bot-go)
- [**接收**消息类型](#%E6%8E%A5%E6%94%B6%E6%B6%88%E6%81%AF%E7%B1%BB%E5%9E%8B)
- [acknowledgement](#acknowledgement)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### wechat-bot-go

wechat bot golang client

### **接收**消息类型

- 群文本消息

```json
{
  "content":"高速老是在修路",
  "id":"20201211160304",
  "receiver":"xxxx@chatroom", 
  "sender":"wxid_xxxxxx",
  "srvid":1,
  "time":"2020-12-11 16:03:04",
  "type":1
}
```

注意到 receiver 为群聊 roomid, sender 为发送的人

- 个人文本消息

```json
{
  "content":"？",
  "id":"20201211160346",
  "receiver":"self",
  "sender":"wxid_xxxx",
  "srvid":1,
  "time":"2020-12-11 16:03:46",
  "type":1
}
```

注意到 receiver 为 self，也就是说接收个人消息的话可以通过这个 key 来进行判断

### acknowledgement

- [wechat-bot](https://github.com/cixingguangming55555/wechat-bot)