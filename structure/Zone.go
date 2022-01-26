package structure

type Zone struct {
	BaseResponse

	ID         int              `json:"id"`         // A unique identifier representing this specific zone.,
	Name       string           `json:"name"`       // The English name of the raid zone.,
	Frozen     bool             `json:"frozen"`     // Whether or not the rankings and statistics for the zone are frozen. If set, then the results for this zone are never going to change, and you don't have to fetch them again.,
	Encounters []ZoneEncounters `json:"encounters"` // The encounters for this zone.,
	Brackets   ZoneBracket      `json:"brackets"`   // (optional) The brackets for this zone. Rankings and statistics are collected for each bracket.
	Partitions []ZonePartition  `json:"partitions"`
}

type ZoneEncounters struct {
	Min    int64   `json:"min"`
	Max    float64 `json:"max"`
	Bucket float64 `json:"bucket"`
	Type   string  `json:"type"`
}

type ZoneBracket struct {
	ID   int    `json:"id"`   // A unique identifier representing this specific rankings bracket.,
	Name string `json:"name"` // An explanation of what the bracket contains.
}

type ZonePartition struct {
	Name         string  `json:"name"`
	Compact      string  `json:"compact"`
	Area         *int64  `json:"area,omitempty"`
	Default      *bool   `json:"default,omitempty"`
	FilteredName *string `json:"filtered_name,omitempty"`
}
