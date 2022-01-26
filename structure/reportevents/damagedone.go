package reportevents

import "github.com/RyuaNerin/go-fflogs/structure"

type DamageDone struct {
	structure.BaseResponse

	Events            []DamageDoneEvent   `json:"events"`
	Count             int64               `json:"count"`
	NextPageTimestamp int64               `json:"nextPageTimestamp"`
	AuraAbilities     []DamageDoneAbility `json:"auraAbilities"`
}

type DamageDoneAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type DamageDoneEvent struct {
	Timestamp           int64                `json:"timestamp"`
	Type                string               `json:"type"`
	SourceID            int64                `json:"sourceID"`
	SourceIsFriendly    bool                 `json:"sourceIsFriendly"`
	TargetID            int64                `json:"targetID"`
	TargetIsFriendly    bool                 `json:"targetIsFriendly"`
	Ability             DamageDoneAbility    `json:"ability"`
	Fight               int64                `json:"fight"`
	Buffs               *string              `json:"buffs,omitempty"`
	HitType             int64                `json:"hitType"`
	Amount              int64                `json:"amount"`
	Multiplier          *float64             `json:"multiplier,omitempty"`
	PacketID            *int64               `json:"packetID,omitempty"`
	SourceResources     *DamageDoneResources `json:"sourceResources,omitempty"`
	TargetResources     DamageDoneResources  `json:"targetResources"`
	DirectHit           *bool                `json:"directHit,omitempty"`
	Tick                *bool                `json:"tick,omitempty"`
	FinalizedAmount     *float64             `json:"finalizedAmount,omitempty"`
	Simulated           *bool                `json:"simulated,omitempty"`
	ExpectedAmount      *int64               `json:"expectedAmount,omitempty"`
	ExpectedCritRate    *int64               `json:"expectedCritRate,omitempty"`
	ActorPotencyRatio   *float64             `json:"actorPotencyRatio,omitempty"`
	GuessAmount         *float64             `json:"guessAmount,omitempty"`
	DirectHitPercentage *float64             `json:"directHitPercentage,omitempty"`
	Absorbed            *int64               `json:"absorbed,omitempty"`
	SourceInstance      *int64               `json:"sourceInstance,omitempty"`
}

type DamageDoneResources struct {
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
