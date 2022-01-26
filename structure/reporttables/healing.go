package reporttables

import "github.com/RyuaNerin/go-fflogs/structure"

type Healing struct {
	structure.BaseResponse

	Entries     []HealingEntry `json:"entries"`
	TotalTime   int64          `json:"totalTime"`
	LogVersion  int64          `json:"logVersion"`
	GameVersion int64          `json:"gameVersion"`
}

type HealingEntry struct {
	Name              string                 `json:"name"`
	ID                int64                  `json:"id"`
	GUID              int64                  `json:"guid"`
	Type              string                 `json:"type"`
	Icon              string                 `json:"icon"`
	Total             int64                  `json:"total"`
	ActiveTime        int64                  `json:"activeTime"`
	ActiveTimeReduced int64                  `json:"activeTimeReduced"`
	Abilities         []HealingAbility       `json:"abilities"`
	DamageAbilities   []HealingDamageAbility `json:"damageAbilities"`
	Targets           []HealingTarget        `json:"targets"`
	Overheal          *int64                 `json:"overheal,omitempty"`
	TotalReduced      *int64                 `json:"totalReduced,omitempty"`
}

type HealingAbility struct {
	Name         string  `json:"name"`
	Total        int64   `json:"total"`
	Type         int64   `json:"type"`
	TotalReduced *int64  `json:"totalReduced,omitempty"`
	PetName      *string `json:"petName,omitempty"`
}

type HealingDamageAbility struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Type  int64  `json:"type"`
}

type HealingTarget struct {
	Name         string `json:"name"`
	Total        int64  `json:"total"`
	Type         string `json:"type"`
	TotalReduced *int64 `json:"totalReduced,omitempty"`
}
