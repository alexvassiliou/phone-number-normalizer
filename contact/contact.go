package contact

import "strings"

// Number is a struct defining the phone numbers table in the database
type Number struct {
	ID          int
	PhoneNumber string
}

// Normalize numbers into the standard format
func (n *Number) Normalize() {
	digits := "0123456789"
	var result []rune
	for _, r := range n.PhoneNumber {
		if strings.ContainsRune(digits, r) {
			result = append(result, r)
		}
	}
	n.PhoneNumber = string(result)
}
