package general

import (
	"fmt"
	"slices"
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
	// patientBillPaymentStackMap := make(map[stack.Stack]string)
	// billPaymentStack := stack.New()

	// noticeIndex := 0
	// paymentIndex := 0

	// for _, b := range bills {
	// 	billPaymentStack, ok := patientBillPaymentStackMap[b.patientId]
	// 	if !ok {
	// 		billPaymentStack = stack.New()
	// 		patientBillPaymentStackMap[b.patientId] = billPaymentStack
	// 	}
	// 	billPaymentStack.Push(b)

	// 	for paymentIndex, payment := range payments {
	// 		if payment.patientId == b.patientId
	// 	}
	// 	// for noticeIndex, notice := range notices {
	// 	// 	if notice.patientId == b.patientId {
	// 	// 		if notice.sentTimestamp <= b.timestamp {
	// 	// 			billPaymentStack.Push(notice)
	// 	// 		} else if not
	// 	// 	}
	// 	// }
	// }
	return emailScore, textScore, paperMailScore
}
