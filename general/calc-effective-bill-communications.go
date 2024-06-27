package general

import (
	"fmt"
	"slices"

	"github.com/golang-collections/collections/stack"
)

type Bill struct {
	patientId string
	amount    float32
	timestamp int32
}

type channelType int

const (
	email     channelType = iota
	text                  = iota
	paperMail             = iota
)

type Notice struct {
	patientId     string
	channel       channelType
	sentTimestamp int32
}

type Payment struct {
	patientId     string
	amount        float32
	paidTimestamp int32
}

type Event interface{}

// handleBillPayment will update the given bill by the given payment
// if the payment amount reduces the bill to zero (0) or less, the
// bill's amount will be 0 and the payment will be updated with the
// remaining amount.
// returns true if the bill was completely paid off
func handleBillPayment(bill Bill, payment Payment) bool {
	if payment.amount >= bill.amount {
		payment.amount = payment.amount - bill.amount
		bill.amount = 0
	} else {
		bill.amount = bill.amount - payment.amount
		payment.amount = 0
	}

	return bill.amount == 0
}

func calculateEffectiveness(bills []Bill, notices []Notice, payments []Payment) (int, int, int) {
	emailScore := 0
	textScore := 0
	paperMailScore := 0

	fmt.Printf("bills: %v, notices: %v, payments: %v\n\n", bills, notices, payments)
	events := make([]Event, len(bills)+len(notices)+len(payments))
	eventIndex := 0
	for _, b := range bills {
		events[eventIndex] = b
		eventIndex += 1
	}
	for _, n := range notices {
		events[eventIndex] = n
		eventIndex += 1
	}
	for _, p := range payments {
		events[eventIndex] = p
		eventIndex += 1
	}
	fmt.Printf("events: %v\n\n", events)
	slices.SortFunc(events, func(a, b Event) int {
		getEventTimestamp := func(event Event) int {
			switch e := event.(type) {
			case Bill:
				return int(e.timestamp)
			case Notice:
				return int(e.sentTimestamp)
			case Payment:
				return int(e.paidTimestamp)
			}
			return 0
		}
		timestampA := getEventTimestamp(a)
		timestampB := getEventTimestamp(b)
		return timestampA - timestampB
	})
	fmt.Printf("events: %v\n", events)

	patientBillPaymentStackMap := make(map[string]*stack.Stack)
	for _, _e := range events {
		patientId := ""
		switch e := _e.(type) {
		case Notice:
			notice_type_str := "email"
			switch e.channel {
			case 1:
				notice_type_str = "text"
			case 2:
				notice_type_str = "paper mail"
			}
			fmt.Printf("processing notice of type %v\n", notice_type_str)
			patientId = e.patientId
		case Bill:
			fmt.Printf("processing bill of %v\n", e.amount)
			patientId = e.patientId
		case Payment:
			fmt.Printf("processing payment of %v\n", e.amount)
			patientId = e.patientId
		}

		if nil == patientBillPaymentStackMap[patientId] {
			patientBillPaymentStackMap[patientId] = stack.New()
		}
		s := patientBillPaymentStackMap[patientId]
		top := s.Peek()
		if top == nil {
			switch e := _e.(type) {
			case Notice:
				continue
			default:
				s.Push(e)
			}
		} else {
			switch e := _e.(type) {
			case Bill:
				switch t := top.(type) {
				case Payment:
					handleBillPayment(e, t)
					if t.amount == 0 {
						s.Pop()
					}
					if e.amount == 0 {
						continue // the bill has been cancelled by a previous credited payment
					}
				}
				s.Push(e)
			case Notice:
				switch top.(type) {
				case Payment:
					continue // discard because it had no influence on the payment
				}
				s.Push(e)
			case Payment:
				payment := e
				for top != nil {
					switch t := top.(type) {
					case Notice:
						notice := t
						switch notice.channel {
						case email:
							emailScore += 1
						case text:
							textScore += 1
						case paperMail:
							paperMailScore += 1
						}
						s.Pop()
						top = s.Peek()
					case Bill:
						// a payment is being made to one or more
						// bills, let's apply the payment to the remaining
						// bills
						bill := t
						handleBillPayment(bill, payment)
						if bill.amount == 0 {
							// the bill has been paid off, let's remove it
							s.Pop()
							top = s.Peek()
						}
						if payment.amount == 0 || bill.amount > 0 {
							// the payment has been exhausted or the existing bill
							// still has money owed, let's stop popping
							top = nil
						}

					case Payment:
						// the top most element is a payment (there are no more bills)
						// add more credit to the patient's account
						t.amount += e.amount
						top = nil
					}
				}
			}
		}
	}

	return emailScore, textScore, paperMailScore
}
