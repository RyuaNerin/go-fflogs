package reporttables

type Deaths struct {
	Entries []DeathsEntry `json:"entries"`
}

type DeathsEntry struct {
	Name        string             `json:"name"`
	ID          int64              `json:"id"`
	GUID        int64              `json:"guid"`
	Type        string             `json:"type"`
	Icon        string             `json:"icon"`
	Timestamp   int64              `json:"timestamp"`
	Damage      DeathsDamageClass  `json:"damage"`
	Healing     DeathsHealing      `json:"healing"`
	Fight       int64              `json:"fight"`
	DeathWindow int64              `json:"deathWindow"`
	Overkill    int64              `json:"overkill"`
	Events      []DeathsEvent      `json:"events"`
	KillingBlow *DeathsKillingBlow `json:"killingBlow,omitempty"`
}

type DeathsDamageClass struct {
	Total             int64                 `json:"total"`
	ActiveTime        int64                 `json:"activeTime"`
	ActiveTimeReduced int64                 `json:"activeTimeReduced"`
	Abilities         []DeathsAbility       `json:"abilities"`
	DamageAbilities   []interface{}         `json:"damageAbilities"`
	Sources           []DeathsSourceElement `json:"sources"`
	Blocked           *int64                `json:"blocked,omitempty"`
}

type DeathsAbility struct {
	Name    string `json:"name"`
	Total   int64  `json:"total"`
	Type    int64  `json:"type"`
	PetName string `json:"petName"`
}

type DeathsSourceElement struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Type  string `json:"type"`
}

type DeathsEvent struct {
	Timestamp        int64              `json:"timestamp"`
	Type             string             `json:"type"`
	SourceID         *int64             `json:"sourceID,omitempty"`
	SourceIsFriendly bool               `json:"sourceIsFriendly"`
	TargetID         int64              `json:"targetID"`
	TargetIsFriendly bool               `json:"targetIsFriendly"`
	Ability          DeathsKillingBlow  `json:"ability"`
	Fight            int64              `json:"fight"`
	HitType          int64              `json:"hitType"`
	Amount           int64              `json:"amount"`
	Overkill         *int64             `json:"overkill,omitempty"`
	PacketID         int64              `json:"packetID"`
	Multiplier       int64              `json:"multiplier"`
	Source           *DeathsEventSource `json:"source,omitempty"`
	SourceInstance   *int64             `json:"sourceInstance,omitempty"`
	Blocked          *int64             `json:"blocked,omitempty"`
}

type DeathsKillingBlow struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
}

type DeathsEventSource struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	GUID int64  `json:"guid"`
	Type string `json:"type"`
	Icon string `json:"icon"`
}

type DeathsHealing struct {
	Total             int64                 `json:"total"`
	ActiveTime        int64                 `json:"activeTime"`
	ActiveTimeReduced int64                 `json:"activeTimeReduced"`
	Abilities         []DeathsAbility       `json:"abilities"`
	DamageAbilities   []DeathsAbility       `json:"damageAbilities"`
	Sources           []DeathsSourceElement `json:"sources"`
	Overheal          *int64                `json:"overheal,omitempty"`
}
