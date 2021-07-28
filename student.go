package main

type Student struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Age              int    `json:"age"`
	FavouriteSubject string `json:"favouriteSubject"`
}
