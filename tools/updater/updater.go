package updater

import (
	"Unofficial_API/bridge/log"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/bytedance/sonic"
)

func RequestApiList(apiPath string) []*ApiGroup {
	url := apiPath
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get api list: %v", err)
		return nil
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		return nil
	}

	var Response struct {
		ApiList []*ApiGroup `json:"resources"`
	}

	if err = sonic.Unmarshal(buf, &Response); err != nil {

		fmt.Printf("failed to unmarshal api list: %v\n", err)
		//fmt.Printf("%s\n", string(buf))
		return nil
	}

	return Response.ApiList
}

func ParseTemplate(pkgName string, folder string, apiPath string) {

	file, err := os.ReadFile("./templates/api.tmpl")
	if err != nil {
		return
	}
	dataTmpl = string(file)

	t, err := template.New("tmpl").Parse(dataTmpl)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// also prepare model template
	m, err := template.New("model").Parse(modelTmpl)
	if err != nil {
		fmt.Printf("Error parsing model template: %v\n", err)
		return
	}

	apis := RequestApiList(apiPath)

	os.MkdirAll(folder, os.ModePerm)
	for _, apiGroup := range apis {
		apiGroup.fixed()
		subFolder := folder + "/" + apiGroup.ApiGroupName
		os.Mkdir(subFolder, os.ModePerm)

		file, err := os.Create(subFolder + "/" + apiGroup.ApiGroupName + ".go")
		if err = t.Execute(file, map[string]any{
			"PkgName":      pkgName,
			"ApiGroupName": apiGroup.ApiGroupName,
			"Apis":         apiGroup.Apis,
		}); err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			return
		}
		file.Close()
		fmt.Printf("generate api %s\n", apiGroup.ApiGroupName)

		// For each API in the group, generate a separate model file named {{.Name}}.go
		for _, api := range apiGroup.Apis {
			modelFilePath := subFolder + "/" + api.Name + ".model.go"
			mf, err := os.Create(modelFilePath)
			if err != nil {
				fmt.Printf("Error creating model file %s: %v\n", modelFilePath, err)
				continue
			}
			if err = m.Execute(mf, map[string]any{
				"PkgName":      pkgName,
				"ApiGroupName": apiGroup.ApiGroupName,
				"Apis":         apiGroup.Apis,
				"Name":         api.Name,
			}); err != nil {
				fmt.Printf("Error executing model template for %s: %v\n", api.Name, err)
				mf.Close()
				continue
			}
			mf.Close()
			fmt.Printf("generate model %s\n", api.Name)
		}
	}
	return
}

func GenerateRouterList(pkgName string, folder string, apiPath string) {
	file, err := os.ReadFile("./templates/routers.tmpl")
	if err != nil {
		return
	}
	Tmpl := string(file)

	t, err := template.New("tmpl").Parse(Tmpl)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}
	/*
		// also prepare model template
		m, err := template.New("model").Parse(modelTmpl)
		if err != nil {
			fmt.Printf("Error parsing model template: %v\n", err)
			return
		}*/

	apis := RequestApiList(apiPath)

	os.MkdirAll(folder, os.ModePerm)
	var apiList []*Api
	for _, apiGroup := range apis {
		apiGroup.fixed()
		apiList = append(apiList, apiGroup.Apis...)
	}

	out, err := os.Create(folder + "/" + pkgName + "." + "router.go")
	if err = t.Execute(out, map[string]any{
		"PkgName":      pkgName,
		"ApiGroupName": "routers",
		"Apis":         apiList,
	}); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}
	out.Close()
	return
}

type GenerateHandler func(pkgName string, folder string, apiList []*ApiGroup)

func read(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
func parse(path string) (*template.Template, error) {
	file, err := read(path)
	if err != nil {
		return nil, err
	}
	log.Info("Parsing template from %s", path)
	log.Debug(file)
	return template.New(path).Parse(file)
}

func GenerateApi(pkgName string, folder string, apiList []*ApiGroup) {
	t, err := parse("./templates/api.tmpl")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
	}

	{
		os.MkdirAll(folder, os.ModePerm)
	}

	for _, apiGroup := range apiList {
		apiGroup.fixed()
		subFolder := folder + "/" + apiGroup.ApiGroupName
		os.Mkdir(subFolder, os.ModePerm)

		file, err := os.Create(subFolder + "/" + apiGroup.ApiGroupName + ".go")
		if err = t.Execute(file, map[string]any{
			"PkgName":      pkgName,
			"ApiGroupName": apiGroup.ApiGroupName,
			"Apis":         apiGroup.Apis,
		}); err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			return
		}
		file.Close()
		fmt.Printf("generate api %s\n", apiGroup.ApiGroupName)
	}

}

func GenerateModels(pkgName string, folder string, apiList []*ApiGroup) {
	t, err := parse("./templates/model.tmpl")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
	}

	{
		os.MkdirAll(folder, os.ModePerm)
	}

	for _, apiGroup := range apiList {
		apiGroup.fixed()
		subFolder := folder + "/" + apiGroup.ApiGroupName
		os.Mkdir(subFolder, os.ModePerm)

		// For each API in the group, generate a separate model file named {{.Name}}.go
		for _, api := range apiGroup.Apis {
			modelFilePath := subFolder + "/" + api.Name + ".model.go"
			mf, err := os.Create(modelFilePath)
			if err != nil {
				fmt.Printf("Error creating model file %s: %v\n", modelFilePath, err)
				continue
			}
			if err = t.Execute(mf, map[string]any{
				"PkgName":      pkgName,
				"ApiGroupName": apiGroup.ApiGroupName,
				"Apis":         apiGroup.Apis,
				"Name":         api.Name,
			}); err != nil {
				fmt.Printf("Error executing model template for %s: %v\n", api.Name, err)
				mf.Close()
				continue
			}
			mf.Close()
			fmt.Printf("generate model %s\n", api.Name)
		}
	}
}

func GenerateRouters(pkgName string, folder string, apiList []*ApiGroup) {
	log.Info("Generating routers...")
	t, err := parse("./templates/routers.tmpl")
	if err != nil {
		log.Warn("Error parsing template: ", err)
	}

	{
		log.Debug("Creating folder: ", folder)
		os.MkdirAll(folder, os.ModePerm)
	}

	for _, apiGroup := range apiList {
		apiGroup.fixed()

		file, err := os.Create(folder + "/" + apiGroup.ApiGroupName + ".go")
		if err = t.Execute(file, map[string]any{
			"PkgName":      pkgName,
			"ApiGroupName": apiGroup.ApiGroupName,
			"Apis":         apiGroup.Apis,
		}); err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			return
		}
		file.Close()
		fmt.Printf("generate router %s\n", apiGroup.ApiGroupName)
	}

}
