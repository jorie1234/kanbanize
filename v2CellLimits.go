package kanbanize

import "fmt"

type CellLimit struct {
	LaneID   int `json:"lane_id"`
	ColumnID int `json:"column_id"`
	Limit    int `json:"limit"`
}

type CellLimits []CellLimit

//GetCellLimits returns the cell limits data of the board
func (c *Client) GetCellLimits(boardid string) (*CellLimits, error) {

	type kanbanizeCellLimits struct {
		Data []CellLimit `json:"data"`
	}

	result := kanbanizeCellLimits{}

	url := c.APIV2URL + fmt.Sprintf("/boards/%s/cellLimits", boardid)
	_, err := c.getRequest().
		SetResult(&result).
		Get(url)
	if err != nil {
		//handle read response error
		return nil, err
	}

	var cells CellLimits
	cells = result.Data
	return &cells, nil
}
