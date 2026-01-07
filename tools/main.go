package main

import (
	"Unofficial_API/tools/updater"
	"encoding/json"
	"os"
	"sync"
)

func updateApi() {
	var wg sync.WaitGroup
	wg.Add(5)

	// retail version
	/*go*/

	var wowdataServices, wowprofileServices []*updater.ApiGroup
	var classicdataServices, classicprofileServices []*updater.ApiGroup
	var D3communityServices, D3dataServices []*updater.ApiGroup
	var HSdataServices []*updater.ApiGroup
	var SCcommunityServices, SCdataServices []*updater.ApiGroup

	func() {
		defer wg.Done()
		wowdataServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/game-data-apis.json")
		wowprofileServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/world-of-warcraft/profile-apis.json")

		//updater.GenerateApi("wow", "./api/wow/DataService/", wowdataServices)
		//updater.GenerateApi("wow", "./api/wow/ProfileService/", wowprofileServices)
		//
		//updater.GenerateModels("wow", "./api/wow/Models/", append(wowdataServices, wowprofileServices...))
	}()

	// Wow Classic version
	go func() {
		defer wg.Done()
		classicdataServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/game-data-apis.json")
		classicprofileServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/world-of-warcraft-classic/profile-apis.json")

		//updater.GenerateApi("wowClassic", "./api/wowClassic/DataService/", classicdataServices)
		//updater.GenerateApi("wowClassic", "./api/wowClassic/ProfileService/", classicprofileServices)

	}()
	// D3 support
	go func() {
		defer wg.Done()
		D3communityServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/diablo-3/community-apis.json")
		D3dataServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/diablo-3/game-data-apis.json")

		//updater.GenerateApi("D3", "./api/D3/Community/", D3communityServices)
		//updater.GenerateApi("D3", "./api/D3/DataService/", D3dataServices)
		// CN data not supported
		//parseTemplate("D3", "./api/D3/Community/CN/", "https://develop.battle.net/api/pages/content/documentation/diablo-3/community-apis-cn.json")
	}()
	//HeartStone support
	go func() {
		defer wg.Done()
		HSdataServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/hearthstone/game-data-apis.json")
		//updater.GenerateApi("HeartStone", "./api/HeartStone/DataService/", HSdataServices)
	}()
	// StarCraftII support
	go func() {
		defer wg.Done()
		// CN data not supported
		//parseTemplate("SC", "./api/StarCraftII/Community/CN/", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/community-apis-cn.json")
		//updater.GenerateRouterList("SC", "./api/StarCraftII/Community/", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/community-apis.json")
		//updater.GenerateRouterList("SC", "./api/StarCraftII/DataService/", "https://develop.battle.net/api/pages/content/documentation/starcraft-2/game-data-apis.json")

		SCcommunityServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/starcraft-2/community-apis.json")
		SCdataServices = updater.RequestApiList("https://develop.battle.net/api/pages/content/documentation/starcraft-2/game-data-apis.json")

		/*updater.GenerateApi("StarCraftII", "./api/StarCraftII/Community/", SCcommunityServices)
		updater.GenerateApi("StarCraftII", "./api/StarCraftII/DataService/", SCdataServices)*/
	}()
	wg.Wait()

	collection := make([]*updater.ApiGroup, 0)
	collection = append(append(collection, wowprofileServices...), classicprofileServices...)
	collection = append(append(collection, wowdataServices...), classicdataServices...)
	collection = append(append(collection, D3communityServices...), D3dataServices...)
	collection = append(collection, HSdataServices...)
	collection = append(append(collection, SCcommunityServices...), SCdataServices...)

	updater.GenerateRouters("routers", "./routers/", collection)

	buf, _ := json.MarshalIndent(collection, "", "  ")
	os.WriteFile("./api_collection.json", buf, 0644)
}

func main() {
	updateApi()
}
