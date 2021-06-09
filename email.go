package main




import (
	"fmt"
	"os/exec"
	// "log"

	// gomail "gopkg.in/gomail.v2"
)

// func sendMail(from string, password string, to string) {
// 	mailer := gomail.NewMessage()
// 	mailer.SetHeader("From", from)
// 	mailer.SetHeader("To", to)
// 	mailer.SetHeader("Subject", "Mail from Go")
// 	mailer.SetBody("text/plain", "PFA")
// 	mailer.Attach("./exported.json")
// 	dialer := gomail.NewDialer("smtp.gmail.com", 587, from, password)
// 	if err := dialer.DialAndSend(mailer); err != nil {
// 		fmt.Println(err)
// 		log.Fatalln("Error occured during sendig mail:", err)
// 	} else {
// 		fmt.Println("Mail sent successfully")
// 	}
// }

func sendMail(recipient string){
//command = " mailx -s \"Exported Json File\" -A exported.json  "+recipient
cmd := exec.Command("mailx","-s","\"Exported Json File\"","-A","exported.json",recipient)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
