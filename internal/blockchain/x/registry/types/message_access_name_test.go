package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAccessName_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAccessName
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAccessName{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAccessName{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}