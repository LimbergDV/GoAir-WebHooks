package infrastructure

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func SendMessageToDiscord(msg string) int {
    const URL = "https://discord.com/api/webhooks/1350530363721650228/D86TVRMcxJONK7peGp7sLeiUiGHBBlQWAVr_NjA9VONysEUdmN2qXKOhTom3QLjB21kj"

    payload := map[string]string{
        "content": msg,
    }

    jsonPayload, _ := json.Marshal(payload)

    _, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonPayload))

    if err != nil {
        return 500
    }

    return 200
}