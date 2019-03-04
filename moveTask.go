package kanbanize

import "fmt"

type moveTask struct {
	Boardid  string `json:"boardid,omitempty"`
	Taskid   string `json:"taskid,omitempty"`
	Column   string `json:"column,omitempty"`
	Lane     string `json:"lane,omitempty"`
	Position string `json:"position,omitempty"`
	client   *Client
}

func (c *Client) MoveTask(TaskID string) *moveTask {
	return &moveTask{
		Taskid: TaskID,
		client: c,
	}
}

func (t *moveTask) ToBoard(BoardID string) *moveTask {
	t.Boardid = BoardID
	return t
}

func (t *moveTask) ToColumn(colName string) *moveTask {
	t.Column = colName
	return t
}

func (t *moveTask) ToLane(laneName string) *moveTask {
	t.Lane = laneName
	return t
}
func (t *moveTask) ToPos(pos string) *moveTask {
	t.Position = pos
	return t
}

//GetBoardStructure returns the structure of all boards
func (t *moveTask) Move() error {

	url := t.client.APIURL + "move_task"
	res, err := t.client.getRequest().
		SetBody(t).
		Post(url)
	fmt.Printf("Move returned %s\n", res.String())
	if err != nil {
		//handle read response error
		return err
	}

	return nil
}
