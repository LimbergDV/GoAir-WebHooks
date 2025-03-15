package infrastructure

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func SendMessageToDiscord(msg string) int {
    const URL = "https://discord.com/api/webhooks/1344047196303261696/gd1Kp2-XW87fXHs3zOAA7noi1-_6Lyf3N8yykCm48YJSWc6opXS-Q7gwGgQqxHvJrFHG"

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