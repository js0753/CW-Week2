package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JiraTestTeam struct {
	Total float64 `json:"total"`
	//Total2 float64
	Issues []JiraIssue `json:"issues"`
}

type JiraIssue struct {
	Id     string               `json:"id"`
	Key    string               `json:"key"`
	Fields map[string]JiraField `json:"fields"`
}

type JiraField struct {
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
}

func readJsonData(filename string) JiraTestTeam {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// opt1 - log err and return call to newDeck()
		// opt2 - log err and quit program
		fmt.Println("ERROR:", err)
		os.Exit(1)

	}
	// var tempJsonData interface{}
	var jtt JiraTestTeam
	json.Unmarshal(bs, &jtt)
	fmt.Println(jtt.Issues[0].Id)
	// m := tempJsonData.(map[string]interface{})
	// var jtt JiraTestTeam
	// jtt.total = m["total"].(float64)
	// im := m["issues"].([]interface{})
	// //fmt.Println(im)
	// for _, x := range im {
	// 	tx := x.(map[string]interface{})
	// 	var tempIssue JiraIssue
	// 	tempIssue.id = tx["id"].(string)
	// 	tempIssue.key = tx["key"].(string)
	// 	fi := tx["fields"].(map[string]interface{})
	// 	for _, y := range fi {
	// 		ty := y.([]interface{})
	// 		fmt.Println(ty)
	// 	}
	// }
	return jtt
}

func jsonFromObj(jtt JiraTestTeam) {
	b, err := json.Marshal(jtt)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("exported.json", b, 0666)
}
