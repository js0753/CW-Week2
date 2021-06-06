package main

func main() {
	jtt := readJsonData("jiraTestTeam.json")
	jsonFromObj(jtt)
	sendMail("dev.js0753@gmail.com", "devgmail55", "dev.js0753@gmail.com")
}
