package kanbanize

//BoardStructure represents the structure of all boards
type BoardStructure struct {
	Columns []struct {
		Position    string `json:"position"`
		Lcname      string `json:"lcname"`
		Section     string `json:"section"`
		Path        string `json:"path"`
		Description string `json:"description"`
		Lcid        string `json:"lcid"`
		Flowtype    string `json:"flowtype"`
		Tasksperrow int    `json:"tasksperrow,omitempty"`
		Children    []struct {
			Position    string `json:"position"`
			Lcname      string `json:"lcname"`
			Section     string `json:"section"`
			Path        string `json:"path"`
			Description string `json:"description"`
			Lcid        string `json:"lcid"`
			Flowtype    string `json:"flowtype"`
			Tasksperrow int    `json:"tasksperrow"`
		} `json:"children,omitempty"`
	} `json:"columns"`
	Lanes []struct {
		Position    string `json:"position"`
		Lcname      string `json:"lcname"`
		Path        string `json:"path"`
		Description string `json:"description"`
		Lcid        string `json:"lcid"`
		Flowtype    string `json:"flowtype"`
		Color       string `json:"color"`
	} `json:"lanes"`
}

//GetBoardStructure returns the structure of all boards
func (c *Client) GetBoardStructure(boardid string) (*BoardStructure, error) {

	result := BoardStructure{}

	type getBoardStructureData struct {
		BoardID string `json:"boardid"`
	}

	url := c.APIURL + "get_full_board_structure"
	_, err := c.getRequest().
		SetBody(getBoardStructureData{
			BoardID: boardid,
		}).
		SetResult(&result).
		Post(url)
	if err != nil {
		//handle read response error
		return nil, err
	}

	return &result, nil
}
