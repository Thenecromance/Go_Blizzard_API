package model

type Statistics struct {
	Links     Links `json:"_links"`
	Health    int   `json:"health"`
	Power     int   `json:"power"`
	PowerType struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"power_type"`
	Speed struct {
		RatingBonus      float64 `json:"rating_bonus"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"speed"`
	Strength struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"strength"`
	Agility struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"agility"`
	Intellect struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"intellect"`
	Stamina struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"stamina"`
	MeleeCrit struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"melee_crit"`
	MeleeHaste struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"melee_haste"`
	Mastery struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"mastery"`
	BonusArmor int `json:"bonus_armor"`
	Lifesteal  struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"lifesteal"`
	Versatility                 int     `json:"versatility"`
	VersatilityDamageDoneBonus  float64 `json:"versatility_damage_done_bonus"`
	VersatilityHealingDoneBonus float64 `json:"versatility_healing_done_bonus"`
	VersatilityDamageTakenBonus float64 `json:"versatility_damage_taken_bonus"`
	Avoidance                   struct {
		RatingBonus      float64 `json:"rating_bonus"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"avoidance"`

	AttackPower       int     `json:"attack_power"`
	MainHandDamageMin float64 `json:"main_hand_damage_min"`
	MainHandDamageMax float64 `json:"main_hand_damage_max"`
	MainHandSpeed     float64 `json:"main_hand_speed"`
	MainHandDps       float64 `json:"main_hand_dps"`
	OffHandDamageMin  float64 `json:"off_hand_damage_min"`
	OffHandDamageMax  float64 `json:"off_hand_damage_max"`
	OffHandSpeed      float64 `json:"off_hand_speed"`
	OffHandDps        float64 `json:"off_hand_dps"`
	SpellPower        int     `json:"spell_power"`
	SpellPenetration  int     `json:"spell_penetration"`
	SpellCrit         struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"spell_crit"`
	ManaRegen       int `json:"mana_regen"`
	ManaRegenCombat int `json:"mana_regen_combat"`
	Armor           struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"armor"`
	Dodge struct {
		RatingBonus      int     `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"dodge"`
	Parry struct {
		RatingBonus      int `json:"rating_bonus"`
		Value            int `json:"value"`
		RatingNormalized int `json:"rating_normalized"`
	} `json:"parry"`
	Block struct {
		RatingBonus      int `json:"rating_bonus"`
		Value            int `json:"value"`
		RatingNormalized int `json:"rating_normalized"`
	} `json:"block"`
	RangedCrit struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"ranged_crit"`
	RangedHaste struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"ranged_haste"`
	SpellHaste struct {
		RatingBonus      float64 `json:"rating_bonus"`
		Value            float64 `json:"value"`
		RatingNormalized int     `json:"rating_normalized"`
	} `json:"spell_haste"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		Id    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Id   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
}
