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
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/openwallet"
)

//var (
//	AddressDecoder = &openwallet.AddressDecoder{
//		PrivateKeyToWIF:    PrivateKeyToWIF,
//		PublicKeyToAddress: PublicKeyToAddress,
//		WIFToPrivateKey:    WIFToPrivateKey,
//		RedeemScriptToAddress: RedeemScriptToAddress,
//	}
//)

type AddressDecoder interface {
	openwallet.AddressDecoder

	//
	ScriptPubKeyToBech32Address(scriptPubKey []byte) (string, error)
}

type addressDecoder struct {
	wm *WalletManager //钱包管理者
}

//NewAddressDecoder 地址解析器
func NewAddressDecoder(wm *WalletManager) *addressDecoder {
	decoder := addressDecoder{}
	decoder.wm = wm
	return &decoder
}

//priv := "ab9715425cb447f24dcd64df7099fea0dd41986fa917c569a255858dcbf26147"
//buf,_ := hex.DecodeString(priv)
//privKey, _ := chainec.Secp256k1.PrivKeyFromBytes(buf)
////algType := chainec.ECTypeSecp256k1
////
//////	netIDTest := [2]byte{25,171}
//wifTest, _ := NewWIF(privKey, &chaincfg.MainNetParams, algType)
//
//fmt.Println(wifTest.String())

//PrivateKeyToWIF 私钥转WIF  HC not support only once blake256
func (decoder *addressDecoder) PrivateKeyToWIF(priv []byte, isTestnet bool) (string, error) {

	//cfg := addressEncoder.HC_mainnetPrivateWIFCompressed
	//if decoder.wm.Config.IsTestNet {
	//	cfg = addressEncoder.HC_mainnetPrivateWIFCompressed
	//}
	//
	//
	//wif := addressEncoder.AddressEncode(priv, cfg)

	return "", nil

}

//PublicKeyToAddress 公钥转地址
func (decoder *addressDecoder) PublicKeyToAddress(pub []byte, isTestnet bool) (string, error) {

	cfg := addressEncoder.HC_mainnetAddressP2PKH
	if decoder.wm.Config.IsTestNet {
		cfg = addressEncoder.HC_testnetAddressP2PKH
	}

	//pkHash := btcutil.Hash160(pub)
	//address, err :=  btcutil.NewAddressPubKeyHash(pkHash, &cfg)
	//if err != nil {
	//	return "", err
	//}

	pkHash := owcrypt.Hash(pub, 0, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)

	//if decoder.wm.Config.RPCServerType == RPCServerCore {
	//	//如果使用core钱包作为全节点，需要导入地址到core，这样才能查询地址余额和utxo
	//	err := decoder.wm.ImportAddress(address, "")
	//	if err != nil {
	//		return "", err
	//	}
	//}

	return address, nil

}

//RedeemScriptToAddress 多重签名赎回脚本转地址
func (decoder *addressDecoder) RedeemScriptToAddress(pubs [][]byte, required uint64, isTestnet bool) (string, error) {

	cfg := addressEncoder.HC_mainnetAddressP2SH
	if decoder.wm.Config.IsTestNet {
		cfg = addressEncoder.HC_testnetAddressP2SH
	}

	redeemScript := make([]byte, 0)

	for _, pub := range pubs {
		redeemScript = append(redeemScript, pub...)
	}

	pkHash := owcrypt.Hash(redeemScript, 0, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)

	return address, nil

}

//WIFToPrivateKey WIF转私钥
func (decoder *addressDecoder) WIFToPrivateKey(wif string, isTestnet bool) ([]byte, error) {

	//cfg := addressEncoder.HC_mainnetPrivateWIFCompressed
	//if decoder.wm.Config.IsTestNet {
	//	cfg = addressEncoder.HC_testnetPrivateWIFCompressed
	//}
	//
	//priv, err := addressEncoder.AddressDecode(wif, cfg)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil

}

//ScriptPubKeyToBech32Address scriptPubKey转Bech32地址
func (decoder *addressDecoder) ScriptPubKeyToBech32Address(scriptPubKey []byte) (string, error) {
	return "", nil //scriptPubKeyToBech32Address(scriptPubKey, decoder.wm.Config.IsTestNet)

}

//ScriptPubKeyToBech32Address scriptPubKey转Bech32地址
func scriptPubKeyToBech32Address(scriptPubKey []byte, isTestNet bool) (string, error) {
	//var (
	//	hash []byte
	//)
	//
	//cfg := addressEncoder.BTC_mainnetAddressBech32V0
	//if isTestNet {
	//	cfg = addressEncoder.BTC_testnetAddressBech32V0
	//}
	//
	//if len(scriptPubKey) == 22 || len(scriptPubKey) == 34 {
	//
	//	hash = scriptPubKey[2:]
	//
	//	address := addressEncoder.AddressEncode(hash, cfg)
	//
	//	return address, nil
	//
	//} else {
	//	return "", fmt.Errorf("scriptPubKey length is invalid")
	//}

	return "", nil

}
