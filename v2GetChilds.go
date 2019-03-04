package kanbanize

import "fmt"

type Children struct {
	CardID   int `json:"card_id"`
	Position int `json:"position"`
}

type Childrens []Children

//GetCellLimits returns the cell limits data of the board
func (c *Client) GetChildren(cardid string) (*Childrens, error) {

	type kanbanizeChildren struct {
		Data []Children `json:"data"`
	}

	result := kanbanizeChildren{}

	url := c.APIV2URL + fmt.Sprintf("/cards/%s/children", cardid)
	_, err := c.getRequest().
		SetResult(&result).
		Get(url)
	if err != nil {
		//handle read response error
		return nil, err
	}

	var childs Childrens
	childs = result.Data
	return &childs, nil
}
