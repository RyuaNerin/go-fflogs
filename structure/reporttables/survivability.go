package reporttables

type Survivability struct {
	Players       []SurvivabilityActor        `json:"players"`
	Fights        []SurvivabilityFight        `json:"fights"`
	Actortotals   []SurvivabilityActor        `json:"actortotals"`
	Abilitytotals []SurvivabilityAbilityTotal `json:"abilitytotals"`
}

type SurvivabilityAbilityTotal struct {
	Name        string `json:"name"`
	GUID        int64  `json:"guid"`
	Type        int64  `json:"type"`
	AbilityIcon string `json:"abilityIcon"`
	Total       *int64 `json:"total,omitempty"`
}

type SurvivabilityEvent struct {
	Timestamp        int64                     `json:"timestamp"`
	Type             string                    `json:"type"`
	SourceID         *int64                    `json:"sourceID,omitempty"`
	SourceInstance   *int64                    `json:"sourceInstance,omitempty"`
	SourceIsFriendly bool                      `json:"sourceIsFriendly"`
	TargetID         int64                     `json:"targetID"`
	TargetIsFriendly bool                      `json:"targetIsFriendly"`
	Ability          SurvivabilityAbilityTotal `json:"ability"`
	Fight            int64                     `json:"fight"`
	HitType          int64                     `json:"hitType"`
	Amount           int64                     `json:"amount"`
	Overkill         *int64                    `json:"overkill,omitempty"`
	PacketID         int64                     `json:"packetID"`
	Multiplier       int64                     `json:"multiplier"`
	Source           *SurvivabilityActor       `json:"source,omitempty"`
	Blocked          *int64                    `json:"blocked,omitempty"`
}

type SurvivabilityDeath struct {
	Timestamp   int64                      `json:"timestamp"`
	Damage      SurvivabilityDamageClass   `json:"damage"`
	Healing     SurvivabilityHealing       `json:"healing"`
	Fight       int64                      `json:"fight"`
	DeathWindow int64                      `json:"deathWindow"`
	Overkill    int64                      `json:"overkill"`
	Events      []SurvivabilityEvent       `json:"events"`
	KillingBlow *SurvivabilityAbilityTotal `json:"killingBlow,omitempty"`
}

type SurvivabilityActortotalFight struct {
	ID     int64                `json:"id"`
	Deaths []SurvivabilityDeath `json:"deaths"`
}

type SurvivabilityActor struct {
	Name   string                         `json:"name"`
	ID     int64                          `json:"id"`
	GUID   int64                          `json:"guid"`
	Type   string                         `json:"type"`
	Icon   string                         `json:"icon"`
	Total  *int64                         `json:"total,omitempty"`
	Fights []SurvivabilityActortotalFight `json:"fights,omitempty"`
}

type SurvivabilityDamageClass struct {
	Total             int64                  `json:"total"`
	ActiveTime        int64                  `json:"activeTime"`
	ActiveTimeReduced int64                  `json:"activeTimeReduced"`
	Abilities         []SurvivabilityAbility `json:"abilities"`
	DamageAbilities   []interface{}          `json:"damageAbilities"`
	Sources           []SurvivabilitySource  `json:"sources"`
	Blocked           *int64                 `json:"blocked,omitempty"`
}

type SurvivabilityAbility struct {
	Name    string `json:"name"`
	Total   int64  `json:"total"`
	Type    int64  `json:"type"`
	PetName string `json:"petName"`
}

type SurvivabilitySource struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Type  string `json:"type"`
}

type SurvivabilityHealing struct {
	Total             int64                  `json:"total"`
	ActiveTime        int64                  `json:"activeTime"`
	ActiveTimeReduced int64                  `json:"activeTimeReduced"`
	Abilities         []SurvivabilityAbility `json:"abilities"`
	DamageAbilities   []SurvivabilityAbility `json:"damageAbilities"`
	Sources           []SurvivabilitySource  `json:"sources"`
	Overheal          *int64                 `json:"overheal,omitempty"`
}

type SurvivabilityFight struct {
	ID    int64 `json:"id"`
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}
