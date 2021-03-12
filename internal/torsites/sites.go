package torsites

type TorSites struct {
	Sites []site `toml:"site"`
}

type site struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}

func (ts *TorSites) Name(name string) *site {
	for _, s := range ts.Sites {
		if s.Name == name {
			return &s
		}
	}
	return nil
}
