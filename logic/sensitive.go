package logic

import (
    "strings"

    "github.com/hd2yao/chat-room/global"
)

func FilterSensitive(content string) string {
    for _, word := range global.SensitiveWords {
        content = strings.ReplaceAll(content, word, "**")
    }
    return content
}
