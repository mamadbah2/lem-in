package models

type File struct {
	Number_of_ants int
	TheRooms []Room
	TheLinks []Link
}

type Room struct {
	Name string
	Level string
	X, Y int
}

type Link struct {
	NodeA string
	NodeB string
	Archived bool
}