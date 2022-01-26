package reportevents

import "fflogs/structure"

type Buffs struct {
	structure.BaseResponse
	Events            []BuffsEvent `json:"events"`
	Count             int64        `json:"count"`
	NextPageTimestamp int64        `json:"nextPageTimestamp"`
}

type BuffsEvent struct {
	Timestamp        int64        `json:"timestamp"`
	Type             string       `json:"type"`
	SourceID         int64        `json:"sourceID"`
	SourceIsFriendly bool         `json:"sourceIsFriendly"`
	TargetID         int64        `json:"targetID"`
	TargetIsFriendly bool         `json:"targetIsFriendly"`
	Ability          BuffsAbility `json:"ability"`
	Fight            int64        `json:"fight"`
	Absorb           *int64       `json:"absorb,omitempty"`
	Stack            *int64       `json:"stack,omitempty"`
	Absorbed         *int64       `json:"absorbed,omitempty"`
	TargetInstance   *int64       `json:"targetInstance,omitempty"`
}

type BuffsAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}
