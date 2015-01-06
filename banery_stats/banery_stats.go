package banery_stats

import (
	"fmt"
	"os"
)

func Run() {
	logger := Logger()


  if len(ApiToken()) == 0 {
    logger("Set KANBANERY_API_TOKEN to your personal kanbanery API token")
    return
  }

	ownUserId := GetOwnUserId()
	workspaces := GetWorkspaces()
	for _, workspace := range workspaces {
		logger("WORKSPACE " + workspace.Name + " #projects: " + fmt.Sprintf("%v", len(workspace.Projects)))
		for _, project := range workspace.Projects {
			project_tasks := GetOwnProjectTasks(workspace.Name, project.Id, ownUserId)
			project_tasks_length := len(project_tasks)
			if project_tasks_length > 0 {
				logger(fmt.Sprintf("%5v", project_tasks_length) + " " + project.Name)
			}
		}
	}
}

func ApiToken() string {
	return os.Getenv("KANBANERY_API_TOKEN")
}
