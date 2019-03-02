package kanbanize

import "fmt"

type Column struct {
	ColumnID       int         `json:"column_id"`
	Workflow       int         `json:"workflow"`
	Section        int         `json:"section"`
	ParentColumnID interface{} `json:"parent_column_id"`
	Position       int         `json:"position"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Color          string      `json:"color"`
	Limit          int         `json:"limit"`
	CardsPerRow    int         `json:"cards_per_row"`
	FlowType       int         `json:"flow_type"`
}

type Columns []Column

//GetColumns returns the columns data of the board
func (c *Client) GetColumns(boardid string) (*Columns, error) {

	type kanbanizeColumns struct {
		Data []Column `json:"data"`
	}

	result := kanbanizeColumns{}

	url := c.APIV2URL + fmt.Sprintf("/boards/%s/columns", boardid)
	_, err := c.getRequest().
		SetResult(&result).
		Get(url)
	if err != nil {
		//handle read response error
		return nil, err
	}

	var cols Columns
	cols = result.Data
	return &cols, nil
}
