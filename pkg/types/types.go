package types

type Money int64

type PaymentCategory string

type PaymentStatus string

type Phone string

const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}
