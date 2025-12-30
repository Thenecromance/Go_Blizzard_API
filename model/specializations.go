package model

type Specializations struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Specializations []struct {
		Specialization struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"specialization"`
		Loadouts []struct {
			IsActive             bool   `json:"is_active"`
			TalentLoadoutCode    string `json:"talent_loadout_code"`
			SelectedClassTalents []struct {
				Id      int `json:"id"`
				Rank    int `json:"rank"`
				Tooltip struct {
					Talent struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
						Id   int    `json:"id"`
					} `json:"talent"`
					SpellTooltip struct {
						Spell struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
							Id   int    `json:"id"`
						} `json:"spell"`
						Description string  `json:"description"`
						CastTime    string  `json:"cast_time"`
						PowerCost   *string `json:"power_cost,omitempty"`
						Range       string  `json:"range,omitempty"`
						Cooldown    string  `json:"cooldown,omitempty"`
					} `json:"spell_tooltip"`
				} `json:"tooltip,omitempty"`
				DefaultPoints int `json:"default_points,omitempty"`
			} `json:"selected_class_talents"`
			SelectedSpecTalents []struct {
				Id      int `json:"id"`
				Rank    int `json:"rank"`
				Tooltip struct {
					Talent struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
						Id   int    `json:"id"`
					} `json:"talent"`
					SpellTooltip struct {
						Spell struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
							Id   int    `json:"id"`
						} `json:"spell"`
						Description string  `json:"description"`
						CastTime    string  `json:"cast_time"`
						PowerCost   *string `json:"power_cost,omitempty"`
						Range       string  `json:"range,omitempty"`
						Cooldown    string  `json:"cooldown,omitempty"`
					} `json:"spell_tooltip"`
				} `json:"tooltip"`
			} `json:"selected_spec_talents,omitempty"`
			SelectedClassTalentTree struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
			} `json:"selected_class_talent_tree,omitempty"`
			SelectedSpecTalentTree struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
			} `json:"selected_spec_talent_tree,omitempty"`
			SelectedHeroTalents []struct {
				Id      int `json:"id"`
				Rank    int `json:"rank"`
				Tooltip struct {
					Talent struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
						Id   int    `json:"id"`
					} `json:"talent"`
					SpellTooltip struct {
						Spell struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
							Id   int    `json:"id"`
						} `json:"spell"`
						Description string `json:"description"`
						CastTime    string `json:"cast_time"`
					} `json:"spell_tooltip"`
				} `json:"tooltip"`
				DefaultPoints int `json:"default_points,omitempty"`
			} `json:"selected_hero_talents,omitempty"`
			SelectedHeroTalentTree struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"selected_hero_talent_tree,omitempty"`
		} `json:"loadouts"`
		PvpTalentSlots []struct {
			Selected struct {
				Talent struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					Id   int    `json:"id"`
				} `json:"talent"`
				SpellTooltip struct {
					Spell struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
						Id   int    `json:"id"`
					} `json:"spell"`
					Description string `json:"description"`
					CastTime    string `json:"cast_time"`
				} `json:"spell_tooltip"`
			} `json:"selected"`
			SlotNumber int `json:"slot_number"`
		} `json:"pvp_talent_slots,omitempty"`
	} `json:"specializations"`
	ActiveSpecialization struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"active_specialization"`
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
	ActiveHeroTalentTree struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"active_hero_talent_tree"`
}
