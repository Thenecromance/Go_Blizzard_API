package wow_CharacterPvP

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Thenecromance/BlizzardAPI/ApiError"
	"github.com/Thenecromance/BlizzardAPI/bridge/log"
	cn "github.com/Thenecromance/BlizzardAPI/cnSupport"
	"github.com/Thenecromance/BlizzardAPI/global"
)

func init() {
	CNHookCharacterPvPSummary = CNHookCharacterPvPSummaryFunc
	CNHookCharacterPvPBracketStatistics = CNHookCharacterPvPBracketStatisticsFunc
}

func CNHookCharacterPvPSummaryFunc(ctx context.Context, field *CharacterPvPSummaryFields) (any, error) {
	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://webapi.blizzard.cn/wow-armory-server/api/do", nil)
	if err != nil {
		return nil, err
	}

	{
		token, ok := cn.TryGetPlayerToken(field.CharacterName, field.RealmSlug)
		if !ok {
			return nil, ApiError.ErrorNoPlayerToken
		}

		q := req.URL.Query()
		q.Add("api", "pvp_summary")
		q.Add("token", token)
		req.URL.RawQuery = q.Encode()

		req.AddCookie(&http.Cookie{
			Name:  "blizzard_user_info",
			Value: global.GetCookie(),
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
		Code    int                         ` json:"code,omitempty"`
		Message string                      `json:"message,omitempty"`
		Data    *CNCharacterPvPSummaryModel `json:"data,omitempty"`
	}
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		log.Infof("request URL:[%s] failed with message:[%s]", req.URL.String(), result.Message)
		return nil, ApiError.ErrorRequestFailed
	}

	/*if !result.Data {
		log.Infof("request URL:[%s] not allowed to access", req.URL.String())
		return nil, ApiError.ErrorPlayerNotAllowedToAccess
	}*/

	return postProcess(result.Data, field), nil
}

func postProcess(data *CNCharacterPvPSummaryModel, field *CharacterPvPSummaryFields) *CharacterPvPSummaryModel {

	/*	buf, _ := json.MarshalIndent(data, "", "  ")
		os.WriteFile(fmt.Sprintf("./%s_%s_equipment_original.json", field.CharacterName, field.RealmSlug), buf, 0777)
		buf, _ = json.MarshalIndent(result, "", "  ")
		os.WriteFile(fmt.Sprintf("./%s_%s_equipment_bnet.json", field.CharacterName, field.RealmSlug), buf, 0777)*/

	result := data.ToBNet().(*CharacterPvPSummaryModel)
	/*profileBasePath := fmt.Sprintf("%s/profile/wow/character/%s/%s", global.GetLocalHost(), field.RealmSlug, field.CharacterName)
	dataBasePath := fmt.Sprintf("%s/data/wow/character/%s/%s", global.GetLocalHost(), field.RealmSlug, field.CharacterName)

	links := map[string]*struct {
		IsProfile bool
		Href      *string
	}{
		"character": {IsProfile: true, Href: &result.Character.Key.Href},
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
	}*/

	return result
}

func CNHookCharacterPvPBracketStatisticsFunc(ctx context.Context, field *CharacterPvPBracketStatisticsFields) (any, error) {
	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://webapi.blizzard.cn/wow-armory-server/api/do", nil)
	if err != nil {
		return nil, err
	}

	{
		token, ok := cn.TryGetPlayerToken(field.CharacterName, field.RealmSlug)
		if !ok {
			return nil, ApiError.ErrorNoPlayerToken
		}

		q := req.URL.Query()
		q.Add("api", "pvp_summary")
		q.Add("token", token)
		req.URL.RawQuery = q.Encode()

		req.AddCookie(&http.Cookie{
			Name:  "blizzard_user_info",
			Value: global.GetCookie(),
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
		Code    int                         ` json:"code,omitempty"`
		Message string                      `json:"message,omitempty"`
		Data    *CNCharacterPvPSummaryModel `json:"data,omitempty"`
	}
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		log.Infof("request URL:[%s] failed with message:[%s]", req.URL.String(), result.Message)
		return nil, ApiError.ErrorRequestFailed
	}

	/*if !result.Data {
		log.Infof("request URL:[%s] not allowed to access", req.URL.String())
		return nil, ApiError.ErrorPlayerNotAllowedToAccess
	}*/

	return result.Data.ToBNet(), nil
}
