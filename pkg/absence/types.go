package absence

import "time"

type User struct {
	ID                     string    `json:"_id"`
	Company                string    `json:"company"`
	Email                  string    `json:"email"`
	FirstName              string    `json:"firstName"`
	LastName               string    `json:"lastName"`
	EmploymentStartDate    time.Time `json:"employmentStartDate"`
	HolidayCountryLanguage string    `json:"holidayCountryLanguage"`
	OauthGoogleImageurl    string    `json:"oauthGoogleImageurl"`
	TeamIds                []string  `json:"teamIds"`
	DepartmentId           string    `json:"departmentId"`
	LocationId             string    `json:"locationId"`
	HolidayIds             []string  `json:"holidayIds"`
	TimeZoneName           string    `json:"timeZone"`
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

type HolidaysResponse struct {
	Skip       int       `json:"skip"`
	Limit      int       `json:"limit"`
	Count      int       `json:"count"`
	TotalCount int       `json:"totalCount"`
	Data       []Holiday `json:"data"`
}

type Holiday struct {
	Dates            []time.Time `json:"dates"`
	DayType          int         `json:"dayType"`
	IsMandatoryLeave bool        `json:"isMandatoryLeave"`
	Id               string      `json:"_id"`
	Name             string      `json:"name"`
	Repeating        bool        `json:"repeating"`
	Date             time.Time   `json:"date"`
	LocationIds      []string    `json:"locationIds"`
}

type AbsencesResponse struct {
	Skip          int       `json:"skip"`
	Limit         int       `json:"limit"`
	Count         int       `json:"count"`
	TotalCount    int       `json:"totalCount"`
	Data          []Absence `json:"data"`
	ResponseModel string    `json:"responseModel"`
}

type Absence struct {
	Id           string    `json:"_id"`
	Status       int       `json:"status"`
	DaysCount    float64   `json:"daysCount"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	ReasonId     string    `json:"reasonId"`
	AssignedToId string    `json:"assignedToId"`
	ApproverId   string    `json:"approverId"`
	CanBeEdited  bool      `json:"canBeEdited,omitempty"`
	CanBeDeleted bool      `json:"canBeDeleted,omitempty"`
}

type ReasonsResponse struct {
	Skip       int      `json:"skip"`
	Limit      int      `json:"limit"`
	Count      int      `json:"count"`
	TotalCount int      `json:"totalCount"`
	Data       []Reason `json:"data"`
}

type Reason struct {
	RequiresApproval       bool          `json:"requiresApproval"`
	ReducesDays            bool          `json:"reducesDays"`
	EmailList              []interface{} `json:"emailList"`
	IsPublic               bool          `json:"isPublic"`
	CountsAsWork           bool          `json:"countsAsWork"`
	IsHourly               bool          `json:"isHourly"`
	RespectsWorkingHours   bool          `json:"respectsWorkingHours"`
	AllowOverlapping       bool          `json:"allowOverlapping"`
	DaysNotice             interface{}   `json:"daysNotice"`
	OverrideMandatoryLeave bool          `json:"overrideMandatoryLeave"`
	RequiresFullDays       bool          `json:"requiresFullDays"`
	ShowDoctorsAppointment bool          `json:"showDoctorsAppointment"`
	Id                     string        `json:"_id"`
	Company                string        `json:"company"`
	Name                   string        `json:"name"`
	SortIndex              int           `json:"sortIndex"`
	Modified               time.Time     `json:"modified"`
	ColorId                string        `json:"colorId"`
	AllowanceTypeId        interface{}   `json:"allowanceTypeId"`
	UserNotificationIds    []interface{} `json:"userNotificationIds"`
	SelectableByIds        []string      `json:"selectableByIds"`
}
