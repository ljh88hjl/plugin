// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"testing"

	"github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	"github.com/33cn/chain33/util"
	auty "github.com/33cn/plugin/plugin/dapp/autonomy/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecLocalRule(t *testing.T) {
	_, sdb, kvdb := util.CreateTestDB()
	au := &Autonomy{}
	au.SetLocalDB(kvdb)
	//TyLogPropRule
	cur := &auty.AutonomyProposalRule{
		PropRule:   &auty.ProposalRule{},
		CurRule:    &auty.RuleConfig{},
		VoteResult: &auty.VoteResult{},
		Status:     auty.AutonomyStatusProposalRule,
		Address:    "11111111111111",
		Height:     1,
		Index:      2,
	}
	receiptRule := &auty.ReceiptProposalRule{
		Prev:    nil,
		Current: cur,
	}
	receipt := &types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogPropRule, Log: types.Encode(receiptRule)},
		},
	}
	set, err := au.execLocalRule(receipt)
	require.NoError(t, err)
	require.NotNil(t, set)
	//save to database
	saveKvs(sdb, set.KV)

	// check
	checkExecLocalRule(t, kvdb, cur)

	// TyLogRvkPropRule
	pre1 := copyAutonomyProposalRule(cur)
	cur.Status = auty.AutonomyStatusRvkPropRule
	receiptRule1 := &auty.ReceiptProposalRule{
		Prev:    pre1,
		Current: cur,
	}
	set, err = au.execLocalRule(&types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogRvkPropRule, Log: types.Encode(receiptRule1)},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, set)
	//save to database
	saveKvs(sdb, set.KV)

	// check
	checkExecLocalRule(t, kvdb, cur)

	// TyLogVotePropRule
	cur.Status = auty.AutonomyStatusProposalRule
	cur.Height = 1
	cur.Index = 2
	pre2 := copyAutonomyProposalRule(cur)
	cur.Status = auty.AutonomyStatusVotePropRule
	cur.Height = 1
	cur.Index = 2
	cur.Address = "2222222222222"
	receiptRule2 := &auty.ReceiptProposalRule{
		Prev:    pre2,
		Current: cur,
	}
	set, err = au.execLocalRule(&types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogVotePropRule, Log: types.Encode(receiptRule2)},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, set)
	//save to database
	saveKvs(sdb, set.KV)
	// check
	checkExecLocalRule(t, kvdb, cur)
}

func TestExecDelLocalRule(t *testing.T) {
	_, sdb, kvdb := util.CreateTestDB()
	au := &Autonomy{}
	au.SetLocalDB(kvdb)
	//TyLogPropRule
	cur := &auty.AutonomyProposalRule{
		PropRule:   &auty.ProposalRule{},
		CurRule:    &auty.RuleConfig{},
		VoteResult: &auty.VoteResult{},
		Status:     auty.AutonomyStatusProposalRule,
		Address:    "11111111111111",
		Height:     1,
		Index:      2,
	}
	receiptRule := &auty.ReceiptProposalRule{
		Prev:    nil,
		Current: cur,
	}
	receipt := &types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogPropRule, Log: types.Encode(receiptRule)},
		},
	}
	// 先执行local然后进行删除
	set, err := au.execLocalRule(receipt)
	require.NoError(t, err)
	require.NotNil(t, set)
	saveKvs(sdb, set.KV)

	set, err = au.execDelLocalRule(receipt)
	require.NoError(t, err)
	require.NotNil(t, set)
	saveKvs(sdb, set.KV)
	// check
	table := NewRuleTable(au.GetLocalDB())
	query := table.GetQuery(kvdb)
	_, err = query.ListIndex("primary", nil, nil, 10, 0)
	assert.Equal(t, err, types.ErrNotFound)
	_, err = query.ListIndex("addr", nil, nil, 10, 0)
	assert.Equal(t, err, types.ErrNotFound)
	_, err = query.ListIndex("status", nil, nil, 10, 0)
	assert.Equal(t, err, types.ErrNotFound)
	_, err = query.ListIndex("addr_status", nil, nil, 10, 0)
	assert.Equal(t, err, types.ErrNotFound)

	// TyLogVotePropRule
	pre1 := copyAutonomyProposalRule(cur)
	cur.Status = auty.AutonomyStatusVotePropRule
	receiptRule2 := &auty.ReceiptProposalRule{
		Prev:    pre1,
		Current: cur,
	}
	// 先执行local然后进行删除
	set, err = au.execLocalRule(&types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogVotePropRule, Log: types.Encode(receiptRule2)},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, set)
	saveKvs(sdb, set.KV)
	// check
	checkExecLocalRule(t, kvdb, cur)

	set, err = au.execDelLocalRule(&types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogVotePropRule, Log: types.Encode(receiptRule2)},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, set)
	saveKvs(sdb, set.KV)
	// check
	checkExecLocalRule(t, kvdb, pre1)
}

