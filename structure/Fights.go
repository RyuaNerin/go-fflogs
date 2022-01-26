package structure

type Fights struct {
	BaseResponse
	Fights             []FightsFight       `json:"fights"`
	Lang               string              `json:"lang"`
	Friendlies         []FightsFriendly    `json:"friendlies"`
	Enemies            []FightsEnemy       `json:"enemies"`
	FriendlyPets       []FightsFriendlyPet `json:"friendlyPets"`
	EnemyPets          []interface{}       `json:"enemyPets"`
	Phases             []FightsPhase       `json:"phases"`
	LogVersion         int64               `json:"logVersion"`
	GameVersion        int64               `json:"gameVersion"`
	Title              string              `json:"title"`
	Owner              string              `json:"owner"`
	Start              int64               `json:"start"`
	End                int64               `json:"end"`
	Zone               int64               `json:"zone"`
	ExportedCharacters []interface{}       `json:"exportedCharacters"`
}

type FightsEnemy struct {
	Name   string             `json:"name"`
	ID     int64              `json:"id"`
	GUID   int64              `json:"guid"`
	Type   string             `json:"type"`
	Icon   string             `json:"icon"`
	Fights []FightsEnemyFight `json:"fights"`
}

type FightsEnemyFight struct {
	ID        int64  `json:"id"`
	Instances int64  `json:"instances"`
	Groups    *int64 `json:"groups,omitempty"`
}

type FightsFight struct {
	ID                            int64       `json:"id"`
	Boss                          int64       `json:"boss"`
	StartTime                     int64       `json:"start_time"`
	EndTime                       int64       `json:"end_time"`
	Name                          string      `json:"name"`
	ZoneID                        int64       `json:"zoneID"`
	ZoneName                      string      `json:"zoneName"`
	Size                          *int64      `json:"size,omitempty"`
	Difficulty                    *int64      `json:"difficulty,omitempty"`
	Kill                          *bool       `json:"kill,omitempty"`
	Partial                       *int64      `json:"partial,omitempty"`
	StandardComposition           *bool       `json:"standardComposition,omitempty"`
	HasEcho                       *bool       `json:"hasEcho,omitempty"`
	BossPercentage                *int64      `json:"bossPercentage,omitempty"`
	FightPercentage               *int64      `json:"fightPercentage,omitempty"`
	LastPhaseForPercentageDisplay *int64      `json:"lastPhaseForPercentageDisplay,omitempty"`
	Maps                          []FightsMap `json:"maps,omitempty"`
	OriginalBoss                  *int64      `json:"originalBoss,omitempty"`
}

type FightsMap struct {
	MapID   int64  `json:"mapID"`
	MapName string `json:"mapName"`
	MapFile string `json:"mapFile"`
}

type FightsFriendly struct {
	Name   string                `json:"name"`
	ID     int64                 `json:"id"`
	GUID   int64                 `json:"guid"`
	Type   string                `json:"type"`
	Server string                `json:"server,omitempty"`
	Icon   string                `json:"icon"`
	Fights []FightsFriendlyFight `json:"fights"`
}

type FightsFriendlyFight struct {
	ID int64 `json:"id"`
}

type FightsFriendlyPet struct {
	Name     string                   `json:"name"`
	ID       int64                    `json:"id"`
	GUID     int64                    `json:"guid"`
	Type     string                   `json:"type"`
	Icon     string                   `json:"icon"`
	PetOwner int64                    `json:"petOwner"`
	Fights   []FightsFriendlyPetFight `json:"fights"`
}

type FightsFriendlyPetFight struct {
	ID        int64 `json:"id"`
	Instances int64 `json:"instances"`
}

type FightsPhase struct {
	Boss   int64    `json:"boss"`
	Phases []string `json:"phases"`
}
