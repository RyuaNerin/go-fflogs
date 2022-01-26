package reporttables

import "github.com/RyuaNerin/go-fflogs/structure"

type DamageTaken struct {
	structure.BaseResponse

	Entries     []DamageTakenEntry `json:"entries"`
	TotalTime   int64              `json:"totalTime"`
	LogVersion  int64              `json:"logVersion"`
	GameVersion int64              `json:"gameVersion"`
}

type DamageTakenEntry struct {
	Name              string               `json:"name"`
	ID                int64                `json:"id"`
	GUID              int64                `json:"guid"`
	Type              string               `json:"type"`
	Icon              string               `json:"icon"`
	Total             int64                `json:"total"`
	ActiveTime        int64                `json:"activeTime"`
	ActiveTimeReduced int64                `json:"activeTimeReduced"`
	Abilities         []DamageTakenAbility `json:"abilities"`
	DamageAbilities   []interface{}        `json:"damageAbilities"`
	Sources           []DamageTakenSource  `json:"sources"`
	Blocked           *int64               `json:"blocked,omitempty"`
}

type DamageTakenAbility struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Type  int64  `json:"type"`
}

type DamageTakenSource struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Type  string `json:"type"`
}
