package utils

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}
