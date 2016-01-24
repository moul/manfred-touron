package manfredtouron

type profile struct {
	Provider string `json:"provider"`
	Handle   string `json:"handle"`
	URL      string `json:"url"`
}

type organization struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Position string `json:"position"`
}

type Person struct {
	Firstname     string                  `json:"firstname"`
	Lastname      string                  `json:"lastname"`
	Fullname      string                  `json:"fullname"`
	Nickname      string                  `json:"nickname"`
	Company       string                  `json:"company"`
	Homepage      string                  `json:"homepage"`
	Blog          string                  `json:"blog"`
	Location      string                  `json:"location"`
	Headline      string                  `json:"headline"`
	Emoji         string                  `json:"emoji"`
	Organizations map[string]organization `json:"organizations"`
	Profiles      map[string]profile      `json:"profiles"`
	PGP           struct {
		Fingerprint  string `json:"fingerprint"`
		F64bit       string `json:"64bit"`
		F32bit       string `json:"32bit"`
		KeyAlgorithm string `json:"key-algorithm"`
		KeyLength    int    `json:"key-length"`
		Name         string `json:"name"`
		URL          string `json:"url"`
	} `json:"pgp"`
}

var Manfred = Person{
	Firstname: "Manfred",
	Lastname:  "Touron",
	Fullname:  "Manfred Touron",
	Nickname:  "moul",
	Company:   "Scaleway",
	Homepage:  "https://m.42.am/",
	Blog:      "https://manfredtouron.tumblr.com",
	Location:  "Rouen, France / Paris, France",
	Headline:  "For passion, madness and glory",
	Emoji:     "👌",
}

func init() {
	// PGP
	Manfred.PGP.Fingerprint = "0859942B474A2B3C90D380490DCB9CE0CABAE1B5"
	Manfred.PGP.F64bit = "0DCB9CE0CABAE1B5"
	Manfred.PGP.F32bit = "CABAE1B5"
	Manfred.PGP.KeyAlgorithm = "RSA"
	Manfred.PGP.KeyLength = 4096
	Manfred.PGP.Name = "4096R/CABAE1B5"
	Manfred.PGP.URL = "https://keybase.io/moul/key.asc"

	// Orgs
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

	// Social profiles
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
		"stackoverflow": {
			Provider: "StackOverflow",
			Handle:   "moul",
			URL:      "http://stackoverflow.com/users/1271690/moul",
		},
		"reddit": {
			Provider: "Reddit",
			Handle:   "manfred42",
			URL:      "https://reddit.com/user/manfred42",
		},
		"hackernews": {
			Provider: "HackerNews",
			Handle:   "moul",
			URL:      "https://news.ycombinator.com/user?id=moul",
		},
		"soundcloud": {
			Provider: "SoundCloud",
			Handle:   "manfred-touron",
			URL:      "https://soundcloud.com/manfred-touron",
		},
		"youtube": {
			Provider: "Youtube",
			Handle:   "m42am",
			URL:      "https://www.youtube.com/user/m42am",
		},
		"slideshare": {
			Provider: "Slideshare",
			Handle:   "manfredtouron",
			URL:      "http://www.slideshare.net/manfredtouron",
		},
		"linkedin": {
			Provider: "LinkedIn",
			Handle:   "manfredtouron",
			URL:      "https://fr.linkedin.com/in/manfredtouron",
		},
	}
}
