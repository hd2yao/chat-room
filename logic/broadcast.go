package logic

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
    messageChannel:  make(chan *User),
}
