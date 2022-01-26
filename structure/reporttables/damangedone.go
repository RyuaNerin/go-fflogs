package reporttables

type DamageDone struct {
	Entries     []DamageDoneEntry `json:"entries"`
	TotalTime   int64             `json:"totalTime"`
	Downtime    int64             `json:"downtime"`
	LogVersion  int64             `json:"logVersion"`
	GameVersion int64             `json:"gameVersion"`
}

type DamageDoneEntry struct {
	Name              string              `json:"name"`
	ID                int64               `json:"id"`
	GUID              int64               `json:"guid"`
	Type              string              `json:"type"`
	Icon              string              `json:"icon"`
	Total             int64               `json:"total"`
	TotalReduced      *int64              `json:"totalReduced,omitempty"`
	ActiveTime        int64               `json:"activeTime"`
	ActiveTimeReduced *int64              `json:"activeTimeReduced,omitempty"`
	Abilities         []DamageDoneAbility `json:"abilities,omitempty"`
	DamageAbilities   []interface{}       `json:"damageAbilities,omitempty"`
	Targets           []DamageDoneTarget  `json:"targets,omitempty"`
	TotalRDPS         float64             `json:"totalRDPS"`
	TotalRDPSTaken    float64             `json:"totalRDPSTaken"`
	TotalRDPSGiven    float64             `json:"totalRDPSGiven"`
	TotalADPS         float64             `json:"totalADPS"`
	TotalNDPS         float64             `json:"totalNDPS"`
	Given             []DamageDoneGiven   `json:"given"`
	Taken             []DamageDoneGiven   `json:"taken"`
	Pets              []DamageDoneEntry   `json:"pets,omitempty"`
}

type DamageDoneAbility struct {
	Name         string  `json:"name"`
	Total        int64   `json:"total"`
	Type         int64   `json:"type"`
	TotalReduced *int64  `json:"totalReduced,omitempty"`
	PetName      *string `json:"petName,omitempty"`
}

type DamageDoneGiven struct {
	GUID        int64   `json:"guid"`
	Name        string  `json:"name"`
	Total       float64 `json:"total"`
	Type        string  `json:"type"`
	AbilityIcon string  `json:"abilityIcon"`
}

type DamageDoneTarget struct {
	Name         string `json:"name"`
	Total        int64  `json:"total"`
	TotalReduced *int64 `json:"totalReduced,omitempty"`
	Type         string `json:"type"`
}
