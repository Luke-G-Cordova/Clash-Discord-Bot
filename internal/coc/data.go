package coc

// Member structs
type Member struct {
	Tag      string `json:"tag"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	ExpLevel int    `json:"expLevel"`

	League struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		IconUrls struct {
			Small  string
			Tiny   string
			Medium string
		}
	} `json:"league"`

	Trophies            int `json:"trophies"`
	BuilderBaseTrophies int `json:"builderBaseTrophies"`
	VersusTrophies      int `json:"versusTrophies"`
	ClanRank            int `json:"clanRank"`
	PreviousClanRank    int `json:"previousClanRank"`
	Donations           int `json:"donations"`
	DonationsReceived   int `json:"donationsReceived"`

	PlayerHouse struct {
		Elements []struct {
			Type string `json:"type"`
			Id   int    `json:"id"`
		}
	} `json:"playerHouse"`

	BuilderBaseLeague struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"builderBaseLeague"`
}

// clan structs
type Clan struct {
	Tag         string `json:"tag"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`

	Location struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		IsCountry   bool   `json:"isCountry"`
		CountryCode string `json:"countryCode"`
	} `json:"location"`

	IsFamilyFriendly bool `json:"isFamilyFriendly"`

	BadgeUrls struct {
		Small  string `json:"small"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
	} `json:"badgeUrls"`

	ClanLevel             int `json:"clanLevel"`
	ClanPoints            int `json:"clanPoints"`
	ClanBuilderBasePoints int `json:"clanBuilderBasePoints"`
	ClanVersusPoints      int `json:"clanVersusPoints"`
	ClanCapitalPoints     int `json:"clanCapitalPoints"`

	CapitalLeague struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"capitalLeague"`

	RequiredTrophies int    `json:"requiredTrophies"`
	WarFrequency     string `json:"warFrequency"`
	WarWinStreak     int    `json:"warWinStreak"`
	WarWins          int    `json:"warWins"`
	WarTies          int    `json:"warTies"`
	WarLosses        int    `json:"warLosses"`
	IsWarLogPublic   bool   `json:"isWarLogPublic"`

	WarLeague struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"warLeague"`

	Members    int `json:"members"`
	MemberList []Member

	ChatLanguage struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		LanguageCode string `json:"languageCode"`
	} `json:"chatLanguage"`
}
