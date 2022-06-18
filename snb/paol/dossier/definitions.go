package dossier

type Assessment struct {
	Levy  float64 `json:"levy"`
	Total int     `json:"totalAssessment"`
	Year  int     `json:"asstYear"`
}

type Sale struct {
	Amount int    `json:"saleAmount"`
	Date   string `json:"saleDate"`
}

type SaleNearby struct {
	LocationNumbers     []int  `json:"locNumbers"`
	LocationStreet      string `json:"locStreet"`
	Pan                 string `json:"pan"`
	PropertyDescription string `json:"propDesc"`
	PropertyStreet      string `json:"propStreet"`
	Sale
}

type SalesNearby struct {
	Key   string       `json:"key"`
	Sales []SaleNearby `json:"sales"`
}

type HistoryAssessments struct {
	Assessments []Assessment `json:"assessments"`
}

type HistorySales struct {
	Sales []Sale `json:"sales"`
}

type Summary struct {
	AssessmentYear    int     `json:"asstYear"`
	CurrentAssessment int     `json:"currAsst"`
	CurrentLevy       float64 `json:"curLevy"`
	Description       string  `json:"description"`
	ImageKey          string  `json:"imageKey"`
	Location          string  `json:"location"`
	Pan               string  `json:"pan"`
	TaxAuthority      string  `json:"taxAuth"`
}

type Results struct {
	History      HistoryAssessments `json:"history"`
	nearbySales  SalesNearby        `json:"nearbySales"`
	Pan          string             `json:"pan"`
	SalesHistory HistorySales       `json:"salesHistory"`
	Summary      Summary            `json:"summary"`
}
