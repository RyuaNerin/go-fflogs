package reportevents

import "fflogs/structure"

type DamageTaken struct {
	structure.BaseResponse
	Events            []DamageTakenEvent   `json:"events"`
	Count             int64                `json:"count"`
	NextPageTimestamp int64                `json:"nextPageTimestamp"`
	AuraAbilities     []DamageTakenAbility `json:"auraAbilities"`
}

type DamageTakenAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type DamageTakenEvent struct {
	Timestamp        int64                 `json:"timestamp"`
	Type             string                `json:"type"`
	SourceID         *int64                `json:"sourceID,omitempty"`
	SourceIsFriendly bool                  `json:"sourceIsFriendly"`
	TargetID         int64                 `json:"targetID"`
	TargetIsFriendly bool                  `json:"targetIsFriendly"`
	Ability          DamageTakenAbility    `json:"ability"`
	Fight            int64                 `json:"fight"`
	Buffs            *string               `json:"buffs,omitempty"`
	HitType          int64                 `json:"hitType"`
	Amount           int64                 `json:"amount"`
	Multiplier       *int64                `json:"multiplier,omitempty"`
	PacketID         *int64                `json:"packetID,omitempty"`
	SourceResources  *DamageTakenResources `json:"sourceResources,omitempty"`
	TargetResources  DamageTakenResources  `json:"targetResources"`
	Source           *DamageTakenSource    `json:"source,omitempty"`
	SourceInstance   *int64                `json:"sourceInstance,omitempty"`
	Blocked          *int64                `json:"blocked,omitempty"`
	Overkill         *int64                `json:"overkill,omitempty"`
}

type DamageTakenSource struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	GUID int64  `json:"guid"`
	Type string `json:"type"`
	Icon string `json:"icon"`
}

type DamageTakenResources struct {
	HitPoints    int64  `json:"hitPoints"`
	MaxHitPoints int64  `json:"maxHitPoints"`
	Mp           int64  `json:"mp"`
	MaxMP        int64  `json:"maxMP"`
	Tp           int64  `json:"tp"`
	MaxTP        int64  `json:"maxTP"`
	X            int64  `json:"x"`
	Y            int64  `json:"y"`
	Facing       int64  `json:"facing"`
	Absorb       *int64 `json:"absorb,omitempty"`
}