func checkExecLocalRule(t *testing.T, kvdb db.KVDB, cur *auty.AutonomyProposalRule) {
	table := NewRuleTable(kvdb)
	query := table.GetQuery(kvdb)

	rows, err := query.ListIndex("primary", nil, nil, 10, 0)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(rows[0].Primary), dapp.HeightIndexStr(1, 2))

	rows, err = query.ListIndex("addr", nil, nil, 10, 0)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(rows))

	rows, err = query.ListIndex("status", nil, nil, 10, 0)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(rows))

	rows, err = query.ListIndex("addr_status", nil, nil, 10, 0)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(rows))

	prop, ok := rows[0].Data.(*auty.AutonomyProposalRule)
	assert.Equal(t, true, ok)
	assert.Equal(t, prop, cur)

}

func TestGetProposalRule(t *testing.T) {
	au := &Autonomy{
		dapp.DriverBase{},
	}
	_, storedb, _ := util.CreateTestDB()
	au.SetStateDB(storedb)
	tx := "1111111111111111111"
	storedb.Set(propRuleID(tx), types.Encode(&auty.AutonomyProposalRule{}))
	rsp, err := au.getProposalRule(&types.ReqString{Data: tx})
	require.NoError(t, err)
	require.NotNil(t, rsp)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalRule).PropRules), 1)
}

func TestListProposalRule(t *testing.T) {
	au := &Autonomy{
		dapp.DriverBase{},
	}
	_, sdb, kvdb := util.CreateTestDB()
	au.SetLocalDB(kvdb)

	type statu struct {
		status int32
		height int64
		index  int64
	}

	testcase1 := []statu{
		{auty.AutonomyStatusRvkPropRule, 10, 2},
		{auty.AutonomyStatusVotePropRule, 15, 1},
		{auty.AutonomyStatusTmintPropRule, 20, 1},
	}
	testcase2 := []statu{
		{auty.AutonomyStatusProposalRule, 10, 1},
		{auty.AutonomyStatusProposalRule, 20, 2},
		{auty.AutonomyStatusProposalRule, 20, 5},
	}
	var testcase []statu
	testcase = append(testcase, testcase1...)
	testcase = append(testcase, testcase2...)
	cur := &auty.AutonomyProposalRule{
		PropRule:   &auty.ProposalRule{},
		CurRule:    &auty.RuleConfig{},
		VoteResult: &auty.VoteResult{},
		Status:     auty.AutonomyStatusProposalRule,
		Address:    "11111111111111",
		Height:     1,
		Index:      2,
	}

	//将数据保存下去
	var kvs []*types.KeyValue
	table := NewRuleTable(kvdb)
	for _, tcase := range testcase {
		cur.Status = tcase.status
		cur.Height = tcase.height
		cur.Index = int32(tcase.index)

		err := table.Replace(cur)
		require.NoError(t, err)
		kv, err := table.Save()
		require.NoError(t, err)
		kvs = append(kvs, kv...)
	}
	saveKvs(sdb, kvs)

	// 反向查找
	req := &auty.ReqQueryProposalRule{
		Status:    auty.AutonomyStatusProposalRule,
		Count:     10,
		Direction: 0,
		Index:     -1,
	}
	rsp, err := au.listProposalRule(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalRule).PropRules), len(testcase2))
	k := 2
	for _, tcase := range testcase2 {
		require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[k].Height, tcase.height)
		require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[k].Index, int32(tcase.index))
		k--
	}

	// 正向查找
	req = &auty.ReqQueryProposalRule{
		Status:    auty.AutonomyStatusProposalRule,
		Count:     10,
		Direction: 1,
		Index:     -1,
	}
	rsp, err = au.listProposalRule(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalRule).PropRules), len(testcase2))
	for i, tcase := range testcase2 {
		require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[i].Height, tcase.height)
		require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[i].Index, int32(tcase.index))
	}

	// 翻页查找
	req = &auty.ReqQueryProposalRule{
		Status:    auty.AutonomyStatusProposalRule,
		Count:     1,
		Direction: 0,
		Index:     -1,
	}
	rsp, err = au.listProposalRule(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalRule).PropRules), 1)
	height := rsp.(*auty.ReplyQueryProposalRule).PropRules[0].Height
	index := rsp.(*auty.ReplyQueryProposalRule).PropRules[0].Index
	require.Equal(t, height, testcase2[2].height)
	require.Equal(t, index, int32(testcase2[2].index))
	//
	req = &auty.ReqQueryProposalRule{
		Status:    auty.AutonomyStatusProposalRule,
		Count:     10,
		Direction: 0,
		Height:    height,
		Index:     index,
	}
	rsp, err = au.listProposalRule(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalRule).PropRules), 2)
	require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[0].Height, testcase2[1].height)
	require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[0].Index, int32(testcase2[1].index))
	require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[1].Height, testcase2[0].height)
	require.Equal(t, rsp.(*auty.ReplyQueryProposalRule).PropRules[1].Index, int32(testcase2[0].index))
}

