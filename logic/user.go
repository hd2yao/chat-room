package logic

import (
    "time"

    "nhooyr.io/websocket"
)

var globalUID uint32 = 0

type User struct {
    UID            int           `json:"uid"`
    NickName       string        `json:"nickname"`
    EnterAt        time.Time     `json:"enter_at"`
    Addr           string        `json:"addr"`
    MessageChannel chan *Message `json:"-"`
    Token          string        `json:"token"`

    conn *websocket.Conn

    isNew bool
}
