package banery_stats

import (
	"fmt"
	"os"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(workspaces))

	for _, workspace := range workspaces {
		go outputWorkspace(logger, workspace, ownUserId, &wg)
	}

	wg.Wait()

}

func ApiToken() string {
	return os.Getenv("KANBANERY_API_TOKEN")
}


func outputWorkspace(logger func(string), workspace Workspace, ownUserId OwnId, wg *sync.WaitGroup) {
	client := ClientWithKeyAndWorkspaceName(ApiToken(), workspace.Name)
	project_chan := make(chan Project)
	projects := make([]Project, 0)


	defer wg.Done()


	for _, project := range workspace.Projects {
		go GetOwnProjectTasks(client, project, ownUserId, project_chan)
	}

	for i := 0; i < len(workspace.Projects); i++ {
		projects = append(projects, <-project_chan)
	}

	logger("WORKSPACE " + workspace.Name + " #projects: " + fmt.Sprintf("%v", len(workspace.Projects)))
	for _, project := range projects {
		project_tasks_length := len(project.Tasks)
		if project_tasks_length > 0 {
			logger(fmt.Sprintf("%5v", project_tasks_length) + " " + project.Name)
		}
	}

}
