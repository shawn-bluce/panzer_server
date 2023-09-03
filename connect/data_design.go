package connect

import (
	"encoding/json"
	"panzer_server/game"
)

type RequestData struct {
	Action    string `json:"action"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Direction string `json:"direction"`
	Move      int    `json:"move"`
}

type selfPanzer struct {
	maxHP              int
	attack             int
	moveSpeed          int
	bulletSpeed        int
	maxBulletCount     int
	shotInterval       int
	status             int
	direction          int
	currentHP          int
	currentBulletCount int
	nextAllowShotTime  int
	x                  int
	y                  int
}

type otherPanzer struct {
	name      string
	maxHP     int
	status    int
	direction int
	x         int
	y         int
}

type ResponseData struct {
	UUID       string        `json:"uuid"`
	Name       string        `json:"name"`
	Direction  int           `json:"direction"`
	Panzer     selfPanzer    `json:"panzer"`
	PanzerList []otherPanzer `json:"panzer_list"`
}

func GetReceivedData(jsonData string) *RequestData {
	var receivedData RequestData
	json.Unmarshal([]byte(jsonData), &receivedData)
	return &receivedData
}

func GenerateResponseData(playerUUID string) []byte {
	selfPz := game.GetPanzerWithUUID(playerUUID)
	otherPzList := game.GetPanzerListExcludedUUID(playerUUID)
	subDataSelfPanzer := selfPanzer{
		maxHP:              selfPz.MaxHP,
		attack:             selfPz.Attack,
		moveSpeed:          selfPz.MoveSpeed,
		bulletSpeed:        selfPz.BulletSpeed,
		maxBulletCount:     selfPz.MaxBulletCount,
		shotInterval:       selfPz.ShotInterval,
		status:             selfPz.Status,
		direction:          selfPz.Direction,
		currentHP:          selfPz.CurrentHP,
		currentBulletCount: selfPz.CurrentBulletCount,
		nextAllowShotTime:  selfPz.NextAllowShotTime,
		x:                  selfPz.X,
		y:                  selfPz.Y,
	}
	otherPzMiniInfoList := []otherPanzer{}
	for _, p := range otherPzList {
		otherPzMiniInfoList = append(otherPzMiniInfoList, otherPanzer{
			name:      p.Name,
			maxHP:     p.MaxHP,
			status:    p.Status,
			direction: p.Direction,
			x:         p.X,
			y:         p.Y,
		})
	}
	respData := &ResponseData{
		UUID:       playerUUID,
		Name:       selfPz.Name,
		Direction:  selfPz.Direction,
		Panzer:     subDataSelfPanzer,
		PanzerList: otherPzMiniInfoList,
	}

	// trans respData to bytes
	bytesData, _ := json.Marshal(respData)
	return bytesData
}
