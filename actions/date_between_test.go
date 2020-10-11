package actions

import (
	"testing"

	"github.com/br4in3x/time-since-google-action/internal/request"
	"github.com/br4in3x/time-since-google-action/internal/response"
	"github.com/stretchr/testify/require"
)

func Test_DateBetween(t *testing.T) {
	type output struct {
		res *response.Simple
		err string
	}

	type test struct {
		name   string
		tm     *timeMock
		input  request.Request
		output output
	}

	tests := []test{
		{
			"invalid from date input",
			&timeMock{},
			request.Request{},
			output{
				nil,
				"invalid from date input",
			},
		},
		{
			"invalid to date input",
			&timeMock{},
			request.Request{
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
			},
			output{
				nil,
				"invalid to date input",
			},
		},
		{
			"same date",
			&timeMock{},
			request.Request{
				Intent: request.Intent{
					Params: request.IntentParams{
						DateFrom: request.IntentParam{
							Resolved: request.Resolved{
								Day:   1,
								Month: 1,
								Year:  2020,
							},
						},
						DateTo: request.IntentParam{
							Resolved: request.Resolved{
								Day:   1,
								Month: 1,
								Year:  2020,
							},
						},
					},
				},
			},
			output{
				response.SimpleResponse("There is no difference. These are same dates."),
				"",
			},
		},
		{
			"success",
			&timeMock{},
			request.Request{
				Intent: request.Intent{
					Params: request.IntentParams{
						DateFrom: request.IntentParam{
							Resolved: request.Resolved{
								Day:   1,
								Month: 1,
								Year:  2020,
							},
							Original: "january 1, 2020",
						},
						DateTo: request.IntentParam{
							Resolved: request.Resolved{
								Day:   3,
								Month: 1,
								Year:  2020,
							},
							Original: "january 13, 2020",
						},
					},
				},
			},
			output{
				response.SimpleResponse("It's 2 days."),
				"",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			// arrange
			a := NewDateBetweenAction(test.tm)

			// act
			res, err := a.Invoke(test.input)

			// assert
			require.Equal(tt, test.output.res, res)

			if test.output.err != "" {
				require.Equal(tt, test.output.err, err.Error())
			}
		})
	}

}
