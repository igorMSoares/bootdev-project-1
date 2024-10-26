package pokeapi

type Pokemon struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Experience int    `json:"base_experience"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Types      []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		Base int `json:"base_stat"`
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}
