package structure

type EncounterRankings struct {
	BaseResponse
	Page         int64              `json:"page"`
	HasMorePages bool               `json:"hasMorePages"`
	Count        int64              `json:"count"`
	Rankings     []EncounterRanking `json:"rankings"`
}
type EncounterRanking struct {
	ServerID     int64   `json:"serverID"`
	ServerName   string  `json:"serverName"`
	RegionName   string  `json:"regionName"`
	Duration     int64   `json:"duration"`
	StartTime    int64   `json:"startTime"`
	ReportStart  int64   `json:"reportStart"`
	DamageTaken  int64   `json:"damageTaken"`
	Deaths       int64   `json:"deaths"`
	FightID      int64   `json:"fightID"`
	ReportID     string  `json:"reportID"`
	Exploit      int64   `json:"exploit"`
	Tanks        int64   `json:"tanks"`
	Healers      int64   `json:"healers"`
	Melee        int64   `json:"melee"`
	Ranged       int64   `json:"ranged"`
	Casters      int64   `json:"casters"`
	GuildFaction int64   `json:"guildFaction"`
	GuildName    string  `json:"guildName"`
	GuildID      int64   `json:"guildID"`
	Bracket      float64 `json:"bracket"`
}
