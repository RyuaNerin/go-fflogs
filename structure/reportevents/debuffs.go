package reportevents

import "github.com/RyuaNerin/go-fflogs/structure"

type Debuffs struct {
	structure.BaseResponse
	Events            []DebuffsEvent `json:"events"`
	Count             int64          `json:"count"`
	NextPageTimestamp int64          `json:"nextPageTimestamp"`
}

type DebuffsEvent struct {
	Timestamp        int64          `json:"timestamp"`
	Type             string         `json:"type"`
	Source           *DebuffsSource `json:"source,omitempty"`
	SourceInstance   *int64         `json:"sourceInstance,omitempty"`
	SourceIsFriendly bool           `json:"sourceIsFriendly"`
	TargetID         int64          `json:"targetID"`
	TargetIsFriendly bool           `json:"targetIsFriendly"`
	Ability          DebuffsAbility `json:"ability"`
	Fight            int64          `json:"fight"`
	SourceID         *int64         `json:"sourceID,omitempty"`
	Stack            *int64         `json:"stack,omitempty"`
}

type DebuffsAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type DebuffsSource struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	GUID int64  `json:"guid"`
	Type string `json:"type"`
	Icon string `json:"icon"`
}
