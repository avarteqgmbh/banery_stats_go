package banery_stats

type (
	Task struct {
		Id      int `json:"id"`
		OwnerId int `json:"owner_id"`
	}

	Project struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	Workspace struct {
		Id       int       `json:"id"`
		Name     string    `json:"name"`
		Projects []Project `json:projects`
	}

	OwnId struct {
		Id int `json:"id"`
	}
)
