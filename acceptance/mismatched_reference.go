package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/mismatch"

//go:generate gostub MismatchedReference

type MismatchedReference interface {
	Mismatched(wrong.Job)
}
