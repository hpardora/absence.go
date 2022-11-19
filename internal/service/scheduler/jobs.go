package scheduler

import (
	"github.com/hpardora/absence.go/pkg/absence"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (s *Scheduler) Calculate() {
	// Once by week, reload all base data
	now := time.Now().UTC()

	// Check if today is a week working day
	if !s.isWorkingWeekDay(now) {
		s.logger.Infof("today %d is not a working week day!!!!!!! %t", now.Weekday(), s.conf.WorkingDays)
		return
	}
	if s.todayHaveAbsence(now) {
		s.logger.Infof("today have absence!!!")
		return
	}
	if s.todayIsHoliday(now) {
		s.logger.Infof("today is holiday!!!")
		return
	}
	s.notifyToTelegram("Sorry but Today you must work!!")

	startDuration, err := calculateDurationFromString(s.conf.StartHour, now)
	if err != nil {
		return
	}
	endDuration, err := calculateDurationFromString(s.conf.EndHour, now)
	if err != nil {
		return
	}
	s.manageClockIn(startDuration)
	s.manageClockOut(endDuration)
}

func (s *Scheduler) isWorkingWeekDay(now time.Time) bool {
	for _, v := range s.conf.WorkingDays {
		if v == now.Weekday() {
			return true
		}
	}
	return false
}

func (s *Scheduler) todayHaveAbsence(now time.Time) bool {
	for _, abs := range s.absences {
		if now.After(abs.Start) && now.Before(abs.End) {
			reason := s.getReasonDetail(abs.ReasonId)
			if reason != nil {
				if !reason.CountsAsWork {
					return true
				}
			}
			return false
		}
	}
	return false
}

func (s *Scheduler) getReasonDetail(reasonID string) *absence.Reason {
	for _, r := range s.reasons {
		if r.Id == reasonID {
			return &r
		}
	}
	return nil
}

func (s *Scheduler) todayIsHoliday(now time.Time) bool {
	for _, h := range s.clearHolidays() {
		y, m, d := now.Date()
		if h.Date.Year() == y && h.Date.Month() == m && h.Date.Day() == d {
			return true
		}
	}
	return false
}

func calculateDurationFromString(baseDuration string, now time.Time) (time.Duration, error) {
	adding := time.Now()

	splited := strings.Split(baseDuration, ":")
	hours, err := strconv.Atoi(splited[0])
	if err != nil {
		return 0, err
	}

	mins, err := strconv.Atoi(splited[1])
	if err != nil {
		return 0, err
	}

	toAddHours := hours - adding.Hour()
	adding.Add(time.Duration(toAddHours) * time.Hour)

	toAddMins := mins - adding.Minute()
	randStartMin := rand.Intn(10 - 1)
	toAddMins *= randStartMin

	adding.Add(time.Duration(toAddMins) * time.Minute)

	resultTime := adding.Sub(now)
	return resultTime, nil
}
