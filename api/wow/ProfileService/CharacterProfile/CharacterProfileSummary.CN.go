package wow_CharacterProfile

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Thenecromance/Go_Blizzard_API/ApiError"
	"github.com/Thenecromance/Go_Blizzard_API/bridge/log"
	cn "github.com/Thenecromance/Go_Blizzard_API/cnSupport"
	"github.com/Thenecromance/Go_Blizzard_API/global"
)

func init() {
	CNHookCharacterProfileSummary = CNHookCharacterProfileSummaryFunc
}

func CNHookCharacterProfileSummaryFunc(ctx context.Context, field *CharacterProfileSummaryFields) (any, error) {
	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://webapi.blizzard.cn/wow-armory-server/api/index", nil)
	if err != nil {
		return nil, err
	}

	{
		q := req.URL.Query()
		q.Add("realm_slug", field.RealmSlug)
		q.Add("role_name", field.CharacterName)
		req.URL.RawQuery = q.Encode()

		req.AddCookie(&http.Cookie{
			Name:  "blizzard_user_info",
			Value: global.GetNetEaseCookie(),
		})
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0")
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Infof("request URL:[%s] failed with status code:[%d]", req.URL.String(), resp.StatusCode)
		return nil, ApiError.ErrorRequestFailed
	}

	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Debug(string(buf))
	var result struct {
		Code    int                             ` json:"code,omitempty"`
		Message string                          `json:"message,omitempty"`
		Data    *CNCharacterProfileSummaryModel `json:"data,omitempty"`
	}
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		log.Infof("request URL:[%s] failed with message:[%s]", req.URL.String(), result.Message)
		return nil, ApiError.ErrorRequestFailed
	}

	if !result.Data.Allow {
		log.Infof("request URL:[%s] not allowed to access", req.URL.String())
		return nil, ApiError.ErrorPlayerNotAllowedToAccess
	}

	return postProcess(result.Data, field), nil
}

func postProcess(data *CNCharacterProfileSummaryModel, field *CharacterProfileSummaryFields) *CharacterProfileSummaryModel {
	cn.AddPlayerToken(field.CharacterName, field.RealmSlug, data.Token)
	result := data.ToBNet().(*CharacterProfileSummaryModel)
	profileBasePath := fmt.Sprintf("%s/profile/wow/character/%s/%s", global.GetLocalHost(), field.RealmSlug, field.CharacterName)
	dataBasePath := fmt.Sprintf("%s/data/wow/data/character/%s/%s", global.GetLocalHost(), field.RealmSlug, field.CharacterName)
	links := map[string]*struct {
		IsProfile bool
		Href      *string
	}{
		"statistics":              {IsProfile: true, Href: &result.Statistics.Href},
		"equipment":               {IsProfile: true, Href: &result.Equipment.Href},
		"mythic-keystone-profile": {IsProfile: true, Href: &result.MythicKeystoneProfile.Href},
		"achievements":            {IsProfile: true, Href: &result.Achievements.Href},
		"titles":                  {IsProfile: true, Href: &result.Titles.Href},
		"pvp-summary":             {IsProfile: true, Href: &result.PvpSummary.Href},
		"encounters":              {IsProfile: true, Href: &result.Encounters.Href},
		"character-media":         {IsProfile: true, Href: &result.Media.Href},
		"specializations":         {IsProfile: true, Href: &result.Specializations.Href},
		"appearance":              {IsProfile: true, Href: &result.Appearance.Href},
		"collections":             {IsProfile: true, Href: &result.Collections.Href},
		"reputations":             {IsProfile: true, Href: &result.Reputations.Href},
		"quests":                  {IsProfile: true, Href: &result.Quests.Href},
		"professions":             {IsProfile: true, Href: &result.Professions.Href},

		"achievements/statistics": {IsProfile: true, Href: &result.AchievementsStatistics.Href},
	}

	for endpoint, hrefPtr := range links {
		if hrefPtr != nil {
			if hrefPtr.Href == nil {
				continue
			}
			if hrefPtr.IsProfile {
				*hrefPtr.Href = fmt.Sprintf("%s/%s?namespace=profile-cn", profileBasePath, endpoint)
			} else {
				*hrefPtr.Href = fmt.Sprintf("%s/%s?namespace=profile-cn", dataBasePath, endpoint)
			}
		}
	}
	return result
}
