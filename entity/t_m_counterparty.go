package entity

import "time"

type T_M_Counterparty struct {
	CounterpartyID        string
	CounterpartyName      string
	CounterpartyShortname string
	CounterpartySpell     *string
	Email                 *string
	Mobile                *string
	Tel                   *string
	Zip                   *string
	Fax                   *string
	PicName               *string
	PicSpell              *string
	Address1              *string
	Address2              *string
	CounterpartyType      string
	DebtsAndCredits       string
	ManageAccount         string
	PaymentCutoffDate     *int
	PaymentMethod         *string
	PaymentCycle          *string
	ReceiptCutoffDate     *int
	ReceiptMethod         *string
	ReceiptCycle          *string
	InsRep                string
	InsDate               time.Time
	UpdRep                string
	UpdDate               time.Time
	DelFlg                string
	Version               int
}
