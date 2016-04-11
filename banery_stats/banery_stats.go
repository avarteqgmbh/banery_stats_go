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

	workspaceClient := ClientWithKeyForWorkspaceApi(ApiToken())
	workspaces := GetWorkspaces(workspaceClient)
	if len(workspaces) == 0 {
		logger("No workspaces present.")
		return
	}

	ownUserIdClient := ClientWithKeyAndWorkspaceName(ApiToken(), workspaces[0].Name)
	ownUserId := GetOwnUserId(ownUserIdClient)

	for _, workspace := range workspaces {
		outputWorkspace(logger, workspace, ownUserId)
	}
}

func ApiToken() string {
	return os.Getenv("KANBANERY_API_TOKEN")
}

func outputWorkspace(logger func(string), workspace Workspace, ownUserId OwnId) {
	logger("WORKSPACE " + workspace.Name + " #projects: " + fmt.Sprintf("%v", len(workspace.Projects)))


	client := ClientWithKeyAndWorkspaceName(ApiToken(), workspace.Name)
	project_chan := make(chan Project)

	for _, project := range workspace.Projects {
		go GetOwnProjectTasks(client, project, ownUserId, project_chan)
		//project_tasks := GetOwnProjectTasks(client, project.Id, ownUserId, project_chan)
		/*project_tasks_length := len(project_tasks)
		if project_tasks_length > 0 {
			logger(fmt.Sprintf("%5v", project_tasks_length) + " " + project.Name)
		}*/
	}
	for {
		project := <-project_chan
		project_tasks_length := len(project.Tasks)
		if project_tasks_length > 0 {
			logger(fmt.Sprintf("%5v", project_tasks_length) + " " + project.Name)
		}
	}

}
