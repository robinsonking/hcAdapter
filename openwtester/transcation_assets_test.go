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

package openwtester

import (
	"github.com/blocktree/openwallet/openw"
	"testing"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract,
	feeSupportAccount *openwallet.FeesSupportAccount) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, feeSupportAccount)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer(t *testing.T) {

	//mi9qsHKMqtrgnbxg7ifdPMk1LsFmen4xNn
	//mjbVUivVbQuh4PdW4Z7mxvQb99FJ62EQdC
	//mjjH75fZrDbAQUg9s3mq9qKwP6jDzryijP
	//mk44BMmg8YPRjxui7SvZ9GAv4VxMcwJhFY
	//mnSPBLZA5X75R8mnYRsjjTnzz38NmL9RPd
	//mqzv5TsejXzCRQB3hTKk6emvjocSoq291Z

	tm := testInitWalletManager()
	walletID := "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC"
	accountID := "FqQBQ8Bn26GogR7UAu6e2ZVhrYYmKUpmBS7CSM1KLTTZ"
	to := "mkUDixUXqgUqGPEGEtxJUoQBvpDeSnjL2Z"

	//accountID := "EPxkNBu6iMospC6aHQppv36UGY4mb1WqUE7oNZ7Xp9Df"
	//to := "mi9qsHKMqtrgnbxg7ifdPMk1LsFmen4xNn"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.02", "", nil)
	if err != nil {
		return
	}

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


func TestTransfer_OMNI(t *testing.T) {

	//mkJrhf8Bp3RWfL5eyatcfqPBDDUUXgHQYm
	//mkUDixUXqgUqGPEGEtxJUoQBvpDeSnjL2Z
	//mkzbZ9557cvXAEpvCSQypDoXRiVb8dCtEV
	//mqUoroU8ttLBNgCatzAgBtvyMYM8M1tN8d
	//mu9ZjobWsHr2KxLG3ELJTRkYJzRLsSqNCX
	//mwawxdBn9w4CPxic961vPnyj9HqDVGnkth

	tm := testInitWalletManager()
	walletID := "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC"
	accountID := "FqQBQ8Bn26GogR7UAu6e2ZVhrYYmKUpmBS7CSM1KLTTZ"
	to := "mkJrhf8Bp3RWfL5eyatcfqPBDDUUXgHQYm"

	contract := openwallet.SmartContract{
		Address:  "2",
		Symbol:   "BTC",
		Name:     "Test Omni",
		Token:    "Omni",
		Decimals: 8,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1000", "", &contract)
	if err != nil {
		return
	}

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

func TestSummary(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC"
	accountID := "EPxkNBu6iMospC6aHQppv36UGY4mb1WqUE7oNZ7Xp9Df"
	summaryAddress := "mi9qsHKMqtrgnbxg7ifdPMk1LsFmen4xNn"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}

func TestSummary_OMNI(t *testing.T) {

	tm := testInitWalletManager()
	walletID := "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC"
	accountID := "EPxkNBu6iMospC6aHQppv36UGY4mb1WqUE7oNZ7Xp9Df"
	summaryAddress := "mi9qsHKMqtrgnbxg7ifdPMk1LsFmen4xNn"

	contract := openwallet.SmartContract{
		Address:  "2",
		Symbol:   "BTC",
		Name:     "Test Omni",
		Token:    "Omni",
		Decimals: 8,
	}

	feesSupport := openwallet.FeesSupportAccount{
		AccountID: "FqQBQ8Bn26GogR7UAu6e2ZVhrYYmKUpmBS7CSM1KLTTZ",
		//FixSupportAmount: "0.01",
		FeesSupportScale: "1",
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, &contract, &feesSupport)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}
