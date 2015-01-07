package banery_stats

import (
	"encoding/json"
)

func GetWorkspaces(c *Client) []Workspace {
	workspaces := make([]Workspace, 0)
	json.Unmarshal(c.FetchBody("user/workspaces.json/"), &workspaces)
	return workspaces
}

func GetOwnProjectTasks(client *Client, projectId int, ownId OwnId) []Task {
	url_path := client.OwnProjectTasksURLPath(projectId, ownId)
	tasks := make([]Task, 0)
	json.Unmarshal(client.FetchBody(url_path), &tasks)
	return FilterTasks(tasks, ownId)
}

func FilterTasks(tasks []Task, ownId OwnId) []Task {
	var ownProjects []Task
	for _, v := range tasks {
		if v.OwnerId == ownId.Id {
			ownProjects = append(ownProjects, v)
		}
	}
	return ownProjects
}

func GetOwnUserId(client *Client) OwnId {
	ownId := OwnId{}
  json.Unmarshal(client.FetchBody("user.json"), &ownId)
	return ownId
}
