package models

type Reg struct {
	Name string            `json:"name"`
	Data map[string]uint16 `json:"data"`
}

type Config struct {
	NameConfig string `json:"nameConfig"`
	IP         string `json:"ip"`
	Reg        []Reg  `json:"reg"`
}
