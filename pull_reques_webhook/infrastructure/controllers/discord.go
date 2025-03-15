package infrastructure

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func SendMessageToDiscord(msg string) int {
    const URL = "https://discord.com/api/webhooks/1350529783964110940/mBGoldkrSyXlo3E-_3IN4VFPSqLOGGlAvK1pINixWD0An1ZiiLu1E8TiNJ_aUV5xYcyX"

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