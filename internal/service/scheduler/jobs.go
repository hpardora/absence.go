package scheduler

import (
	"fmt"
	"github.com/hpardora/absence.go/pkg/absence"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func (s *Scheduler) Calculate() {
	// Once by week, reload all base data
	now := time.Now()

	// Check if today is a week working day
	if !s.isWorkingWeekDay(now) {
		s.logger.Info("today %d is not a working week day!!!!!!! %t", now.Weekday(), s.conf.WorkingDays)
		return
	}
	if s.todayHaveAbsence(now) {
		s.logger.Info("today have absence!!!")
		return
	}
	if s.todayIsHoliday(now) {
		s.logger.Info("today is holiday!!!")
		return
	}

	// Calculate durations
	startDuration, err := calculateDurationFromString(s.conf.StartHour, now, s.conf.RandomMinutes)
	if err != nil {
		return
	}
	endDuration, err := calculateDurationFromString(s.conf.EndHour, now, s.conf.RandomMinutes)
	if err != nil {
		return
	}
	willWorkFrom := time.Now().Add(startDuration)
	willWorkTo := time.Now().Add(endDuration)
	diff := willWorkTo.Sub(willWorkFrom)
	msg := fmt.Sprintf("hello %s! you will work from %s to %s for a total of %v",
		s.user.FirstName,
		willWorkFrom.Format("2006-01-02 15:04:05"),
		willWorkTo.Format("2006-01-02 15:04:05"),
		diff,
	)
	s.logger.Info(msg)
	s.notifyToTelegram(msg)

	wg.Add(2)
	s.manageClockIn(startDuration)
	s.manageClockOut(endDuration)
	wg.Wait()
	s.logger.Info("for today all is done!")
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

func calculateDurationFromString(baseDuration string, now time.Time, randomMinutes int) (time.Duration, error) {
	splited := strings.Split(baseDuration, ":")
	hours, err := strconv.Atoi(splited[0])
	if err != nil {
		return 0, err
	}

	mins, err := strconv.Atoi(splited[1])
	if err != nil {
		return 0, err
	}

	randStartMin := rand.Intn(randomMinutes)
	mins += randStartMin

	requestedTime := time.Date(now.Year(), now.Month(), now.Day(), hours, mins, rand.Intn(60), 0, now.Location())

	resultTime := requestedTime.Sub(now)
	return resultTime, nil
}
