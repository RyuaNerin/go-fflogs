package reportevents

import "fflogs/structure"

type Summary struct {
	structure.BaseResponse
	Events            []SummaryEvent `json:"events"`
	Count             int64          `json:"count"`
	NextPageTimestamp int64          `json:"nextPageTimestamp"`
}

type SummaryEvent struct {
	Timestamp           int64             `json:"timestamp"`
	Type                string            `json:"type"`
	SourceID            *int64            `json:"sourceID,omitempty"`
	SourceIsFriendly    *bool             `json:"sourceIsFriendly,omitempty"`
	TargetID            *int64            `json:"targetID,omitempty"`
	TargetIsFriendly    *bool             `json:"targetIsFriendly,omitempty"`
	Ability             *SummaryAbility   `json:"ability,omitempty"`
	Fight               int64             `json:"fight"`
	HitType             *int64            `json:"hitType,omitempty"`
	Amount              *int64            `json:"amount,omitempty"`
	Multiplier          *float64          `json:"multiplier,omitempty"`
	PacketID            *int64            `json:"packetID,omitempty"`
	SourceResources     *SummaryResources `json:"sourceResources,omitempty"`
	TargetResources     *SummaryResources `json:"targetResources,omitempty"`
	DirectHit           *bool             `json:"directHit,omitempty"`
	Stack               *int64            `json:"stack,omitempty"`
	Melee               *bool             `json:"melee,omitempty"`
	Value               *int64            `json:"value,omitempty"`
	Bars                *int64            `json:"bars,omitempty"`
	Absorbed            *int64            `json:"absorbed,omitempty"`
	Absorb              *int64            `json:"absorb,omitempty"`
	Target              *SummarySource    `json:"target,omitempty"`
	Unpaired            *bool             `json:"unpaired,omitempty"`
	Overheal            *int64            `json:"overheal,omitempty"`
	Tick                *bool             `json:"tick,omitempty"`
	FinalizedAmount     *float64          `json:"finalizedAmount,omitempty"`
	Simulated           *bool             `json:"simulated,omitempty"`
	ExpectedAmount      *int64            `json:"expectedAmount,omitempty"`
	ExpectedCritRate    *int64            `json:"expectedCritRate,omitempty"`
	ActorPotencyRatio   *float64          `json:"actorPotencyRatio,omitempty"`
	GuessAmount         *float64          `json:"guessAmount,omitempty"`
	DirectHitPercentage *float64          `json:"directHitPercentage,omitempty"`
	Source              *SummarySource    `json:"source,omitempty"`
	TargetInstance      *int64            `json:"targetInstance,omitempty"`
}

type SummaryAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type SummarySource struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	GUID int64  `json:"guid"`
	Type string `json:"type"`
	Icon string `json:"icon"`
}

type SummaryResources struct {
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
