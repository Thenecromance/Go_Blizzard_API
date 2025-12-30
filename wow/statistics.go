package wow

import (
	"Unofficial_API/internal"
	"Unofficial_API/model"
	"Unofficial_API/utils"
	"context"
	"encoding/json"
)

// /profile/wow/character/{realmSlug}/{characterName}/statistics

type Statistics struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Agility struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"agility"`
		Armor struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"armor"`
		AttackPower int `json:"attack_power"`
		Avoidance   struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
		} `json:"avoidance"`
		Block struct {
			RatingBonus      int `json:"rating_bonus"`
			RatingNormalized int `json:"rating_normalized"`
			Value            int `json:"value"`
		} `json:"block"`
		BonusArmor int `json:"bonus_armor"`
		Character  struct {
			Name  string `json:"name"`
			Realm struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"realm"`
		} `json:"character"`
		Dodge struct {
			RatingBonus      int     `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"dodge"`
		Health    int `json:"health"`
		Intellect struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"intellect"`
		Lifesteal struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"lifesteal"`
		MainHandDamageMax float64 `json:"main_hand_damage_max"`
		MainHandDamageMin float64 `json:"main_hand_damage_min"`
		MainHandDps       float64 `json:"main_hand_dps"`
		MainHandSpeed     float64 `json:"main_hand_speed"`
		ManaRegen         int     `json:"mana_regen"`
		ManaRegenCombat   int     `json:"mana_regen_combat"`
		Mastery           struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"mastery"`
		MeleeCrit struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"melee_crit"`
		MeleeHaste struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"melee_haste"`
		OffHandDamageMax float64 `json:"off_hand_damage_max"`
		OffHandDamageMin float64 `json:"off_hand_damage_min"`
		OffHandDps       float64 `json:"off_hand_dps"`
		OffHandSpeed     float64 `json:"off_hand_speed"`
		Parry            struct {
			RatingBonus      int `json:"rating_bonus"`
			RatingNormalized int `json:"rating_normalized"`
			Value            int `json:"value"`
		} `json:"parry"`
		Power     int `json:"power"`
		PowerType struct {
			Name string `json:"name"`
		} `json:"power_type"`
		RangedCrit struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"ranged_crit"`
		RangedHaste struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"ranged_haste"`
		Speed struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
		} `json:"speed"`
		SpellCrit struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"spell_crit"`
		SpellHaste struct {
			RatingBonus      float64 `json:"rating_bonus"`
			RatingNormalized int     `json:"rating_normalized"`
			Value            float64 `json:"value"`
		} `json:"spell_haste"`
		SpellPenetration int `json:"spell_penetration"`
		SpellPower       int `json:"spell_power"`
		Stamina          struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"stamina"`
		Strength struct {
			Base      int `json:"base"`
			Effective int `json:"effective"`
		} `json:"strength"`
		Versatility                 int     `json:"versatility"`
		VersatilityDamageDoneBonus  float64 `json:"versatility_damage_done_bonus"`
		VersatilityDamageTakenBonus float64 `json:"versatility_damage_taken_bonus"`
		VersatilityHealingDoneBonus float64 `json:"versatility_healing_done_bonus"`
	} `json:"data"`
}

