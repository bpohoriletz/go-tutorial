package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(11)}
		err := wallet.Withdraw(1)

		assertBalance(t, wallet)

		if nil != err {
			t.Errorf("Expected no error")
		}
	})

	t.Run("Overdraft", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0)}
		err := wallet.Withdraw(Bitcoin(10))

		if nil == err {
			t.Errorf("Expected overdraft error got nil")
		}
	})
}

func assertBalance(t *testing.T, wallet Wallet) {
	if wallet.Balance() != Bitcoin(10) {
		t.Errorf("expected balance to be 10 BTC, got %s", wallet.Balance())
	}
}
