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
