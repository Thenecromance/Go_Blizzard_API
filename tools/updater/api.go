package updater

import (
	"Unofficial_API/bridge/log"
	"fmt"
	"strings"
)

type ApiGroup struct {
	ApiGroupName string `json:"name"`    // just like "Achievement API"
	Apis         []*Api `json:"methods"` // just like "Achievement API"
}

func (ap *ApiGroup) fixed() {
	if strings.Contains(ap.ApiGroupName, " API") {
		fixedName := ap.ApiGroupName[:len(ap.ApiGroupName)-4]
		fixedName = strings.ReplaceAll(fixedName, " ", "")
		ap.ApiGroupName = fixedName
	}
	if strings.Contains(ap.ApiGroupName, " ") {
		ap.ApiGroupName = strings.ReplaceAll(ap.ApiGroupName, " ", "") // Remove spaces
	}
	for i := 0; i < len(ap.Apis); i++ {
		ap.Apis[i].fixed()
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

func (a *Api) fixed() {
	if strings.Contains(a.Name, " Card") {
		log.Info(a.Name)
	}
	a.Name = strings.ReplaceAll(a.Name, " ", "_") // Remove spaces
	a.Name = strings.ReplaceAll(a.Name, " ", "_") // Remove dashes
	if strings.Contains(a.Name, "(US,EU,KR,TW)") {
		a.Name = strings.ReplaceAll(a.Name, "(US,EU,KR,TW)", "")
	}
	if strings.Contains(a.Name, "(CN)") {
		a.Name = strings.ReplaceAll(a.Name, "(CN)", "CN")
	}

	a.Description = a.Name + " " + a.Description
	a.fixParams()
	a.fixPath()

}

func (a *Api) fixParams() {
	var filtered []*Parameters
	for i, p := range a.Params {
		/*if strings.Contains(p.Name, "locale") {
			continue
		} else if strings.Contains(p.Name, "namespace") {
			a.NameSpace = strings.Split(p.DefaultValue.(string), "-")[0] // just need this api is static or dynamic
		}else*/if strings.Contains(p.Type, "(deprecated)") {
			a.Params[i] = nil // so crazy
			continue
		} else if strings.Contains(p.Name, ".") {
			p.Name = strings.ReplaceAll(p.Name, ".", "")
			fmt.Printf("wrong params need to be fixed %s\n", p.Name)
			filtered = append(filtered, p)
		} else {
			filtered = append(filtered, p)
		}
	}
	a.Params = filtered

	for i := 0; i < len(a.Params); i++ {
		a.Params[i].fixed()
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
	Name         string `json:"name"`       // original name
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

	if strings.Contains(p.Name, "{") && strings.Contains(p.Name, "}") {
		fmt.Printf("found binding uri param: %s\n", p.Name)
		p.Name = strings.ReplaceAll(p.Name, "{", "") // remove
		p.Name = strings.ReplaceAll(p.Name, "}", "")
		p.IsBindingUri = true
	}

	if strings.HasPrefix(p.Name, ":") {
		p.Name = strings.ReplaceAll(p.Name, ":", "")
		p.IsBindingUri = true
	}
	p.ParamName = p.Name
	p.Name = strings.Title(p.Name) // capitalize the first letter
	// todo: parse Type to support golang types
	switch p.Type {
	case "integer":
		p.Type = "int"
	case "number":
		p.Type = "int"
	case "numbers":
		p.Type = "[]int"
	}
	// todo: fix the DefaultValue ( which just like static-us or dynamic-us)
}
