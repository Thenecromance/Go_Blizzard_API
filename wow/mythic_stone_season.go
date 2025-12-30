package wow

import (
	"Unofficial_API/internal"
	"Unofficial_API/model"
	"Unofficial_API/utils"
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
)

// /profile/wow/character/{realmSlug}/{characterName}/mythic-keystone-profile/season/{seasonId}
type MythicStoneSeason struct {
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

func (m MythicStoneSeason) ToBNet() *model.MythicStoneSeason {

	s := &model.MythicStoneSeason{}
	copier.Copy(&s, m)
	//s.Season.Id = m.Season.Id
	//// best runs
	//{
	//	s.BestRuns = make([]model.BestRun, len(m.BestRuns))
	//	for i, br := range m.BestRuns {
	//		s.BestRuns[i].CompletedTimestamp = br.CompletedTimestamp
	//		s.BestRuns[i].Duration = br.Duration
	//		s.BestRuns[i].KeystoneLevel = br.KeystoneLevel
	//
	//		s.BestRuns[i].KeystoneAffixes = make([]model.KeystoneAffix, len(br.KeystoneAffixes))
	//		for j, ka := range br.KeystoneAffixes {
	//			s.BestRuns[i].KeystoneAffixes[j].Id = ka.Id
	//			s.BestRuns[i].KeystoneAffixes[j].Name = ka.Name
	//		}
	//		s.BestRuns[i].Members = make([]model.Member, len(br.Members))
	//		for j, mb := range br.Members {
	//			{
	//				s.BestRuns[i].Members[j].Character.Name = mb.Character.Name
	//				s.BestRuns[i].Members[j].Character.Id = mb.Character.Id
	//			}
	//			{
	//				s.BestRuns[i].Members[j].Character.Realm.Id = mb.Character.Realm.Id
	//				s.BestRuns[i].Members[j].Character.Realm.Slug = mb.Character.Realm.Slug
	//			}
	//			{
	//				s.BestRuns[i].Members[i].Specialization.Name = mb.Specialization.Name
	//				s.BestRuns[i].Members[i].Specialization.Id = mb.Specialization.Id
	//			}
	//			{
	//				s.BestRuns[i].Members[i].Race.Name = mb.Race.Name
	//				s.BestRuns[i].Members[i].Race.Id = mb.Race.Id
	//			}
	//			s.BestRuns[i].Members[j].EquippedItemLevel = mb.EquippedItemLevel
	//		}
	//
	//		{
	//			s.BestRuns[i].Dungeon.Id = br.Dungeon.Id
	//			s.BestRuns[i].Dungeon.Name = br.Dungeon.Name
	//		}
	//		s.BestRuns[i].IsCompletedWithinTime = br.IsCompletedWithinTime
	//		s.BestRuns[i].MythicRating.Rating = br.MythicRating.Rating
	//		{
	//			s.BestRuns[i].MythicRating.Color.R = br.MythicRating.Color.R
	//			s.BestRuns[i].MythicRating.Color.G = br.MythicRating.Color.G
	//			s.BestRuns[i].MythicRating.Color.B = br.MythicRating.Color.B
	//			s.BestRuns[i].MythicRating.Color.A = br.MythicRating.Color.A
	//		}
	//		s.BestRuns[i].MapRating.Rating = br.MapRating.Rating
	//		{
	//			s.BestRuns[i].MapRating.Color.R = br.MapRating.Color.R
	//			s.BestRuns[i].MapRating.Color.G = br.MapRating.Color.G
	//			s.BestRuns[i].MapRating.Color.B = br.MapRating.Color.B
	//			s.BestRuns[i].MapRating.Color.A = br.MapRating.Color.A
	//		}
	//
	//	}
	//
	//}
	//
	//s.MythicRating.Rating = m.MythicRating.Rating
	//{
	//	s.MythicRating.Color.R = m.MythicRating.Color.R
	//	s.MythicRating.Color.G = m.MythicRating.Color.G
	//	s.MythicRating.Color.B = m.MythicRating.Color.B
	//	s.MythicRating.Color.A = m.MythicRating.Color.A
	//}
	return s
}

func StringMythicStoneSeason(ctx context.Context, name string, realm_slug string) (string, error) {
	client := utils.NewRequest()
	token := internal.TryToGetToken(name, realm_slug)
	if token == "" {
		summary := CNPlayerSummary(ctx, name, realm_slug)
		if summary == nil {
			return "", nil
		}
		token = internal.TryToGetToken(name, realm_slug)
	}

	return client.GET("https://webapi.blizzard.cn/wow-armory-server/api/do", "api", "mythic_keystone_season", "token", realm_slug)
}

func CNMythicStoneSeason(ctx context.Context, name string, realm_slug string) *MythicStoneSeason {
	obj, err := StringMythicStoneSeason(ctx, name, realm_slug)
	if err != nil {
		return nil
	}
	var MS struct {
		Code    int                `json:"code"`
		Message string             `json:"message"`
		Data    *MythicStoneSeason `json:"data"`
	}

	err = json.Unmarshal([]byte(obj), &MS)
	if err != nil {
		return nil
	}
	return MS.Data
}

func BNetMythicStoneSeason(ctx context.Context, name string, realm_slug string) *model.MythicStoneSeason {
	obj := CNMythicStoneSeason(ctx, name, realm_slug)
	if obj == nil {
		return nil
	}
	return obj.ToBNet()
}
