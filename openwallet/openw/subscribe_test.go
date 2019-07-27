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
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
	"path/filepath"
	"testing"
	"time"
)

type subscriber struct {
}

func init() {
	//tm.Init()
}

//BlockScanNotify 新区块扫描完成通知
func (sub *subscriber) BlockScanNotify(header *openwallet.BlockHeader) error {
	log.Debug("header:", header)
	return nil
}

//BlockTxExtractDataNotify 区块提取结果通知
func (sub *subscriber) BlockTxExtractDataNotify(account *openwallet.AssetsAccount, data *openwallet.TxExtractData) error {
	log.Debug("account:", account)
	log.Debug("data:", data)

	return nil
}

func TestSubscribe(t *testing.T) {
	tm := testInitWalletManager()
	var (
		endRunning = make(chan bool, 1)
	)

	sub := subscriber{}
	tm.AddObserver(&sub)
	//tm.SetRescanBlockHeight("QTUM", 236098)
	log.Debug("SupportAssets:", tm.cfg.SupportAssets)
	<-endRunning
}

////////////////////////// 测试单个扫描器 //////////////////////////

type subscriberSingle struct {
}

//BlockScanNotify 新区块扫描完成通知
func (sub *subscriberSingle) BlockScanNotify(header *openwallet.BlockHeader) error {
	log.Notice("header:", header)
	return nil
}

//BlockTxExtractDataNotify 区块提取结果通知
func (sub *subscriberSingle) BlockExtractDataNotify(sourceKey string, data *openwallet.TxExtractData) error {
	log.Notice("account:", sourceKey)

	for i, input := range data.TxInputs {
		log.Std.Notice("data.TxInputs[%d]: %+v", i, input)
	}

	for i, output := range data.TxOutputs {
		log.Std.Notice("data.TxOutputs[%d]: %+v", i, output)
	}

	log.Std.Notice("data.Transaction: %+v", data.Transaction)

	return nil
}

