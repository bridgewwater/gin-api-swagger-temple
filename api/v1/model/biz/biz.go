package biz

type Biz struct {
	Info   string `json:"info,omitempty" example:"input info here" default:"info"`
	Id     string `json:"id,omitempty"  example:"id123zqqeeadg24qasd" default:"id123zqqeeadg24qasd"`
	Offset int    `json:"offset,omitempty" example:"0" default:"0"`
	Limit  int    `json:"limit,omitempty" example:"10" default:"10"`
}
