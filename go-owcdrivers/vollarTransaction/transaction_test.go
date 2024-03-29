package vollarTransaction

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_transaction(t *testing.T) {
	// 案例复现了链上交易
	// id为： d217c2d2047b8c7004c6a99fbd3fb97a80bc8eaf3b2f219c02ac866632c1944b
	// 为了使结果与链上数据一致，测试过程中修改了签名结果和公钥

	// utxo
	txid := "3ba0c1c5e8ce46391287942ef4dadd3edfddadbf5a61d3fb2cfd8638a2cbdeb4"
	vout := uint32(1)
	lockscript := "76a91481772013db2bf0d22428053c68bb24868c381b9988ac"

	in := Vin{txid, vout, lockscript}

	// 交易输出
	out1 := Vout{"VcezzxyRG8y5vXhaUUibrqygm4QSPbc6Z7F", uint64(690000000)}
	out2 := Vout{"VcimwAzkqBDdFhzTunS2hQFrzBuXGmBHYYB", uint64(44990000)}

	// 同时获取空交易单和待签哈希
	emptyTrans, hashes, err := CreateEmptyRawTransactionAndHash([]Vin{in}, []Vout{out1, out2})

	if err != nil {
		t.Error("failed to create transaction!")
	} else {
		fmt.Println("空交易单: \n", emptyTrans)
		for i, hash := range hashes {
			fmt.Println("第 ", i, "个hash为:\n", hash)
		}
	}

	prikey := []byte{0x84, 0x98, 0x23, 0xd2, 0x2d, 0x81, 0xe4, 0x9e, 0xb7, 0x19, 0x06, 0x6b, 0xcf, 0x7e, 0xd1, 0x73, 0xe6, 0x09, 0x48, 0x22, 0xb0, 0xea, 0x4e, 0x79, 0x3f, 0x1d, 0x85, 0x97, 0xa5, 0x06, 0x0d, 0x27}

	signature, err := SignRawTransaction(hashes[0], prikey)
	if err != nil {
		t.Error("failed to sign transaction!")
	} else {
		// only for test
		// 此处修改为与该笔交易对应的签名
		signature = []byte{0x3b, 0xeb, 0x79, 0x1e, 0x2d, 0x79, 0x4e, 0xcc, 0xad, 0xd1, 0x9b, 0x23, 0x62, 0xa5, 0xa1, 0x71, 0x75, 0xf1, 0x4e, 0x54, 0x7d, 0x8d, 0xd8, 0xee, 0xd2, 0xbc, 0xed, 0xa9, 0x0a, 0x3c, 0xe2, 0x00, 0x38, 0x6b, 0xe2, 0x6d, 0x86, 0xfc, 0x3a, 0x41, 0x2e, 0x08, 0xbc, 0xa9, 0x26, 0x71, 0x8f, 0x84, 0x28, 0xba, 0xba, 0x6f, 0xbd, 0x57, 0xeb, 0x00, 0xf7, 0xaa, 0x82, 0xcd, 0xc5, 0xfc, 0xe9, 0x60}
		fmt.Println("签名值为: \n", hex.EncodeToString(signature))
	}

	pubkey := []byte{0x03, 0x57, 0xcb, 0x81, 0xef, 0xed, 0x4e, 0x5c, 0x86, 0x9d, 0xb1, 0x51, 0xaf, 0x95, 0x6b, 0xbf, 0x89, 0x56, 0xa8, 0x17, 0x77, 0x15, 0xb1, 0x3a, 0xdb, 0x6c, 0x93, 0x16, 0x94, 0xaa, 0x8e, 0x7a, 0x37}

	sigPub := SigPub{
		Pubkey:    pubkey,
		Signature: signature,
	}

	// 同时获取验签结果和最终合并交易单
	pass, signedTrans, err := VerifyAndCombineRawTransaction(emptyTrans, []SigPub{sigPub})

	if err != nil {
		t.Error("verify failed!")
	} else {
		fmt.Println(pass)
		fmt.Println(signedTrans)
	}
}