func TestExecLocalCommentProp(t *testing.T) {
	au := &Autonomy{}
	propID := "11111111111111"
	Repcmt := "2222222222"
	comment := "3333333333"
	receiptCmt := &auty.ReceiptProposalComment{
		Cmt: &auty.Comment{
			ProposalID: propID,
			RepHash:    Repcmt,
			Comment:    comment,
		},
		Height: 11,
		Index:  1,
	}
	receipt := &types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogCommentProp, Log: types.Encode(receiptCmt)},
		},
	}
	set, err := au.execLocalCommentProp(receipt)
	require.NoError(t, err)
	require.NotNil(t, set)
	require.Equal(t, set.KV[0].Key, calcCommentHeight(propID,
		dapp.HeightIndexStr(receiptCmt.Height, int64(receiptCmt.Index))))
	require.NotNil(t, set.KV[0].Value)
}

func TestExecDelLocalCommentProp(t *testing.T) {
	au := &Autonomy{}
	propID := "11111111111111"
	Repcmt := "2222222222"
	comment := "3333333333"
	receiptCmt := &auty.ReceiptProposalComment{
		Cmt: &auty.Comment{
			ProposalID: propID,
			RepHash:    Repcmt,
			Comment:    comment,
		},
		Height: 11,
		Index:  1,
	}
	receipt := &types.ReceiptData{
		Logs: []*types.ReceiptLog{
			{Ty: auty.TyLogCommentProp, Log: types.Encode(receiptCmt)},
		},
	}
	set, err := au.execDelLocalCommentProp(receipt)
	require.NoError(t, err)
	require.NotNil(t, set)
	require.Equal(t, set.KV[0].Key, calcCommentHeight(propID,
		dapp.HeightIndexStr(receiptCmt.Height, int64(receiptCmt.Index))))
	require.Nil(t, set.KV[0].Value)
}

func TestListProposalComment(t *testing.T) {
	au := &Autonomy{
		dapp.DriverBase{},
	}
	_, _, kvdb := util.CreateTestDB()
	au.SetLocalDB(kvdb)

	type statu struct {
		propID string
		height int64
		index  int64
	}

	propID := "3333333333"
	propID1 := "2222222"
	propID2 := "111111111111"

	testcase1 := []statu{
		{propID, 10, 2},
		{propID1, 15, 1},
		{propID, 20, 1},
	}
	testcase2 := []statu{
		{propID2, 10, 1},
		{propID2, 20, 2},
		{propID2, 20, 5},
	}
	var testcase []statu
	testcase = append(testcase, testcase1...)
	testcase = append(testcase, testcase2...)
	cur := &auty.RelationCmt{
		RepHash: "aaaaaa",
		Comment: "bbbbbbbbbb",
	}
	for _, tcase := range testcase {
		key := calcCommentHeight(tcase.propID,
			dapp.HeightIndexStr(tcase.height, int64(tcase.index)))
		cur.Height = tcase.height
		cur.Index = int32(tcase.index)
		value := types.Encode(cur)
		kvdb.Set(key, value)
	}

	// 反向查找
	req := &auty.ReqQueryProposalComment{
		ProposalID: propID2,
		Count:      10,
		Direction:  0,
		Index:      -1,
	}
	rsp, err := au.listProposalComment(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalComment).RltCmt), len(testcase2))
	k := 2
	for _, tcase := range testcase2 {
		require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[k].Height, tcase.height)
		require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[k].Index, int32(tcase.index))
		k--
	}

	// 正向查找
	req = &auty.ReqQueryProposalComment{
		ProposalID: propID2,
		Count:      10,
		Direction:  1,
		Index:      -1,
	}
	rsp, err = au.listProposalComment(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalComment).RltCmt), len(testcase2))
	for i, tcase := range testcase2 {
		require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[i].Height, tcase.height)
		require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[i].Index, int32(tcase.index))
	}

	// 翻页查找
	req = &auty.ReqQueryProposalComment{
		ProposalID: propID2,
		Count:      1,
		Direction:  0,
		Index:      -1,
	}
	rsp, err = au.listProposalComment(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalComment).RltCmt), 1)
	height := rsp.(*auty.ReplyQueryProposalComment).RltCmt[0].Height
	index := rsp.(*auty.ReplyQueryProposalComment).RltCmt[0].Index
	require.Equal(t, height, testcase2[2].height)
	require.Equal(t, index, int32(testcase2[2].index))
	//
	req = &auty.ReqQueryProposalComment{
		ProposalID: propID2,
		Count:      10,
		Direction:  0,
		Height:     height,
		Index:      index,
	}
	rsp, err = au.listProposalComment(req)
	require.NoError(t, err)
	require.Equal(t, len(rsp.(*auty.ReplyQueryProposalComment).RltCmt), 2)
	require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[0].Height, testcase2[1].height)
	require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[0].Index, int32(testcase2[1].index))
	require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[1].Height, testcase2[0].height)
	require.Equal(t, rsp.(*auty.ReplyQueryProposalComment).RltCmt[1].Index, int32(testcase2[0].index))
}
