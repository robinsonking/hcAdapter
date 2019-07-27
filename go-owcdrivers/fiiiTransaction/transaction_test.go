package fiiiTransaction

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// 交易费计算
// ((262*inputCount + 101*outputCount + 68)/1024.0) * FeePerKB

func Test_send(t *testing.T) {
	//本案例旨在复现链上交易 CE427FC1CF51EF1C72BAF0C97D7547A47D8E564EA662F36AA57D09F958164CED
	//如果想要结果完全一致，需要将 transaction.go 文件中的第36行替换为如下：
	//txMsg.Timestamp = int64(1546443912218)
	// 使用prefix区分主网和测试网
	prefix := []byte{0x40, 0xe7, 0xe9, 0x15}

	txid := "5E36D1C2879780E8ABBE4B451DE9202F66D1F037BF3D4372BA463522459B4BD4"
	vout := 1

	to1 := "fiiimH1KXvvVxfuNAH97u3hzfFEmojvTdkJZXg"
	amount1 := int64(624000000000)
	to2 := "fiiimRS3TUtbAyTX3SYLxZ7rq5wB2bcT5Y2KFb"
	amount2 := int64(510934733923614)

	//构建输入
	inputs := []Vin{Vin{txid, vout}}

	//构建输出
	outputs := []Vout{Vout{prefix, to1, amount1}, Vout{prefix, to2, amount2}}

	//其他参数
	version := int(1)
	locktime := int64(0)
	expiredtime := int64(0)

	//创建交易单与待签消息
	emptyTrans, messages, err := CreateEmptyTransactionAndMessage(inputs, outputs, version, locktime, expiredtime)

	if err != nil {
		t.Error("Failed to create transaction! with error: \n", err)
	} else {
		fmt.Println("空交易单: \n", emptyTrans, "\n ")
		fmt.Println("待签消息: \n", messages[0], "\n ")
	}

	// 客户端签名
	prikey := []byte{0x84, 0x98, 0x23, 0xd2, 0x2d, 0x81, 0xe4, 0x9e, 0xb7, 0x19, 0x06, 0x6b, 0xcf, 0x7e, 0xd1, 0x73, 0xe6, 0x09, 0x48, 0x22, 0xb0, 0xea, 0x4e, 0x79, 0x3f, 0x1d, 0x85, 0x97, 0xa5, 0x06, 0x0d, 0x27}

	signature, err := SignTransactionMessage(messages[0], prikey)
	if err != nil {
		t.Error("Failed to sign! with error : \n", err)
	} else {
		// only for test，构建与链上一致数据
		signature = []byte{0xCF, 0x1B, 0x46, 0x15, 0x7E, 0xB1, 0xBA, 0x6E, 0x39, 0xCD, 0xA6, 0xF4, 0xE6, 0x42, 0x17, 0xCC, 0x34, 0xF6, 0xE0, 0xAE, 0xFE, 0x91, 0x07, 0xAA, 0xFF, 0x98, 0x47, 0x4B, 0xA6, 0x71, 0xBC, 0xB4, 0x8C, 0x7A, 0x7D, 0xE6, 0x74, 0x60, 0x9C, 0xA6, 0xC9, 0xEF, 0x1C, 0x14, 0x2A, 0x3F, 0x66, 0x60, 0x18, 0xFE, 0xFA, 0xB3, 0xEB, 0x78, 0xB2, 0x1C, 0x2E, 0xC0, 0x3B, 0xAB, 0xD9, 0x31, 0xB5, 0x0D}
		fmt.Println("签名结果: \n", hex.EncodeToString(signature))
	}

	// 验证与交易单合并
	// 获取对应公钥
	pubkey := []byte{0x22, 0x35, 0x88, 0xFB, 0xFA, 0xF1, 0x04, 0x08, 0xF0, 0x42, 0x46, 0x90, 0x31, 0xDD, 0x6D, 0x06, 0x28, 0xCC, 0x9E, 0xDE, 0x9F, 0x31, 0x2A, 0xD1, 0x28, 0xA4, 0x31, 0xF4, 0x20, 0x49, 0x22, 0x42}

	// 填充结构体
	sigPubs := []SigPub{SigPub{signature, pubkey}}

	pass, signedTrans, err := VerifyAndCombineTransaction(emptyTrans, sigPubs)
	if err != nil {
		t.Error("Failed to combine!")
	} else {
		if pass != true {
			t.Error("Verify failed!")
		} else {
			fmt.Println("待发送交易单： \n", signedTrans)
		}
	}
}
