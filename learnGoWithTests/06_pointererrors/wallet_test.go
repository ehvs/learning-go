package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))

		//got := wallet.Balance()

		//fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		//want := Bitcoin(10)

		/*if got != want {
			t.Errorf("got %s want %s", got, want)
		}*/
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))
		checkBalance(t, wallet, Bitcoin(15))
		checkNoError(t, err)

		/*got := wallet.Balance()

		want := Bitcoin(15)

		fmt.Printf("Balance is %d \n", got)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}*/
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		checkError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, startingBalance)
	})
}

func checkError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func checkBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func checkNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but did'nt want one")
	}
}
