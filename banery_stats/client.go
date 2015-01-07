package banery_stats

import (
  "fmt"
	"io/ioutil"
  "net/http"
)

type Client struct {
  // kanbanery API key
  Key string

  // kanbanery API base. e.g. "https://WORKSPACE.kanbanery.com/api/v1/"
  BaseURL string

  // Requests are transported through this client
  HTTPClient *http.Client
}

func ClientWithKeyAndWorkspaceName(key string, workspace_name string) *Client {
  return &Client{
    Key: key,
    HTTPClient: &http.Client{},
    BaseURL: "https://" + workspace_name + ".kanbanery.com/api/v1/",
  }
}

func ClientWithKeyForWorkspaceApi(key string) *Client {
  return &Client{
    Key: key,
    HTTPClient: &http.Client{},
    BaseURL: "https://kanbanery.com/api/v1/",
  }
}

func (c *Client)OwnProjectTasksURLPath(projectId int, ownId OwnId) string {
  return "projects/" + fmt.Sprintf("%v", projectId) + "/tasks.json"
}

func (c *Client)FetchBody(urlPath string) []byte {
	req, _ := http.NewRequest("GET", c.BaseURL + urlPath, nil)
	req.Header.Set("X-Kanbanery-ApiToken", c.Key)
  resp, _ := c.HTTPClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
