package coingecko

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchCoinLists(t *testing.T) {
	tests := []struct {
		name              string
		lengthGreaterThan int
	}{
		{
			name:              "just_get_coin_lists",
			lengthGreaterThan: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchCoinLists()

			t.Logf("CoinLists=%+v; err=%+v", got, err)

			require.NoError(t, err)
			require.Equal(t, true, (len(*got) > tt.lengthGreaterThan))
		})
	}
}
