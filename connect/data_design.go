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
	MaxHP              int
	Attack             int
	MoveSpeed          int
	BulletSpeed        int
	MaxBulletCount     int
	ShotInterval       int
	Status             int
	Direction          int
	CurrentHP          int
	CurrentBulletCount int
	NextAllowShotTime  int
	X                  int
	Y                  int
}

type otherPanzer struct {
	Name      string
	MaxHP     int
	Status    int
	Direction int
	X         int
	Y         int
}

type ResponseData struct {
	OK         bool          `json:"ok"`
	Message    string        `json:"message"`
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
		MaxHP:              selfPz.MaxHP,
		Attack:             selfPz.Attack,
		MoveSpeed:          selfPz.MoveSpeed,
		BulletSpeed:        selfPz.BulletSpeed,
		MaxBulletCount:     selfPz.MaxBulletCount,
		ShotInterval:       selfPz.ShotInterval,
		Status:             selfPz.Status,
		Direction:          selfPz.Direction,
		CurrentHP:          selfPz.CurrentHP,
		CurrentBulletCount: selfPz.CurrentBulletCount,
		NextAllowShotTime:  selfPz.NextAllowShotTime,
		X:                  selfPz.X,
		Y:                  selfPz.Y,
	}
	otherPzMiniInfoList := []otherPanzer{}
	for _, p := range otherPzList {
		otherPzMiniInfoList = append(otherPzMiniInfoList, otherPanzer{
			Name:      p.Name,
			MaxHP:     p.MaxHP,
			Status:    p.Status,
			Direction: p.Direction,
			X:         p.X,
			Y:         p.Y,
		})
	}
	respData := &ResponseData{
		OK:         true,
		Message:    "OK",
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
