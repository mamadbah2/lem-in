package main

import (
	"lemin/colony"
)

func main() {
	contentFile := colony.GetDataFromFile("./repository/example.txt")
	dataFile := colony.DecodeFile(contentFile)
	/* fmt.Println("number of ants : \n", dataFile.Number_of_ants)
	fmt.Println("All rooms : \n", dataFile.TheRooms)
	fmt.Println("All links : \n", dataFile.TheLinks) */
	colony.FindPaths(&dataFile)
}
