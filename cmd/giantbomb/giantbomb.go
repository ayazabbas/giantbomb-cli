package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/ayazabbas/giantbombcli/internal/commands"
	"github.com/ayazabbas/giantbombcli/internal/model"
)

func main() {
	model.APIKey = os.Getenv("GB_API_KEY")
	model.BaseURL = "http://www.giantbomb.com"

	app := cli.NewApp()
	app.Name = "Giantbomb CLI"
	app.Usage = "Let's you search for and describe games in the Giant Bomb database"
	app.HideVersion = true
	app.Action = commands.NoArgs

	searchFlags := []cli.Flag{
		cli.StringFlag {
			Name: "name",
			Usage: "Required -- name of the game (can be just a keyword in the name).",
		},
	}

	describeGameFlags := []cli.Flag{
		cli.StringFlag {
			Name: "guid",
			Usage: "Required -- GUID of the game, use search-games to find this first.",
		},
		cli.BoolFlag {
			Name: "show-dlc",
			Usage: "Display DLC details",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "search-games",
			Usage: "Search for games in the Giant Bomb database",
			Flags: searchFlags,
			Action: commands.SearchGames,
		},
		{
			Name: "describe-game",
			Usage: "Retrieve info for a specific game in the Giant Bomb database",
			Flags: describeGameFlags,
			Action: commands.DescribeGame,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
