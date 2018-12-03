// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package types

import (
	"reflect"

	"github.com/33cn/chain33/types"
)

func init() {
	types.AllowUserExec = append(types.AllowUserExec, []byte(MultiSigX))
	types.RegistorExecutor(MultiSigX, NewType())
	types.RegisterDappFork(MultiSigX, "Enable", 1)
}

// MultiSigType multisig合约结构体
type MultiSigType struct {
	types.ExecTypeBase
}

// NewType new一个新的multisig合约实例
func NewType() *MultiSigType {
	c := &MultiSigType{}
	c.SetChild(c)
	return c
}

//GetPayload 获取交易payload的结构体信息：也就是multisig.pb.go中定义的具体交易类型结构体
func (m *MultiSigType) GetPayload() types.Message {
	return &MultiSigAction{}
}

//GetName 获取合约name
func (m *MultiSigType) GetName() string {
	return MultiSigX
}

//GetTypeMap 获取处理具体交易的接口函数，也就是exec.go中的函数实现，去掉EXEC_
func (m *MultiSigType) GetTypeMap() map[string]int32 {
	return map[string]int32{
		"MultiSigAccCreate":        ActionMultiSigAccCreate,
		"MultiSigOwnerOperate":     ActionMultiSigOwnerOperate,
		"MultiSigAccOperate":       ActionMultiSigAccOperate,
		"MultiSigConfirmTx":        ActionMultiSigConfirmTx,
		"MultiSigExecTransferTo":   ActionMultiSigExecTransferTo,
		"MultiSigExecTransferFrom": ActionMultiSigExecTransferFrom,
	}
}

//GetLogMap 获取具体执行Receiptlog对应的结构体：
func (m *MultiSigType) GetLogMap() map[int64]*types.LogInfo {
	return map[int64]*types.LogInfo{
		TyLogMultiSigAccCreate: {reflect.TypeOf(MultiSig{}), "LogMultiSigAccCreate"},

		TyLogMultiSigOwnerAdd:     {reflect.TypeOf(ReceiptOwnerAddOrDel{}), "LogMultiSigOwnerAdd"},
		TyLogMultiSigOwnerDel:     {reflect.TypeOf(ReceiptOwnerAddOrDel{}), "LogMultiSigOwnerDel"},
		TyLogMultiSigOwnerModify:  {reflect.TypeOf(ReceiptOwnerModOrRep{}), "LogMultiSigOwnerModify"},
		TyLogMultiSigOwnerReplace: {reflect.TypeOf(ReceiptOwnerModOrRep{}), "LogMultiSigOwnerReplace"},

		TyLogMultiSigAccWeightModify:     {reflect.TypeOf(ReceiptWeightModify{}), "LogMultiSigAccWeightModify"},
		TyLogMultiSigAccDailyLimitAdd:    {reflect.TypeOf(ReceiptDailyLimitOperate{}), "LogMultiSigAccDailyLimitAdd"},
		TyLogMultiSigAccDailyLimitModify: {reflect.TypeOf(ReceiptDailyLimitOperate{}), "LogMultiSigAccDailyLimitModify"},

		TyLogMultiSigConfirmTx:       {reflect.TypeOf(ReceiptConfirmTx{}), "LogMultiSigConfirmTx"},
		TyLogMultiSigConfirmTxRevoke: {reflect.TypeOf(ReceiptConfirmTx{}), "LogMultiSigConfirmTxRevoke"},

		TyLogDailyLimitUpdate: {reflect.TypeOf(ReceiptAccDailyLimitUpdate{}), "LogAccDailyLimitUpdate"},
		TyLogMultiSigTx:       {reflect.TypeOf(ReceiptMultiSigTx{}), "LogMultiSigAccTx"},
		TyLogTxCountUpdate:    {reflect.TypeOf(ReceiptTxCountUpdate{}), "LogTxCountUpdate"},
	}
}

//DecodePayload 解码交易的Payload信息
func (m MultiSigType) DecodePayload(tx *types.Transaction) (types.Message, error) {
	var action MultiSigAction
	err := types.Decode(tx.Payload, &action)
	if err != nil {
		return nil, err
	}
	return &action, nil
}

//ActionName 获取actionid对应的name
func (m MultiSigType) ActionName(tx *types.Transaction) string {
	var g MultiSigAction
	err := types.Decode(tx.Payload, &g)
	if err != nil {
		return "unkown-MultiSig-action-err"
	}
	if g.Ty == ActionMultiSigAccCreate && g.GetMultiSigAccCreate() != nil {
		return "MultiSigAccCreate"
	} else if g.Ty == ActionMultiSigOwnerOperate && g.GetMultiSigOwnerOperate() != nil {
		return "MultiSigOwnerOperate"
	} else if g.Ty == ActionMultiSigAccOperate && g.GetMultiSigAccOperate() != nil {
		return "MultiSigAccOperate"
	} else if g.Ty == ActionMultiSigConfirmTx && g.GetMultiSigConfirmTx() != nil {
		return "MultiSigTxConfirm"
	} else if g.Ty == ActionMultiSigExecTransferTo && g.GetMultiSigExecTransferTo() != nil {
		return "MultiSigExecTransfer"
	} else if g.Ty == ActionMultiSigExecTransferFrom && g.GetMultiSigExecTransferFrom() != nil {
		return "MultiSigAccExecTransfer"
	}
	return "unkown"
}
