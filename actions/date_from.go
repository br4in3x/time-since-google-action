package actions

import (
	"errors"
	"fmt"
	"time"

	"github.com/br4in3x/google-action-time-since/internal/request"
	"github.com/br4in3x/google-action-time-since/internal/response"
	"github.com/br4in3x/google-action-time-since/internal/util"
)

type DateFromToAction struct {
	TimeWrapper util.Time
}

func NewDateFromToAction(timeWrapper util.Time) *DateFromToAction {
	return &DateFromToAction{timeWrapper}
}

func (a *DateFromToAction) Invoke(r request.Request) (*response.Simple, error) {
	day := r.Intent.Params.DateFrom.Resolved.Day
	month := r.Intent.Params.DateFrom.Resolved.Month
	year := r.Intent.Params.DateFrom.Resolved.Year

	if day == 0 || month == 0 || year == 0 {
		return nil, errors.New("invalid date input")
	}

	now := a.TimeWrapper.Now()
	loc := now.Location()

	tts := ""
	resDays := 0.0
	timeFrom := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
	timeNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	if timeFrom.Before(timeNow) {
		resDays = timeNow.Sub(timeFrom).Hours() / 24
		tts = fmt.Sprintf("%d days have passed since that time.", int(resDays))
	}

	if timeFrom.After(timeNow) {
		resDays = timeFrom.Sub(timeNow).Hours() / 24
		tts = fmt.Sprintf("%d days until %s", int(resDays), r.Intent.Params.DateFrom.Original)
	}

	if timeFrom.Equal(timeNow) {
		tts = fmt.Sprintf("0 days. The %s is today.", r.Intent.Params.DateFrom.Original)
	}

	return response.SimpleResponse(tts), nil
}