func (s Statistics) ToBNet() *model.Statistics {

	m := &model.Statistics{}
	m.Health = s.Data.Health
	m.Power = s.Data.Power
	{
		m.PowerType.Name = s.Data.PowerType.Name
		m.PowerType.Id = 0
	}
	{
		m.Speed.RatingBonus = s.Data.Speed.RatingBonus
		m.Speed.RatingNormalized = s.Data.Speed.RatingNormalized
	}
	{
		m.Strength.Base = s.Data.Strength.Base
		m.Strength.Effective = s.Data.Strength.Effective
	}
	{
		m.Agility.Base = s.Data.Agility.Base
		m.Agility.Effective = s.Data.Agility.Effective
	}
	{
		m.Intellect.Base = s.Data.Intellect.Base
		m.Intellect.Effective = s.Data.Intellect.Effective
	}
	{
		m.Stamina.Base = s.Data.Stamina.Base
		m.Stamina.Effective = s.Data.Stamina.Effective
	}
	{
		m.MeleeCrit.RatingBonus = s.Data.MeleeCrit.RatingBonus
		m.MeleeCrit.Value = s.Data.MeleeCrit.Value
		m.MeleeCrit.RatingNormalized = s.Data.MeleeCrit.RatingNormalized
	}
	{
		m.MeleeHaste.RatingBonus = s.Data.MeleeHaste.RatingBonus
		m.MeleeHaste.Value = s.Data.MeleeHaste.Value
		m.MeleeHaste.RatingNormalized = s.Data.MeleeHaste.RatingNormalized
	}
	{
		m.Mastery.RatingBonus = s.Data.Mastery.RatingBonus
		m.Mastery.Value = s.Data.Mastery.Value
		m.Mastery.RatingNormalized = s.Data.Mastery.RatingNormalized
	}
	m.BonusArmor = s.Data.BonusArmor
	{
		m.Lifesteal.RatingBonus = s.Data.Lifesteal.RatingBonus
		m.Lifesteal.Value = s.Data.Lifesteal.Value
		m.Lifesteal.RatingNormalized = s.Data.Lifesteal.RatingNormalized
	}
	m.Versatility = s.Data.Versatility
	m.VersatilityDamageDoneBonus = s.Data.VersatilityDamageDoneBonus
	m.VersatilityHealingDoneBonus = s.Data.VersatilityHealingDoneBonus
	m.VersatilityDamageTakenBonus = s.Data.VersatilityDamageTakenBonus
	{
		m.Avoidance.RatingBonus = s.Data.Avoidance.RatingBonus
		m.Avoidance.RatingNormalized = s.Data.Avoidance.RatingNormalized
	}
	m.AttackPower = s.Data.AttackPower
	m.MainHandDamageMin = s.Data.MainHandDamageMin
	m.MainHandDamageMax = s.Data.MainHandDamageMax
	m.MainHandSpeed = s.Data.MainHandSpeed
	m.MainHandDps = s.Data.MainHandDps
	m.OffHandDamageMin = s.Data.OffHandDamageMin
	m.OffHandDamageMax = s.Data.OffHandDamageMax
	m.OffHandSpeed = s.Data.OffHandSpeed
	m.OffHandDps = s.Data.OffHandDps
	m.SpellPower = s.Data.SpellPower
	m.SpellPenetration = s.Data.SpellPenetration
	{
		m.SpellCrit.RatingBonus = s.Data.SpellCrit.RatingBonus
		m.SpellCrit.Value = s.Data.SpellCrit.Value
		m.SpellCrit.RatingNormalized = s.Data.SpellCrit.RatingNormalized
	}
	m.ManaRegen = s.Data.ManaRegen
	m.ManaRegenCombat = s.Data.ManaRegenCombat
	{
		m.Armor.Base = s.Data.Armor.Base
		m.Armor.Effective = s.Data.Armor.Effective
	}
	{
		m.Dodge.RatingBonus = s.Data.Dodge.RatingBonus
		m.Dodge.Value = s.Data.Dodge.Value
		m.Dodge.RatingNormalized = s.Data.Dodge.RatingNormalized
	}
	{
		m.Parry.RatingBonus = s.Data.Parry.RatingBonus
		m.Parry.Value = s.Data.Parry.Value
		m.Parry.RatingNormalized = s.Data.Parry.RatingNormalized
	}
	{
		m.Block.RatingBonus = s.Data.Block.RatingBonus
		m.Block.Value = s.Data.Block.Value
		m.Block.RatingNormalized = s.Data.Block.RatingNormalized
	}
	{
		m.RangedCrit.RatingBonus = s.Data.RangedCrit.RatingBonus
		m.RangedCrit.Value = s.Data.RangedCrit.Value
		m.RangedCrit.RatingNormalized = s.Data.RangedCrit.RatingNormalized
	}
	{
		m.RangedHaste.RatingBonus = s.Data.RangedHaste.RatingBonus
		m.RangedHaste.Value = s.Data.RangedHaste.Value
		m.RangedHaste.RatingNormalized = s.Data.RangedHaste.RatingNormalized
	}
	{
		m.SpellHaste.RatingBonus = s.Data.SpellHaste.RatingBonus
		m.SpellHaste.Value = s.Data.SpellHaste.Value
		m.SpellHaste.RatingNormalized = s.Data.SpellHaste.RatingNormalized
	}
	return m
}

func StringStatisticsSummary(ctx context.Context, name string, realm_slug string) (string, error) {
	client := utils.NewRequest()
	token := internal.TryToGetToken(name, realm_slug)
	if token == "" {
		summary := CNPlayerSummary(ctx, name, realm_slug)
		if summary == nil {
			return "", nil
		}
		token = internal.TryToGetToken(name, realm_slug)
	}

	return client.GET("https://webapi.blizzard.cn/wow-armory-server/api/do", "api", "statistics", "token", realm_slug)
}

func CNStatisticsSummary(ctx context.Context, name string, realm_slug string) *Statistics {
	obj, err := StringStatisticsSummary(ctx, name, realm_slug)
	if err != nil {
		return nil
	}
	stats := &Statistics{}
	err = json.Unmarshal([]byte(obj), stats)
	if err != nil {
		return nil
	}
	return stats
}

func BNetStatisticsSummary(ctx context.Context, name string, realm_slug string) *model.Statistics {
	obj := CNStatisticsSummary(ctx, name, realm_slug)
	if obj == nil {
		return nil
	}
	return obj.ToBNet()
}
