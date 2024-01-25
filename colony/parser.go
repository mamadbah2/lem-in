package colony

import (
	"lemin/models"
	"lemin/pkg"
	"os"
	"strconv"
	"strings"
)

func GetDataFromFile(path string) string {
	byte, err := os.ReadFile(path)
	pkg.HandleError("Ouverture Fichier", err)
	return string(byte)
}

func DecodeFile(contentFile string) models.File {
	var mainData models.File
	tableContent := strings.Split(contentFile, "\n")

	// Recuperation du nombre de fourmi
	nbreFourmi, err := strconv.Atoi(tableContent[0])
	pkg.HandleError("number of ants invalid", err)
	if nbreFourmi == 0 {
		pkg.HandleError("number of ants is null", err)
	}
	mainData.Number_of_ants = nbreFourmi

	// Recuperation des chambres
	var rooms []models.Room
	for i := 1; i < len(tableContent); i++ {
		var room models.Room
		line := strings.Split(tableContent[i], " ")
		if len(line) == 3 {
			x, err := strconv.Atoi(line[1])
			pkg.HandleError("bad coordinate", err)
			y, err := strconv.Atoi(line[1])
			pkg.HandleError("bad coordinate", err)
			room.Name = line[0]
			if tableContent[i-1] == "##start" {
				room.Level = "start"
			} else if tableContent[i-1] == "##end" {
				room.Level = "end"
			} else {
				room.Level = "content"
			} 
			room.X = x 
			room.Y = y
			rooms = append(rooms, room)
		}
	}
	mainData.TheRooms = rooms

	var links []models.Link
	for i := 1; i < len(tableContent); i++ {
		var link models.Link
		line := strings.Split(tableContent[i], "-")
		if len(line) == 2 {
			link.NodeA = line[0]
			link.NodeB = line[1]
			link.Archived = false
			links = append(links, link)
		}
	}
	mainData.TheLinks = links

	return mainData
}
