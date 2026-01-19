package validations

import (
	"regexp"
	"slices"

	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/domain_err"
	"github.com/marcocesar1/Go-Service-Omnicloud/src/internal/domain/models"
)

func ValidateStatus(status models.PeopleStatus) error {
	validStatuses := []models.PeopleStatus{
		models.StatusIn,
		models.StatusOut,
	}
	if !slices.Contains(validStatuses, status) {
		return domain_err.InvalidStatus
	}

	return nil
}

func ValidatePeople(people *models.People) error {
	if people.Name == "" {
		return domain_err.InvalidPeopleFieldError(domain_err.ErrNameRequired)
	}

	if len(people.Name) < 3 || len(people.Name) > 80 {
		return domain_err.InvalidPeopleFieldError(domain_err.ErrNameInvalidLength)
	}

	if people.Email == "" {
		return domain_err.InvalidPeopleFieldError(domain_err.ErrEmailRequired)
	}

	if !isEmailValid(people.Email) {
		return domain_err.InvalidPeopleFieldError(domain_err.ErrEmailInvalid)
	}

	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
