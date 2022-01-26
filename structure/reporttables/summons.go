package reporttables

// Todo
type Summons struct {
	Entries     []interface{} `json:"entries"`
	TotalTime   int64         `json:"totalTime"`
	LogVersion  int64         `json:"logVersion"`
	GameVersion int64         `json:"gameVersion"`
}
