package banery_stats

import (
	"encoding/json"
	"fmt"
)

const (
  WORKSPACE_URL string = "https://kanbanery.com/api/v1/user/workspaces.json/"
  USER_URL string = "https://avarteq.kanbanery.com/api/v1/user.json"
)

func GetWorkspaces() []Workspace {
	workspaces := make([]Workspace, 0)
	json.Unmarshal(FetchBody(WORKSPACE_URL), &workspaces)

	return workspaces
}

func GetOwnProjectTasks(workspaceName string, projectId int, ownId OwnId) []Task {
	url := "https://" + workspaceName + ".kanbanery.com/api/v1/projects/" + fmt.Sprintf("%v", projectId) + "/tasks.json"

	tasks := make([]Task, 0)
	json.Unmarshal(FetchBody(url), &tasks)
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

func GetOwnUserId() OwnId {
	ownId := OwnId{}
	json.Unmarshal(FetchBody(USER_URL), &ownId)
	return ownId
}
