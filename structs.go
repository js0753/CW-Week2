package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JIRATestTeam struct {
	Total  float64     `json:"total"`
	Issues []JIRAIssue `json:"issues"`
}

type JIRAIssue struct {
	Id     string     `json:"id"`
	Key    string     `json:"key"`
	Fields JIRAFields `json:"fields"`
}

type JIRAFields struct {
	EngineeringManager []JIRAUser `json:"customfield_14526"`
	ScrumMaster        []JIRAUser `json:"customfield_14517"`
	Developer          []JIRAUser `json:"customfield_14518"`
	ProductOwners      []JIRAUser `json:"customfield_14516"`
	QualityEngineer    []JIRAUser `json:"customfield_14519"`
	ProductManager     JIRAUser   `json:"customfield_14504"`

	Summary string `json:"summary"`
}

type JIRAUser struct {
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
}

func readJsonData(filename string) JIRATestTeam {
	jsonFileIp, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	var jtt JIRATestTeam
	json.Unmarshal(jsonFileIp, &jtt)
	return jtt
}

func jsonFromObj(jtt JIRATestTeam) {
	b, err := json.Marshal(jtt)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("exported.json", b, 0666)
}

/*
# JIRA fields mapping required:
"customfield_14526": "Engineering Managers",
"customfield_14517": "Scrum Master(s)",
"customfield_14518": "Developers",
"customfield_14516": "Product Owner(s)",
"customfield_14519": "Quality Engineers",
"customfield_14504": "Product Manager",
"summary": "Summary"
*/
