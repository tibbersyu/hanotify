package main

import (
	"flag"
	"fmt"
	"gopkg.in/chanxuehong/wechat.v1/corp"
	"gopkg.in/chanxuehong/wechat.v1/corp/message/send"
)

var (
	title     = flag.String("t", "", "title")
	receivers = flag.String("r", "", "receivers")
	message   = flag.String("m", "", "message")
	source    = flag.String("s", "", "from_source")
)

var AccessTokenServer = corp.NewDefaultAccessTokenServer("", "", nil)
var corpClient = corp.NewClient(AccessTokenServer, nil)

func main() {
	flag.Parse()

	sendClient := (*send.Client)(corpClient)

	header := send.MessageHeader{
		AgentId: 1000002,
		ToParty: "1",
	}

	if title != nil && *title != "" && message != nil && *message != "" {
		header.MsgType = send.MsgTypeNews
		news := []send.NewsArticle{
			{
				Title:       *title,
				Description: *message,
				URL:         "https://ha.tibbers.cc:8443/",
			},
		}
		msg := &send.News{
			MessageHeader: header,
			News: struct {
				Articles []send.NewsArticle `json:"articles,omitempty"`
			}{
				news,
			},
		}
		_, err := sendClient.SendNews(msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}
	} else if title != nil && *title != "" || message != nil && *message != "" {
		header.MsgType = send.MsgTypeText
		msg := &send.Text{
			MessageHeader: header,
		}
		if title != nil && *title != "" {
			msg.Text.Content = *title
		} else {
			msg.Text.Content = *message
		}
		_, err := sendClient.SendText(msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}
	} else {
		fmt.Println("param error")
	}
}
