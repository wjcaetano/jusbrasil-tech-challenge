package http

import "jusbrasil-tech-challenge/app/scrapper"

type legalCase struct {
	CaseNumber      string `json:"case_number"`
	Summary         string `json:"summary"`
	Reporter        string `json:"reporter"`
	Court           string `json:"court"`
	JudgingBody     string `json:"judging_body"`
	JudgementDate   string `json:"judgement_date"`
	CaseClass       string `json:"case_class"`
	PublicationDate string `json:"publication_date"`
}

func (p legalCase) ToEntity() scrapper.LegalCase {
	return scrapper.LegalCase{
		CaseNumber:      p.CaseNumber,
		Summary:         p.Summary,
		Reporter:        p.Reporter,
		Court:           p.Court,
		JudgingBody:     p.JudgingBody,
		JudgementDate:   p.JudgementDate,
		PublicationDate: p.PublicationDate,
	}
}
