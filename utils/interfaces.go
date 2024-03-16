package utils

type Integer32 int32

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}
