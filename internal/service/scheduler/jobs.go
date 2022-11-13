package scheduler

import (
	"github.com/hpardora/absence.go/pkg/absence"
	"time"
)

func (s *Scheduler) Calculate() {
	// Once by week, reload all base data
	now := time.Now()
	if now.Weekday() == 1 {
		s.populateUser()
		s.populateCompany()
		s.populateHolidays()
		s.populateReasons()
	}

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

	// TODO Add timers to start and stop

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
