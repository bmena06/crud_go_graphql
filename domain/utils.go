package domain

import "regexp"

// ValidateEmail verifica si el email proporcionado es válido.
func ValidateEmail(email string) bool {
	match := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return match.MatchString(email)
}

// ContainsNumbers verifica si la cadena de texto contiene números.
// validateName verifica si el nombre no contiene números.
func ValidateName(name string) bool {
	match, _ := regexp.MatchString(`^[^0-9]+$`, name)
	return match
}
