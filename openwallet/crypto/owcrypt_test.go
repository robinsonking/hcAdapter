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

package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/blocktree/go-owcrypt"
)

func slide_cmp_equ(a []byte, b []byte, length uint16) bool {
	i := uint16(0)
	for i = 0; i < length; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_sm2_keyagreement(*testing.T) {
	////////////////////////////////////////////////////////////////////发起方 - initiator////////////////////////////////////////////////////////////////////
	//发起方标识符
	IDinitiator := [4]byte{0x11, 0x22, 0x33, 0x44}
	//发起方私钥
	prikeyInitiator := [32]byte{0xF5, 0xA6, 0xF5, 0x9C, 0x50, 0xCD, 0x3D, 0x57, 0xCF, 0xC7, 0xC8, 0xB2, 0xA9, 0x9C, 0x98, 0x3C, 0x3B, 0xC0, 0x9A, 0xB6, 0x6E, 0x86, 0xAB, 0x64, 0x35, 0xC5, 0x18, 0x5B, 0x70, 0x15, 0xEA, 0x37}
	//发起方共钥
	pubkeyInitiator := [64]byte{0xDE, 0x6E, 0xB0, 0xD4, 0x70, 0x42, 0xF9, 0x51, 0x61, 0xC6, 0xC7, 0x18, 0x75, 0xEE, 0x62, 0x7D, 0xE1, 0xDE, 0x49, 0xE4, 0x23, 0x61, 0xF5, 0x3B, 0x2A, 0x15, 0x13, 0x2D, 0xA9, 0x8A, 0xEB, 0x2E, 0xFE, 0xA1, 0xDF, 0xC8, 0x2C, 0xCE, 0xF3, 0x0A, 0xB1, 0xB0, 0xAD, 0x8F, 0xB3, 0x05, 0x90, 0x55, 0x25, 0xF8, 0x4A, 0x8C, 0xEC, 0x81, 0x85, 0xB8, 0x0A, 0x51, 0x63, 0x8B, 0x5A, 0x10, 0x05, 0xB7}
	//发起方临时私钥 tmpPrikeyInitiator := [32]byte{}
	//发起方临时公钥tmpPubkeyInitiator := [64]byte{}
	//发起方向响应方发送的验证值 SA := [32]byte{}
	//发起方协商的结果retA := [32]byte{}

	////////////////////////////////////////////////////////////////////响应方 - responder////////////////////////////////////////////////////////////////////
	//响应方标识符
	IDresponder := [4]byte{0x55, 0x66, 0x77, 0x88}
	//相应方私钥
	prikeyResponder := [32]byte{0xB3, 0x93, 0xF6, 0xDB, 0xAB, 0x4E, 0xB4, 0x7C, 0x89, 0x03, 0xB3, 0x3A, 0xAA, 0x6E, 0x36, 0x70, 0xE1, 0xAE, 0x1A, 0xFD, 0xE3, 0x7F, 0x44, 0x1B, 0x7C, 0x78, 0xF1, 0x9E, 0x68, 0xEA, 0x7A, 0x90}
	//响应方公钥
	pubkeyResponder := [64]byte{0x01, 0x99, 0xB9, 0x57, 0x9B, 0x44, 0x83, 0x95, 0x62, 0x91, 0x12, 0x36, 0xA1, 0x44, 0x8E, 0x1B, 0xF2, 0xFF, 0x7B, 0xC0, 0xAE, 0xD9, 0x77, 0xFD, 0x88, 0x67, 0x1B, 0x16, 0x21, 0x13, 0x59, 0x73, 0x4D, 0x3F, 0x9A, 0xC4, 0xC1, 0x11, 0x2B, 0x4B, 0xE8, 0x8B, 0x30, 0x93, 0x84, 0x9F, 0xB8, 0x3E, 0x8D, 0xAB, 0xD0, 0xCE, 0x6F, 0xA4, 0x5F, 0x90, 0x41, 0xC5, 0x38, 0x16, 0xB2, 0x6B, 0x14, 0xBB}
	//响应方临时公钥 tmpPubkeyResponder := [64]byte{}
	//响应方本地验证S 2 := [32]byte{}
	//响应方向发起方发送的验证值 SB := [32]byte{}
	//响应方协商的结果 retB := [32]byte{}

	///////////////////////////////////////////////////////////////////////////协商开始////////////////////////////////////////////////////////////////////
	//协商开始前，发起方掌握的信息有：  自身的私钥、公钥、响应方的公钥，以及提前约定好的曲线参数
	//          响应方掌握的信息有：  自身的私钥、公钥、发送方的公钥，以及提前约定好的曲线参数
	//第一步：
	//1.1 发起方在本地产生一组临时的公私钥对，然后发起协商
	fmt.Println("--------------------------发起方第一步--------------------------")
	tmpPrikeyInitiator, tmpPubkeyInitiator := owcrypt.KeyAgreement_initiator_step1(owcrypt.ECC_CURVE_SM2_STANDARD)

	fmt.Println("发起方产生临时公私钥对，产生结果为：")
	fmt.Println("发起方临时私钥：", hex.EncodeToString(tmpPrikeyInitiator[:]))
	fmt.Println("发起方临时公钥：", hex.EncodeToString(tmpPubkeyInitiator[:]))

	//1.2 发起方将临时私钥保存在本地，用于第二步操作的输入
	//1.3 发起方将临时公钥发送给响应方来发起协商，同时会指定协商的具体长度

	//第二步：
	//2.1 响应方获得发送方的临时公钥和协商长度，然后开始进行协商计算
	fmt.Println("--------------------------响应方第一步--------------------------")
	retB, tmpPubkeyResponder, S2, SB, ret := owcrypt.KeyAgreement_responder_step1(IDinitiator[:],
		4,
		IDresponder[:],
		4,
		prikeyResponder[:],
		pubkeyResponder[:],
		pubkeyInitiator[:],
		tmpPubkeyInitiator[:],
		32,
		owcrypt.ECC_CURVE_SM2_STANDARD)
	if ret != owcrypt.SUCCESS {
		fmt.Println("响应方协商第一步出错！")
		return

	} else {
		fmt.Println("响应方产生临时公钥 ：", hex.EncodeToString(tmpPubkeyResponder[:]))
		fmt.Println("响应方本地校验值： ", hex.EncodeToString(S2[:]))
		fmt.Println("响应方发送给发起方的校验值： ", hex.EncodeToString(SB[:]))
		fmt.Println("响应方获得的协商结果： ", hex.EncodeToString(retB[:]))
	}

	//2.2 响应方此时获得临时公钥、用于本地校验的S2、用于发送给发起方的校验值SB， 协商结果
	//2.3 响应方将S2和协商保存在本地，用于第二步的校验
	//2.4 响应方将临时公钥和校验值SB发送给发起方

	//第三步：
	//发起方获得响应方的临时公钥和校验值，开始进行协商计算
	fmt.Println("--------------------------发起方第二步--------------------------")
	retA, SA, err := owcrypt.KeyAgreement_initiator_step2(IDinitiator[:],
		4,
		IDresponder[:],
		4,
		prikeyInitiator[:],
		pubkeyInitiator[:],
		pubkeyResponder[:],
		tmpPrikeyInitiator[:],
		tmpPubkeyInitiator[:],
		tmpPubkeyResponder[:],
		SB[:],
		32,
		owcrypt.ECC_CURVE_SM2_STANDARD)
	if err != owcrypt.SUCCESS {
		fmt.Println("发起方协商第一步出错！")
		return
	} else {
		fmt.Println("发起方发送给响应方的校验值： ", hex.EncodeToString(SA[:]))
		fmt.Println("发起方获得的协商结果： ", hex.EncodeToString(retA[:]))
	}

	//此时，发起方已经获得协商结果，如果接口返回SUCCESS，则说明接口内部已经与响应方发来的校验值完成校验
	//即：发起方的协商流程已经完成
	//然后，发起方需要将输出的校验值SA发送给响应方进行校验

	//第四步：
	//响应方拿到发起方发来的最终校验值SA， 与之前本地保存的校验值S2进行比对，返回SUCCESS则响应方协商通过
	fmt.Println("--------------------------响应方第二步--------------------------")
	if owcrypt.SUCCESS != owcrypt.KeyAgreement_responder_step2(SA[:], S2[:], owcrypt.ECC_CURVE_SM2_STANDARD) {
		fmt.Println("响应方校验未通过")
		return
	} else {
		fmt.Println("响应方校验通过")
	}

	if slide_cmp_equ(retA[:], retB[:], 32) {
		fmt.Println("双方协商结果一致")
	} else {
		fmt.Println("双方协商结果不一致")
	}

}

func Test_sm2_genpubkey(t *testing.T) {
	prikey_0 := [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	prikey_illegal := [32]byte{0xFF, 0xFF, 0xFF, 0xFf, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xff, 0xff, 0xff, 0xff, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFe}
	prikey := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}
	//pubkey_ecdsasecp256r1 := [64]byte{0xB2,0xC0,0xC3,0x07,0x28,0xAB,0x3C,0x10,0xCC,0x28,0x9B,0x78,0x0D,0x79,0xFB,0x92,0xF0,0x62,0xB7,0x2E,0x7D,0x3A,0x69,0xEC,0x0C,0x88,0xF1,0xE1,0x49,0x1D,0xDF,0xC1,0x24,0x8B,0x13,0xB2,0x6C,0xCD,0x60,0x2A,0xDA,0x23,0x29,0x69,0x65,0x70,0x94,0x01,0x33,0x53,0xFE,0xB7,0x94,0xDA,0x04,0xDD,0xF3,0xF7,0xF4,0x5C,0x26,0x70,0xFD,0x7B}
	pubkey_sm2 := [64]byte{0x6C, 0x1E, 0xFB, 0x83, 0xBD, 0xEC, 0xBC, 0x47, 0x99, 0xCE, 0x03, 0xAB, 0xE2, 0x64, 0x3C, 0x18, 0x4B, 0x75, 0x0A, 0x28, 0xA9, 0x33, 0x7A, 0x22, 0x1C, 0x58, 0x1B, 0x7B, 0x11, 0x61, 0xDA, 0x01, 0xB0, 0x5E, 0x98, 0x96, 0x58, 0x61, 0xC8, 0x78, 0x16, 0xFB, 0x5D, 0x57, 0xFA, 0xD6, 0xB6, 0x30, 0x6F, 0x98, 0x2A, 0x36, 0x97, 0xA0, 0x11, 0x80, 0x7A, 0x5C, 0x4C, 0xBB, 0xD4, 0xEF, 0x8A, 0xC2}
	//pubkey := [64]byte{}

	//私钥全0的情况
	pubkey, err := owcrypt.GenPubkey(prikey_0[:], owcrypt.ECC_CURVE_SM2_STANDARD)
	if err != owcrypt.ECC_PRIKEY_ILLEGAL {
		fmt.Println("sm2产生公钥接口未对全零私钥正确检查")
		return
	} else {
		fmt.Println("sm2产生公钥接口全零私钥检查通过")
	}
	pubkey, err = owcrypt.GenPubkey(prikey_illegal[:], owcrypt.ECC_CURVE_SM2_STANDARD)
	if err != owcrypt.ECC_PRIKEY_ILLEGAL {
		fmt.Println("sm2产生公钥接口未对非法私钥正确检查")
		return
	} else {
		fmt.Println("sm2产生公钥接口非法私钥检查通过")
	}
	pubkey, err = owcrypt.GenPubkey(prikey[:], owcrypt.ECC_CURVE_SM2_STANDARD)
	if err != owcrypt.SUCCESS {
		fmt.Println("sm2产生公钥接口返回值错误")
		return
	} else {
		fmt.Println("sm2产生公钥接口返回值正确")
	}
	if slide_cmp_equ(pubkey_sm2[:], pubkey[:], 64) != true {
		fmt.Println("sm2产生了错误的公钥")
		return
	} else {
		fmt.Println("sm2 genPubkey pass")
	}
}

func Test_sm2_encdec(t *testing.T) {
	prikey := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}
	pubkey := [64]byte{0x6C, 0x1E, 0xFB, 0x83, 0xBD, 0xEC, 0xBC, 0x47, 0x99, 0xCE, 0x03, 0xAB, 0xE2, 0x64, 0x3C, 0x18, 0x4B, 0x75, 0x0A, 0x28, 0xA9, 0x33, 0x7A, 0x22, 0x1C, 0x58, 0x1B, 0x7B, 0x11, 0x61, 0xDA, 0x01, 0xB0, 0x5E, 0x98, 0x96, 0x58, 0x61, 0xC8, 0x78, 0x16, 0xFB, 0x5D, 0x57, 0xFA, 0xD6, 0xB6, 0x30, 0x6F, 0x98, 0x2A, 0x36, 0x97, 0xA0, 0x11, 0x80, 0x7A, 0x5C, 0x4C, 0xBB, 0xD4, 0xEF, 0x8A, 0xC2}

	plain := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}

	i := uint16(0)

	for i = 1; i < 32; i++ {
		cipher, ret := owcrypt.Encryption(pubkey[:], plain[:i], owcrypt.ECC_CURVE_SM2_STANDARD)

		if ret != owcrypt.SUCCESS || uint16(len(cipher)) != i+97 {
			fmt.Println("sm2加密错误")
			return
		} else {
			fmt.Println("sm2 enc pass")
			fmt.Println("sm2 cipher value : ", hex.EncodeToString(cipher[:i+97]))
		}

		check, ret := owcrypt.Decryption(prikey[:], cipher[:], owcrypt.ECC_CURVE_SM2_STANDARD)

		if ret != owcrypt.SUCCESS || uint16(len(check)) != i {
			fmt.Println("sm2解密错误")
			return
		}

		if slide_cmp_equ(plain[:], check[:], i) != true {
			fmt.Println("sm2解密失败")
			return
		} else {
			fmt.Println("sm2 dec pass")
			fmt.Println("sm2 source plain value    : ", hex.EncodeToString(plain[:i]))
			fmt.Println("sm2 decrypted plain value : ", hex.EncodeToString(check[:i]))
		}
	}
}

