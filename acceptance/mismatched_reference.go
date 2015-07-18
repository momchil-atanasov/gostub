package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/mismatch"

//go:generate gostub MismatchedReference

type MismatchedReference interface {
	Mismatched(wrong.Job) wrong.Job
	Array([3]wrong.Job) [3]wrong.Job
	Slice([]wrong.Job) []wrong.Job
}