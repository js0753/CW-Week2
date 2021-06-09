package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	jtt := readJsonData("jiraTestTeam.json")
	jsonFromObj(jtt)
	fmt.Println(jtt)
	var sender string
	fmt.Print("\nEnter Sender :")
	fmt.Scanln(&sender)
	fmt.Print("Enter List of recipient mail(space separated) :")
	scanner.Scan()
	recipients:=strings.Split(scanner.Text(), " ")
	sendMail(sender,recipients)
}
