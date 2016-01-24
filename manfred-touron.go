package manfredtouron

type profile struct {
	Provider string
	Handle   string
	URL      string
}

type organization struct {
	Name     string
	URL      string
	Position string
}

type Person struct {
	Firstname     string
	Lastname      string
	Fullname      string
	Nickname      string
	Company       string
	Homepage      string
	Location      string
	Headline      string
	Emoji         string
	Organizations map[string]organization
	Profiles      map[string]profile
}

var Manfred = Person{
	Firstname: "Manfred",
	Lastname:  "Touron",
	Fullname:  "Manfred Touron",
	Nickname:  "moul",
	Company:   "Scaleway",
	Homepage:  "http://m.42.am/",
	Location:  "Rouen, France / Paris, France",
	Headline:  "For passion, madness and glory",
	Emoji:     "👌",
}

func init() {
	Manfred.Organizations = map[string]organization{
		"scaleway": {
			Name:     "Scaleway",
			URL:      "https://www.scaleway.com/",
			Position: "CTO",
		},
		"pathwar": {
			Name:     "Pathwar",
			URL:      "https://www.pathwar.net/",
			Position: "Co-founder",
		},
		"while42": {
			Name:     "while42",
			URL:      "http://while42.org",
			Position: "Paris Staff",
		},
		"epitech": {
			Name:     "Epitech",
			URL:      "http://www.epitech.eu/",
			Position: "Alumnus",
		},
		"onlinenet": {
			Name:     "Online.net",
			URL:      "https://www.online.net/en",
			Position: "Collaborator",
		},
		"camembertaulaitcrew": {
			Name:     "Camembert au lait crew",
			URL:      "http://soundcloud.com/camembert-au-lait-crew",
			Position: "Crew member",
		},
	}
	Manfred.Profiles = map[string]profile{
		"github": {
			Provider: "GitHub",
			Handle:   "moul",
			URL:      "https://github.com/moul",
		},
		"gravatar": {
			Provider: "Gravatar",
			Handle:   "da14d5cef42c8142d3d40286f28f29bd",
			URL:      "https://www.gravatar.com/avatar/da14d5cef42c8142d3d40286f28f29bd?s=800",
		},
		"twitter": {
			Provider: "Twitter",
			Handle:   "moul",
			URL:      "https://twitter.com/moul",
		},
		"flickr": {
			Provider: "Flickr",
			Handle:   "38994875@N06",
			URL:      "https://www.flickr.com/photos/38994875@N06/",
		},
		"keybase": {
			Provider: "Keybase",
			Handle:   "moul",
			URL:      "https://keybase.io/moul",
		},
		"googleplus": {
			Provider: "Google+",
			Handle:   "ManfredTouron",
			URL:      "https://plus.google.com/u/0/+ManfredTouron",
		},
		"instagram": {
			Provider: "Instagram",
			Handle:   "m42am",
			URL:      "https://www.instagram.com/m42am/",
		},
	}
}
