package wow

import (
	"Unofficial_API/internal"
	"Unofficial_API/model"
	"Unofficial_API/utils"
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
)

type MythicStoneProfile struct {
	BestRuns []struct {
		CompletedTimestamp int64 `json:"completed_timestamp"`
		Dungeon            struct {
			Icon string `json:"icon"`
			Id   int    `json:"id"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"dungeon"`
		Duration              int  `json:"duration"`
		IsCompletedWithinTime bool `json:"is_completed_within_time"`
		KeystoneAffixes       []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
			Id          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"keystone_affixes"`
		KeystoneLevel int `json:"keystone_level"`
		MapRating     struct {
			Color struct {
				A int `json:"a"`
				B int `json:"b"`
				G int `json:"g"`
				R int `json:"r"`
			} `json:"color"`
			Rating float64 `json:"rating"`
		} `json:"map_rating"`
		Members []struct {
			Character struct {
				CharacterId string `json:"character_id"`
				ClassId     int    `json:"class_id"`
				Id          int    `json:"id"`
				Name        string `json:"name"`
				Realm       struct {
					Id   int    `json:"id"`
					Slug string `json:"slug"`
				} `json:"realm"`
			} `json:"character"`
			EquippedItemLevel int `json:"equipped_item_level"`
			Race              struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"race"`
			Specialization struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"specialization"`
		} `json:"members"`
		MythicRating struct {
			Color struct {
				A int `json:"a"`
				B int `json:"b"`
				G int `json:"g"`
				R int `json:"r"`
			} `json:"color"`
			Rating float64 `json:"rating"`
		} `json:"mythic_rating"`
	} `json:"best_runs"`
	Character struct {
		Id  int `json:"id"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		Realm struct {
			Id  int `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	MythicRating struct {
		Color struct {
			A int `json:"a"`
			B int `json:"b"`
			G int `json:"g"`
			R int `json:"r"`
		} `json:"color"`
		Rating float64 `json:"rating"`
	} `json:"mythic_rating"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"Key"`
		Id        int    `json:"id"`
		SeasonStr string `json:"season_str"`
	} `json:"season"`
}

func (m *MythicStoneProfile) ToBNet() *model.MythicStoneProfile {
	if m == nil {
		return nil
	}
	ms := &model.MythicStoneProfile{}

	copier.Copy(&ms, m)
	return ms
}

func StringMythicStoneProfile(ctx context.Context, name string, realm_slug string) (string, error) {
	client := utils.NewRequest()
	token := internal.TryToGetToken(name, realm_slug)
	if token == "" {
		summary := CNPlayerSummary(ctx, name, realm_slug)
		if summary == nil {
			return "", nil
		}
		token = internal.TryToGetToken(name, realm_slug)
	}

	return client.GET("https://webapi.blizzard.cn/wow-armory-server/api/do", "api", "mythic_keystone_profile", "token", realm_slug)
}

func CNMythicStoneProfile(ctx context.Context, name string, realm_slug string) *MythicStoneProfile {
	obj, err := StringMythicStoneProfile(ctx, name, realm_slug)
	if err != nil {
		return nil
	}
	var MS struct {
		Code    int                 `json:"code"`
		Message string              `json:"message"`
		Data    *MythicStoneProfile `json:"data"`
	}

	err = json.Unmarshal([]byte(obj), &MS)
	if err != nil {
		return nil
	}
	return MS.Data
}

func BNetMythicStoneProfile(ctx context.Context, name string, realm_slug string) *model.MythicStoneProfile {
	obj := CNMythicStoneProfile(ctx, name, realm_slug)
	if obj == nil {
		return nil
	}
	return obj.ToBNet()
}
