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
	"encoding/hex"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
	"testing"
)

func TestAddressDecoder_Encode(t *testing.T) {

	//pubKeyBytes
	//hash, _ := hex.DecodeString("03278750407ffa19a3dfd1bb9248d53694810a192b5f63be6577979e844ff90f30")

	hash, _ := hex.DecodeString("6391dad744d45500476a0fa32a64a0048398cef9")

	cfg := addressEncoder.HC_testnetAddressP2PKH
	addr := addressEncoder.AddressEncode(hash, cfg)

	t.Logf("addr: %s", addr)
}

func TestAddressDecoder_Decode(t *testing.T) {

	addr := "Tsa6c6ivPG3fjRnJwKjGPUyMauPYDvK6Ckk"

	cfg := addressEncoder.HC_testnetAddressP2PKH

	hash, _ := addressEncoder.AddressDecode(addr, cfg)

	t.Logf("hash: %s", hex.EncodeToString(hash))
}

/*func TestAddressDecoder_PrivateKeyToWIF(t *testing.T) {
	priv := "ab9715425cb447f24dcd64df7099fea0dd41986fa917c569a255858dcbf26147"
	buf, _ := hex.DecodeString(priv)
	//privKey, _ := chainec.Secp256k1.PrivKeyFromBytes(buf)

	cfg := addressEncoder.HC_mainnetPrivateWIFCompressed

	wif := addressEncoder.AddressEncode(buf, cfg)
	t.Logf("wif: %s", wif)

}
*/
//func TestAddressDecoder_ScriptPubKeyToBech32Address(t *testing.T) {
//
//	scriptPubKey, _ := hex.DecodeString("002079db247b3da5d5e33e036005911b9341a8d136768a001e9f7b86c5211315e3e1")
//
//	addr, err := scriptPubKeyToBech32Address(scriptPubKey, true)
//	if err != nil {
//		t.Errorf("ScriptPubKeyToBech32Address failed unexpected error: %v\n", err)
//		return
//	}
//	t.Logf("addr: %s", addr)
//
//	t.Logf("addr: %s", addr)
//}

func TestAddressDecoder_WIFToP2WPKH_nested_in_P2SH(t *testing.T) {
	wif := "KwFE3SQqgADPAwkWc2A15Wh68rg7Xn2oAa9rwCF2pCb7KFKru4Mo"

	privkey, err := addressEncoder.AddressDecode(wif, addressEncoder.BTC_mainnetPrivateWIFCompressed)
	if err != nil {
		t.Errorf("AddressDecode failed unexpected error: %v\n", err)
		return
	}
	t.Logf("privkey: %s", hex.EncodeToString(privkey))

	pubkey, _ := owcrypt.GenPubkey(privkey, owcrypt.ECC_CURVE_SECP256K1)
	pubkey = owcrypt.PointCompress(pubkey, owcrypt.ECC_CURVE_SECP256K1)

	t.Logf("pubkey: %s", hex.EncodeToString(pubkey))

	hash := owcrypt.Hash(pubkey, 0, owcrypt.HASH_ALG_HASH160)

	//scriptSig = <0 <keyhash>>
	hash = append([]byte{0x00, 0x14}, hash...)
	hash = owcrypt.Hash(hash, 0, owcrypt.HASH_ALG_HASH160)

	t.Logf("hash: %s", hex.EncodeToString(hash))

	addr := addressEncoder.AddressEncode(hash, addressEncoder.BTC_mainnetAddressP2SH)

	t.Logf("addr: %s", addr)
}
