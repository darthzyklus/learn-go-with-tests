package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("first deposit should be get the same amount of balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("multiple deposits should get a balance equal to the sum of the deposit amounts", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		wallet.Deposit(10)

		got := wallet.Balance()
		want := 20

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
