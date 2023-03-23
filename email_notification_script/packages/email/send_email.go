package email

import (
	"fmt"
	"strings"

	"github.com/mailjet/mailjet-apiv3-go"
)
var api_key = "cbf937e1765fff26bc4653408cf71bf0"
var secret_key = "990858a423e816c5bc884fb060eb36eb"
var recipient = "aryapb2017@gmail.com"
var sender = "arya.pb@aptonworks.com"

func SendMailForDown(result string) {
	mj := mailjet.NewMailjetClient(api_key, secret_key)
	message := &mailjet.InfoSendMail{
		FromEmail: sender,
		Recipients: []mailjet.Recipient{
			mailjet.Recipient{
				Email: recipient,
			},
		},
			Subject:  "Deployment Down",
			TextPart: "The following deployments are down:\n",
			HTMLPart: "<h3>" + strings.Replace(result, "\n", "<br>", -1) + "</h3>",
	}

	// Send email
	_, err := mj.SendMail(message)

	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}

}

func SendMailForUp(result string) {
	mj := mailjet.NewMailjetClient(api_key, secret_key)
	message := &mailjet.InfoSendMail{
		FromEmail: sender,
		Recipients: []mailjet.Recipient{
			mailjet.Recipient{
				Email: recipient,
			},
		},
			Subject:  "Deployment up",
			TextPart: "The following deployments are up:\n",
			HTMLPart: "<h3>" + strings.Replace(result, "\n", "<br>", -1) + "</h3>",
	}

	// Send email
	_, err := mj.SendMail(message)

	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}

}

