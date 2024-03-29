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
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/log"
	"github.com/codeskyblue/go-sh"
	"github.com/shopspring/decimal"
	"math"
	"path/filepath"
	"testing"
	"time"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", "HC.ini")
	//	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		panic(err)
	}
	wm.LoadAssetsConfig(c)
	//wm.ExplorerClient.Debug = true
	//wm.WalletClient.Debug = true
	wm.OnmiClient.Debug = true
	return wm
}

func TestGetCoreWalletinfo(t *testing.T) {
	res, err := tw.GetCoreWalletinfo()
	t.Log(res, err)
}

func TestKeyPoolRefill(t *testing.T) {

	//解锁钱包
	err := tw.UnlockWallet("111111", 12000)
	if err != nil {
		t.Errorf("KeyPoolRefill failed unexpected error: %v\n", err)
	}

	err = tw.KeyPoolRefill(10000)
	if err != nil {
		t.Errorf("KeyPoolRefill failed unexpected error: %v\n", err)
	}
}

func TestCreateReceiverAddress(t *testing.T) {

	tests := []struct {
		account string
		tag     string
	}{
		{
			account: "default",
			tag:     "normal",
		},
	}

	for i, test := range tests {

		a, err := tw.CreateReceiverAddress(test.account)
		if err != nil {
			t.Errorf("CreateReceiverAddress[%d] failed unexpected error: %v", i, err)
		} else {
			t.Logf("CreateReceiverAddress[%d] address = %v", i, a)
		}

	}

}

func TestWalletManager_CreateMultiSig(t *testing.T) {
	addr, redeemScript, err := tw.AddMultiSigAddress(2, []string{"", ""})
	if err != nil {
		t.Errorf("CreateMultiSig failed unexpected error: %v", err)
		return
	}
	fmt.Printf("address: %s\n", addr)
	fmt.Printf("redeemScript: %s\n", redeemScript)
}

