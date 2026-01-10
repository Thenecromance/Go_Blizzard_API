package automodel

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/Thenecromance/Go_Blizzard_API/global"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/jtacoma/uritemplates"
	log "github.com/sirupsen/logrus"
)

var Params map[string]any

var client *http.Client

func init() {
	buf, err := os.ReadFile("./params.json")
	if err != nil {
		return
	}

	json.Unmarshal(buf, &Params)

	// Setup oauth2 config
	conf := &clientcredentials.Config{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		TokenURL:     "https://us.battle.net/oauth/token",
	}
	client = conf.Client(context.Background())
}

type Fields struct {
	Method string
	Path   string
	Params map[string]any

	StructName  string
	PackageName string
}

func Request(field *Fields) (string, error) {
	log.Debug("Requesting:", field.Path)

	req, err := http.NewRequest(field.Method, global.PortalURL, nil)
	if err != nil {
		return "", err
	}

	tpl, err := uritemplates.Parse(field.Path)
	// pathValues := make(map[string]any)
	/* 	for key, value := range field.Params {
	   		if strings.Contains(key, "{") {
	   			key = strings.TrimLeft(key, "{")
	   			key = strings.TrimRight(key, "}")
	   			pathValues[key] = value
	   			log.Debug("Path Param:", key, "Value:", value)
	   		}
	   	}
	   	for k, v := range pathValues {
	   		log.Debug(k, v)
	   	} */
	u, _ := tpl.Expand(Params)
	log.Debug(u)
	req.URL.Path = u
	{
		q := req.URL.Query()
		for key, value := range field.Params {
			if !strings.Contains(key, "{") {
				q.Add(key, value.(string))
			}
		}
		req.URL.RawQuery = q.Encode()
	}
	log.Debug(req.URL.String())

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		log.Errorf("%s Request failed with status: %s", req.URL.String(), resp.Status)
		return "", nil /* errors.New("Request failed with status: " + resp.Status) */
	}
	return fromJsonToGoStruct(
		resp.Body,
		field.StructName,
		field.PackageName)
}
