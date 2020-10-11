package actions

import (
	"fmt"
	"math"
	"time"

	"github.com/br4in3x/time-since-google-action/internal/request"
	"github.com/br4in3x/time-since-google-action/internal/response"
	"github.com/br4in3x/time-since-google-action/internal/util"
)

type DateBetweenAction struct {
	TimeWrapper util.Time
}

func NewDateBetweenAction(timeWrapper util.Time) *DateBetweenAction {
	return &DateBetweenAction{timeWrapper}
}

func (a *DateBetweenAction) Invoke(r request.Request) (*response.Simple, error) {
	dFrom := r.Intent.Params.DateFrom.Resolved.Day
	mFrom := r.Intent.Params.DateFrom.Resolved.Month
	yFrom := r.Intent.Params.DateFrom.Resolved.Year

	if dFrom == 0 || mFrom == 0 || yFrom == 0 {
		return response.ClarifyDate(), nil
	}

	dTo := r.Intent.Params.DateTo.Resolved.Day
	mTo := r.Intent.Params.DateTo.Resolved.Month
	yTo := r.Intent.Params.DateTo.Resolved.Year

	if dTo == 0 || mTo == 0 || yTo == 0 {
		return response.ClarifyDate(), nil
	}

	loc := a.TimeWrapper.Now().Location()
	timeFrom := time.Date(yFrom, time.Month(mFrom), dFrom, 0, 0, 0, 0, loc)
	timeTo := time.Date(yTo, time.Month(mTo), dTo, 0, 0, 0, 0, loc)

	if timeFrom.Equal(timeTo) {
		tts := "There is no difference. These are same dates."
		return response.SimpleResponse(tts), nil
	}

	diff := timeFrom.Sub(timeTo).Hours() / 24
	resDays := math.Abs(diff)
	tts := fmt.Sprintf(
		"It's %d days.",
		int(resDays),
	)

	return response.SimpleResponse(tts), nil
}
