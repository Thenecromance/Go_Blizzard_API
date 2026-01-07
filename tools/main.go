package main

import (
	"Unofficial_API/tools/updater"
	"encoding/json"
	"os"
	"sync"
)

type Fields struct {
	generateApi    bool
	generateModel  bool
	generateRouter bool

	generateWow     bool
	generateClassic bool
	generateD3      bool
	generateHS      bool
	generateSC2     bool

	writeToLocal bool
}

func updateApi(f *Fields) {

	type apiTask struct {
		Game     string
		Category string
		URL      string
		Enabled  bool
		Result   []*updater.ApiGroup
	}

	tasks := []*apiTask{
		// Wow Retail
		{"wow", "DataService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/game-data-apis.json", f.generateWow, nil},
		{"wow", "ProfileService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/profile-apis.json", f.generateWow, nil},
		// Wow Classic
		{"wowClassic", "DataService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/game-data-apis.json", f.generateClassic, nil},
		{"wowClassic", "ProfileService", "https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/profile-apis.json", f.generateClassic, nil},
		// Diablo 3
		{"D3", "Community", "https://develop.battle.net/api/pages/content/documentation/diablo-3/community-apis.json", f.generateD3, nil},
		{"D3", "DataService", "https://develop.battle.net/api/pages/content/documentation/diablo-3/game-data-apis.json", f.generateD3, nil},
		// Hearthstone
		{"HeartStone", "DataService", "https://develop.battle.net/api/pages/content/documentation/hearthstone/game-data-apis.json", f.generateHS, nil},
		// StarCraft II
		{"StarCraftII", "Community", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/community-apis.json", f.generateSC2, nil},
		{"StarCraftII", "DataService", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/game-data-apis.json", f.generateSC2, nil},
	}

	var wg sync.WaitGroup
	for i := range tasks {
		wg.Add(1)
		go func(t *apiTask) {
			defer wg.Done()
			t.Result = updater.RequestApiList(t.Game, t.URL)
		}(tasks[i])

	}
	wg.Wait()

	for i := range tasks {
		for j := range tasks[i].Result {
			tasks[i].Result[j].Fixed()
		}
	}

	collection := make([]*updater.ApiGroup, 0)

	for _, task := range tasks {
		if task.Result != nil {
			collection = append(collection, task.Result...)
		}

		outPath := "./api/" + task.Game + "/" + task.Category + "/"

		if f.generateApi && task.Enabled {
			updater.GenerateApi(task.Game, outPath, task.Result)
		}

		if f.generateModel && task.Enabled {
			updater.GenerateModels(task.Game, outPath, task.Result)
		}
	}

	if f.generateRouter {
		updater.GenerateRouters("routers", "./routers/", collection)
	}

	if f.writeToLocal {
		buf, _ := json.MarshalIndent(collection, "", "  ")
		os.WriteFile("./api_collection.json", buf, 0644)
	}

}

func main() {
	updateApi(&Fields{
		generateApi:    true,
		generateModel:  true,
		generateRouter: true,

		generateWow:     true,
		generateClassic: false,
		generateD3:      false,
		generateHS:      false,
		generateSC2:     false,

		writeToLocal: true,
	})
}
