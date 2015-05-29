package entities

type Move struct {
	Name        string `json:"name"`
	Typing      string `json:"type"`
	Category    string `json:"category"`
	Power       int    `json:"power"`
	Accuracy    int    `json:"accuracy"`
	Pp          int    `json:"pp"`
	Description string `json:"description"`
}

type NameBased struct {
	Name string `json:"name"`
}

type Ability struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Pokemon struct {
	Name      string      `json:"name"`
	Abilities []Ability   `json:"abilities"`
	Attack    int         `json:"attack"`
	Defense   int         `json:"defense"`
	Speed     int         `json:"speed"`
	Hp        int         `json:"hp"`
	SpAtk     int         `json:"sp_atk"`
	SpDef     int         `json:"sp_def"`
	Total     int         `json:"total"`
	Types     []NameBased `json:"types"`
	Moves     []Move      `json:"moves"`
}
