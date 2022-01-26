package structure

type CharacterRanking struct {
	BaseResponse
	EncounterID    int     `json:"encounterID"`
	EncounterName  string  `json:"encounterName"`
	Class          string  `json:"class"`
	Spec           string  `json:"spec"`
	Rank           int64   `json:"rank"`
	OutOf          int64   `json:"outOf"`
	Duration       int64   `json:"duration"`
	StartTime      int64   `json:"startTime"`
	ReportID       string  `json:"reportID"`
	FightID        int     `json:"fightID"`
	Difficulty     int     `json:"difficulty"`
	CharacterID    int64   `json:"characterID"`
	CharacterName  string  `json:"characterName"`
	Server         string  `json:"server"`
	Percentile     float64 `json:"percentile"`
	IlvlKeyOrPatch float64 `json:"ilvlKeyOrPatch"`
	Total          float64 `json:"total"`
	Estimated      *bool   `json:"estimated,omitempty"`
}
