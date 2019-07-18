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

package hc

import (
	"github.com/asdine/storm"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/pborman/uuid"
	"path/filepath"
	"testing"
)

func TestGetHCBlockHeight(t *testing.T) {
	height, err := tw.GetBlockHeight()
	if err != nil {
		t.Errorf("GetBlockHeight failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetBlockHeight height = %d \n", height)
}

func TestHCBlockScanner_GetCurrentBlockHeight(t *testing.T) {
	bs := NewHCBlockScanner(tw)
	header, _ := bs.GetCurrentBlockHeader()
	t.Logf("GetCurrentBlockHeight height = %d \n", header.Height)
	t.Logf("GetCurrentBlockHeight hash = %v \n", header.Hash)
}

func TestGetBlockHeight(t *testing.T) {
	height, _ := tw.GetBlockHeight()
	t.Logf("GetBlockHeight height = %d \n", height)
}

func TestGetLocalNewBlock(t *testing.T) {
	height, hash := tw.GetLocalNewBlock()
	t.Logf("GetLocalBlockHeight height = %d \n", height)
	t.Logf("GetLocalBlockHeight hash = %v \n", hash)
}

func TestSaveLocalBlockHeight(t *testing.T) {
	bs := NewHCBlockScanner(tw)
	header, _ := bs.GetCurrentBlockHeader()
	t.Logf("SaveLocalBlockHeight height = %d \n", header.Height)
	t.Logf("GetLocalBlockHeight hash = %v \n", header.Hash)
	tw.SaveLocalNewBlock(header.Height, header.Hash)
}

func TestGetBlockHash(t *testing.T) {
	//height := GetLocalBlockHeight()
	hash, err := tw.GetBlockHash(1354918)
	if err != nil {
		t.Errorf("GetBlockHash failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetBlockHash hash = %s \n", hash)
}

func TestGetBlock(t *testing.T) {
	raw, err := tw.GetBlock("000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506")
	if err != nil {
		t.Errorf("GetBlock failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetBlock = %v \n", raw)
}

func TestGetTransaction(t *testing.T) {
	raw, err := tw.GetTransaction("7e0c0cdcc101afb7efb94e1539b95e403dddb74b4472967efd2a44dfca68b0b6")
	if err != nil {
		t.Errorf("GetTransaction failed unexpected error: %v\n", err)
		return
	}

	t.Logf("BlockHash = %v \n", raw.BlockHash)
	t.Logf("BlockHeight = %v \n", raw.BlockHeight)
	t.Logf("Blocktime = %v \n", raw.Blocktime)
	t.Logf("Fees = %v \n", raw.Fees)

	t.Logf("========= vins ========= \n")

	for i, vin := range raw.Vins {
		t.Logf("TxID[%d] = %v \n", i, vin.TxID)
		t.Logf("Vout[%d] = %v \n", i, vin.Vout)
		t.Logf("Addr[%d] = %v \n", i, vin.Addr)
		t.Logf("Value[%d] = %v \n", i, vin.Value)
	}

	t.Logf("========= vouts ========= \n")

	for i, out := range raw.Vouts {
		t.Logf("ScriptPubKey[%d] = %v \n", i, out.ScriptPubKey)
		t.Logf("Addr[%d] = %v \n", i, out.Addr)
		t.Logf("Value[%d] = %v \n", i, out.Value)
	}
}

func TestGetTxOut(t *testing.T) {
	raw, err := tw.GetTxOut("7768a6436475ed804344a3711e90e7f10f7db42da8918580c8b669dd63d64cc3", 0)
	if err != nil {
		t.Errorf("GetTxOut failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetTxOut = %v \n", raw)
}

func TestGetTxIDsInMemPool(t *testing.T) {
	txids, err := tw.GetTxIDsInMemPool()
	if err != nil {
		t.Errorf("GetTxIDsInMemPool failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetTxIDsInMemPool = %v \n", txids)
}

func TestHCBlockScanner_scanning(t *testing.T) {

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "miphUAzHHeM1VXGSFnw6owopsQW3jAQZAk"

	//wallet, err := tw.GetWalletInfo(accountID)
	//if err != nil {
	//	t.Errorf("HCBlockScanner_scanning failed unexpected error: %v\n", err)
	//	return
	//}

	bs := NewHCBlockScanner(tw)

	//bs.DropRechargeRecords(accountID)

	bs.SetRescanBlockHeight(1384586)
	//tw.SaveLocalNewBlock(1355030, "00000000000000125b86abb80b1f94af13a5d9b07340076092eda92dade27686")

	//bs.AddAddress(address, accountID)

	bs.ScanBlockTask()
}

func TestHCBlockScanner_Run(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
	)

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "msnYsBdBXQZqYYqNNJZsjShzwCx9fJVSin"

	//wallet, err := tw.GetWalletInfo(accountID)
	//if err != nil {
	//	t.Errorf("HCBlockScanner_Run failed unexpected error: %v\n", err)
	//	return
	//}

	bs := NewHCBlockScanner(tw)

	//bs.DropRechargeRecords(accountID)

	//bs.SetRescanBlockHeight(1384586)

	//bs.AddAddress(address, accountID)

	bs.Run()

	<-endRunning

}

func TestHCBlockScanner_ScanBlock(t *testing.T) {

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "msnYsBdBXQZqYYqNNJZsjShzwCx9fJVSin"

	bs := tw.Blockscanner
	//bs.AddAddress(address, accountID)
	bs.ScanBlock(1384961)
}

func TestHCBlockScanner_ExtractTransaction(t *testing.T) {

	//accountID := "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek"
	//address := "msHemmfSZ3au6h9S1annGcTGrTVryRbSFV"

	//bs := tw.Blockscanner
	////bs.AddAddress(address, accountID)
	//bs.ExtractTransaction(
	//	1435497,
	//	"00000000e271b8234ed2271cb80f1cf2701469a4e02b0536fdce4f4306ff7852",
	//	"c550ae3ffafdda46c13217797dd0aa8ee870727d3e8cab1551d6a3f5e3f7ace0", bs.GetSourceKeyByAddress)

}

func TestWallet_GetRecharges(t *testing.T) {
	accountID := "WFvvr5q83WxWp1neUMiTaNuH7ZbaxJFpWu"
	wallet, err := tw.GetWalletInfo(accountID)
	if err != nil {
		t.Errorf("GetRecharges failed unexpected error: %v\n", err)
		return
	}

	recharges, err := wallet.GetRecharges(false)
	if err != nil {
		t.Errorf("GetRecharges failed unexpected error: %v\n", err)
		return
	}

	t.Logf("recharges.count = %v", len(recharges))
	//for _, r := range recharges {
	//	t.Logf("rechanges.count = %v", len(r))
	//}
}

//func TestHCBlockScanner_DropRechargeRecords(t *testing.T) {
//	accountID := "W4ruoAyS5HdBMrEeeHQTBxo4XtaAixheXQ"
//	bs := NewHCBlockScanner(tw)
//	bs.DropRechargeRecords(accountID)
//}

func TestGetUnscanRecords(t *testing.T) {
	list, err := tw.GetUnscanRecords()
	if err != nil {
		t.Errorf("GetUnscanRecords failed unexpected error: %v\n", err)
		return
	}

	for _, r := range list {
		t.Logf("GetUnscanRecords unscan: %v", r)
	}
}

func TestHCBlockScanner_RescanFailedRecord(t *testing.T) {
	bs := NewHCBlockScanner(tw)
	bs.RescanFailedRecord()
}

func TestFullAddress(t *testing.T) {

	dic := make(map[string]string)
	for i := 0; i < 20000000; i++ {
		dic[uuid.NewUUID().String()] = uuid.NewUUID().String()
	}
}

func TestHCBlockScanner_GetTransactionsByAddress(t *testing.T) {
	coin := openwallet.Coin{
		Symbol:     "HC",
		IsContract: false,
	}
	txExtractDatas, err := tw.Blockscanner.GetTransactionsByAddress(0, 50, coin, "2N7Mh6PLX39japSF76r2MAf7wT7WKU5TdpK")
	if err != nil {
		t.Errorf("GetTransactionsByAddress failed unexpected error: %v\n", err)
		return
	}

	for _, ted := range txExtractDatas {
		t.Logf("tx = %v", ted.Transaction)
	}

}

func TestGetLocalBlock(t *testing.T) {
	db, err := storm.Open(filepath.Join(tw.Config.dbPath, tw.Config.BlockchainFile))
	if err != nil {
		return
	}
	defer db.Close()

	var blocks []*Block
	err = db.All(&blocks)
	if err != nil {
		log.Error("no find")
		return
	}
	log.Info("blocks = ", len(blocks))
}