func TestSubscribeAddress_BTC(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "BTC"
		accountID  = "W4VUMN3wxQcwVEwsRvoyuhrJ95zhyc4zRW"
		addrs      = map[string]string{
			"mh84vzGGA6T5s39BovBJe2dQ3r9KkWQPML":         accountID,
			"tb1qk7ya4ntkmnc6mrzj7a2ly60524hjue440egv5k": accountID,
			"mpS3FzHr1WsXWnVw15uiFzLQ9wRmo7rBJZ":         accountID,
			"mgbZ61Hpbgd55rBk3vvk3r2um1htTxdrrr":         accountID, //usdt
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	scanner.SetRescanBlockHeight(1365800)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	//time.Sleep(12 * time.Second)
	//
	//log.Notice("scanner.Pause()")
	//scanner.Pause()
	//
	//time.Sleep(12 * time.Second)
	//
	//log.Notice("scanner.Restart()")
	//scanner.Restart()
	//
	//time.Sleep(12 * time.Second)
	//
	//log.Notice("scanner.Stop()")
	//scanner.Stop()
	//
	//time.Sleep(12 * time.Second)
	//
	//log.Notice("scanner.Run()")
	//scanner.Run()

	<-endRunning
}

func TestSubscribeAddress_ETH(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "ETH"
		accountID  = "W4VUMN3wxQcwVEwsRvoyuhrJ95zhyc4zRW"
		addrs      = map[string]string{
			"0x558ef7a2b56611ef352b2ecf5d2dd2bf548afecc": accountID,
			"0x95576498d5c2971bea1986ee92a3971de0747fc0": accountID,
			"0x73995d52f20d9d40cbc339d5d2772d9cde6b6858": accountID,
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	assetsLogger := assetsMgr.GetAssetsLogger()
	if assetsLogger != nil {
		assetsLogger.SetLogFuncCall(true)
	}

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	scanner.SetRescanBlockHeight(6518561)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	//time.Sleep(10 * time.Second)
	//log.Notice("scanner.Pause()")
	//scanner.Pause()
	//
	//time.Sleep(30 * time.Second)
	//log.Notice("scanner.Restart()")
	//scanner.Restart()

	<-endRunning
}

func TestSubscribeAddress_QTUM(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "QTUM"
		accountID  = "W4VUMN3wxQcwVEwsRvoyuhrJ95zhyc4zRW"
		addrs      = map[string]string{
			"Qf6t5Ww14ZWVbG3kpXKoTt4gXeKNVxM9QJ": accountID, //合约收币
			//"QWSTGRwdScLfdr6agUqR4G7ow4Mjc4e5re",	//合约发币
			//"QbTQBADMqSuHM6wJk2e8w1KckqK5RRYrQ6",	//主链转账
			//"QREUcesH46vMeF6frLy92aR1QC22tADNda", 	//主链转账
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	scanner.SetRescanBlockHeight(330000)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	time.Sleep(6 * time.Second)

	log.Notice("scanner.Pause()")
	scanner.Pause()

	time.Sleep(6 * time.Second)

	log.Notice("scanner.Restart()")
	scanner.Restart()

	time.Sleep(6 * time.Second)

	log.Notice("scanner.Stop()")
	scanner.Stop()

	time.Sleep(6 * time.Second)

	log.Notice("scanner.Run()")
	scanner.Run()

	<-endRunning
}

func TestSubscribeAddress_LTC(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "LTC"
		//accountID  = "W4VUMN3wxQcwVEwsRvoyuhrJ95zhyc4zRW"
		addrs      = map[string]string{
			"mgCzMJDyJoqa6XE3RSdNGvD5Bi5VTWudRq": "A3Mxhqm65kTgS2ybHLenNrZzZNtLGVobDFYdpc1ge4eK", //主链转账
			"mkdStRouBPVrDVpYmbE5VUJqhBgxJb3dSS": "3i26MQmtuWVVnw8GnRCVopG3pi8MaYU6RqWVV2E1hwJx", //主链转账
			//"mzZYMMuTNEHTQ1zqpSGkgpZvv8g3ZYWoGX": accountID,
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	scanner.SetRescanBlockHeight(973484)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning
}

func TestSubscribeAddress_NAS(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "NAS"
		accountID  = "7VftKuNoDtwZ3mn3wDA4smTDMz4iqCg3fNna1fXicVDg"
		addrs      = map[string]string{
			"n1FYB93yKATg42MWrWDuNGYzKiVcRwH4NMX": accountID,
			"n1GEpN9ZdaRzjBkVWbTagg4efWkmTtGeuYB": accountID,
			"n1HU5Zxfch2K6J1MvskGDADajzb2NYZ1tZb": accountID,
			"n1JjXv57Cztbfsyf6UwcwPxWNQm2ejd7vFu": accountID,
			"n1U8RmQDZvidVeeSTeyRX4CgqYpADfUdddb": accountID,
			"n1UKFCBCLrQZQ4eENPwRoSt74uJaP51cKN3": accountID,
			"n1a4f5F1pzjFhg78rt2LpER8vfUYErWo9yo": accountID,
			"n1cT1JhXUQFxDSyBYwcn3ZBmd93nin5yrfB": accountID,
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	//scanner.SetRescanBlockHeight(1972169)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning
}

func TestSubscribeAddress_BCH(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "BCH"
		accountID  = "W4VUMN3wxQcwVEwsRvoyuhrJ95zhyc4zRW"
		addrs      = map[string]string{
			"mh84vzGGA6T5s39BovBJe2dQ3r9KkWQPML":         accountID,
			"tb1qk7ya4ntkmnc6mrzj7a2ly60524hjue440egv5k": accountID,
			"mpS3FzHr1WsXWnVw15uiFzLQ9wRmo7rBJZ":         accountID,
			"mgbZ61Hpbgd55rBk3vvk3r2um1htTxdrrr":         accountID, //usdt
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	scanner.SetRescanBlockHeight(558337)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning
}

func TestSubscribeAddress_TRON(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "TRX"
		//accountID  = "CfRjWjct569qp7oygSA2LrsAoTrfEB8wRk3sHGUj9Erm"
		accountID = "6msrcfed9rA7njVNDtY1Ppo9XQdX5p3SFPc1zxWgd8ut"
		addrs     = map[string]string{
			"TT44ohw23WGNv1jQCAUN3etUWND1KXN2Eq": accountID,
			"TJLypjev8iLdQR3X63rSMeZK8GKwkeSH1Y": accountID,
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	scanner.SetRescanBlockHeight(7848677)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning

}

func TestSubscribeAddress_ONT(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "ONT"
		accountID  = "W4VUMN3wxQcwVEwsRvoyuhrJ95zhyc4zRW"
		addrs      = map[string]string{
			"AHt5qGrP9HG8zqXAEH5gqUwv2FY9pWqRyE": accountID,
			"AQ9FxqKbXnZTqf22zy2NFcXhi13fipg48S": accountID,
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	scanner.SetRescanBlockHeight(1184421)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning
}


func TestSubscribeAddress_VSYS(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "VSYS"
		accountID  = "FUAKFujfVwdWJn79DFB4ZZQ6LRZS5cXfrGC9er2T5TSt"
		addrs      = map[string]string{
			"ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11": accountID,
			"AREkgFxYhyCdtKD9JSSVhuGQomgGcacvQqM": accountID,
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	assetsLogger := assetsMgr.GetAssetsLogger()
	if assetsLogger != nil {
		assetsLogger.SetLogFuncCall(true)
	}

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	//scanner.SetRescanBlockHeight(1993837)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning
}


func TestSubscribeAddress_TRUE(t *testing.T) {

	var (
		endRunning = make(chan bool, 1)
		symbol     = "TRUE"
		accountID  = "W4VUMN3wxQcwVEwsRvoyuhrJ95zhyc4zRW"
		addrs      = map[string]string{
			"0x558ef7a2b56611ef352b2ecf5d2dd2bf548afecc": accountID,
			"0x95576498d5c2971bea1986ee92a3971de0747fc0": accountID,
			"0x73995d52f20d9d40cbc339d5d2772d9cde6b6858": accountID,
		}
	)

	//GetSourceKeyByAddress 获取地址对应的数据源标识
	scanAddressFunc := func(address string) (string, bool) {
		key, ok := addrs[address]
		if !ok {
			return "", false
		}
		return key, true
	}

	assetsMgr, err := GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}

	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)

	assetsLogger := assetsMgr.GetAssetsLogger()
	if assetsLogger != nil {
		assetsLogger.SetLogFuncCall(true)
	}

	//log.Debug("already got scanner:", assetsMgr)
	scanner := assetsMgr.GetBlockScanner()
	//scanner.SetRescanBlockHeight(6518561)

	if scanner == nil {
		log.Error(symbol, "is not support block scan")
		return
	}

	scanner.SetBlockScanAddressFunc(scanAddressFunc)

	sub := subscriberSingle{}
	scanner.AddObserver(&sub)

	scanner.Run()

	<-endRunning
}