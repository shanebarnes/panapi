package search

type TaxAuth struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Pid    string `json:"pid"`
}

type Result struct {
	Location string  `json:"location"`
	Pan      string  `json:"pan"`
	TaxAuth  TaxAuth `json:"taxAuth"`
}

type Results struct {
	ResultsByGroup []ResultsByGroup `json:"searchResultsByGroup"`
	Term           string           `json:"searchTerm"`
}

type ResultsByGroup struct {
	HasMore     bool     `json:"hasMore"`
	ResultGroup string   `json:"resultGroup"`
	Results     []Result `json:"results"`
}
