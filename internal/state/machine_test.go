package state_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	s "github.com/sonr-io/core/internal/state"
)

// Boilerplate for State Machine
const (
	// States
	CreatingOrder     s.StateType = "CreatingOrder"
	OrderFailed       s.StateType = "OrderFailed"
	OrderPlaced       s.StateType = "OrderPlaced"
	ChargingCard      s.StateType = "ChargingCard"
	TransactionFailed s.StateType = "TransactionFailed"
	OrderShipped      s.StateType = "OrderShipped"

	// Events
	CreateOrder     s.EventType = "CreateOrder"
	FailOrder       s.EventType = "FailOrder"
	PlaceOrder      s.EventType = "PlaceOrder"
	ChargeCard      s.EventType = "ChargeCard"
	FailTransaction s.EventType = "FailTransaction"
	ShipOrder       s.EventType = "ShipOrder"
)

type OrderCreationContext struct {
	items []string
	err   error
}

func (c *OrderCreationContext) String() string {
	return fmt.Sprintf("OrderCreationContext [ items: %s, err: %v ]",
		strings.Join(c.items, ","), c.err)
}

type OrderShipmentContext struct {
	cardNumber string
	address    string
	err        error
}

func (c *OrderShipmentContext) String() string {
	return fmt.Sprintf("OrderShipmentContext [ cardNumber: %s, address: %s, err: %v ]",
		c.cardNumber, c.address, c.err)
}

type CreatingOrderAction struct{}

func (a *CreatingOrderAction) Execute(eventCtx s.EventContext) s.EventType {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Validating, order:", order)
	if len(order.items) == 0 {
		order.err = errors.New("Insufficient number of items in order")
		return FailOrder
	}
	return PlaceOrder
}

type OrderFailedAction struct{}

func (a *OrderFailedAction) Execute(eventCtx s.EventContext) s.EventType {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Order failed, err:", order.err)
	return s.NoOp
}

type OrderPlacedAction struct{}

func (a *OrderPlacedAction) Execute(eventCtx s.EventContext) s.EventType {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Order placed, items:", order.items)
	return s.NoOp
}

type ChargingCardAction struct{}

func (a *ChargingCardAction) Execute(eventCtx s.EventContext) s.EventType {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Validating card, shipment:", shipment)
	if shipment.cardNumber == "" {
		shipment.err = errors.New("Card number is invalid")
		return FailTransaction
	}
	return ShipOrder
}

type TransactionFailedAction struct{}

func (a *TransactionFailedAction) Execute(eventCtx s.EventContext) s.EventType {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Transaction failed, err:", shipment.err)
	return s.NoOp
}

type OrderShippedAction struct{}

func (a *OrderShippedAction) Execute(eventCtx s.EventContext) s.EventType {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Order shipped, address:", shipment.address)
	return s.NoOp
}

func newOrderFSM() *s.StateMachine {
	return &s.StateMachine{
		States: s.States{
			s.Default: s.State{
				Events: s.Events{
					CreateOrder: CreatingOrder,
				},
			},
			CreatingOrder: s.State{
				Action: &CreatingOrderAction{},
				Events: s.Events{
					FailOrder:  OrderFailed,
					PlaceOrder: OrderPlaced,
				},
			},
			OrderFailed: s.State{
				Action: &OrderFailedAction{},
				Events: s.Events{
					CreateOrder: CreatingOrder,
				},
			},
			OrderPlaced: s.State{
				Action: &OrderPlacedAction{},
				Events: s.Events{
					ChargeCard: ChargingCard,
				},
			},
			ChargingCard: s.State{
				Action: &ChargingCardAction{},
				Events: s.Events{
					FailTransaction: TransactionFailed,
					ShipOrder:       OrderShipped,
				},
			},
			TransactionFailed: s.State{
				Action: &TransactionFailedAction{},
				Events: s.Events{
					ChargeCard: ChargingCard,
				},
			},
			OrderShipped: s.State{
				Action: &OrderShippedAction{},
			},
		},
	}
}
func TestOrderFSM(t *testing.T) {
	orderFsm := newOrderFSM()

	// Define the context for order creation.
	creationCtx := &OrderCreationContext{
		items: []string{},
	}

	// Try to create an order with invalid set of items.
	err := orderFsm.SendEvent(CreateOrder, creationCtx)
	if err != nil {
		t.Errorf("Failed to send create order event, err: %v", err)
	}

	// The state machine should enter the OrderFailed state.
	if orderFsm.Current != OrderFailed {
		t.Errorf("Expected the FSM to be in the OrderFailed state, actual: %s",
			orderFsm.Current)
	}

	// Let's fix the order creation context.
	creationCtx = &OrderCreationContext{
		items: []string{"foo", "bar"},
	}

	// Let's now retry the same order with a valid set of items.
	err = orderFsm.SendEvent(CreateOrder, creationCtx)
	if err != nil {
		t.Errorf("Failed to send create order event, err: %v", err)
	}

	// The state machine should enter the OrderPlaced state.
	if orderFsm.Current != OrderPlaced {
		t.Errorf("Expected the FSM to be in the OrderPlaced state, actual: %s",
			orderFsm.Current)
	}

	// Let's now define the context for shipping the order.
	shipmentCtx := &OrderShipmentContext{
		cardNumber: "",
		address:    "123 Foo Street, Bar Baz, QU 45678, USA",
	}

	// Try to charge the card using an invalid card number.
	err = orderFsm.SendEvent(ChargeCard, shipmentCtx)
	if err != nil {
		t.Errorf("Failed to send charge card event, err: %v", err)
	}

	// The state machine should enter the TransactionFailed state.
	if orderFsm.Current != TransactionFailed {
		t.Errorf("Expected the FSM to be in the TransactionFailed state, actual: %s",
			orderFsm.Current)
	}

	// Let's fix the shipment context.
	shipmentCtx = &OrderShipmentContext{
		cardNumber: "0000-0000-0000-0000",
		address:    "123 Foo Street, Bar Baz, QU 45678, USA",
	}

	// Let's now retry the transaction with a valid card number.
	err = orderFsm.SendEvent(ChargeCard, shipmentCtx)
	if err != nil {
		t.Errorf("Failed to send charge card event, err: %v", err)
	}

	// The state machine should enter the OrderShipped state.
	if orderFsm.Current != OrderShipped {
		t.Errorf("Expected the FSM to be in the OrderShipped state, actual: %s",
			orderFsm.Current)
	}

	// Let's try charging the card again for the same order. We should see the
	// event getting rejected as the card has been already charged once and the
	// state machine is in the OrderShipped state.
	err = orderFsm.SendEvent(ChargeCard, shipmentCtx)
	if err != s.ErrEventRejected {
		t.Errorf("Expected the FSM to return a rejected event error")
	}
}
