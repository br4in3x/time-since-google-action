package actions

import (
	"testing"
	"time"

	"github.com/br4in3x/time-since-google-action/internal/request"
	"github.com/br4in3x/time-since-google-action/internal/response"
	"github.com/stretchr/testify/require"
)

func Test_DateFrom_DateNotParsed_ReturnsError(t *testing.T) {
	// Arrange
	req := request.Request{}
	act := NewDateFromToAction(&timeMock{})

	// Act
	res, err := act.Invoke(req)

	// assert
	require.NoError(t, err)
	require.Equal(t, response.ClarifyDate(), res)
}

func Test_DateFrom_TimeBeforeCurrentDate(t *testing.T) {
	// arrange
	req := request.Request{
		Intent: request.Intent{
			Params: request.IntentParams{
				DateFrom: request.IntentParam{
					Resolved: request.Resolved{
						Day:   1,
						Month: 1,
						Year:  2020,
					},
				},
			},
		},
	}
	act := NewDateFromToAction(&timeMock{ResTime: time.Date(2020, time.January, 7, 0, 0, 0, 0, time.Now().Location())})

	// act
	res, err := act.Invoke(req)

	// assert
	require.NoError(t, err)
	require.Equal(t, "6 days have passed since that time.", res.Prompt.FirstSimple.Speech)
}

func Test_DateFrom_TimeAfterCurrentDate(t *testing.T) {
	// arrange
	req := request.Request{
		Intent: request.Intent{
			Params: request.IntentParams{
				DateFrom: request.IntentParam{
					Resolved: request.Resolved{
						Day:   7,
						Month: 1,
						Year:  2020,
					},
					Original: "January 7th, 2020",
				},
			},
		},
	}
	act := NewDateFromToAction(&timeMock{ResTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Now().Location())})

	// act
	res, err := act.Invoke(req)

	// assert
	require.NoError(t, err)
	require.Equal(t, "It's 6 days until January 7th, 2020.", res.Prompt.FirstSimple.Speech)
}

func Test_DateFrom_ZeroDays(t *testing.T) {
	// arrange
	req := request.Request{
		Intent: request.Intent{
			Params: request.IntentParams{
				DateFrom: request.IntentParam{
					Resolved: request.Resolved{
						Day:   7,
						Month: 1,
						Year:  2020,
					},
					Original: "January 7th, 2020",
				},
			},
		},
	}
	act := NewDateFromToAction(&timeMock{ResTime: time.Date(2020, time.January, 7, 0, 0, 0, 0, time.Now().Location())})

	// act
	res, err := act.Invoke(req)

	// assert
	require.NoError(t, err)
	require.Equal(t, "0 days. The January 7th, 2020 is today.", res.Prompt.FirstSimple.Speech)
}

type timeMock struct {
	ResTime time.Time
}

func (m *timeMock) Now() time.Time {
	return m.ResTime
}
