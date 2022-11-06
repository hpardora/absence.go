package printer

import "errors"

var (
	ErrRetrieveUserInformation    = errors.New("unable to retrieve user information")
	ErrRetrieveCompanyInformation = errors.New("unable to retrieve company information")
)
