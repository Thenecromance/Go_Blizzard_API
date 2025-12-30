package wow

import (
	"Unofficial_API/internal"
	"Unofficial_API/model"
	"Unofficial_API/utils"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
)

type Summary struct {
	CharacterSummary struct {
		Name   string `json:"name"`
		Gender struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"gender"`
		Faction struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
		Race struct {
			Name string `json:"name"`
		} `json:"race"`
		CharacterClass struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"character_class"`
		ActiveSpec struct {
			Name string `json:"name"`
		} `json:"active_spec"`
		Realm struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"realm"`
		Guild struct {
			Name  string `json:"name"`
			Realm struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"realm"`
			Faction struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"faction"`
		} `json:"guild"`
		Level              int   `json:"level"`
		Experience         int   `json:"experience"`
		AchievementPoints  int   `json:"achievement_points"`
		LastLoginTimestamp int64 `json:"last_login_timestamp"`
		AverageItemLevel   int   `json:"average_item_level"`
		EquippedItemLevel  int   `json:"equipped_item_level"`
		ActiveTitle        struct {
			Name          string `json:"name"`
			DisplayString string `json:"display_string"`
		} `json:"active_title"`
		CovenantProgress struct {
			ChosenCovenant struct {
				Name string `json:"name"`
			} `json:"chosen_covenant"`
			RenownLevel int `json:"renown_level"`
			Soulbinds   struct {
			} `json:"soulbinds"`
		} `json:"covenant_progress"`
		NameSearch string `json:"name_search"`
	} `json:"character_summary"`
	List                    []string `json:"list"`
	Token                   string   `json:"token"`
	IsSelf                  bool     `json:"is_self"`
	Allow                   bool     `json:"allow"`
	EncountersJson          string   `json:"encounters_json"`
	AchievementCategoryJson string   `json:"achievement_category_json"`
	CharacterMedia          struct {
		Assets []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"assets"`
	} `json:"character_media"`
}

func (s *Summary) ToBNet() model.Player {
	m := model.Player{}
	m.Name = s.CharacterSummary.Name

	copier.Copy(&m, &s.CharacterSummary)

	////s.Id = 0 // This field need to use another way to get, I can not public this method now
	//{
	//	m.Gender.Name = s.CharacterSummary.Gender.Name
	//	m.Gender.Type = s.CharacterSummary.Gender.Type
	//}
	//{
	//	m.Faction.Name = s.CharacterSummary.Faction.Name
	//	m.Faction.Type = s.CharacterSummary.Faction.Type
	//}
	//{
	//	m.Race.Id = 0 // just request race list from battle.net api then remap it here
	//	m.Race.Name = s.CharacterSummary.Race.Name
	//}
	//{
	//	m.CharacterClass.Id = s.CharacterSummary.CharacterClass.Id
	//	m.CharacterClass.Name = s.CharacterSummary.CharacterClass.Name
	//}
	//{
	//	m.ActiveSpec.Name = s.CharacterSummary.ActiveSpec.Name
	//	m.ActiveSpec.Id = 0 // just request spec list from battle.net api then remap it here
	//}
	//{
	//	m.Realm.Name = s.CharacterSummary.Realm.Name
	//	m.Realm.Slug = s.CharacterSummary.Realm.Slug
	//	m.Realm.Id = 0 // check server api to get it
	//}
	//{
	//	m.Guild.Name = s.CharacterSummary.Guild.Name
	//	{
	//		m.Guild.Faction.Name = s.CharacterSummary.Guild.Faction.Name
	//		m.Guild.Faction.Type = s.CharacterSummary.Guild.Faction.Type
	//	}
	//	{
	//		m.Guild.Realm.Name = s.CharacterSummary.Guild.Realm.Name
	//		m.Guild.Realm.Slug = s.CharacterSummary.Guild.Realm.Slug
	//		m.Guild.Realm.Id = 0 // This field need to use another way to get, I can not public this method now
	//	}
	//	m.Guild.Id = 0 // This field need to use another way to get, I can not public this method now
	//}
	//m.Level = s.CharacterSummary.Level
	//m.Experience = s.CharacterSummary.Experience
	//m.AchievementPoints = s.CharacterSummary.AchievementPoints
	//m.LastLoginTimestamp = s.CharacterSummary.LastLoginTimestamp
	//m.AverageItemLevel = s.CharacterSummary.AverageItemLevel
	//m.EquippedItemLevel = s.CharacterSummary.EquippedItemLevel
	//// Other fields are ignored for now
	return m
}

func StringPlayerSummary(ctx context.Context, name string, realm_slug string) (string, error) {
	client := utils.NewRequest()
	cookies := make(map[string]string)
	//cookies["blizzard_user_info"] =
	return client.GETWithCookies("https://webapi.blizzard.cn/wow-armory-server/api/index", cookies, "role_name", name, "realm_slug", realm_slug)
}
func CNPlayerSummary(ctx context.Context, name string, realm_slug string) *Summary {
	obj, err := StringPlayerSummary(ctx, name, realm_slug)
	if err != nil {
		return nil
	}
	/*	summary := &Summary{}*/
	var summary struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Data    *Summary `json:"data"`
	}
	err = json.Unmarshal([]byte(obj), &summary)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if summary.Data.Allow == false {
		return nil
	}
	//if == "" {
	internal.StoreToken(name, realm_slug, summary.Data.Token)
	return summary.Data
}

// BNetPlayerSummary gets the player summary in Battle.net format
//
// this method should return the same structure from /profile/wow/character/{realmSlug}/{characterName}
func BNetPlayerSummary(ctx context.Context, name string, realm_slug string) *model.Player {
	summary := CNPlayerSummary(ctx, name, realm_slug)
	if summary == nil {
		return nil
	}
	player := summary.ToBNet()
	return &player
}
