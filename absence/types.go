package absence

import "time"

type User struct {
	ID        string `json:"_id"`
	Company   string `json:"company"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Company struct {
	Id                  string `json:"_id"`
	Name                string `json:"name"`
	Ident               int    `json:"ident"`
	Email               string `json:"email"`
	Region              string `json:"region"`
	Country             string `json:"country"`
	DefaultVacationDays int    `json:"defaultVacationDays"`
	PaymentFirstName    string `json:"paymentFirstName"`
	PaymentLastName     string `json:"paymentLastName"`
	PaymentEmail        string `json:"paymentEmail"`
	Address1            string `json:"address1"`
	Address2            string `json:"address2"`
	City                string `json:"city"`
	PostalCode          string `json:"postalCode"`
	PaymentCountry      string `json:"paymentCountry"`
	PaymentTaxid        string `json:"paymentTaxid"`
	Phone               string `json:"phone"`
	CurrentCompanyYear  int    `json:"currentCompanyYear"`
}

type Holidays struct {
	Skip       int             `json:"skip"`
	Limit      int             `json:"limit"`
	Count      int             `json:"count"`
	TotalCount int             `json:"totalCount"`
	Data       []HolidayDetail `json:"data"`
}

type HolidayDetail struct {
	Dates            []time.Time `json:"dates"`
	DayType          int         `json:"dayType"`
	IsMandatoryLeave bool        `json:"isMandatoryLeave"`
	Id               string      `json:"_id"`
	Name             string      `json:"name"`
	Repeating        bool        `json:"repeating"`
	Date             time.Time   `json:"date"`
	LocationIds      []string    `json:"locationIds"`
}

type HolidayRegion struct {
	Skip       int                   `json:"skip"`
	Limit      int                   `json:"limit"`
	Count      int                   `json:"count"`
	TotalCount int                   `json:"totalCount"`
	Data       []HolidayRegionDetail `json:"data"`
}

type HolidayRegionDetail struct {
	Id         string   `json:"_id"`
	Name       string   `json:"name"`
	HolidayIds []string `json:"holidayIds"`
	Src        string   `json:"src,omitempty"`
}
