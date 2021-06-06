package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JiraTestTeam struct {
	Total  float64     `json:"total"`
	Issues []JiraIssue `json:"issues"`
}

type JiraIssue struct {
	Id     string                 `json:"id"`
	Key    string                 `json:"key"`
	Fields map[string][]JiraField `json:"fields"`
}

type JiraField struct {
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
}

func readJsonData(filename string) JiraTestTeam {
	jsonFileIp, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	var jtt JiraTestTeam
	json.Unmarshal(jsonFileIp, &jtt)
	return jtt
}

func jsonFromObj(jtt JiraTestTeam) {
	b, err := json.Marshal(jtt)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("exported.json", b, 0666)
}
