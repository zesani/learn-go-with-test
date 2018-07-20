package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// wallet := Wallet{}
	// wallet.Deposit(Bitcoin(10))
	// got := wallet.Balance()
	// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	// want := Bitcoin(20)

	// if got != want {
	// 	t.Errorf("got %d want %d", got, want)
	// }

	// if got != want {
	// 	t.Errorf("got %s want %s", got, want)
	// }
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Error("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}
	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Error("didn't get an error but wanted one")
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
}
