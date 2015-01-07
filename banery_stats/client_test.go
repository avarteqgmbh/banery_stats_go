package banery_stats

import "testing"

func TestClientWithKeyAndWorkspaceName(t *testing.T) {
  client := ClientWithKeyAndWorkspaceName("API_KEY", "WORKSPACE")

  if client.Key != "API_KEY" {
    t.Error("Expected Key to be API_KEY, got", client.Key)
  }

  if client.BaseURL != "https://WORKSPACE.kanbanery.com/api/v1/" {
    t.Error("Expected base url to be https://WORKSPACE.kanbanery.com/api/v1/, got", client.BaseURL)
  }
}
