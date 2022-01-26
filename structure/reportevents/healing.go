package reportevents

import "github.com/RyuaNerin/go-fflogs/structure"

type Healing struct {
	structure.BaseResponse

	Events            []HealingEvent   `json:"events"`
	Count             int64            `json:"count"`
	NextPageTimestamp int64            `json:"nextPageTimestamp"`
	AuraAbilities     []HealingAbility `json:"auraAbilities"`
}

type HealingAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type HealingEvent struct {
	Timestamp           int64             `json:"timestamp"`
	Type                string            `json:"type"`
	SourceID            int64             `json:"sourceID"`
	SourceIsFriendly    bool              `json:"sourceIsFriendly"`
	TargetID            int64             `json:"targetID"`
	TargetIsFriendly    bool              `json:"targetIsFriendly"`
	Ability             HealingAbility    `json:"ability"`
	Fight               int64             `json:"fight"`
	Absorbed            *int64            `json:"absorbed,omitempty"`
	Absorb              *int64            `json:"absorb,omitempty"`
	HitType             *int64            `json:"hitType,omitempty"`
	Amount              *int64            `json:"amount,omitempty"`
	Multiplier          *float64          `json:"multiplier,omitempty"`
	PacketID            *int64            `json:"packetID,omitempty"`
	SourceResources     *HealingResources `json:"sourceResources,omitempty"`
	TargetResources     *HealingResources `json:"targetResources,omitempty"`
	Unpaired            *bool             `json:"unpaired,omitempty"`
	Overheal            *int64            `json:"overheal,omitempty"`
	Tick                *bool             `json:"tick,omitempty"`
	Buffs               string            `json:"buffs,omitempty"`
	SourceInstance      *int64            `json:"sourceInstance,omitempty"`
	FinalizedAmount     *float64          `json:"finalizedAmount,omitempty"`
	Simulated           *bool             `json:"simulated,omitempty"`
	ExpectedAmount      *int64            `json:"expectedAmount,omitempty"`
	ExpectedCritRate    *int64            `json:"expectedCritRate,omitempty"`
	ActorPotencyRatio   *float64          `json:"actorPotencyRatio,omitempty"`
	GuessAmount         *float64          `json:"guessAmount,omitempty"`
	DirectHitPercentage *int64            `json:"directHitPercentage,omitempty"`
}

type HealingResources struct {
	HitPoints    int64 `json:"hitPoints"`
	MaxHitPoints int64 `json:"maxHitPoints"`
	Mp           int64 `json:"mp"`
	MaxMP        int64 `json:"maxMP"`
	Tp           int64 `json:"tp"`
	MaxTP        int64 `json:"maxTP"`
	X            int64 `json:"x"`
	Y            int64 `json:"y"`
	Facing       int64 `json:"facing"`
	Absorb       int64 `json:"absorb"`
}
