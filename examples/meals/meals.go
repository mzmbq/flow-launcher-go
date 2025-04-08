package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"

	flow "github.com/mzmbq/flow-launcher-go"
)

const (
	randomMealURL = "https://www.themealdb.com/api/json/v1/1/random.php"
	searchMealURL = "https://www.themealdb.com/api/json/v1/1/search.php?s="
)

type Meals struct {
	Meals []struct {
		Title    string `json:"strMeal"`
		Category string `json:"strCategory"`
		Link     string `json:"strYoutube"`
	} `json:"meals"`
}

func fetchRandomMeal() (*Meals, error) {
	resp, err := http.Get(randomMealURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ms Meals
	err = json.NewDecoder(resp.Body).Decode(&ms)
	if err != nil {
		return nil, err
	}

	// Reroll if recipe doesn't have a url
	if ms.Meals[0].Link == "" {
		return fetchRandomMeal()
	}

	return &ms, nil
}

func searchMealByName(name string) (*Meals, error) {
	resp, err := http.Get(searchMealURL + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ms Meals
	err = json.NewDecoder(resp.Body).Decode(&ms)
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

func handleQuery(req *flow.Request) *flow.Response {
	var meals *Meals
	var err error
	if req.Parameters[0] == "" {
		meals, err = fetchRandomMeal()
	} else {
		meals, err = searchMealByName(req.Parameters[0])
	}

	if err != nil {
		res := flow.NewResponse(req)
		res.DebugMessage = "failed to fetch API"
		return res
	}

	resp := flow.NewResponse(req)
	for _, m := range meals.Meals {
		result := flow.Result{
			Title:    m.Title,
			SubTitle: m.Category,
			IcoPath:  "icon.png",
			RpcAction: &flow.JsonRpcAction{
				Method:     "open",
				Parameters: []string{m.Link},
			},
		}
		resp.AddResult(&result)
	}
	return resp
}

func handleOpenAction(params []string) *flow.Response {
	url := params[0]
	if url == "" {
		url = "https://telegra.ph/Error-04-05-449"
	}

	err := exec.Command("cmd", "/c", "start", url).Start()
	if err != nil {
		return flow.ErrorResponse(err.Error())
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		panic("Usage: meals.exe <query>")
	}

	p := flow.NewPlugin()
	p.Query(handleQuery)
	p.Action("open", handleOpenAction)
	if err := p.HandleRPC(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
