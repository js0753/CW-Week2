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

func sendMail(sender string,recipients []string){
//command = " mailx -s \"Exported Json File\" -A exported.json  "+recipient
var cmd_args []string
if sender == ""{
	cmd_args=[]string{"-s","\"Exported Json File\"","-A","exported.json"}
}else{
cmd_args=[]string{"-s","\"Exported Json File\"","-A","exported.json","-r","\"JSON-Exporter<"+sender+">\""}
}
cmd_args=append(cmd_args,recipients...)
cmd := exec.Command("mailx",cmd_args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
