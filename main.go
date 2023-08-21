package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	xena "github.com/zarkones/xena-client"
)

const C2_HOST = "http://127.0.0.1:8080"
const C2_TIMEOUT = time.Minute

func main() {
	xena.Init(C2_HOST, C2_TIMEOUT)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "_unknown"
	}

	agentID := ""

	for {
		if agentID == "" {
			id, err := xena.Identify(hostname, runtime.GOOS, runtime.GOARCH)
			if err != nil {
				fmt.Println(err)
				continue
			}
			agentID = id
		}

		messages, err := xena.FetchMessages(agentID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, msg := range messages {
			output := "unknown message"

			// Add your message interpreter here.
			if msg.Request == "/ping" {
				output = "pong"
			}

			if err := xena.RespondToMessage(msg.ID, output); err != nil {
				fmt.Println(err)
				continue
			}
		}

		time.Sleep(time.Second * 10)
	}
}
