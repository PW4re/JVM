package class_file

import "fmt"

type VerificationError struct {
	reason string // some text that can be inserted after the "because"
}

func (e VerificationError) Error() string {
	return fmt.Sprintf("Error occured because %s", e.reason)
}
