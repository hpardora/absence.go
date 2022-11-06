package printer

import (
	"github.com/hpardora/absence.go/pkg/absence"
	"sort"
	"time"
)

func (p *Printer) printUser(user *absence.User) {
	p.logger.Infof("")
	p.logger.Infof("USER INFORMATION")
	p.logger.Infof("----------------")
	p.logger.Infof("\tID:\t\t%s", user.ID)
	p.logger.Infof("\tName:\t\t%s %s", user.FirstName, user.LastName)
	p.logger.Infof("\tEmail:\t\t%s", user.Email)
	p.logger.Infof("\tCompanyID:\t%s", user.Company)
	p.logger.Infof("\tTimeZone:\t%s", user.TimeZoneName)
	p.logger.Infof("")
}

func (p *Printer) printCompany(company *absence.Company) {
	p.logger.Infof("")
	p.logger.Infof("COMPANY INFORMATION")
	p.logger.Infof("----------------")
	p.logger.Infof("\tID:\t\t%s", company.Id)
	p.logger.Infof("\tName:\t\t%s", company.Name)
	p.logger.Infof("\tCurrentYear:\t%d", company.CurrentCompanyYear)
	p.logger.Infof("")
	p.logger.Infof("\tAddress 1:\t%s", company.Address1)
	p.logger.Infof("\tAddress 2:\t%s", company.Address2)
	p.logger.Infof("\tPostal Code:\t%s", company.PostalCode)
	p.logger.Infof("\tCity:\t\t%s", company.City)
	p.logger.Infof("\tCountry:\t%s", company.Country)
	p.logger.Infof("")
	p.logger.Infof("\tEmail:\t\t%s", company.Email)
	p.logger.Infof("\tPhone:\t\t%s", company.Phone)
	p.logger.Infof("")
}

func (p *Printer) printHolidays(holidays []absence.Holiday) {
	now := time.Now()
	var dates []*Holiday
	p.logger.Infof("")
	p.logger.Infof("HOLIDAYS")
	p.logger.Infof("----------------")
	p.logger.Infof("\t%s\t%s\t\t%s", "ID", "Date", "Name")
	//p.logger.Infof("\t-------------------------------------------")

	for _, h := range holidays {
		date := time.Time{}
		if len(h.Dates) == 1 {
			date = h.Dates[0]
		} else {
			for _, d := range h.Dates {
				if d.Year() == now.Year() {
					date = d
					break
				}
			}
		}
		dates = append(dates, &Holiday{
			ID:   h.Id,
			Date: date,
			Name: h.Name,
		})
	}
	sort.SliceStable(dates, func(i, j int) bool {
		return dates[i].Date.Before(dates[j].Date)
	})
	for i, d := range dates {
		p.logger.Infof("\t%d\t%s\t%s", i, d.Date.Format("2006-01-02"), d.Name)
	}
	p.logger.Infof("")
}

func (p *Printer) printReasons() {
	p.logger.Infof("")
	p.logger.Infof("REASONS")
	p.logger.Infof("-------")

	for _, r := range p.reasons {
		p.logger.Infof("\t%s\t%t\t%s", r.Id, r.CountsAsWork, r.Name)
	}
}

func (p *Printer) printAbsences(absences []absence.Absence) {
	p.logger.Infof("")
	p.logger.Infof("HOLIDAYS")
	p.logger.Infof("--------")
	for i, a := range absences {
		p.logger.Infof("\t%d\t%s\t%s\t%s", i, a.ReasonId, a.Start.Format("2006-01-02"), a.End.Format("2006-01-02"))
	}
	p.logger.Infof("")
}
