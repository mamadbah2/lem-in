package colony

import (
	"fmt"
	"sort"
	"strings"
)

type Chemin struct {
	Name             string
	NombreDeChambres int
	NombreDeFourmis  int
}

func maxLenInTable(table [][]string, ind int) int {
	var maxLen []int
	for i := 0; i < len(table); i++ {
		maxLen = append(maxLen, len(table[i]))
	}
	sort.Ints(maxLen)
	var result []int
	for i := 0; i < len(maxLen); i++ {
		if i == 0 || maxLen[i] != maxLen[i-1] {
			result = append(result, maxLen[i])
		}
	}
	if ind >= len(result) {
		ind = len(result) - 1
	}
	return result[ind]
}

func NumberOfAntsByPaths(paths []string, numberOfAnts int) []int {
	var matricePaths [][]string
	for _, path := range paths {
		tempTable := strings.Split(path, "-")
		matricePaths = append(matricePaths, tempTable)
	}

	tabTail := make([]int, len(matricePaths))
	maxLenInTable(matricePaths, 25)
	j := 1
	for {
		count := 0
		for i := 0; i < len(matricePaths); i++ {
			if maxLenInTable(matricePaths, j) <= (len(matricePaths[i]) + tabTail[i]) {
				count++
				continue
			}
			tabTail[i] += 1
			numberOfAnts -= 1
			if numberOfAnts == 0 {
				return tabTail
			}
		}
		if count == len(matricePaths) {
			j++
		}
		if j == len(matricePaths) {
			break
		}
	}
	if j == len(matricePaths) {
		rest := numberOfAnts / len(matricePaths)
		for i := 0; i < len(matricePaths); i++ {
			tabTail[i] += rest
			numberOfAnts -= rest
		}
		if numberOfAnts == 0 {
			return tabTail
		}
		for i := 0; i < len(matricePaths); i++ {
			tabTail[i] += 1
			numberOfAnts -= 1
			if numberOfAnts == 0 {
				return tabTail
			}
		}
	}

	return tabTail
}

func Move(chemins []Chemin) {
	k := 1

	tabs := [][][]string{}

	for _, v := range chemins {
		chambre := strings.Split(v.Name, "-")

		tab := make([][]string, len(chambre)+v.NombreDeFourmis+10)
		for i := range tab {
			tab[i] = make([]string, v.NombreDeFourmis+10)
		}

		i, j := 0, 0
		for v.NombreDeFourmis != 0 {
			j = i
			for _, c := range chambre {
				tab[j][i] = fmt.Sprintf("L%v-%v ", k, c)
				j++
			}
			k++
			i++
			v.NombreDeFourmis--
		}

		tabs = append(tabs, tab)
	}

	Associate(tabs)
}

func Associate(tab [][][]string) {
	maxLignes := 0
	for _, tab2D := range tab {
		if len(tab2D) > maxLignes {
			maxLignes = len(tab2D)
		}
	}

	r := ""

	for i := 0; i < maxLignes; i++ {
		for _, tab2D := range tab {
			if i < len(tab2D) {
				for _, element := range tab2D[i] {
					r += fmt.Sprint(element)
				}
			}
		}
		r += fmt.Sprintln()
	}

	fmt.Println(strings.TrimSpace(r))
}
