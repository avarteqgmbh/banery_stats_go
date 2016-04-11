package banery_stats

import (
	"encoding/json"
)

func GetWorkspaces(c *Client) []Workspace {
	workspaces := make([]Workspace, 0)
	json.Unmarshal(c.FetchBody("user/workspaces.json/"), &workspaces)
	return workspaces
}

func GetOwnProjectTasks(client *Client, project Project, ownId OwnId, project_chan chan Project) Project {
	url_path := client.OwnProjectTasksURLPath(project.Id, ownId)
	tasks := make([]Task, 0)
	json.Unmarshal(client.FetchBody(url_path), &tasks)
	filtered_tasks := FilterTasks(tasks, ownId)
	project.Tasks = filtered_tasks
	project_chan <- project
	return project
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
