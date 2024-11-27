package scrapper

type LegalCase struct {
	CaseNumber      string `json:"case_number"`
	Summary         string `json:"summary"`
	Reporter        string `json:"reporter"`
	JudgingBody     string `json:"judging_body"`
	PublicationDate string `json:"publication_date"`
	JudgementDate   string `json:"judgement_date"`
	CaseClass       string `json:"case_class"`
	Court           string `json:"court"`
}
