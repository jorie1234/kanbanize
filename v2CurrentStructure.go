package kanbanize

import "fmt"

type CurrentStructure struct {
	Version       string `json:"version"`
	ProjectID     int    `json:"project_id"`
	BoardID       int    `json:"board_id"`
	Name          string `json:"name"`
	IsArchived    int    `json:"is_archived"`
	CardsWorkflow struct {
		Exists        int    `json:"exists"`
		IsEnabled     int    `json:"is_enabled"`
		Name          string `json:"name"`
		TopLanes      []int  `json:"top_lanes"`
		BottomLanes   []int  `json:"bottom_lanes"`
		TopColumns    []int  `json:"top_columns"`
		BottomColumns []int  `json:"bottom_columns"`
	} `json:"cards_workflow"`
	InitiativesWorkflow struct {
		Exists        int    `json:"exists"`
		IsEnabled     int    `json:"is_enabled"`
		IsCollapsible int    `json:"is_collapsible"`
		Name          string `json:"name"`
		TopLanes      []int  `json:"top_lanes"`
		BottomLanes   []int  `json:"bottom_lanes"`
		TopColumns    []int  `json:"top_columns"`
		BottomColumns []int  `json:"bottom_columns"`
	} `json:"initiatives_workflow"`
	Lanes map[string]struct {
		Workflow     int    `json:"workflow"`
		ParentLaneID int    `json:"parent_lane_id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Color        string `json:"color"`
	}

	Columns map[string]struct {
		Workflow       int    `json:"workflow"`
		Section        int    `json:"section"`
		ParentColumnID int    `json:"parent_column_id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Color          string `json:"color"`
		Limit          int    `json:"limit"`
		CardsPerRow    int    `json:"cards_per_row"`
		FlowType       int    `json:"flow_type"`
	} `json:"columns"`
	CellLimits            map[string]map[string]int `json:"cell_limits"`
	SizeType              int                       `json:"size_type"`
	AllowExceeding        int                       `json:"allow_exceeding"`
	AutoarchiveCardsAfter int                       `json:"autoarchive_cards_after"`

	Revision int `json:"revision"`
}

//GetCurrentStructure returns the cell limits data of the board
func (c *Client) GetCurrentStructure(boardid string) (*CurrentStructure, error) {

	type kanbanizeCurrentStructure struct {
		Data CurrentStructure `json:"data"`
	}

	result := kanbanizeCurrentStructure{}

	url := c.APIV2URL + fmt.Sprintf("/boards/%s/currentStructure", boardid)
	_, err := c.getRequest().
		SetResult(&result).
		Get(url)
	if err != nil {
		//handle read response error
		return nil, err
	}

	return &result.Data, nil
}
