package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/tjtjtjtj/go-slack/slack"
)

const (
	slackAPIURL = "https://slack.com/api"
)

func main() {
	app := cli.NewApp()
	app.Name = "slack cli"
	app.Usage = "cli for slack api"
	app.Version = "0.0.1"

	// グローバルオプション設定
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token, t",
			Value:  "slackapitoken",
			Usage:  "slack api token",
			EnvVar: "SLACK_API_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		// コマンド設定
		{
			Name:    "history",
			Aliases: []string{"h"},
			Usage:   "show channle messages",
			Action:  showMessages,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "channel, c", Value: "channleid"},
				cli.StringFlag{Name: "number, n", Value: "10"},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		// 開始前の処理をここに書く
		fmt.Println("開始")
		var err error
		slack.SlackClient, err = slack.NewClient(slackAPIURL, c.String("token"))
		if err != nil {
			return err
		}
		return nil // error を返すと処理全体が終了
	}

	app.After = func(c *cli.Context) error {
		// 終了時の処理をここに書く
		fmt.Println("終了")
		return nil
	}

	app.Run(os.Args)
}

func showMessages(c *cli.Context) error {
	fmt.Println("show massages")
	fmt.Printf("parm:%s,%v", c.String("channel"), c.String("number"))
	var cancel context.CancelFunc
	slack.Ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second) // 30秒後にキャンセル
	defer cancel()

	history, err := slack.SlackClient.GetChannlesHistory(slack.Ctx, c.String("channel"), c.String("number"))
	fmt.Printf("history:%v", history)
	fmt.Printf("err:%v", err)

	fmt.Printf("parm:%s,%v", c.String("channel"), c.String("number"))
	if err != nil {
		return err
	}
	fmt.Printf("history:%v", history)
	return nil
}
