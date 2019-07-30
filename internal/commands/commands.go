package commands

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	"github.com/ayazabbas/giantbombcli/internal/model"
	"github.com/jedib0t/go-pretty/table"
)

func NoArgs(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return cli.NewExitError("No commands provided", 1)
}

func SearchGames(c *cli.Context) error {
	searchName := c.String("name")
	
	if searchName == "" {
		return cli.NewExitError("Missing flag. Please provide search term using --name", 2)
	}

	// Make search request to GiantBomb api
	path := "/api/search"
	params := model.SearchGamesRequest{
		Query: searchName,
		Resources: "game",
	}
	req, err := NewRequest("GET", path, params)
	if err != nil {
		return err
	}

	// Create variable to store results
	var games *model.SearchGameResponse

	// Perform the request
	err = DoRequest(req, &games)
	if err != nil {
		return err
	}

	// Create a nice looking table on the screen
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	//t.SetAllowedRowLength(150)  // TODO: Let user set a length limit and decide if remainder is either cut or put on new line
	t.AppendHeader(table.Row{"GUID", "Name", "Release Yr"})
	for _, game := range games.Results {
		t.AppendRow([]interface{}{game.GUID, game.Name, game.ExpectedReleaseYear})
	}
	fmt.Println("GAMES")
	t.Render()

	return nil
}

func DescribeGame(c *cli.Context) error {
	guid := c.String("guid")

	if guid == "" {
		return cli.NewExitError("Missing flag. Please provide the game's guid using --guid", 3)
	}

	// Make game request to GiantBomb api
	path := fmt.Sprintf("/api/game/%s", guid)
	var params struct{}
	req, err := NewRequest("GET", path, params)
	if err != nil {
		return err
	}

	// Create variable to store results
	var game *model.GameResponse

	// Perform the request
	err = DoRequest(req, &game)
	if err != nil {
		return err
	}
	
	// Create a nice looking table on the screen
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	//t.SetAllowedRowLength(150)  // TODO: Let user set a length limit and decide if remainder is either cut or put on new line
	t.AppendHeader(table.Row{"GUID", "Name", "Release Yr", "Deck"})
	t.AppendRow([]interface{}{
		game.Results.GUID, game.Results.Name, game.Results.ExpectedReleaseYear, game.Results.Deck})
	fmt.Println("GAMES")
	t.Render()

	if c.Bool("show-dlc") {
		// Make DLCs request to GiantBomb api
		path := "/api/dlcs"
		params := model.DLCsRequest{
			Filter: fmt.Sprintf("game:%d", game.Results.ID),
			Sort: "release_date:asc",
		}
		req, err := NewRequest("GET", path, params)
		if err != nil {
			return err
		}

		// Create variable to store results
		var dlcs *model.DLCsResponse

		// Perform the request
		err = DoRequest(req, &dlcs)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		//t.SetAllowedRowLength(150)  // TODO: Let user set a length limit and decide if remainder is either cut or put on new line
		t.AppendHeader(table.Row{"GUID", "Name", "Platform", "Release Date", "Deck"})
		for _, dlc := range dlcs.Results {
			t.AppendRow([]interface{}{dlc.GUID, dlc.Name, dlc.Platform.Name, dlc.ReleaseDate, dlc.Deck})
		}
		fmt.Println("DLCs")
		t.Render()
	}

	return nil
}
