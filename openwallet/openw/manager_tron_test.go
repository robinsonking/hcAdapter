/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openw

import (
	"fmt"
	"testing"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

func TestTronWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO KITTY", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("wallet:", nw)
	fmt.Println("key:", key)

}

func TestTronWalletManager_GetWalletInfo(t *testing.T) {
	tm := testInitWalletManager()
	wallet, err := tm.GetWalletInfo(testApp, "W1eRr8nRrawkQ1Ayf1XKPCjmvKk8aLGExu")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	fmt.Println("wallet:", wallet)
}

func TestTronWalletManager_CreateAssetsAccount(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W1eRr8nRrawkQ1Ayf1XKPCjmvKk8aLGExu"
	account := &openwallet.AssetsAccount{Alias: "HELLO KITTY", WalletID: walletID, Required: 1, Symbol: "TRX", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println("account:", account)
	fmt.Println("account PublicKey:", account.PublicKey)
	fmt.Println("address:", address)

	tm.CloseDB(testApp)
}

// account: &{W33vxQiNcgjJgMvowsNerXao6LZjwR61zp HELLO KITTY GEGdASep1uA7RBarNNZuJjgnE8T3DyJGTRGz4JfNE4Me 1 m/44'/88'/1' owpubeyoV6FsMp1uKkCW8tJ3pECGDhvRqwABVHoUVBLrxG4KKBhm6gjK1t6G1qLGEcZpuvUH9rbuZBYjo8FTCVn8K4aR24KUXryvj5qtJJ7d3zT3Cfs7CN [owpubeyoV6FsMp1uKkCW8tJ3pECGDhvRqwABVHoUVBLrxG4KKBhm6gjK1t6G1qLGEcZpuvUH9rbuZBYjo8FTCVn8K4aR24KUXryvj5qtJJ7d3zT3Cfs7CN]  1 TRON 0  true  0 <nil>} [0m
// address: &{GEGdASep1uA7RBarNNZuJjgnE8T3DyJGTRGz4JfNE4Me TDG9rDZfoHqdT5CotrUQ5ukSfZujsBKWTu 02851a8fface19bda036dcbbf18e3c73f6c322cc3c70d86b28ed478f5af9a949ac   0 m/44'/88'/1'/0/0 false tron  false  1542107768 false  <nil>}

func TestTronWalletManager_GetAssetsAccountList(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W1eRr8nRrawkQ1Ayf1XKPCjmvKk8aLGExu"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		fmt.Println("\taccount[", i, "] :", w)
	}
	fmt.Println("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestTronWalletManager_CreateAddress(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W1eRr8nRrawkQ1Ayf1XKPCjmvKk8aLGExu"
	accountID := "8XPSHP5cR16D4b1V225xig3sgNa45e8Y3P5AbeCzR5gr" //"GEGdASep1uA7RBarNNZuJjgnE8T3DyJGTRGz4JfNE4Me"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 3)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println("address:", address)

	tm.CloseDB(testApp)
}

func TestTronWalletManager_GetAddressList(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "CfRjWjct569qp7oygSA2LrsAoTrfEB8wRk3sHGUj9Erm"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		fmt.Println("\taddress[", i, "] :", w)
	}
	fmt.Println("address count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_GetTRC20TokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "CfRjWjct569qp7oygSA2LrsAoTrfEB8wRk3sHGUj9Erm"

	contract := openwallet.SmartContract{
		Address:  "THvZvKPLHKLJhEFYKiyqj6j8G8nGgfg7ur",
		Symbol:   "TRX",
		Name:     "DICE",
		Token:    "DICE",
		Decimals: 0,
		Protocol: "trc20",
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}


func TestWalletManager_GetTRC10TokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "CfRjWjct569qp7oygSA2LrsAoTrfEB8wRk3sHGUj9Erm"

	contract := openwallet.SmartContract{
		Address:  "1002000",
		Symbol:   "TRX",
		Name:     "BitTorrent",
		Token:    "BTT",
		Decimals: 6,
		Protocol: "trc10",
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}


func TestTransfer_TRC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "CfRjWjct569qp7oygSA2LrsAoTrfEB8wRk3sHGUj9Erm"
	to := "TRJJ9Mq4aMjdmKWpTDJAgbYNoY2P9Facg5"

	contract := openwallet.SmartContract{
		Address:  "THvZvKPLHKLJhEFYKiyqj6j8G8nGgfg7ur",
		Symbol:   "TRX",
		Name:     "TRONdice",
		Token:    "DICE",
		Decimals: 6,
		Protocol: "trc20",
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1", "", &contract)
	if err != nil {
		return
	}
	//log.Infof("rawHex: %+v", rawTx.RawHex)
	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

}


func TestTransfer_TRC10(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "CfRjWjct569qp7oygSA2LrsAoTrfEB8wRk3sHGUj9Erm"
	to := "TRJJ9Mq4aMjdmKWpTDJAgbYNoY2P9Facg5"

	contract := openwallet.SmartContract{
		Address:  "1002000",
		Symbol:   "TRX",
		Name:     "BitTorrent",
		Token:    "BTT",
		Decimals: 6,
		Protocol: "trc10",
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "4", "", &contract)
	if err != nil {
		return
	}
	log.Infof("rawHex: %+v", rawTx.RawHex)
	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

}


func TestTransfer_TRX(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "CfRjWjct569qp7oygSA2LrsAoTrfEB8wRk3sHGUj9Erm"
	to := "TRJJ9Mq4aMjdmKWpTDJAgbYNoY2P9Facg5"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1", "", nil)
	if err != nil {
		return
	}
	log.Infof("rawHex: %+v", rawTx.RawHex)
	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

}