func TestGetAddressesByAccount(t *testing.T) {
	addresses, err := tw.GetAddressesByAccount("")
	if err != nil {
		t.Errorf("GetAddressesByAccount failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addresses {
		t.Logf("GetAddressesByAccount address[%d] = %s\n", i, a)
	}
}

func TestCreateBatchAddress(t *testing.T) {
	_, _, err := tw.CreateBatchAddress("WDHupMjR3cR2wm97iDtKajxSPCYEEddoek", "1234qwer", 12500)
	if err != nil {
		t.Errorf("CreateBatchAddress failed unexpected error: %v\n", err)
		return
	}
}

func TestEncryptWallet(t *testing.T) {
	err := tw.EncryptWallet("111111")
	if err != nil {
		t.Errorf("EncryptWallet failed unexpected error: %v\n", err)
		return
	}
}

func TestUnlockWallet(t *testing.T) {
	err := tw.UnlockWallet("111111", 1200000)
	if err != nil {
		t.Errorf("UnlockWallet failed unexpected error: %v\n", err)
		return
	}
}

//func TestCreateNewWallet(t *testing.T) {
//	_, _, err := tw.CreateNewWallet("HC", "123456")
//	if err != nil {
//		t.Errorf("CreateNewWallet failed unexpected error: %v\n", err)
//		return
//	}
//}

func TestGetWalletKeys(t *testing.T) {
	wallets, err := tw.GetWallets()
	if err != nil {
		t.Errorf("GetWalletKeys failed unexpected error: %v\n", err)
		return
	}

	for i, w := range wallets {
		t.Logf("GetWalletKeys wallet[%d] = %v", i, w)
	}
}

func TestGetWalletBalance(t *testing.T) {

	tests := []struct {
		name string
		tag  string
	}{
		{
			name: "W5aQbddWii7ozwWGCzEiej8oagVNqChEoM",
			tag:  "first",
		},

	}

	for i, test := range tests {
		balance := tw.GetWalletBalance(test.name)
		t.Logf("GetWalletBalance[%d] %s balance = %s \n", i, test.name, balance)
	}

}

func TestCreateNewPrivateKey(t *testing.T) {

/*	test := struct {
		name     string
		password string
		tag      string
	}{
		name:     "WDHupMjR3cR2wm97iDtKajxSPCYEEddoek",
		password: "1234qwer",
	}

	count := 100

	w, err := tw.GetWalletInfo(test.name)
	if err != nil {
		t.Errorf("CreateNewPrivateKey failed unexpected error: %v\n", err)
		return
	}

	key, err := w.HDKey(test.password)
	if err != nil {
		t.Errorf("CreateNewPrivateKey failed unexpected error: %v\n", err)
		return
	}

	timestamp := 1
	t.Logf("CreateNewPrivateKey timestamp = %v \n", timestamp)

	derivedPath := fmt.Sprintf("%s/%d", key.RootPath, timestamp)
	childKey, _ := key.DerivedKeyWithPath(derivedPath, tw.Config.CurveType)

	for i := 0; i < count; i++ {

		wif, a, err := tw.CreateNewPrivateKey(key.KeyID, childKey, derivedPath, uint64(i))
		if err != nil {
			t.Errorf("CreateNewPrivateKey[%d] failed unexpected error: %v\n", i, err)
			continue
		}

		t.Logf("CreateNewPrivateKey[%d] wif = %v \n", i, wif)
		t.Logf("CreateNewPrivateKey[%d] address = %v \n", i, a.Address)
	}*/
}

func TestGetWalleInfo(t *testing.T) {
	w, err := tw.GetWalletInfo("WDHupMjR3cR2wm97iDtKajxSPCYEEddoek")
	if err != nil {
		t.Errorf("GetWalletInfo failed unexpected error: %v\n", err)
		return
	}

	t.Logf("GetWalletInfo wallet = %v \n", w)
}

//func TestCreateBatchPrivateKey(t *testing.T) {
//
//	w, err := tw.GetWalletInfo("Zhiquan Test")
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	key, err := w.HDKey("1234qwer")
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	wifs, err := tw.CreateBatchPrivateKey(key, 10000)
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	for i, wif := range wifs {
//		t.Logf("CreateBatchPrivateKey[%d] wif = %v \n", i, wif)
//	}
//
//}

//func TestImportMulti(t *testing.T) {
//
//	addresses := []string{
//		"1CoRcQGjPEyWmB1ZyG6CEDN3SaMsaD3ERa",
//		"1ESGCsXkNr3h5wvWScdCpVHu2GP3KJtCdV",
//	}
//
//	keys := []string{
//		"L5k8VYSvuZxC5FCczGVC8MmnKKix3Mcs6t185eUJVKTzZb1f6bsX",
//		"L3RVDjPVBSc7DD4WtmzbHkAHJW4kDbyXbw4vBppZ4DRtPt5u8Naf",
//	}
//
//	UnlockWallet("1234qwer", 120)
//	failed, err := ImportMulti(addresses, keys, "Zhiquan Test")
//	if err != nil {
//		t.Errorf("ImportMulti failed unexpected error: %v\n", err)
//	} else {
//		t.Errorf("ImportMulti result: %v\n", failed)
//	}
//}

func TestBackupWallet(t *testing.T) {

	backupFile, err := tw.BackupWallet("W5aQbddWii7ozwWGCzEiej8oagVNqChEoM")
	if err != nil {
		t.Errorf("BackupWallet failed unexpected error: %v\n", err)
	} else {
		t.Errorf("BackupWallet filePath: %v\n", backupFile)
	}
}

func TestBackupWalletData(t *testing.T) {
	tw.Config.WalletDataPath = "/home/www/btc/testdata/testnet3/"
	tmpWalletDat := fmt.Sprintf("tmp-walllet-%d.dat", time.Now().Unix())
	backupFile := filepath.Join(tw.Config.WalletDataPath, tmpWalletDat)
	err := tw.BackupWalletData(backupFile)
	if err != nil {
		t.Errorf("BackupWallet failed unexpected error: %v\n", err)
	} else {
		t.Errorf("BackupWallet filePath: %v\n", backupFile)
	}
}

func TestDumpWallet(t *testing.T) {
	//tw.UnlockWallet("1234qwer", 120)
	file := filepath.Join(".", "dump.txt")
	err := tw.DumpPrivate("address")
	if err != nil {
		t.Errorf("DumpWallet failed unexpected error: %v\n", err)
	} else {
		t.Errorf("DumpWallet filePath: %v\n", file)
	}
}

func TestGOSH(t *testing.T) {
	//text, err := sh.Command("go", "env").Output()
	//text, err := sh.Command("wmd", "version").Output()
	text, err := sh.Command("wmd", "Config", "see", "-s", "btm").Output()
	if err != nil {
		t.Errorf("GOSH failed unexpected error: %v\n", err)
	} else {
		t.Errorf("GOSH output: %v\n", string(text))
	}
}

func TestGetBlockChainInfo(t *testing.T) {
	b, err := tw.GetBlockChainInfo()
	if err != nil {
		t.Errorf("GetBlockChainInfo failed unexpected error: %v\n", err)
	} else {
		log.Infof("GetBlockChainInfo info: %v\n", b)
	}
}

func TestListUnspent(t *testing.T) {
	//msHemmfSZ3au6h9S1annGcTGrTVryRbSFV
	//mtHT3JkeKnJZCejqp6nxScxxvbW6Wn8e92
	utxos, err := tw.ListUnspent(0, "TsREAnJ2Ht3BMcWwXHDG8vKKC9NF9kbb94c")
	if err != nil {
		t.Errorf("ListUnspent failed unexpected error: %v\n", err)
		return
	}

	for _, u := range utxos {
		t.Logf("ListUnspent %s: %s = %s\n", u.Address, u.AccountID, u.Amount)
	}
}

//func TestGetAddressesFromLocalDB(t *testing.T) {
//	addresses, err := tw.GetAddressesFromLocalDB("WDHupMjR3cR2wm97iDtKajxSPCYEEddoek", 0, -1)
//	if err != nil {
//		t.Errorf("GetAddressesFromLocalDB failed unexpected error: %v\n", err)
//		return
//	}
//
//	for i, a := range addresses {
//		t.Logf("GetAddressesFromLocalDB address[%d] = %v\n", i, a)
//	}
//}

func TestRebuildWalletUnspent(t *testing.T) {

	err := tw.RebuildWalletUnspent("W5aQbddWii7ozwWGCzEiej8oagVNqChEoM")
	if err != nil {
		t.Errorf("RebuildWalletUnspent failed unexpected error: %v\n", err)
		return
	}

	t.Logf("RebuildWalletUnspent successfully.\n")
}

func TestListUnspentFromLocalDB(t *testing.T) {
	utxos, err := tw.ListUnspentFromLocalDB("W5aQbddWii7ozwWGCzEiej8oagVNqChEoM")
	if err != nil {
		t.Errorf("ListUnspentFromLocalDB failed unexpected error: %v\n", err)
		return
	}
	t.Logf("ListUnspentFromLocalDB totalCount = %d\n", len(utxos))
	total := decimal.New(0, 0)
	for _, u := range utxos {
		amount, _ := decimal.NewFromString(u.Amount)
		total = total.Add(amount)
		t.Logf("ListUnspentFromLocalDB %v: %s = %s\n", u.HDAddress, u.AccountID, u.Amount)
	}
	t.Logf("ListUnspentFromLocalDB total = %s\n", total.String())
}

func TestBuildTransaction(t *testing.T) {
	walletID := "W5aQbddWii7ozwWGCzEiej8oagVNqChEoM"
	utxos, err := tw.ListUnspentFromLocalDB(walletID)
	if err != nil {
		t.Errorf("BuildTransaction failed unexpected error: %v\n", err)
		return
	}

	txRaw, _, err := tw.BuildTransaction(utxos, []string{"mrThNMQ6bMf1YNPjBj9jYXmYYzw1Rt8GFU"}, "n33cHpEc9qAvECM9pFgabZ6ktJimLSeWdy", []decimal.Decimal{decimal.NewFromFloat(0.2)}, decimal.NewFromFloat(0.00002))
	if err != nil {
		t.Errorf("BuildTransaction failed unexpected error: %v\n", err)
		return
	}

	t.Logf("BuildTransaction txRaw = %s\n", txRaw)

	//hex, err := SignRawTransaction(txRaw, walletID, "1234qwer", utxos)
	//if err != nil {
	//	t.Errorf("BuildTransaction failed unexpected error: %v\n", err)
	//	return
	//}
	//
	//t.Logf("BuildTransaction signHex = %s\n", hex)
}

func TestEstimateFee(t *testing.T) {
	feeRate, _ := tw.EstimateFeeRate()
	t.Logf("EstimateFee feeRate = %s\n", feeRate.String())
	fees, _ := tw.EstimateFee(10, 2, feeRate)
	t.Logf("EstimateFee fees = %s\n", fees.String())
}

func TestSendTransaction(t *testing.T) {

	sends := []string{
		"mifDUXy3bzKVbKDHb4FiQTCXwDjWCixKWm",
	}

	tw.RebuildWalletUnspent("Tsa6c6ivPG3fjRnJwKjGPUyMauPYDvK6Ckk")

	for _, to := range sends {

		txIDs, err := tw.SendTransaction("W5aQbddWii7ozwWGCzEiej8oagVNqChEoM", to, decimal.NewFromFloat(1), "111111", false)

		if err != nil {
			t.Errorf("SendTransaction failed unexpected error: %v\n", err)
			return
		}

		t.Logf("SendTransaction txid = %v\n", txIDs)

	}

}

func TestSendBatchTransaction(t *testing.T) {

	sends := []string{
		"mqwis1h9GqmMkMjmkQEeYbz68RTC1QPvb9",
	}

	amounts := []decimal.Decimal{
		decimal.NewFromFloat(0.07),
		decimal.NewFromFloat(0.03),
		decimal.NewFromFloat(0.08),
		decimal.NewFromFloat(0.02),
	}

	tw.RebuildWalletUnspent("W5aQbddWii7ozwWGCzEiej8oagVNqChEoM")

	txID, err := tw.SendBatchTransaction("W5aQbddWii7ozwWGCzEiej8oagVNqChEoM", sends, amounts, "111111")

	if err != nil {
		t.Errorf("TestSendBatchTransaction failed unexpected error: %v\n", err)
		return
	}

	t.Logf("SendTransaction txid = %v\n", txID)

}

func TestMath(t *testing.T) {
	piece := int64(math.Ceil(float64(67) / float64(30)))

	t.Logf("ceil = %d", piece)
}

func TestGetNetworkInfo(t *testing.T) {
	tw.GetNetworkInfo()
}

func TestPrintConfig(t *testing.T) {
	tw.Config.PrintConfig()
}

func TestRestoreWallet(t *testing.T) {
	keyFile := "D:/gopath/src/github.com/blocktree/hc-adapter/openwtester/openw_data/key/HELLO HC-W5aQbddWii7ozwWGCzEiej8oagVNqChEoM.key"
	dbFile := "D:/gopath/src/github.com/blocktree/hc-adapter/openwtester/openw_data/db/hc-adapter.db"
	datFile := "D:/gopath/bin/openwallet.exe"
	tw.LoadConfig()
	err := tw.RestoreWallet(keyFile, dbFile, datFile, "111111")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
	}

}

func TestWalletManager_ImportAddress(t *testing.T) {
	addr := "TsREAnJ2Ht3BMcWwXHDG8vKKC9NF9kbb94c"
	err := tw.ImportAddress(addr, "")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
		return
	}
	log.Info("imported success")
}
