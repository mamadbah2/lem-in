package colony

import (
	"fmt"
	"lemin/models"
)

func FindPaths(data *models.File) {
	mapLevel := make(map[string][]string)
	var nodeStart, nodeEnd string
	for _, room := range data.TheRooms {
		if room.Level == "start" {
			nodeStart = room.Name
		} else if room.Level == "end" {
			nodeEnd = room.Name
		}
		if nodeEnd != "" && nodeStart != "" {
			break
		}
	}
	fmt.Println("node Start: ", nodeStart)
	fmt.Println("node End: ", nodeEnd)
	var temporalLevel []string
	for i := 0; i < len(data.TheLinks); i++ {
		if data.TheLinks[i].Archived == false {
			if data.TheLinks[i].NodeA == nodeStart {
				temporalLevel = append(temporalLevel, data.TheLinks[i].NodeB)
				data.TheLinks[i].Archived = true
			} else if data.TheLinks[i].NodeB == nodeStart {
				temporalLevel = append(temporalLevel, data.TheLinks[i].NodeA)
				data.TheLinks[i].Archived = true
			}
		}
	}
	mapLevel[nodeStart] = temporalLevel
	fmt.Println("map ap : ", mapLevel)
	// fmt.Println("lev. : ", data.TheLinks)
	// fmt.Println("len : ", len(data.TheLinks))
	solveMap(mapLevel, nodeStart, data, 0)
}

func solveMap(mapLevel map[string][]string, nodePush string, data *models.File, i int) {
	// if i == 7 {
	// 	fmt.Println("level : ", data.TheLinks)
	// 	return
	// }
	for _, node := range mapLevel[nodePush] {
		temporalLevel := []string{}
		for i := 0; i < len(data.TheLinks); i++ {
			if data.TheLinks[i].Archived == false {
				if data.TheLinks[i].NodeA == node {
					temporalLevel = append(temporalLevel, data.TheLinks[i].NodeB)
					data.TheLinks[i].Archived = true
				} else if data.TheLinks[i].NodeB == node {
					temporalLevel = append(temporalLevel, data.TheLinks[i].NodeA)
					data.TheLinks[i].Archived = true
				}
			}
		}
		mapLevel[node] = temporalLevel
	}
	for _, node := range mapLevel[nodePush] {
		fmt.Println(mapLevel)
		solveMap(mapLevel, node, data, i+1)
	}
}
