package reportevents

import "fflogs/structure"

type Casts struct {
	structure.BaseResponse
	Events            []CastsEvent `json:"events"`
	Count             int64        `json:"count"`
	NextPageTimestamp int64        `json:"nextPageTimestamp"`
}

type CastsEvent struct {
	Timestamp        int64        `json:"timestamp"`
	Type             string       `json:"type"`
	SourceID         int64        `json:"sourceID"`
	SourceIsFriendly bool         `json:"sourceIsFriendly"`
	TargetID         *int64       `json:"targetID,omitempty"`
	TargetIsFriendly bool         `json:"targetIsFriendly"`
	Ability          CastsAbility `json:"ability"`
	Fight            int64        `json:"fight"`
	Target           *CastsTarget `json:"target,omitempty"`
}

type CastsAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type CastsTarget struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	GUID int64  `json:"guid"`
	Type string `json:"type"`
	Icon string `json:"icon"`
}
