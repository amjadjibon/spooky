package cmd

import (
	"fmt"
	"github.com/amjadjibon/spooky/pkg/constant"
	"github.com/amjadjibon/spooky/pkg/db"
	"github.com/amjadjibon/spooky/pkg/dictionary"
	"github.com/amjadjibon/spooky/pkg/fakeit"
	"github.com/amjadjibon/spooky/pkg/ipapi"
	"github.com/amjadjibon/spooky/pkg/spooky"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func Run() {
	log.Debug("spooky cli app")
	app := cli.App{
		Name:  "spooky",
		Usage: "Spooky",
		Action: func(context *cli.Context) error {
			fmt.Println("Spooky Command Line Application")
			return cli.ShowAppHelp(context)
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Usage:   "Shows the version of running spooky",
				Aliases: []string{"v"},
				Action: func(context *cli.Context) error {
					fmt.Println(constant.Version)
					return nil
				},
			},
			{
				Name:    "get",
				Usage:   "Get status code for webpages. Example: 'spooky get facebook.com google.com'",
				Aliases: []string{"g"},
				Action: func(context *cli.Context) error {
					return spooky.GetStatusCode(context)
				},
			},

			{
				Name:    "time",
				Usage:   "Time",
				Aliases: []string{"t"},
				Action: func(context *cli.Context) error {
					fmt.Println(time.Now())
					return nil
				},

				Subcommands: []*cli.Command{
					{
						Name: "utc",
						Action: func(context *cli.Context) error {
							fmt.Println(time.Now().UTC().Format(time.RFC3339Nano))
							return nil
						},
					},
				},
			},

			{
				Name:  "password",
				Usage: "Get password hashes and verify passwords with the hash",
				Subcommands: []*cli.Command{
					{
						Name: "hash",
						Usage: "Get password hashes",
						HelpName: "hash",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "algorithm", Aliases: []string{"a"}},
						},
						Action: func(context *cli.Context) error {
							return spooky.GenerateHashes(context)
						},
					},

					{
						Name: "check",
						Usage: "Verify password hashes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "password"},
							&cli.StringFlag{Name: "hash"},
						},
						Action: func(context *cli.Context) error {
							return spooky.CheckPassword(context)
						},
					},
				},
			},

			{
				Name: "dictionary",
				Usage: "Dictionary",
				Aliases: []string{"d"},
				Action: func(context *cli.Context) error {
					return dictionary.Dictionary(context)
				},
			},

			{
				Name: "ulid",
				Usage: "ULID",
				Action: func(context *cli.Context) error {
					return spooky.GetULID(context)
				},
			},

			{
				Name: "uuid",
				Usage: "UUID",
				Action: func(context *cli.Context) error {
					return spooky.GetUUID(context)
				},
			},
			{
				Name: "pubip",
				Usage: "My Public IP Address",
				Action: func(context *cli.Context) error {
					return spooky.GetPubIP(context)
				},
			},

			{
				Name: "ip",
				Usage: "My Local Machine IP Address",
				Action: func(context *cli.Context) error {
					return spooky.MyLocalIP(context)
				},
			},

			{
				Name: "infoip",
				Usage: "Information of IP",
				Action: func(context *cli.Context) error {
					return ipapi.GetIPInformation(context)
				},
			},

			{
				Name: "battery",
				Usage: "Shows battery information of the device",
				Action: func(context *cli.Context) error {
					return spooky.GetBatteryInfo(context)
				},
			},


			{
				Name: "fakeit",
				Usage: "Give Fake Data",
				Action: func(context *cli.Context) error {
					return fakeit.Fake(context)
				},
			},

			{
				Name: "postgres",
				Usage: "postgres",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "user"},
					&cli.StringFlag{Name: "pass"},
					&cli.StringFlag{Name: "host"},
					&cli.StringFlag{Name: "port"},
					&cli.StringFlag{Name: "db_name"},
					&cli.StringFlag{Name: "sslmode"},
				},

				Subcommands: []*cli.Command{
					{
						Name: "ping",
						Usage: "ping postgres database",
						Action: func(context *cli.Context) error {
							return db.PingPostgres(context)
						},
					},

					{
						Name: "version",
						Usage: "print postgres database version",
						Action: func(context *cli.Context) error {
							return db.PostgresVersion(context)
						},
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Error(err)
	}
}
