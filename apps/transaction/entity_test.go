package transaction

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSetSubTotal(t *testing.T) {
	var trx = Transaction{
		ProductPrice: 10_000,
		Amount:       10,
	}
	expected := uint(100_000)

	trx.SetSubTotal()
	trx.SetSubTotal()
	trx.SetSubTotal()
	trx.SetSubTotal()
	trx.SetSubTotal()

	require.Equal(t, expected, trx.SubTotal)
}
func TestGrandTotal(t *testing.T) {
	t.Run("without set sub total first", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}
		expected := uint(100_000)
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})
	t.Run("without platform fee", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}
		expected := uint(100_000)

		trx.SetSubTotal()
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})
	t.Run("with platform fee", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 10_000,
			Amount:       10,
			PlatformFee:  1_000,
		}
		expected := uint(101_000)

		trx.SetSubTotal()
		trx.SetGrandTotal()

		require.Equal(t, expected, trx.GrandTotal)
	})
}

func TestProductJSON(t *testing.T) {
	product := Product{
		Id:    1,
		SKU:   uuid.NewString(),
		Name:  "Product 1",
		Price: 10_000,
	}
	var trx = Transaction{}
	err := trx.SetProductJSON(product)
	require.Nil(t, err)
	require.NotNil(t, trx.ProductJSON)

	productFromTrx, err := trx.GetProduct()
	require.Nil(t, err)
	require.NotEmpty(t, productFromTrx)

	require.Equal(t, product, productFromTrx)

}

func TestTransactionStatus(t *testing.T) {
	type tabletest struct {
		title    string
		expected string
		trx      Transaction
	}

	var tableTests = []tabletest{
		{
			title:    "status created",
			trx:      Transaction{Status: TransactionStatus_Created},
			expected: TRX_CREATED,
		},
		{
			title:    "status on progress",
			trx:      Transaction{Status: TransactionStatus_Progress},
			expected: TRX_ON_PROGRESS,
		},
		{
			title:    "status in delivery",
			trx:      Transaction{Status: TransactionStatus_InDelivery},
			expected: TRX_IN_DELIVERY,
		},
		{
			title:    "status completed",
			trx:      Transaction{Status: TransactionStatus_Completed},
			expected: TRX_COMPLETED,
		},
		{
			title:    "status unknown",
			trx:      Transaction{Status: 0},
			expected: TRX_UNKNOWN,
		},
	}
	for _, test := range tableTests {
		t.Run(test.title, func(t *testing.T) {
			require.Equal(t, test.expected, test.trx.GetStatus())
		})
	}
}
