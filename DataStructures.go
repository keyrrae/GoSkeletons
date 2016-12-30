package main

type Album struct {
	Id        int    `json:"id" sql:"id"`
	Title     string `json:"title" sql:"title"`
	Artist    string `json:"artist" sql:"artist"`
	Url       string `json:"url" sql:"user"`
	Image     string `json:"image" sql:"image"`
	Thumbnail string `json:"thumbnail_image" sql:"thumbnail"`
}
