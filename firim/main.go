package main

import (
	"drone_firim"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var Version = "0.0.2"

func main() {
	app := cli.NewApp()
	app.Name = "Drone fir.im Plugin"
	app.Usage = "Push files to fir.im"
	app.Copyright = "© 2019 kingzcheung"
	app.Authors = []cli.Author{
		{
			Name:  "Kingz Cheung",
			Email: "i@kingzcheung.com",
		},
	}
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "type",
			Usage:  "ios 或者 android（发布新应用时必填）",
			EnvVar: "PLUGIN_TYPE",
		},
		cli.StringFlag{
			Name:   "bundle_id",
			Usage:  "App 的 bundleId（发布新应用时必填）",
			EnvVar: "PLUGIN_BUNDLE_ID",
		},
		cli.StringFlag{
			Name:   "api_token",
			Usage:  "长度为 32, 用户在 fir 的 api_token",
			EnvVar: "PLUGIN_API_TOKEN",
		},
		cli.StringFlag{
			Name:   "app.file",
			Usage:  "安装包文件",
			EnvVar: "PLUGIN_FILE",
		},
		cli.StringFlag{
			Name:   "app.name",
			Usage:  "应用名称（上传 ICON 时不需要）",
			EnvVar: "PLUGIN_NAME",
		},
		cli.StringFlag{
			Name:   "app.version",
			Usage:  "版本号（上传 ICON 时不需要）",
			EnvVar: "PLUGIN_VERSION",
		},
		cli.StringFlag{
			Name:   "app.build",
			Usage:  "Build 号（上传 ICON 时不需要）",
			EnvVar: "PLUGIN_BUILD",
		},
		cli.StringFlag{
			Name:   "app.release_type",
			Usage:  "打包类型，只针对 iOS (Adhoc, Inhouse)（上传 ICON 时不需要）",
			EnvVar: "PLUGIN_RELEASE_TYPE",
		},
		cli.StringFlag{
			Name:   "app.changelog",
			Usage:  "更新日志（上传 ICON 时不需要）",
			EnvVar: "PLUGIN_CHANGELOG",
		},
	}
	app.Action = run

	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}

func run(c *cli.Context) {
	plugin := &drone_firim.Plugin{
		Firim: drone_firim.NewFirim(
			c.String("type"),
			c.String("bundle_id"),
			c.String("api_token"),
			c.String("app.file"),
			c.String("app.name"),
			c.String("app.version"),
			c.String("app.build"),
			c.String("app.release_type"),
			c.String("app.changelog"),
		),
	}

	err := plugin.Firim.Exec()

	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("上传 fir.im 成功")
}
