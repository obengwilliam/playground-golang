package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		// want := Bitcoin(10)
		var want Bitcoin = 10
		assertBalance(t, wallet, want)
	})
	t.Run("withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.WithDraw(Bitcoin(20))
		want := Bitcoin(0)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingFunds := Bitcoin(20)
		wallet := Wallet{startingFunds}
		err := wallet.WithDraw(Bitcoin(200))
		assertBalance(t, wallet, startingFunds)
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s instead of %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Error("got error instead of nil")
	}
}
func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("expected error but got nil")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
