package wow

import (
	"Unofficial_API/internal"
	"Unofficial_API/model"
	"Unofficial_API/utils"
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
)

// /profile/wow/character/{realmSlug}/{characterName}/specializations
type Specializations struct {
	ActiveHeroTalentTree struct {
		Icon string `json:"icon"`
		Id   int    `json:"id"`
		Key  struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"active_hero_talent_tree"`
	ActiveSpecialization struct {
		Icon string `json:"icon"`
		Id   int    `json:"id"`
		Key  struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"active_specialization"`
	Specializations []struct {
		Loadouts []struct {
			IsActive                bool `json:"is_active"`
			SelectedClassTalentTree struct {
				Icon string `json:"icon"`
				Id   int    `json:"id"`
				Key  struct {
				} `json:"key"`
				Name string `json:"name"`
			} `json:"selected_class_talent_tree"`
			SelectedClassTalents []struct {
				Id      int `json:"id"`
				Rank    int `json:"rank"`
				Tooltip struct {
					SpellTooltip struct {
						CastTime    string `json:"cast_time"`
						Description string `json:"description"`
						Spell       struct {
							Icon string `json:"icon"`
							Id   int    `json:"id"`
							Key  struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
						} `json:"spell"`
						Cooldown  string `json:"cooldown,omitempty"`
						Range     string `json:"range,omitempty"`
						PowerCost string `json:"power_cost,omitempty"`
					} `json:"spell_tooltip"`
					Talent struct {
						Id  int `json:"id"`
						Key struct {
						} `json:"key"`
						Name string `json:"name"`
					} `json:"talent"`
				} `json:"tooltip"`
				DefaultPoints int `json:"default_points,omitempty"`
			} `json:"selected_class_talents"`
			SelectedHeroTalentTree struct {
				Icon string `json:"icon"`
				Id   int    `json:"id"`
				Key  struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
			} `json:"selected_hero_talent_tree"`
			SelectedHeroTalents []struct {
				Id      int `json:"id"`
				Rank    int `json:"rank"`
				Tooltip struct {
					SpellTooltip struct {
						CastTime    string `json:"cast_time"`
						Description string `json:"description"`
						Spell       struct {
							Icon string `json:"icon"`
							Id   int    `json:"id"`
							Key  struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
						} `json:"spell"`
					} `json:"spell_tooltip"`
					Talent struct {
						Id  int `json:"id"`
						Key struct {
						} `json:"key"`
						Name string `json:"name"`
					} `json:"talent"`
				} `json:"tooltip"`
				DefaultPoints int `json:"default_points,omitempty"`
			} `json:"selected_hero_talents"`
			SelectedSpecTalentTree struct {
				Icon string `json:"icon"`
				Id   int    `json:"id"`
				Key  struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
			} `json:"selected_spec_talent_tree"`
			SelectedSpecTalents []struct {
				Id      int `json:"id"`
				Rank    int `json:"rank"`
				Tooltip struct {
					SpellTooltip struct {
						CastTime    string `json:"cast_time"`
						Description string `json:"description"`
						Spell       struct {
							Icon string `json:"icon"`
							Id   int    `json:"id"`
							Key  struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
						} `json:"spell"`
						Cooldown  string `json:"cooldown,omitempty"`
						PowerCost string `json:"power_cost,omitempty"`
						Range     string `json:"range,omitempty"`
					} `json:"spell_tooltip"`
					Talent struct {
						Id  int `json:"id"`
						Key struct {
						} `json:"key"`
						Name string `json:"name"`
					} `json:"talent"`
				} `json:"tooltip"`
			} `json:"selected_spec_talents"`
			TalentLoadoutCode string `json:"talent_loadout_code"`
		} `json:"loadouts"`
		PvpTalentSlots []struct {
			Selected struct {
				SpellTooltip struct {
					CastTime    string `json:"cast_time"`
					Cooldown    string `json:"cooldown,omitempty"`
					Description string `json:"description"`
					Spell       struct {
						Icon string `json:"icon"`
						Id   int    `json:"id"`
						Key  struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
					} `json:"spell"`
				} `json:"spell_tooltip"`
				Talent struct {
					Id  int `json:"id"`
					Key struct {
					} `json:"key"`
					Name string `json:"name"`
				} `json:"talent"`
			} `json:"selected"`
			SlotNumber int `json:"slot_number"`
		} `json:"pvp_talent_slots"`
		Specialization struct {
			Icon string `json:"icon"`
			Id   int    `json:"id"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"specialization"`
	} `json:"specializations"`
}

func (s *Specializations) ToBNet() *model.Specializations {
	if s == nil {
		return nil
	}
	ms := &model.Specializations{}
	copier.Copy(&ms, s)
	return ms
}

func StringSpecializations(ctx context.Context, name string, realm_slug string) (string, error) {
	client := utils.NewRequest()
	token := internal.TryToGetToken(name, realm_slug)
	if token == "" {
		summary := CNPlayerSummary(ctx, name, realm_slug)
		if summary == nil {
			return "", nil
		}
		token = internal.TryToGetToken(name, realm_slug)
	}

	return client.GET("https://webapi.blizzard.cn/wow-armory-server/api/do", "api", "specializations", "token", token)
}

func CNSpecializations(ctx context.Context, name string, realm_slug string) *Specializations {
	obj, err := StringSpecializations(ctx, name, realm_slug)

	if err != nil {
		return nil
	}
	var MS struct {
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Data    *Specializations `json:"data"`
	}

	err = json.Unmarshal([]byte(obj), &MS)
	if err != nil {
		return nil
	}
	return MS.Data
}

func BNetSpecializations(ctx context.Context, name string, realm_slug string) *model.Specializations {
	obj := CNSpecializations(ctx, name, realm_slug)
	if obj == nil {
		return nil
	}
	return obj.ToBNet()
}

//area-52
//affinitym
