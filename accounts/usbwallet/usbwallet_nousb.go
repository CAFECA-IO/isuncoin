// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//go:build !usb

package usbwallet

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
)

// ErrTrezorPINNeeded is returned if opening the trezor requires a PIN code.
var ErrTrezorPINNeeded = errors.New("trezor: pin needed")

// ErrTrezorPassphraseNeeded is returned if opening the trezor requires a passphrase
var ErrTrezorPassphraseNeeded = errors.New("trezor: passphrase needed")

// Hub is a stub for the USB wallet hub.
type Hub struct{}

// NewLedgerHub creates a new stub hardware wallet manager for Ledger devices.
func NewLedgerHub() (*Hub, error) {
	return nil, errors.New("usb wallet disabled by nousb build tag")
}

// NewTrezorHubWithHID creates a new stub hardware wallet manager for Trezor devices.
func NewTrezorHubWithHID() (*Hub, error) {
	return nil, errors.New("usb wallet disabled by nousb build tag")
}

// NewTrezorHubWithWebUSB creates a new stub hardware wallet manager for Trezor devices.
func NewTrezorHubWithWebUSB() (*Hub, error) {
	return nil, errors.New("usb wallet disabled by nousb build tag")
}

// Wallets implements accounts.Backend, returning an empty list.
func (hub *Hub) Wallets() []accounts.Wallet {
	return nil
}

// Subscribe implements accounts.Backend, returning a dummy subscription.
func (hub *Hub) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})
}
