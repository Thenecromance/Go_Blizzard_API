package model

type MythicStoneProfile struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	CurrentPeriod struct {
		Period struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id int `json:"id"`
		} `json:"period"`
		BestRuns []struct {
			CompletedTimestamp int64 `json:"completed_timestamp"`
			Duration           int   `json:"duration"`
			KeystoneLevel      int   `json:"keystone_level"`
			KeystoneAffixes    []struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"keystone_affixes"`
			Members []struct {
				Character struct {
					Name  string `json:"name"`
					Id    int    `json:"id"`
					Realm struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Id   int    `json:"id"`
						Slug string `json:"slug"`
					} `json:"realm"`
				} `json:"character"`
				Specialization struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					Id   int    `json:"id"`
				} `json:"specialization"`
				Race struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					Id   int    `json:"id"`
				} `json:"race"`
				EquippedItemLevel int `json:"equipped_item_level"`
			} `json:"members"`
			Dungeon struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"dungeon"`
			IsCompletedWithinTime bool `json:"is_completed_within_time"`
			MythicRating          struct {
				Color struct {
					R int `json:"r"`
					G int `json:"g"`
					B int `json:"b"`
					A int `json:"a"`
				} `json:"color"`
				Rating float64 `json:"rating"`
			} `json:"mythic_rating"`
			MapRating struct {
				Color struct {
					R int `json:"r"`
					G int `json:"g"`
					B int `json:"b"`
					A int `json:"a"`
				} `json:"color"`
				Rating float64 `json:"rating"`
			} `json:"map_rating"`
		} `json:"best_runs"`
	} `json:"current_period"`
	Seasons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id int `json:"id"`
	} `json:"seasons"`
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
	CurrentMythicRating struct {
		Color struct {
			R int `json:"r"`
			G int `json:"g"`
			B int `json:"b"`
			A int `json:"a"`
		} `json:"color"`
		Rating float64 `json:"rating"`
	} `json:"current_mythic_rating"`
}
