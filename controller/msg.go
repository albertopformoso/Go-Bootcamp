package controller

import "fmt"

type Msg struct {
	Msg string `json:"message"`
}

type allMsgs []Msg

func Message(msg string) (myMessage allMsgs) {
	myMessage = allMsgs{
		{
			Msg: msg,
		},
	}
	return
}

func ErrMessage(m string, err error) (myMessage allMsgs) {
	msg := fmt.Sprintf("%v %v", m, err)
	myMessage = allMsgs{
		{
			Msg: msg,
		},
	}
	return
}
