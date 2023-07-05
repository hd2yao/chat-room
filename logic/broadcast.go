package logic

import (
    "expvar"
    "fmt"

    "github.com/hd2yao/chat-room/global"
)

func init() {
    expvar.Publish("message_queue", expvar.Func(calcMessageQueueLen))
}

func calcMessageQueueLen() interface{} {
    fmt.Println("===len=", len(Broadcaster.messageChannel))
    return len(Broadcaster.messageChannel)
}

// broadcaster 广播器
type broadcaster struct {
    // 所有聊天室用户
    users map[string]*User

    // 所有 channel 统一管理，可以避免外部乱用

    enteringChannel chan *User
    leavingChannel  chan *User
    messageChannel  chan *User

    // 判断该昵称用户是否可以进入聊天室（重复与否）：true 能，false 不能
    checkUserChannel      chan string
    checkUserCanInChannel chan bool

    // 获取用户列表
    requestUserChannel chan struct{}
    usersChannel       chan []*User
}

var Broadcaster = &broadcaster{
    users: make(map[string]*User),

    enteringChannel: make(chan *User),
    leavingChannel:  make(chan *User),
    messageChannel:  make(chan *User, global.MessageQueueLen),

    checkUserChannel:      make(chan string),
    checkUserCanInChannel: make(chan bool),

    requestUserChannel: make(chan struct{}),
    usersChannel:       make(chan []*User),
}

// Start 启动广播器
// 需要在一个新的 goroutine 中运行，因为它不会返回
func (b *broadcaster) Start() {
    for {
        select {
        case user := <-b.enteringChannel:
            // 新用户进入
            b.users[user.NickName] = user
        }
    }
}
