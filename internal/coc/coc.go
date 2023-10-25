package coc

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func GetClans(tag string) (*Clan, error) {

	url := "https://api.clashofclans.com/v1/clans/%23" + tag
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("COC_TOKEN"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	clan := &Clan{}
	err = json.Unmarshal(body, clan)

	if err != nil {
		return nil, err
	}
	return clan, nil
}
