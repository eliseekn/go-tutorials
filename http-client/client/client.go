package client

import (
	"encoding/json"
	"http-client/types"
	"net/http"
)

func Get(url string) (types.Response, error) {
	client, err := http.Get(url)
	if err != nil {
		return types.Response{}, err
	}

	defer client.Body.Close()

	var res types.Response
	jsonErr := json.NewDecoder(client.Body).Decode(&res)

	if jsonErr != nil {
		return types.Response{}, err
	}

	return res, nil
}
