package http

import (
	entity "jusbrasil-tech-challenge/app/scrapper"
)

type (
	LegalCaseDTO struct {
		CaseNumber      string `json:"case_number"`
		Summary         string `json:"summary"`
		Reporter        string `json:"reporter"`
		Court           string `json:"court"`
		JudgingBody     string `json:"judging_body"`
		JudgmentDate    string `json:"judgment_date"`
		CaseClass       string `json:"case_type"`
		PublicationDate string `json:"publication_date"`
	}

	UserAgentResponse struct {
		UserAgent string `json:"user-agent"`
	}
)

func (lc *LegalCaseDTO) ToEntity() entity.LegalCase {
	return entity.LegalCase{
		CaseNumber:      lc.CaseNumber,
		Summary:         lc.Summary,
		Reporter:        lc.Reporter,
		Court:           lc.Court,
		JudgingBody:     lc.JudgingBody,
		JudgementDate:   lc.JudgmentDate,
		CaseClass:       lc.CaseClass,
		PublicationDate: lc.PublicationDate,
	}
}
