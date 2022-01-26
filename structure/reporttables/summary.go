package reporttables

import "github.com/RyuaNerin/go-fflogs/structure"

type Summary struct {
	structure.BaseResponse

	TotalTime      int64                `json:"totalTime"`
	ItemLevel      int64                `json:"itemLevel"`
	LogVersion     int64                `json:"logVersion"`
	GameVersion    int64                `json:"gameVersion"`
	DamageDowntime int64                `json:"damageDowntime"`
	Composition    []SummaryComposition `json:"composition"`
	DamageDone     []SummaryDone        `json:"damageDone"`
	HealingDone    []SummaryDone        `json:"healingDone"`
	DamageTaken    []SummaryDamageTaken `json:"damageTaken"`
	DeathEvents    []SummaryDeathEvent  `json:"deathEvents"`
}

type SummaryComposition struct {
	Name  string        `json:"name"`
	ID    int64         `json:"id"`
	GUID  int64         `json:"guid"`
	Type  string        `json:"type"`
	Specs []SummarySpec `json:"specs"`
}

type SummarySpec struct {
	Spec string `json:"spec"`
	Role string `json:"role"`
}

type SummaryDone struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	GUID  int64  `json:"guid"`
	Type  string `json:"type"`
	Icon  string `json:"icon"`
	Total int64  `json:"total"`
}

type SummaryDamageTaken struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
	Total       *int64 `json:"total,omitempty"`
}

type SummaryDeathEvent struct {
	Name      string              `json:"name"`
	ID        int64               `json:"id"`
	GUID      int64               `json:"guid"`
	Type      string              `json:"type"`
	Icon      string              `json:"icon"`
	DeathTime int64               `json:"deathTime"`
	Ability   *SummaryDamageTaken `json:"ability,omitempty"`
}
