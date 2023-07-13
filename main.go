package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5527719707623-5551986331986-8PExriTjyrOAcYWYw55W4qN8")
	os.Setenv("CHANNEL_ID", "C05G0LQPH36")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"upload-file-example.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s\n", file.Name)
	}
}
