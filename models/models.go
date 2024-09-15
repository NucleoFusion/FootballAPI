package models

type ClubData struct {
	Team         string  `json:"Team"`
	Tournament   string  `json:"Tournament"`
	Goals        int32   `json:"Goals"`
	ShotsPG      float64 `json:"ShotsPG"`
	Yellow_cards int32   `json:"Yellow_cards"`
	Red_cards    int32   `json:"Red_cards"`
	Possession   float64 `json:"Possession"`
	Pass         float64 `json:"Pass"`
	AerialsWon   float64 `json:"AerialsWon"`
	Rating       float64 `json:"Rating"`
}

type StadiumData struct {
	Confederation string `json:"Confederation"`
	Stadium       string `json:"Stadium"`
	City          string `json:"City"`
	HomeTeams     string `json:"HomeTeams"`
	Capacity      int32  `json:"Capacity"`
	Country       string `json:"Country"`
	Population    int32  `json:"Population"`
}

type UserData struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Email string `json:"Email"`
}

type PlayerData struct {
	Rk              int32   `json:"Rk"`
	Player          string  `json:"Player"`
	Nation          string  `json:"Nation"`
	Squad           string  `json:"Squad"`
	Pos             string  `json:"Pos"`
	Age             int32   `json:"Age"`
	Comp            string  `json:"Comp"`
	Born            int32   `json:"Born"`
	MP              int32   `json:"MP"`
	Starts          int32   `json:"Starts"`
	Min             int32   `json:"Min"`
	NinetyMinPlayed float64 `json:"90s"`
	Gls             int32   `json:"Gls"`
	Ast             int32   `json:"Ast"`
	GplusA          int32   `json:"G+A"`
	PK              int32   `json:"PK"`
	PKatt           int32   `json:"PKatt"`
	CrdY            int32   `json:"CrdY"`
	CrdR            int32   `json:"CrdR"`
	Gls_90          float64 `json:"Gls_90"`
	Ast_90          float64 `json:"Ast_90"`
}
