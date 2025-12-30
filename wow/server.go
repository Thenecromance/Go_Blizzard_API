package wow

import (
	"Unofficial_API/model"
	"Unofficial_API/utils"
	"context"
	"encoding/json"
)

type ServerStatus struct {
	Category       string `json:"category"`
	Id             int    `json:"id"`
	Locale         string `json:"locale"`
	Name           string `json:"name"`
	PopulationName string `json:"population_name"`
	PopulationType string `json:"population_type"`
	Region         string `json:"region"`
	Slug           string `json:"slug"`
	StatusName     string `json:"status_name"`
	StatusType     string `json:"status_type"`
	Timezone       string `json:"timezone"`
	TypeName       string `json:"type_name"`
	TypeType       string `json:"type_type"`
}

func (s ServerStatus) ToBNet() model.ConnectedRealm {
	m := model.ConnectedRealm{}
	m.Id = s.Id
	m.HasQueue = false
	{
		m.Status.Name = s.StatusName
		m.Status.Type = s.StatusType
	}
	{

		m.Population.Name = s.PopulationName
		m.Population.Type = s.PopulationType
	}

	m.Realms = make([]model.Realm, 1)
	m.Realms[0].Id = s.Id
	m.Realms[0].Region.Name = s.Region
	m.Realms[0].Name = s.Name
	m.Realms[0].Category = s.Category
	m.Realms[0].Locale = s.Locale
	m.Realms[0].Timezone = s.Timezone
	m.Realms[0].Slug = s.Slug
	{
		m.Realms[0].Type.Name = s.TypeName
		m.Realms[0].Type.Type = s.TypeType
	}
	m.Realms[0].IsTournament = false

	return m
}

// StringUpdateMainLineServer Request the latest mainline server status
func StringUpdateMainLineServer(ctx context.Context) (string, error) {
	return requestServer(ctx, "wow_mainline")
}

func UpdateMainLineServer(ctx context.Context) []ServerStatus {
	obj, err := StringUpdateMainLineServer(ctx)
	if err != nil {
		return nil
	}

	var StatusList struct {
		Code int `json:"code"`
		Data struct {
			List []ServerStatus
		} `json:"data"`
	}
	err = json.Unmarshal([]byte(obj), &StatusList)
	if err != nil {
		return nil
	}

	return StatusList.Data.List
}

func OfficialUpdateMainLineServer(ctx context.Context) []model.ConnectedRealm {
	statuses := UpdateMainLineServer(ctx)
	if statuses == nil {
		return nil
	}
	realms := make([]model.ConnectedRealm, 0, len(statuses))
	for _, status := range statuses {
		realms = append(realms, status.ToBNet())
	}
	return realms
}

// StringUpdateClassicServer Request the latest classic server status (Current is MoP)
func StringUpdateClassicServer(ctx context.Context) (string, error) {
	return requestServer(ctx, "wow_classic")
}

func UpdateClassicServer(ctx context.Context) []ServerStatus {
	obj, err := StringUpdateClassicServer(ctx)
	if err != nil {
		return nil
	}

	var StatusList struct {
		Code int `json:"code"`
		Data struct {
			List []ServerStatus
		} `json:"data"`
	}
	err = json.Unmarshal([]byte(obj), &StatusList)
	if err != nil {
		return nil
	}

	return StatusList.Data.List
}

func OfficialUpdateClassicServer(ctx context.Context) []model.ConnectedRealm {
	statuses := UpdateClassicServer(ctx)
	if statuses == nil {
		return nil
	}
	realms := make([]model.ConnectedRealm, 0, len(statuses))
	for _, status := range statuses {
		realms = append(realms, status.ToBNet())
	}
	return realms
}

// StringUpdate60Classic Request the classic server(60 level) status
func StringUpdate60Classic(ctx context.Context) (string, error) {
	return requestServer(ctx, "wow_classic1x")
}

func Update60Classic(ctx context.Context) []ServerStatus {
	obj, err := StringUpdate60Classic(ctx)
	if err != nil {
		return nil
	}

	var StatusList struct {
		Code int `json:"code"`
		Data struct {
			List []ServerStatus
		} `json:"data"`
	}
	err = json.Unmarshal([]byte(obj), &StatusList)
	if err != nil {
		return nil
	}

	return StatusList.Data.List
}

func OfficialUpdate60Classic(ctx context.Context) []model.ConnectedRealm {
	statuses := Update60Classic(ctx)
	if statuses == nil {
		return nil
	}
	realms := make([]model.ConnectedRealm, 0, len(statuses))
	for _, status := range statuses {
		realms = append(realms, status.ToBNet())
	}
	return realms
}

// requestServer is a helper function to request server status
func requestServer(ctx context.Context, svr string) (string, error) {
	cli := utils.NewRequest() //TODO: client should be moved to package level
	return cli.GET("https://webapi.blizzard.cn/wow-armory-server/api/server_status", "server_type", svr)
}
