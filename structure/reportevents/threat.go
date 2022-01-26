package reportevents

import "github.com/RyuaNerin/go-fflogs/structure"

type Threat struct {
	structure.BaseResponse
	Events            []ThreatEvent `json:"events"`
	Count             int64         `json:"count"`
	NextPageTimestamp int64         `json:"nextPageTimestamp"`
}

type ThreatEvent struct {
	Timestamp        int64         `json:"timestamp"`
	Type             string        `json:"type"`
	SourceID         int64         `json:"sourceID"`
	SourceIsFriendly bool          `json:"sourceIsFriendly"`
	TargetID         int64         `json:"targetID"`
	TargetIsFriendly bool          `json:"targetIsFriendly"`
	Ability          ThreatAbility `json:"ability"`
	Fight            int64         `json:"fight"`
	Melee            *bool         `json:"melee,omitempty"`
	TargetInstance   *int64        `json:"targetInstance,omitempty"`
}

type ThreatAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}
