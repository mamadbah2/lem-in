package colony

import (
	"lemin/models"
	"strings"
)

func FindPaths(data *models.File) []string {
	tabTemp := make(map[string][]string)
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

	var allPaths []string
	strTemp := []string{}
	for j := 0; j < len(data.TheLinks); j++ {
		if data.TheLinks[j].Archived == false {
			if data.TheLinks[j].NodeA == nodeStart {
				data.TheLinks[j].Archived = true
				snd := data.TheLinks[j].NodeB
				if data.TheLinks[j].NodeB == nodeEnd {
					allPaths = append(allPaths, nodeStart+"-"+snd)
				}
				strTemp = append(strTemp, snd)
			} else if data.TheLinks[j].NodeB == nodeStart {
				data.TheLinks[j].Archived = true
				snd := data.TheLinks[j].NodeA
				if data.TheLinks[j].NodeA == nodeEnd {
					allPaths = append(allPaths, nodeStart+"-"+snd)
				}
				strTemp = append(strTemp, snd)
			}
		}
	}

	tabTemp[nodeStart] = strTemp
	var inside string
	inside += strings.Join(strTemp, "")
	for i := 0; i < len(tabTemp[nodeStart]); i++ {
		inside += "/"
		finder(data, tabTemp[nodeStart][i], nodeEnd, &inside, 0)
	}

	insideTable := strings.Split(inside, "/")[1:]

	if len(insideTable) > 0 {
		for i := 0; i < len(tabTemp[nodeStart]); i++ {
			if strings.Contains(insideTable[i], "$") {
				trace := nodeStart + "-" + tabTemp[nodeStart][i] + "-" + insideTable[i][:len(insideTable[i])-1] + nodeEnd
				allPaths = append(allPaths, trace)
			}
		}
	}

	return allPaths
}

func finder(data *models.File, node, end string, inside *string, i int) {
	tabTemp := make(map[string][]string)
	strTemp := []string{}
	for j := 0; j < len(data.TheLinks); j++ {
		if data.TheLinks[j].Archived == false {
			if data.TheLinks[j].NodeA == node {
				if len(strTemp) == 0 {
					data.TheLinks[j].Archived = true
				}
				strTemp = append(strTemp, data.TheLinks[j].NodeB)
			} else if data.TheLinks[j].NodeB == node {
				if len(strTemp) == 0 {
					data.TheLinks[j].Archived = true
				}
				strTemp = append(strTemp, data.TheLinks[j].NodeA)
			}
		}
	}
	lastTemp := []string{}
	for k := 0; k < len(strTemp); k++ {
		if !strings.Contains(*inside, string(strTemp[k])) || strTemp[k] == end {
			lastTemp = append(lastTemp, strTemp[k])
		}
	}
	if strings.Contains(strings.Join(lastTemp, ""), end) {
		tabTemp = map[string][]string{}
		tabTemp[node] = []string{end + " is end"}
		*inside += "$"
		for _, lastnode := range lastTemp {
			for j := 0; j < len(data.TheLinks); j++ {
				if data.TheLinks[j].NodeA == lastnode {
					data.TheLinks[j].Archived = false
				} else if data.TheLinks[j].NodeB == lastnode {
					data.TheLinks[j].Archived = false
				}
			}
		}
		return
	}
	tabTemp[node] = lastTemp
	if len(lastTemp) > 0 {
		*inside += lastTemp[0] + "-"
	}

	for _, onNode := range tabTemp[node] {
		finder(data, onNode, end, inside, i+1)
		break
	}
}
