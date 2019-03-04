package kanbanize

type getAllTasks struct {
	Boardid         string `json:"boardid,omitempty"`
	Subtasks        string `json:"subtasks,omitempty"`
	Comments        string `json:"comments,omitempty"`
	Container       string `json:"container,omitempty"`
	Fromdate        string `json:"fromdate,omitempty"`
	Todate          string `json:"todate,omitempty"`
	Showinitiatives string `json:"showinitiatives,omitempty"`
	Version         string `json:"version,omitempty"`
	Pagee           string `json:"page,omitempty"`
	Textformat      string `json:"textformat,omitempty"`
	Column          string `json:"column,omitempty"`
	Lane            string `json:"lane,omitempty"`
	Section         string `json:"section,omitempty"`
	client          *Client
}

func (c *Client) GetAllTasks(boardid string) *getAllTasks {
	return &getAllTasks{
		Boardid: boardid,
		client:  c,
	}
}

func (t *getAllTasks) WithSubtasks() *getAllTasks {
	t.Subtasks = "yes"
	return t
}

func (t *getAllTasks) WithComments() *getAllTasks {
	t.Comments = "yes"
	return t
}

func (t *getAllTasks) FromArchive() *getAllTasks {
	t.Container = "archive"
	return t
}

func (t *getAllTasks) FromDate(date string) *getAllTasks {
	t.Fromdate = date
	return t
}

func (t *getAllTasks) ToDate(date string) *getAllTasks {
	t.Todate = date
	return t
}

func (t *getAllTasks) OnlyInitiatives(date string) *getAllTasks {
	t.Showinitiatives = "1"
	return t
}

func (t *getAllTasks) WithVersion(version string) *getAllTasks {
	t.Version = version
	return t
}

func (t *getAllTasks) Page(page string) *getAllTasks {
	t.Pagee = page
	return t
}

func (t *getAllTasks) FormatText() *getAllTasks {
	t.Textformat = "plain"
	return t
}

func (t *getAllTasks) FormatHTML() *getAllTasks {
	t.Textformat = "html"
	return t
}

func (t *getAllTasks) FromColumn(col string) *getAllTasks {
	t.Column = col
	return t
}

func (t *getAllTasks) FromLane(lane string) *getAllTasks {
	t.Lane = lane
	return t
}

func (t *getAllTasks) FromSection(section string) *getAllTasks {
	t.Section = section
	return t
}

//GetBoardStructure returns the structure of all boards
func (t *getAllTasks) Get() (*Tasks, error) {

	type modTask struct {
		Task
		BlockedString string `json:"blocked,omitempty"`
		BlockedInt    int    `json:"blocked,omitempty"`
	}
	result := []modTask{}

	url := t.client.APIURL + "get_all_tasks"
	_, err := t.client.getRequest().
		SetBody(t).
		SetResult(&result).
		Post(url)
	if err != nil {
		//handle read response error
		return nil, err
	}

	var res Tasks
	for _, v := range result {
		tsk := v.Task
		tsk.Boardid = t.Boardid
		if v.BlockedInt == 1 || v.BlockedString == "1" {
			tsk.Blocked = "1"
		} else {
			tsk.Blocked = "0"
		}
		res = append(res, tsk)
	}
	return &res, nil
}
