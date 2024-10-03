/*
 * Copyright (c) 2024 Ruiyuan "mizumoto-cn" Xu
 *
 * This file is part of "github.com/mizumoto-cn/mizubot".
 *
 * Licensed under the Mizumoto General Public License v1.5 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://github.com/mizumoto-cn/mizubot/blob/main/LICENSE
 *     https://github.com/mizumoto-cn/mizubot/blob/main/licensing
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os"

	"github.com/mizumoto-cn/mizubot/core"
	"github.com/mizumoto-cn/mizubot/post"

	"github.com/urfave/cli/v2"
)

func main() {
	// print("Hello, World!")

	app := &cli.App{
		Name:  "MizBot",
		Usage: "A Discord bot for Mizumoto",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "webhook-url",
				Aliases:  []string{"w"},
				Usage:    "Discord webhook URL",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "tmp",
				Aliases:  []string{"t"},
				Usage:    "Template File Path",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "to-mizumoto",
				Aliases:  []string{"m"},
				Usage:    "To Mizumoto",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "content",
				Aliases:  []string{"c"},
				Usage:    "Content",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "chan1",
				Aliases:  []string{"c1"},
				Usage:    "Channel 1",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "chan2",
				Aliases:  []string{"c2"},
				Usage:    "Channel 2",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "user1mail",
				Aliases:  []string{"u1"},
				Usage:    "User 1 Mail",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "user2mail",
				Aliases:  []string{"u2"},
				Usage:    "User 2 Mail",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			content, err := core.CreateReport(c.String("tmp"), c.String("content"))
			if err != nil {
				return err
			}
			poster := post.NewPoster(
				c.String("webhook-url"),
				content,
				c.String("chan1"),
				c.String("chan2"),
				c.String("user1mail"),
				c.String("user2mail"),
			)
			err = poster.Post(c.Context)
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
