package gonavitia

type PtDisplayInfo struct {
	Direction      string   `json:"direction"`
	Code           string   `json:"code"`
	Network        string   `json:"network"`
	Color          string   `json:"color"`
	Name           string   `json:"name"`
	Label          string   `json:"label"`
	TextColor      string   `json:"text_color"`
	CommercialMode string   `json:"commercial_mode"`
	Links          []Link   `json:"links"`
	Description    *string  `json:"description,omitempty"`
	PhysicalMode   *string  `json:"physical_mode,omitempty"`
	Headsign       *string  `json:"headsign,omitempty"`
	Headsigns      []string `json:"headsigns,omitempty"`
	Equipments     []string `json:"equipments,omitempty"`
}
