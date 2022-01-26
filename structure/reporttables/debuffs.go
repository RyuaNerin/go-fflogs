package reporttables

import "github.com/RyuaNerin/go-fflogs/structure"

type Debuffs struct {
	structure.BaseResponse

	Auras      []DebuffsAura `json:"auras"`
	UseTargets bool          `json:"useTargets"`
	TotalTime  int64         `json:"totalTime"`
	StartTime  int64         `json:"startTime"`
	EndTime    int64         `json:"endTime"`
	LogVersion int64         `json:"logVersion"`
}

type DebuffsAura struct {
	Name        string        `json:"name"`
	GUID        int64         `json:"guid"`
	Type        int64         `json:"type"`
	AbilityIcon string        `json:"abilityIcon"`
	TotalUptime int64         `json:"totalUptime"`
	TotalUses   int64         `json:"totalUses"`
	Bands       []DebuffsBand `json:"bands"`
}

type DebuffsBand struct {
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}
