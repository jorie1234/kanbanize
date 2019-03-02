package kanbanize

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

const (
	CardWorkflow        = 0
	InitiativesWorkflow = 1
)

type Board struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

//ProjectsAndBoards represents all boards and projects
type ProjectsAndBoards struct {
	Projects []struct {
		Name   string  `json:"name"`
		ID     string  `json:"id"`
		Boards []Board `json:"boards"`
	} `json:"projects"`
}

//Task represents the Kanbanize Task
type Task struct {
	Taskid                 string        `json:"taskid,omitempty"`
	Boardid                string        `json:"boardid,omitempty"`
	Title                  string        `json:"title,omitempty"`
	Description            string        `json:"description,omitempty"`
	Type                   string        `json:"type,omitempty"`
	Assignee               string        `json:"assignee,omitempty"`
	Subtasks               string        `json:"subtasks,omitempty"`
	Subtaskscomplete       string        `json:"subtaskscomplete,omitempty"`
	Color                  string        `json:"color,omitempty"`
	Priority               string        `json:"priority,omitempty"`
	Size                   string        `json:"size,omitempty"`
	Deadline               string        `json:"deadline,omitempty"`
	Deadlineoriginalformat string        `json:"deadlineoriginalformat,omitempty"`
	Extlink                string        `json:"extlink,omitempty"`
	Tags                   string        `json:"tags,omitempty"`
	Leadtime               int           `json:"leadtime,omitempty"`
	Blocked                string        `json:"blocked,omitempty"`
	Blockedreason          string        `json:"blockedreason,omitempty"`
	Columnname             string        `json:"columnname,omitempty"`
	Column                 string        `json:"column,omitempty"`
	Lanename               string        `json:"lanename,omitempty"`
	Subtaskdetails         []interface{} `json:"subtaskdetails,omitempty"`
	Columnid               string        `json:"columnid,omitempty"`
	Laneid                 string        `json:"laneid,omitempty"`
	Position               string        `json:"position,omitempty"`
	Workflow               int           `json:"workflow,omitempty"`
	Attachments            struct {
	} `json:"attachments,omitempty"`
	Columnpath   string        `json:"columnpath,omitempty"`
	Loggedtime   int           `json:"loggedtime,omitempty"`
	Customfields []interface{} `json:"customfields,omitempty"`
	Updatedat    string        `json:"updatedat,omitempty"`
}

//Link represents all the relations of a single task
type Link struct {
	Taskid       string      `json:"taskid"`
	Parent       interface{} `json:"parent"`
	Children     []string    `json:"children"`
	Mirrors      []string    `json:"mirrors"`
	Relatives    []string    `json:"relatives"`
	Successors   []string    `json:"successors"`
	Predecessors []string    `json:"predecessors"`
}

//Links represents all relations of a task
type Links map[string]Link

type Tasks []Task

func (t *Tasks) GetColumnNameByColumnId(id int) string {
	for _, d := range *t {
		if d.Columnid == fmt.Sprintf("%d", id) {
			return d.Columnname
		}
	}
	return "-"
}

//Client is the Kanbanize Client
type Client struct {
	APIURL   string
	APIV2URL string
	APIKey   string
	client   *resty.Client
}

//NewClient returns an new Client for the Kanbanize API
func NewClient(APIRUL, APIKey string) *Client {
	rc := resty.New()
	return &Client{
		APIKey:   APIKey,
		APIURL:   APIRUL + "/index.php/api/kanbanize/",
		APIV2URL: APIRUL + "/api/v2/",
		client:   rc,
	}
}

func (c *Client) getRequest() *resty.Request {
	return c.client.R().
		SetHeader("apikey", c.APIKey).
		SetHeader("Accept", "application/json")
}

//GetBoards returns a struct containing all Projects and Boards
func (c *Client) GetBoards() (*ProjectsAndBoards, error) {
	boards := ProjectsAndBoards{}

	url := c.APIURL + "get_projects_and_boards" + "/format/json"
	_, err := c.getRequest().
		SetResult(&boards).
		Post(url)
	if err != nil {
		return nil, err
	}

	return &boards, nil
}

//CreateTask creates a Task in Kanbanize
func (c *Client) CreateTask(t *Task) (*Task, error) {

	type tSend struct {
		Task
		ReturnTaskDetails string `json:"returntaskdetails,omitempty"`
	}
	type returnedTask struct {
		ID      int `json:"id"`
		Details struct {
			Task
		} `json:"details"`
	}
	var resultTask returnedTask
	var tt tSend
	tt.Task = *t
	tt.ReturnTaskDetails = "1"
	url := c.APIURL + "create_new_task"
	_, err := c.getRequest().
		SetBody(tt).
		SetResult(&resultTask).
		Post(url)
	if err != nil {
		return nil, err
	}
	return &resultTask.Details.Task, nil
}

//GetTask returns the task with taskid from the board boardid
func (c *Client) GetTask(boardid, taskid string) (*Task, error) {

	type GetTaskData struct {
		Boardid string `json:"boardid"`
		Taskid  string `json:"taskid"`
	}

	var t Task
	url := c.APIURL + "get_task_details" + "/format/json"
	_, err := c.getRequest().
		SetBody(GetTaskData{
			Boardid: boardid,
			Taskid:  taskid,
		}).
		SetResult(&t).
		Post(url)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

//DeleteTask deletes the Task with taskid on the board boardid
func (c *Client) DeleteTask(boardid, taskid string) error {

	type DeleteTaskData struct {
		Boardid string `json:"boardid"`
		Taskid  string `json:"taskid"`
	}

	url := c.APIURL + "delete_task"
	_, err := c.getRequest().
		SetBody(DeleteTaskData{
			Boardid: boardid,
			Taskid:  taskid,
		}).
		Post(url)
	if err != nil {
		return err
	}
	return nil
}

//GetLinks gets all the links of the Task with taskid
func (c *Client) GetLinks(taskid string) (*Links, error) {

	result := make(Links)

	type GetLinksData struct {
		Taskid string `json:"taskid"`
	}

	url := c.APIURL + "get_links"
	_, err := c.getRequest().
		SetBody(GetLinksData{
			Taskid: taskid,
		}).
		SetResult(&result).
		Post(url)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
