package connect

import "encoding/json"

type ReceivedData struct {
	Action    string `json:"action"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Direction string `json:"direction"`
	Move      int    `json:"move"`
}

func GetReceivedData(jsonData string) *ReceivedData {
	var receivedData ReceivedData
	json.Unmarshal([]byte(jsonData), &receivedData)
	return &receivedData
}
