package attributes

import "fmt"

type VerificationError struct {
	reason string // some reason description that can be inserted after the "because"
}

func (e VerificationError) Error() string {
	return fmt.Sprintf("Error occured because %s", e.reason)
}
