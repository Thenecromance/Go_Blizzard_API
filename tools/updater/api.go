package updater

import (
	"Unofficial_API/bridge/log"
	"fmt"
	"strings"
)

type ApiGroup struct {
	Game         string `json:"package"` // just like "Achievement"
	ApiGroupName string `json:"name"`    // just like "Achievement API"
	Apis         []*Api `json:"methods"` // just like "Achievement API"
}

func (ap *ApiGroup) Fixed() {
	if strings.HasSuffix(ap.ApiGroupName, " API") {
		ap.ApiGroupName = strings.TrimSuffix(ap.ApiGroupName, " API")
	}
	ap.ApiGroupName = strings.ReplaceAll(ap.ApiGroupName, " ", "")

	for _, api := range ap.Apis {
		api.fixed()
	}
}

type Api struct {
	Name        string        `json:"name"`        // just like "Achievement API"
	Description string        `json:"description"` // what this api do
	Path        string        `json:"path"`        // just like "/wow/achievement"
	GinPath     string        `json:"gin_path"`    // the real gin path with :id
	Method      string        `json:"httpMethod"`  // just like "GET"
	CnRegion    bool          `json:"cnRegion"`    // just like true
	Params      []*Parameters `json:"parameters"`  // just like "id"
	NameSpace   string        `json:"name_space"`  // the real place which will write to api file
}

func (a *Api) HasURIBinding() bool {
	for _, p := range a.Params {
		if p != nil && p.IsBindingUri {
			return true
		}
	}
	return false
}

func (a *Api) fixed() {
	if strings.Contains(a.Name, " Card") {
		log.Info(a.Name)
	}
	a.Name = strings.ReplaceAll(a.Name, " ", "") // Remove spaces
	a.Name = strings.ReplaceAll(a.Name, "(US,EU,KR,TW)", "")
	a.Name = strings.ReplaceAll(a.Name, "(CN)", "CN")

	a.Description = a.Name + " " + a.Description
	a.fixParams()
	a.fixPath()
}

func (a *Api) fixParams() {
	var filtered []*Parameters
	for _, p := range a.Params {
		// 优先处理弃用字段，直接跳过不添加到 filtered 中即为删除
		if strings.Contains(p.SourceName, "(deprecated)") {
			log.Infof("found deprecated param: %s at API:%s", p.SourceName, a.Name)
			continue
		}

		if strings.Contains(p.SourceName, ".") {
			p.SourceName = strings.ReplaceAll(p.SourceName, ".", "")
			fmt.Printf("wrong params need to be fixed %s\n", p.SourceName)
		}

		filtered = append(filtered, p)
	}
	a.Params = filtered

	for _, p := range a.Params {
		p.fixed()
	}
}

func (a *Api) formatPath(template string, params map[string]string) string {
	var args = make([]string, 0)
	for key, value := range params {
		placeholder := "{" + key + "}"
		switch value {
		case "string":
			template = strings.Replace(template, placeholder, "%s", -1)
			args = append(args, key)
		case "int":
			template = strings.Replace(template, placeholder, "%d", -1)
			args = append(args, key)
		case "float":
			template = strings.Replace(template, placeholder, "%f", -1)
			args = append(args, key)
		default:
			panic(fmt.Errorf("unsupported type: %s", value))
		}
	}

	fmt.Sprintf("fmt.Sprinf(template, args...) ")

	result := "fmt.Sprintf(\"" + template + "\","
	for _, arg := range args {
		result += arg + ","
	}
	result = strings.TrimRight(result, ",")
	result += ")"

	return result
}

func (a *Api) fixPath() {
	lists := strings.Split(a.Path, "/")
	for i, segment := range lists {
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			paramName := strings.ReplaceAll(strings.ReplaceAll(segment, "{", ""), "}", "")
			lists[i] = ":" + paramName // just mark it
		}
	}
	a.GinPath = strings.Join(lists, "/")
}

type Parameters struct {
	SourceName   string `json:"name"`
	Name         string `json:"title"`      // original name
	ParamName    string `json:"param_name"` // fixed name for golang
	IsBindingUri bool   `json:"is_binding_uri"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	DefaultValue any    `json:"defaultValue"`
}

func (p *Parameters) fixed() {
	if p == nil {
		return
	}

	if strings.Contains(p.SourceName, "{") && strings.Contains(p.SourceName, "}") {
		fmt.Printf("found binding uri param: %s\n", p.SourceName)
		p.SourceName = strings.ReplaceAll(p.SourceName, "{", "") // remove
		p.SourceName = strings.ReplaceAll(p.SourceName, "}", "")
		p.IsBindingUri = true
	}

	if strings.HasPrefix(p.SourceName, ":") {
		p.SourceName = strings.ReplaceAll(p.SourceName, ":", "")
		p.IsBindingUri = true
	}
	p.ParamName = p.SourceName
	p.Name = strings.Title(p.SourceName) // capitalize the first letter
	// todo: parse Type to support golang types
	switch p.Type {
	case "integer":
		p.Type = "int"
	case "number":
		p.Type = "int"
	case "numbers":
		p.Type = "[]int"
	}
}
