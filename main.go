package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

func main() {
	var (
		//port      = os.Getenv("PORT")
		//publicURL = os.Getenv("PUBLIC_URL")
		token = "810595026:AAFZ2BcIbEOnPw9eyAIkoDSd2VPcJRyRnZ4"
	)

	//webhook := &tb.Webhook{
	//	Listen:   ":" + port,
	//	Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	//}

	pref := tb.Settings{
		Token: token,
		//Poller: webhook,
	}

	_, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	//b.Handle("/hello", func(m *tb.Message) {
	//	b.Send(m.Sender, "Hi!")
	//})

}
