package reporttables

import "github.com/RyuaNerin/go-fflogs/structure"

// Todo
type Summons struct {
	structure.BaseResponse

	Entries     []interface{} `json:"entries"`
	TotalTime   int64         `json:"totalTime"`
	LogVersion  int64         `json:"logVersion"`
	GameVersion int64         `json:"gameVersion"`
}
