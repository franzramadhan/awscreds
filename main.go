package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:     "stscreds",
		Version:  "v0.2.0",
		Compiled: time.Now(),
		Usage:    "Set AWS Credential from custom STS provider",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Frans Caisar Ramadhan",
				Email: "frans.ramadhan@traveloka.com",
			},
		},
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "URL to retrieve STS credential.",
				EnvVars:  []string{"STS_API_URL"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "assumed_role_arn",
				Aliases:  []string{"r", "role_arn"},
				Usage:    "ARN of assumed IAM Role.",
				EnvVars:  []string{"ASSUMED_ROLE_ARN"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "external_id",
				Aliases: []string{"e"},
				Usage:   "External ID configured for assumed IAM Role.",
				EnvVars: []string{"STS_EXTERNAL_ID"},
			},
			&cli.IntFlag{
				Name:    "duration",
				Aliases: []string{"d", "token_duration"},
				Value:   3600,
				Usage:   "Token duration in seconds.",
				EnvVars: []string{"STS_TOKEN_DURATION"},
			},
			&cli.IntFlag{
				Name:    "expiration_window",
				Aliases: []string{"w", "exp", "expiry"},
				Value:   0,
				Usage:   "Expiry window for STS token duration.",
				EnvVars: []string{"STS_TOKEN_WINDOW"},
			},
		},
		Action: func(c *cli.Context) error {
			data, err := getCredential(c.String("url"), c.String("assumed_role_arn"), c.String("external_id"), c.Int("duration"), c.Int("expiry_window"))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(data))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
