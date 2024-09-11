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
