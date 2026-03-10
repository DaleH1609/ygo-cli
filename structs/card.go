package structs

type Card struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Attribute string `json:"attribute"`
	ATK       int    `json:"atk"`
	DEF       int    `json:"def"`
	Desc      string `json:"desc"`
}

type CardResponse struct {
	Data []Card `json:"data"`
}