func Test_sm2_signverify(t *testing.T) {
	prikey := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}
	pubkey := [64]byte{0x6C, 0x1E, 0xFB, 0x83, 0xBD, 0xEC, 0xBC, 0x47, 0x99, 0xCE, 0x03, 0xAB, 0xE2, 0x64, 0x3C, 0x18, 0x4B, 0x75, 0x0A, 0x28, 0xA9, 0x33, 0x7A, 0x22, 0x1C, 0x58, 0x1B, 0x7B, 0x11, 0x61, 0xDA, 0x01, 0xB0, 0x5E, 0x98, 0x96, 0x58, 0x61, 0xC8, 0x78, 0x16, 0xFB, 0x5D, 0x57, 0xFA, 0xD6, 0xB6, 0x30, 0x6F, 0x98, 0x2A, 0x36, 0x97, 0xA0, 0x11, 0x80, 0x7A, 0x5C, 0x4C, 0xBB, 0xD4, 0xEF, 0x8A, 0xC2}

	message := [32]byte{0xF6, 0xAA, 0xD7, 0x09, 0x86, 0x01, 0x19, 0x2E, 0x0E, 0x1C, 0xB3, 0x6B, 0x5F, 0x12, 0xF7, 0x3B, 0xBA, 0x39, 0xD7, 0x59, 0x6A, 0x10, 0x0F, 0x41, 0x84, 0xC0, 0xFE, 0x43, 0xFC, 0x3C, 0x11, 0x0A}

	ID := [4]byte{1, 2, 3, 4}

	i := uint16(0)

	for i = 1; i < 32; i++ {
		signature, err := owcrypt.Signature(prikey[:], ID[:], 4, message[:], i, owcrypt.ECC_CURVE_SM2_STANDARD)
		if err != owcrypt.SUCCESS {
			fmt.Println("sm2签名错误")
			return
		} else {
			fmt.Println("sm2 sign pass")
			fmt.Println("sm2 sign value : ", hex.EncodeToString(signature[:]))
		}
		if owcrypt.Verify(pubkey[:], ID[:], 4, message[:], i, signature[:], owcrypt.ECC_CURVE_SM2_STANDARD) != 0x0001 {
			fmt.Println("sm2验签错误")
			return
		} else {
			fmt.Println("sm2 verify pass")
		}
	}
}

