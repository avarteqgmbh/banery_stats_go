package banery_stats

import (
	"io/ioutil"
	"net/http"
)

func FetchBody(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Kanbanery-ApiToken", ApiToken())
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
