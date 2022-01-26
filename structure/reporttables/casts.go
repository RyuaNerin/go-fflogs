package reporttables

type Casts struct {
	Entries     []CastsEntry `json:"entries"`
	TotalTime   int64        `json:"totalTime"`
	LogVersion  int64        `json:"logVersion"`
	GameVersion int64        `json:"gameVersion"`
}

type CastsEntry struct {
	Name              string         `json:"name"`
	ID                int64          `json:"id"`
	GUID              int64          `json:"guid"`
	Type              string         `json:"type"`
	Icon              string         `json:"icon"`
	Total             int64          `json:"total"`
	ActiveTime        int64          `json:"activeTime"`
	ActiveTimeReduced int64          `json:"activeTimeReduced"`
	Abilities         []CastsAbility `json:"abilities"`
	DamageAbilities   []interface{}  `json:"damageAbilities"` // TODO
	Targets           []CastsTarget  `json:"targets"`
}

type CastsAbility struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Type  int64  `json:"type"`
}

type CastsTarget struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Type  string `json:"type"`
}
