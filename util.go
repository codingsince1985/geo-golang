package geo

import (
	"io/ioutil"
	"net/http"
)

func ResponseData(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	return data
}
