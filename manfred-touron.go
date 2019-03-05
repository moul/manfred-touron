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
	SSHPubKey string `json:"ssh-pub-key"`
}

var Manfred = Person{
	Firstname: "Manfred",
	Lastname:  "Touron",
	Fullname:  "Manfred Touron",
	Nickname:  "moul",
	Company:   "Berty",
	Homepage:  "https://manfredtouron.com/",
	Blog:      "https://manfredtouron.com/blog",
	Location:  "Paris, France",
	Headline:  "Passion & Madness",
	Emoji:     "👌",
	SSHPubKey: "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEApvPvDbWDY50Lsx4WyUInw407379iERte63OTTNae6+JgAeYsn52Z43Oeks/2qC0gxweq+sRY9ccqhfReie+r+mvl756T4G8lxX1ND8m6lZ9kM30Rvk0piZn3scF45spmLNzCNXza/Hagxy53P82ej2vq2ewXtjVdvW20G3cMHVLkcdgKJN+2s+UkSYlASW6enUj3no+bukT+6M8lJtlT0/0mZtnBRJtqCCvF0cm9xU0uxILrhIfdYAJ1XqaoqIQLFSDLVo5lILMzDNwV+CfAotRMWIKvWomCszhVQYHCQo2Z+b2Gs0TL4DRb23fRMdeaRufnVhh5ZMlNkb2ajaL6sw== m",
}

func init() {
	// PGP
	Manfred.PGP.Fingerprint = "4B2B95D630B5429595AC62286D4DED2EAB123456"
	Manfred.PGP.F64bit = "6D4DED2EAB123456"
	Manfred.PGP.F32bit = "AB123456"
	Manfred.PGP.KeyAlgorithm = "RSA"
	Manfred.PGP.KeyLength = 4096
	Manfred.PGP.Name = "4096R/AB123456"
	Manfred.PGP.URL = "https://manfredtouron.com/manfred.asc"

	// Orgs
	Manfred.Organizations = map[string]organization{
		"bertytech": {
			Name:     "Berty Technologies",
			URL:      "https://berty.tech",
			Position: "Founder",
		},
		"united-drivers": {
			Name:     "United Drivers",
			URL:      "https://www.united-drivers.org",
			Position: "Founder",
		},
		"ultreme": {
			Name:     "Ultreme",
			URL:      "https://ultre.me",
			Position: "Co-founder",
		},
		"scaleway": {
			Name:     "Scaleway",
			URL:      "https://www.scaleway.com/",
			Position: "CTO",
		},
		"pathwar": {
			Name:     "Pathwar",
			URL:      "https://www.pathwar.net/",
			Position: "Founder",
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
		"anjunabeats": {
			Name:     "Anjunabeats",
			URL:      "http://www.anjunabeats.com",
			Position: "IT staff",
		},
		"42am": {
			Name:     "42.am",
			URL:      "http://www.42.am",
			Position: "Co-founder",
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
		"angellist": {
			Provider: "AngelList",
			Handle:   "moul",
			URL:      "https://angel.co/moul",
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
		"dockerhub": {
			Provider: "Docker",
			Handle:   "moul",
			URL:      "https://hub.docker.com/u/moul/",
		},
		"meetup": {
			Provider: "Meetup",
			Handle:   "96602082",
			URL:      "http://www.meetup.com/members/96602082/",
		},
		"tumblr": {
			Provider: "Tumblr",
			Handle:   "manfredtouron",
			URL:      "https://manfredtouron.tumblr.com",
		},
		"aboutme": {
			Provider: "about.me",
			Handle:   "manfredtouron",
			URL:      "https://about.me/manfredtouron",
		},
		"couchsurfing": {
			Provider: "Couchsurfing",
			Handle:   "manfredtouron",
			URL:      "https://www.couchsurfing.com/people/manfredtouron",
		},
		"openhub": {
			Provider: "OpenHUB",
			Handle:   "moul",
			URL:      "https://www.openhub.net/accounts/moul",
		},
		"bountysource": {
			Provider: "Bountysource",
			Handle:   "moul",
			URL:      "https://www.bountysource.com/people/29827",
		},
		"bitbucket": {
			Provider: "Bitbucket",
			Handle:   "moul",
			URL:      "https://bitbucket.org/moul/",
		},
		"gratipay": {
			Provider: "Gratipay",
			Handle:   "moul",
			URL:      "https://gratipay.com/~moul/",
		},
		"openstreetmap": {
			Provider: "OpenStreetMap",
			Handle:   "moul",
			URL:      "http://www.openstreetmap.org/user/moul",
		},
		"coinbase": {
			Provider: "coinbase",
			Handle:   "moul",
			URL:      "https://coinbase.com/moul",
		},
		"gitlabcom": {
			Provider: "GitLab.com",
			Handle:   "moul",
			URL:      "https://gitlab.com/u/moul",
		},
		"wikipedia": {
			Provider: "Wikipedia",
			Handle:   "m42am",
			URL:      "https://fr.wikipedia.org/wiki/Utilisateur:M42am",
		},
		"kickstarter": {
			Provider: "Kickstarter",
			Handle:   "moul",
			URL:      "https://www.kickstarter.com/profile/moul",
		},
		"thingiverse": {
			Provider: "Thingiverse",
			Handle:   "moul",
			URL:      "http://www.thingiverse.com/moul/about",
		},
		"lemarathondelasemaine": {
			Provider: "Le marathon de la semaine",
			Handle:   "moul",
			URL:      "http://www.lemarathondelasemaine.fr/joggeur/view/33",
		},
		"openagenda": {
			Provider: "OpenAgenda",
			Handle:   "manfred",
			URL:      "https://openagenda.com/manfred",
		},
		"robozzle": {
			Provider: "RoboZZle",
			Handle:   "moul",
			URL:      "http://robozzle.com/user.aspx?name=moul",
		},
		"medium": {
			Provider: "Medium",
			Handle:   "moul",
			URL:      "https://medium.com/@moul",
		},
		// Spotify
	}
}
