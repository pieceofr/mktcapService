package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/slack"
)

func slackSendService(message <-chan string, stop <-chan bool) {
	glog.Info("slackSendService has started")
	for {
		select {
		case send := <-message:
			slackSendMessage(send)
		case <-stop:
			break
		}
	}

}

func slackSendMessage(msg string) {
	slackapi := slack.New(srvConfig.SlackToken)
	glog.V(2).Infoln("slack is going to send - ", msg)
	msg = fmt.Sprintf("%s %s", "mktcap notify:", msg)
	_, _, err := slackapi.PostMessage(srvConfig.SlackChannel, msg, slack.PostMessageParameters{})
	if err != nil {
		glog.Error(err)
	}
}
