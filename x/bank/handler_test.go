package bank

import (
	"strings"
	"testing"

	sdk "github.com/evdatsion/cusp-sdk/types"
	abci "github.com/evdatsion/aphelion-dpos-bft/abci/types"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	h := NewHandler(nil)

	res := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, "unrecognized bank message type"))
}