func Test_getcurveorder(t *testing.T) {

	ret := owcrypt.GetCurveOrder(owcrypt.ECC_CURVE_SECP256K1)
	sret := hex.EncodeToString(ret[:])
	if sret == "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141" {
		fmt.Println("曲线secp256k1的阶是：")
		fmt.Println(sret)
	} else {
		fmt.Println("secp256k1获取失败！")
		return
	}

	ret = owcrypt.GetCurveOrder(owcrypt.ECC_CURVE_SECP256R1)
	sret = hex.EncodeToString(ret[:])
	if sret == "ffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551" {
		fmt.Println("曲线secp256r1的阶是：")
		fmt.Println(sret)
	} else {
		fmt.Println("secp256r1获取失败！")
		return
	}

	ret = owcrypt.GetCurveOrder(owcrypt.ECC_CURVE_SM2_STANDARD)
	sret = hex.EncodeToString(ret[:])
	if sret == "fffffffeffffffffffffffffffffffff7203df6b21c6052b53bbf40939d54123" {
		fmt.Println("曲线sm2_std的阶是：")
		fmt.Println(sret)
	} else {
		fmt.Println("sm2_std获取失败！")
		return
	}

	ret = owcrypt.GetCurveOrder(owcrypt.ECC_CURVE_ED25519_NORMAL)
	sret = hex.EncodeToString(ret[:])
	if sret == "1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed" {
		fmt.Println("曲线ed25519的阶是：")
		fmt.Println(hex.EncodeToString(ret[:]))
	} else {
		fmt.Println("ed25519获取失败！")
		return
	}
}

