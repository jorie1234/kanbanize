package kanbanize

import "fmt"

type editTask struct {
	Boardid     string `json:"boardid,omitempty"`
	Taskid      string `json:"taskid,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Assignee    string `json:"assignee,omitempty"`
	Size        string `json:"size,omitempty"`
	Tags        string `json:"tags,omitempty"`
	Deadline    string `json:"deadline,omitempty"`
	Extlink     string `json:"extlink,omitempty"`
	Type        string `json:"type,omitempty"`
	client      *Client
}

func (c *Client) EditTask(BoardID, TaskID string) *editTask {
	return &editTask{
		Taskid:  TaskID,
		Boardid: BoardID,
		client:  c,
	}
}

func (t *editTask) SetTitle(Title string) *editTask {
	t.Title = Title
	return t
}

func (t *editTask) SetDescription(desc string) *editTask {
	t.Description = desc
	return t
}

func (t *editTask) SetPriority(prio string) *editTask {
	t.Priority = prio
	return t
}
func (t *editTask) SetAssignee(assignee string) *editTask {
	t.Assignee = assignee
	return t
}

func (t *editTask) SetSize(size string) *editTask {
	t.Size = size
	return t
}

func (t *editTask) SetTags(tags string) *editTask {
	t.Tags = tags
	return t
}

func (t *editTask) SetDeadline(dead string) *editTask {
	t.Deadline = dead
	return t
}

func (t *editTask) SetExtLink(link string) *editTask {
	t.Extlink = link
	return t
}

func (t *editTask) SetType(typ string) *editTask {
	t.Type = typ
	return t
}

//GetBoardStructure returns the structure of all boards
func (t *editTask) Edit() error {

	url := t.client.APIURL + "edit_task"
	res, err := t.client.getRequest().
		SetBody(t).
		Post(url)
	fmt.Printf("Edit Task returned %s\n", res.String())
	if err != nil {
		//handle read response error
		return err
	}

	return nil
}
