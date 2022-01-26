package reportevents

import "fflogs/structure"

type Deaths struct {
	structure.BaseResponse
	Events            []DeathsEvent `json:"events"`
	Count             int64         `json:"count"`
	NextPageTimestamp int64         `json:"nextPageTimestamp"`
}

type DeathsEvent struct {
	Timestamp        int64          `json:"timestamp"`
	Type             string         `json:"type"`
	SourceID         *int64         `json:"sourceID,omitempty"`
	SourceIsFriendly bool           `json:"sourceIsFriendly"`
	TargetID         *int64         `json:"targetID,omitempty"`
	TargetIsFriendly bool           `json:"targetIsFriendly"`
	Ability          DeathsAbility  `json:"ability"`
	Fight            int64          `json:"fight"`
	KillerID         *int64         `json:"killerID,omitempty"`
	KillingAbility   *DeathsAbility `json:"killingAbility,omitempty"`
	Source           *DeathsSource  `json:"source,omitempty"`
	SourceInstance   *int64         `json:"sourceInstance,omitempty"`
	Killer           *DeathsSource  `json:"killer,omitempty"`
	KillerInstance   *int64         `json:"killerInstance,omitempty"`
	Target           *DeathsSource  `json:"target,omitempty"`
}

type DeathsAbility struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type DeathsSource struct {
	Name   string  `json:"name"`
	ID     int64   `json:"id"`
	GUID   int64   `json:"guid"`
	Type   string  `json:"type"`
	Icon   string  `json:"icon"`
	Server *string `json:"server,omitempty"`
}
