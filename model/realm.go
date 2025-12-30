package model

type Realm struct {
	Id     int `json:"id"`
	Region struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"region"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	Type     struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"type"`
	IsTournament bool   `json:"is_tournament"`
	Slug         string `json:"slug"`
}
type ConnectedRealm struct {
	/*	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`*/ //I can not implement this field

	Id       int  `json:"id"`
	HasQueue bool `json:"has_queue"`
	Status   struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"status"`
	Population struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"population"`
	Realms []Realm `json:"realms"`

	// These fields are not implemented
	MythicLeaderboards struct {
		Href string `json:"href"`
	} `json:"mythic_leaderboards"`
	Auctions struct {
		Href string `json:"href"`
	} `json:"auctions"`
}
