package main

import (
	"fmt"
	"lemin/colony"
	"strings"
)

func main() {
	contentFile := colony.GetDataFromFile("./repository/example.txt")
	fmt.Print(contentFile, "\n\n")
	dataFile := colony.DecodeFile(contentFile)
	paths := colony.FindPaths(&dataFile)
	if len(paths) == 0 {
		fmt.Println("no valid paths found")
		return
	}
	var matricePaths [][]string
	for _, path := range paths {
		tempTable := strings.Split(path, "-")
		matricePaths = append(matricePaths, tempTable)
	}
	var chemins []colony.Chemin
	for _, v := range matricePaths {
		chemins = append(chemins, colony.Chemin{Name: strings.Join(v[1:], "-"), NombreDeChambres: len(v), NombreDeFourmis: 0})
	}
	for i := 0; i < len(chemins); i++ {
		chemins[i].NombreDeFourmis = colony.NumberOfAntsByPaths(paths, dataFile.Number_of_ants)[i]
	}
	colony.Move(chemins)

}
