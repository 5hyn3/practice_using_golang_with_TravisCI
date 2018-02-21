package nem

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type Heartbeat struct {
	Code    int    `json:"code"`
	Type    int    `json:"type"`
	Message string `json:"message"`
}

func GetHeartbeat() (heartbeat Heartbeat,err error){
    resp, err := http.Get("http://alice2.nem.ninja:7890/heartbeat")

    if err != nil {
        return heartbeat,err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return heartbeat,err
    }

    if err := json.Unmarshal(body, &heartbeat); err != nil {
        return heartbeat,err
	}
	return heartbeat,nil
}