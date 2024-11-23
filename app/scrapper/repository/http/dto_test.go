package http

import (
	entity "jusbrasil-tech-challenge/app/scrapper"
	"jusbrasil-tech-challenge/tests/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLegalCaseDTO_ToEntity(t *testing.T) {
	t.Run("should return a legal case entity", func(t *testing.T) {
		expectedLegalCase := buildLegalCase()

		legalcaseDTO := LegalCaseDTO{
			CaseNumber:      expectedLegalCase.CaseNumber,
			Summary:         expectedLegalCase.Summary,
			Reporter:        expectedLegalCase.Reporter,
			Court:           expectedLegalCase.Court,
			JudgingBody:     expectedLegalCase.JudgingBody,
			JudgmentDate:    expectedLegalCase.JudgementDate,
			CaseClass:       expectedLegalCase.CaseClass,
			PublicationDate: expectedLegalCase.PublicationDate,
		}

		result := legalcaseDTO.ToEntity()

		assert.Equal(t, expectedLegalCase, result)
	})
}

func buildLegalCase() entity.LegalCase {
	return entity.LegalCase{
		CaseNumber:      mock.RandomString(),
		Summary:         mock.RandomString(),
		Reporter:        mock.RandomString(),
		Court:           mock.RandomString(),
		JudgingBody:     mock.RandomString(),
		JudgementDate:   mock.RandomString(),
		CaseClass:       mock.RandomString(),
		PublicationDate: mock.RandomString(),
	}
}
