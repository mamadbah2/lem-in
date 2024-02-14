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
	if err != nil {
		pkg.HandleError("Ouverture Fichier", err)
	}
	return string(byte)
}

func DecodeFile(contentFile string) models.File {
	var mainData models.File
	tableContentBrute := strings.Split(contentFile, "\n")
	var tableContent []string
	for _, content := range tableContentBrute {
		if !strings.Contains(content, "#") || content == "##start" || content == "##end" {
			tableContent = append(tableContent, content)
		}
	}

	for _, content := range tableContentBrute {
		cas1 := len(strings.Split(content, " "))
		cas2 := len(strings.Split(content, "-"))
		if cas1 != 1 && cas1 != 3 || cas2 != 2 && cas2 != 1 {
			pkg.HandleError("Error in the data format", nil)
		}
	}

	// Recuperation du nombre de fourmi
	nbreFourmi, err := strconv.Atoi(tableContent[0])
	if nbreFourmi == 0 {
		pkg.HandleError("number of ants is null", err)
	}
	if err != nil {
		pkg.HandleError("number of ants ", err)
	}
	mainData.Number_of_ants = nbreFourmi

	// Recuperation des chambres
	var (
		countStart, countEnd, countContent int
	)
	var rooms []models.Room
	for i := 1; i < len(tableContent); i++ {
		var room models.Room
		line := strings.Split(tableContent[i], " ")
		if len(line) == 3 {
			x, err := strconv.Atoi(line[1])
			if err != nil {
				pkg.HandleError("bad coordinate", err)
			}
			y, err := strconv.Atoi(line[2])
			if err != nil {
				pkg.HandleError("bad coordinate", err)
			}
			room.Name = line[0]
			for _, roomAdded := range rooms {
				if roomAdded.Name == room.Name {
					pkg.HandleError("bad room", nil)
				}
			}
			if tableContent[i-1] == "##start" {
				countStart++
				room.Level = "start"
			} else if tableContent[i-1] == "##end" {
				countEnd++
				room.Level = "end"
			} else {
				countContent++
				room.Level = "content"
			}
			room.X = x
			room.Y = y
			rooms = append(rooms, room)
		}
	}
	if countStart != 1 || countEnd != 1 || countContent == 0 {
		pkg.HandleError("No good start or end or content", nil)
	}

	mainData.TheRooms = rooms

	var links []models.Link
	for i := 1; i < len(tableContent); i++ {
		var link models.Link
		line := strings.Split(tableContent[i], "-")
		if len(line) == 2 {
			if line[0] == line[1] {
				pkg.HandleError("Link room error", nil)
			}
			if !isPresentRoom(mainData.TheRooms, line[0]) || !isPresentRoom(mainData.TheRooms, line[1]) {
				pkg.HandleError("invalid room", nil)
			}
			link.NodeA = line[0]
			link.NodeB = line[1]
			link.Archived = false
			links = append(links, link)
		}
	}
	mainData.TheLinks = links

	return mainData
}

func isPresentRoom(tabRoom []models.Room, nodeName string) bool {
	for _, registNode := range tabRoom {
		if registNode.Name == nodeName {
			return true
		}
	}
	return false
}