func Test_pointcompress(t *testing.T) {
	tmp1 := [65]byte{0x04, 0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	tmp2 := owcrypt.PointCompress(tmp1[:], owcrypt.ECC_CURVE_SECP256R1)
	fmt.Println(hex.EncodeToString(tmp2[:]))

	tmp3 := [64]byte{0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	tmp4 := owcrypt.PointCompress(tmp3[:], owcrypt.ECC_CURVE_SECP256R1)
	fmt.Println(hex.EncodeToString(tmp4[:]))

	// tmp5 := [65]byte{0x04, 0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	// tmp6 := PointCompress(tmp5[:], ECC_CURVE_SECP256R1)
	// fmt.Println(hex.EncodeToString(tmp6[:]))

	// tmp7 := [64]byte{0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73, 0x31, 0x85, 0x38, 0x77, 0x24, 0x08, 0x9D, 0x0E, 0x84, 0xC1, 0x7D, 0x54, 0x70, 0x5B, 0xD1, 0x23, 0xEB, 0x3A, 0x82, 0x54, 0xDA, 0x96, 0x43, 0x9E, 0xF6, 0x4B, 0x45, 0x07, 0x41, 0x2A, 0x94, 0x28}
	// tmp8 := PointCompress(tmp7[:], ECC_CURVE_SECP256R1)
	// fmt.Println(hex.EncodeToString(tmp8[:]))
}

func Test_pointdecompress(t *testing.T) {
	tmp1 := [33]byte{0x02, 0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73}
	tmp2 := owcrypt.PointDecompress(tmp1[:], owcrypt.ECC_CURVE_SECP256R1)
	fmt.Println(hex.EncodeToString(tmp2[:]))
}
func Test_hashset(t *testing.T) {
	tmp1 := [33]byte{0x02, 0x2C, 0x69, 0xDF, 0x94, 0x4F, 0xCD, 0xE9, 0xA7, 0xCB, 0xE7, 0xF9, 0x2F, 0x15, 0x6F, 0x0A, 0x90, 0xC1, 0x53, 0x29, 0x3F, 0x9C, 0xDC, 0x09, 0xCA, 0x64, 0x7E, 0xDC, 0x38, 0xE8, 0xD7, 0x70, 0x73}
	tmp2 := owcrypt.Hash(tmp1[:], 0, owcrypt.HASH_ALG_SHA1)
	fmt.Println(hex.EncodeToString(tmp2[:]))
}
