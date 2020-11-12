package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type TxnOperation string

const (
	TxnOpDeposit    = TxnOperation("DEPOSIT")
	TxnOpWithdrawal = TxnOperation("WITHDRAWAL")
	TxnOpTransfer   = TxnOperation("TRANSFER")
	// TxnOpBalance    = TxnOperation("BALANCE_ENQUIRY")
)

type TxnState string

const (
	TxStateCreated = TxnState("CREATED")
	TxStateFailed  = TxnState("FAILED")
)

type Transaction struct {
	ID        uuid.UUID
	Operation TxnOperation
	Timestamp time.Time
	Amount    float64
	UserID    uuid.UUID
	AccountID uuid.UUID
}

// TxnEvent is a description of a transaction operation event. We have defined operations
// as deposit, withdrawal and transfer. In the end all 3 operations can be modelled as one;
// "transfer" operations.
//
// Example:
// 1. During a deposit, money is moved from an agent's account to the depositor's account
// 2. During a withdrawal, money is moved from the customer withdrawing to the agent's account.
//
// A transfer operation/transaction, needs to have a source and destination and the amount being
// transferred.
type TxnEvent struct {
	Source      TxnCustomer // where money is coming from
	Destination TxnCustomer // where money is going
	// we can further use this field to describe the specific type of transaction/transfer
	TxnType TxnOperation
	// transaction state to track the transaction
	TxnState TxnState
	// amount of money if shillings being transacted
	Amount Shillings
}

// TxnCustomer is a description of a customer involved in a transaction. We can describe them
// by their user id and user type; We have defined a customer being an agent, merchant or subscriber.
type TxnCustomer struct {
	UserID   uuid.UUID
	UserType UserType
}
