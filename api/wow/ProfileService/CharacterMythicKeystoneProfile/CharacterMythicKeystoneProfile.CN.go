package wow_CharacterMythicKeystoneProfile

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
	CNHookCharacterMythicKeystoneProfileIndex = CNHookCharacterMythicKeystoneProfileIndexFunc
	CNHookCharacterMythicKeystoneSeasonDetails = CNHookCharacterMythicKeystoneSeasonDetailsFunc
}

func CNHookCharacterMythicKeystoneProfileIndexFunc(ctx context.Context, field *CharacterMythicKeystoneProfileIndexFields) (any, error) {
	postProcess := func(data *CNCharacterMythicKeystoneProfileIndexModel, field *CharacterMythicKeystoneProfileIndexFields) *CharacterMythicKeystoneProfileIndexModel {

		result := data.ToBNet().(*CharacterMythicKeystoneProfileIndexModel)

		return result
	}

	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://webapi.blizzard.cn/wow-armory-server/api/index", nil)
	if err != nil {
		return nil, err
	}

	{
		token, ok := cn.TryGetPlayerToken(field.CharacterName, field.RealmSlug)
		if !ok {
			return nil, ApiError.ErrorNoPlayerToken
		}

		q := req.URL.Query()
		q.Add("api", "mythic_keystone_profile")
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
		Code    int                                         ` json:"code,omitempty"`
		Message string                                      `json:"message,omitempty"`
		Data    *CNCharacterMythicKeystoneProfileIndexModel `json:"data,omitempty"`
	}
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		log.Infof("request URL:[%s] failed with message:[%s]", req.URL.String(), result.Message)
		return nil, ApiError.ErrorRequestFailed
	}

	return postProcess(result.Data, field), nil
}

func CNHookCharacterMythicKeystoneSeasonDetailsFunc(ctx context.Context, field *CharacterMythicKeystoneSeasonDetailsFields) (any, error) {
	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://webapi.blizzard.cn/wow-armory-server/api/index", nil)
	if err != nil {
		return nil, err
	}

	{
		token, ok := cn.TryGetPlayerToken(field.CharacterName, field.RealmSlug)
		if !ok {
			return nil, ApiError.ErrorNoPlayerToken
		}

		q := req.URL.Query()
		q.Add("api", "mythic_keystone_profile")
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
		Code    int                                          ` json:"code,omitempty"`
		Message string                                       `json:"message,omitempty"`
		Data    *CNCharacterMythicKeystoneSeasonDetailsModel `json:"data,omitempty"`
	}
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		log.Infof("request URL:[%s] failed with message:[%s]", req.URL.String(), result.Message)
		return nil, ApiError.ErrorRequestFailed
	}

	postProcess := func(data *CNCharacterMythicKeystoneSeasonDetailsModel, field *CharacterMythicKeystoneSeasonDetailsFields) *CharacterMythicKeystoneSeasonDetailsModel {
		result := data.ToBNet().(*CharacterMythicKeystoneSeasonDetailsModel)
		return result
	}
	return postProcess(result.Data, field), nil
}
