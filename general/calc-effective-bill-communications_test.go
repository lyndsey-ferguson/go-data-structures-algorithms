package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	BILLS = []Bill{
		{
			patientId: "p1",
			amount:    100.00,
			timestamp: 160310 * 10000,
		},
		{
			patientId: "p2",
			amount:    500.00,
			timestamp: 160320 * 10000,
		},
		{
			patientId: "p2",
			amount:    350.00,
			timestamp: 160320 * 10000,
		},
		{
			patientId: "p3",
			amount:    450.00,
			timestamp: 160340 * 10000,
		},
		{
			patientId: "p1",
			amount:    100.00,
			timestamp: 160500 * 10000,
		},
	}

	COMMUNICATIONS = []Notice{
		{patientId: "p1", channel: email, sentTimestamp: 160325 * 10000},
		{patientId: "p2", channel: text, sentTimestamp: 160320 * 10000},
		{patientId: "p2", channel: paperMail, sentTimestamp: 160330 * 10000},
		{patientId: "p2", channel: paperMail, sentTimestamp: 160340 * 10000},
		{patientId: "p3", channel: paperMail, sentTimestamp: 160345 * 10000},
	}

	PAYMENTS = []Payment{
		{patientId: "p2", amount: 350.00, paidTimestamp: 160345 * 10000},
		{patientId: "p3", amount: 450.00, paidTimestamp: 160360 * 10000},
		{patientId: "p1", amount: 100.00, paidTimestamp: 160600 * 10000},
	}
)

func TestCalculateEffectiveness(t *testing.T) {
	emailScore, textScore, paperMailScore := calculateEffectiveness(BILLS, COMMUNICATIONS, PAYMENTS)
	assert.Equal(t, 0, emailScore)
	assert.Equal(t, 0, textScore)
	assert.Equal(t, 3, paperMailScore)
}
