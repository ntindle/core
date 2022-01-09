package exchange

import (
	"time"

	"github.com/sonr-io/core/types/go/node/motor/v1"
)

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *motor.OnTransmitDecisionResponse {
	return &motor.OnTransmitDecisionResponse{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *motor.OnTransmitInviteResponse {
	return &motor.OnTransmitInviteResponse{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}
