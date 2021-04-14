package panda

import "encoding/json"

type Content struct {
	Author          string `json:"author"`
	AuthorID        string `json:"authorId"`
	Container       string `json:"container"`
	CopyRightAlert  string `json:"copyrightAlert"`
	Description     string `json:"description"`
	EndDate         string `json:"endDate"`
	FromDate        string `json:"fromDate"`
	ModifiedDate    string `json:"modifiedDate"`
	NumChildren     int    `json:"numChildren"`
	Quota           string `json:"quota"`
	Size            int    `json:"size"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	URL             string `json:"url"`
	Usage           string `json:"usage"`
	Hidden          bool   `json:"hidden"`
	Visible         bool   `json:"visible"`
	EntityReference string `json:"entityReference"`
	EntityURL       string `json:"entityURL"`
	EntityTitle     string `json:"entityTitle"`
}

func (p *Handler) GetOwnInfo() (content []Content) {
	res, err := p.get(BaseURI+"/direct/content/my.json", nil)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var data Data

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	content = data.ContentCollection

	return
}

func (p *Handler) GetContent(siteID string) (content []Content) {
	res, err := p.get(BaseURI+"/direct/content/site/"+siteID+".json", nil)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var data Data

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	content = data.ContentCollection

	return
}
