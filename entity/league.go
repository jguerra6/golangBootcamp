package entity

type League struct{
	Country string `json:"country"`
	Name string `json:"name"`
	Current_season_id int64 `json:"current_season_id"`
}