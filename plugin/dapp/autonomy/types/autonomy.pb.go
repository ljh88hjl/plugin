// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autonomy.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// message for execs.Autonomy
type AutonomyAction struct {
	// Types that are valid to be assigned to Value:
	//	*AutonomyAction_PropBoard
	//	*AutonomyAction_RvkPropBoard
	//	*AutonomyAction_VotePropBoard
	//	*AutonomyAction_TmintPropBoard
	//	*AutonomyAction_PropProject
	//	*AutonomyAction_RvkPropProject
	//	*AutonomyAction_VotePropProject
	//	*AutonomyAction_PubVotePropProject
	//	*AutonomyAction_TmintPropProject
	//	*AutonomyAction_PropRule
	//	*AutonomyAction_RvkPropRule
	//	*AutonomyAction_VotePropRule
	//	*AutonomyAction_TmintPropRule
	//	*AutonomyAction_Transfer
	//	*AutonomyAction_CommentProp
	Value                isAutonomyAction_Value `protobuf_oneof:"value"`
	Ty                   int32                  `protobuf:"varint,16,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AutonomyAction) Reset()         { *m = AutonomyAction{} }
func (m *AutonomyAction) String() string { return proto.CompactTextString(m) }
func (*AutonomyAction) ProtoMessage()    {}
func (*AutonomyAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_0246b47df8434d60, []int{0}
}

func (m *AutonomyAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutonomyAction.Unmarshal(m, b)
}
func (m *AutonomyAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutonomyAction.Marshal(b, m, deterministic)
}
func (m *AutonomyAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutonomyAction.Merge(m, src)
}
func (m *AutonomyAction) XXX_Size() int {
	return xxx_messageInfo_AutonomyAction.Size(m)
}
func (m *AutonomyAction) XXX_DiscardUnknown() {
	xxx_messageInfo_AutonomyAction.DiscardUnknown(m)
}

var xxx_messageInfo_AutonomyAction proto.InternalMessageInfo

type isAutonomyAction_Value interface {
	isAutonomyAction_Value()
}

type AutonomyAction_PropBoard struct {
	PropBoard *ProposalBoard `protobuf:"bytes,1,opt,name=propBoard,proto3,oneof"`
}

type AutonomyAction_RvkPropBoard struct {
	RvkPropBoard *RevokeProposalBoard `protobuf:"bytes,2,opt,name=rvkPropBoard,proto3,oneof"`
}

type AutonomyAction_VotePropBoard struct {
	VotePropBoard *VoteProposalBoard `protobuf:"bytes,3,opt,name=votePropBoard,proto3,oneof"`
}

type AutonomyAction_TmintPropBoard struct {
	TmintPropBoard *TerminateProposalBoard `protobuf:"bytes,4,opt,name=tmintPropBoard,proto3,oneof"`
}

type AutonomyAction_PropProject struct {
	PropProject *ProposalProject `protobuf:"bytes,5,opt,name=propProject,proto3,oneof"`
}

type AutonomyAction_RvkPropProject struct {
	RvkPropProject *RevokeProposalProject `protobuf:"bytes,6,opt,name=rvkPropProject,proto3,oneof"`
}

type AutonomyAction_VotePropProject struct {
	VotePropProject *VoteProposalProject `protobuf:"bytes,7,opt,name=votePropProject,proto3,oneof"`
}

type AutonomyAction_PubVotePropProject struct {
	PubVotePropProject *PubVoteProposalProject `protobuf:"bytes,8,opt,name=pubVotePropProject,proto3,oneof"`
}

type AutonomyAction_TmintPropProject struct {
	TmintPropProject *TerminateProposalProject `protobuf:"bytes,9,opt,name=tmintPropProject,proto3,oneof"`
}

type AutonomyAction_PropRule struct {
	PropRule *ProposalRule `protobuf:"bytes,10,opt,name=propRule,proto3,oneof"`
}

type AutonomyAction_RvkPropRule struct {
	RvkPropRule *RevokeProposalRule `protobuf:"bytes,11,opt,name=rvkPropRule,proto3,oneof"`
}

type AutonomyAction_VotePropRule struct {
	VotePropRule *VoteProposalRule `protobuf:"bytes,12,opt,name=votePropRule,proto3,oneof"`
}

type AutonomyAction_TmintPropRule struct {
	TmintPropRule *TerminateProposalRule `protobuf:"bytes,13,opt,name=tmintPropRule,proto3,oneof"`
}

type AutonomyAction_Transfer struct {
	Transfer *TransferFund `protobuf:"bytes,14,opt,name=transfer,proto3,oneof"`
}

type AutonomyAction_CommentProp struct {
	CommentProp *Comment `protobuf:"bytes,15,opt,name=commentProp,proto3,oneof"`
}

func (*AutonomyAction_PropBoard) isAutonomyAction_Value() {}

func (*AutonomyAction_RvkPropBoard) isAutonomyAction_Value() {}

func (*AutonomyAction_VotePropBoard) isAutonomyAction_Value() {}

func (*AutonomyAction_TmintPropBoard) isAutonomyAction_Value() {}

func (*AutonomyAction_PropProject) isAutonomyAction_Value() {}

func (*AutonomyAction_RvkPropProject) isAutonomyAction_Value() {}

func (*AutonomyAction_VotePropProject) isAutonomyAction_Value() {}

func (*AutonomyAction_PubVotePropProject) isAutonomyAction_Value() {}

func (*AutonomyAction_TmintPropProject) isAutonomyAction_Value() {}

func (*AutonomyAction_PropRule) isAutonomyAction_Value() {}

func (*AutonomyAction_RvkPropRule) isAutonomyAction_Value() {}

func (*AutonomyAction_VotePropRule) isAutonomyAction_Value() {}

func (*AutonomyAction_TmintPropRule) isAutonomyAction_Value() {}

func (*AutonomyAction_Transfer) isAutonomyAction_Value() {}

func (*AutonomyAction_CommentProp) isAutonomyAction_Value() {}

func (m *AutonomyAction) GetValue() isAutonomyAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AutonomyAction) GetPropBoard() *ProposalBoard {
	if x, ok := m.GetValue().(*AutonomyAction_PropBoard); ok {
		return x.PropBoard
	}
	return nil
}

func (m *AutonomyAction) GetRvkPropBoard() *RevokeProposalBoard {
	if x, ok := m.GetValue().(*AutonomyAction_RvkPropBoard); ok {
		return x.RvkPropBoard
	}
	return nil
}

func (m *AutonomyAction) GetVotePropBoard() *VoteProposalBoard {
	if x, ok := m.GetValue().(*AutonomyAction_VotePropBoard); ok {
		return x.VotePropBoard
	}
	return nil
}

func (m *AutonomyAction) GetTmintPropBoard() *TerminateProposalBoard {
	if x, ok := m.GetValue().(*AutonomyAction_TmintPropBoard); ok {
		return x.TmintPropBoard
	}
	return nil
}

func (m *AutonomyAction) GetPropProject() *ProposalProject {
	if x, ok := m.GetValue().(*AutonomyAction_PropProject); ok {
		return x.PropProject
	}
	return nil
}

func (m *AutonomyAction) GetRvkPropProject() *RevokeProposalProject {
	if x, ok := m.GetValue().(*AutonomyAction_RvkPropProject); ok {
		return x.RvkPropProject
	}
	return nil
}

func (m *AutonomyAction) GetVotePropProject() *VoteProposalProject {
	if x, ok := m.GetValue().(*AutonomyAction_VotePropProject); ok {
		return x.VotePropProject
	}
	return nil
}

func (m *AutonomyAction) GetPubVotePropProject() *PubVoteProposalProject {
	if x, ok := m.GetValue().(*AutonomyAction_PubVotePropProject); ok {
		return x.PubVotePropProject
	}
	return nil
}

func (m *AutonomyAction) GetTmintPropProject() *TerminateProposalProject {
	if x, ok := m.GetValue().(*AutonomyAction_TmintPropProject); ok {
		return x.TmintPropProject
	}
	return nil
}

func (m *AutonomyAction) GetPropRule() *ProposalRule {
	if x, ok := m.GetValue().(*AutonomyAction_PropRule); ok {
		return x.PropRule
	}
	return nil
}

func (m *AutonomyAction) GetRvkPropRule() *RevokeProposalRule {
	if x, ok := m.GetValue().(*AutonomyAction_RvkPropRule); ok {
		return x.RvkPropRule
	}
	return nil
}

func (m *AutonomyAction) GetVotePropRule() *VoteProposalRule {
	if x, ok := m.GetValue().(*AutonomyAction_VotePropRule); ok {
		return x.VotePropRule
	}
	return nil
}

func (m *AutonomyAction) GetTmintPropRule() *TerminateProposalRule {
	if x, ok := m.GetValue().(*AutonomyAction_TmintPropRule); ok {
		return x.TmintPropRule
	}
	return nil
}

func (m *AutonomyAction) GetTransfer() *TransferFund {
	if x, ok := m.GetValue().(*AutonomyAction_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *AutonomyAction) GetCommentProp() *Comment {
	if x, ok := m.GetValue().(*AutonomyAction_CommentProp); ok {
		return x.CommentProp
	}
	return nil
}

func (m *AutonomyAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AutonomyAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AutonomyAction_OneofMarshaler, _AutonomyAction_OneofUnmarshaler, _AutonomyAction_OneofSizer, []interface{}{
		(*AutonomyAction_PropBoard)(nil),
		(*AutonomyAction_RvkPropBoard)(nil),
		(*AutonomyAction_VotePropBoard)(nil),
		(*AutonomyAction_TmintPropBoard)(nil),
		(*AutonomyAction_PropProject)(nil),
		(*AutonomyAction_RvkPropProject)(nil),
		(*AutonomyAction_VotePropProject)(nil),
		(*AutonomyAction_PubVotePropProject)(nil),
		(*AutonomyAction_TmintPropProject)(nil),
		(*AutonomyAction_PropRule)(nil),
		(*AutonomyAction_RvkPropRule)(nil),
		(*AutonomyAction_VotePropRule)(nil),
		(*AutonomyAction_TmintPropRule)(nil),
		(*AutonomyAction_Transfer)(nil),
		(*AutonomyAction_CommentProp)(nil),
	}
}

func _AutonomyAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AutonomyAction)
	// value
	switch x := m.Value.(type) {
	case *AutonomyAction_PropBoard:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PropBoard); err != nil {
			return err
		}
	case *AutonomyAction_RvkPropBoard:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RvkPropBoard); err != nil {
			return err
		}
	case *AutonomyAction_VotePropBoard:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VotePropBoard); err != nil {
			return err
		}
	case *AutonomyAction_TmintPropBoard:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TmintPropBoard); err != nil {
			return err
		}
	case *AutonomyAction_PropProject:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PropProject); err != nil {
			return err
		}
	case *AutonomyAction_RvkPropProject:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RvkPropProject); err != nil {
			return err
		}
	case *AutonomyAction_VotePropProject:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VotePropProject); err != nil {
			return err
		}
	case *AutonomyAction_PubVotePropProject:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PubVotePropProject); err != nil {
			return err
		}
	case *AutonomyAction_TmintPropProject:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TmintPropProject); err != nil {
			return err
		}
	case *AutonomyAction_PropRule:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PropRule); err != nil {
			return err
		}
	case *AutonomyAction_RvkPropRule:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RvkPropRule); err != nil {
			return err
		}
	case *AutonomyAction_VotePropRule:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VotePropRule); err != nil {
			return err
		}
	case *AutonomyAction_TmintPropRule:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TmintPropRule); err != nil {
			return err
		}
	case *AutonomyAction_Transfer:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *AutonomyAction_CommentProp:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CommentProp); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AutonomyAction.Value has unexpected type %T", x)
	}
	return nil
}

func _AutonomyAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AutonomyAction)
	switch tag {
	case 1: // value.propBoard
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProposalBoard)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_PropBoard{msg}
		return true, err
	case 2: // value.rvkPropBoard
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RevokeProposalBoard)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_RvkPropBoard{msg}
		return true, err
	case 3: // value.votePropBoard
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VoteProposalBoard)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_VotePropBoard{msg}
		return true, err
	case 4: // value.tmintPropBoard
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TerminateProposalBoard)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_TmintPropBoard{msg}
		return true, err
	case 5: // value.propProject
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProposalProject)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_PropProject{msg}
		return true, err
	case 6: // value.rvkPropProject
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RevokeProposalProject)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_RvkPropProject{msg}
		return true, err
	case 7: // value.votePropProject
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VoteProposalProject)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_VotePropProject{msg}
		return true, err
	case 8: // value.pubVotePropProject
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PubVoteProposalProject)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_PubVotePropProject{msg}
		return true, err
	case 9: // value.tmintPropProject
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TerminateProposalProject)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_TmintPropProject{msg}
		return true, err
	case 10: // value.propRule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProposalRule)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_PropRule{msg}
		return true, err
	case 11: // value.rvkPropRule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RevokeProposalRule)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_RvkPropRule{msg}
		return true, err
	case 12: // value.votePropRule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VoteProposalRule)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_VotePropRule{msg}
		return true, err
	case 13: // value.tmintPropRule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TerminateProposalRule)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_TmintPropRule{msg}
		return true, err
	case 14: // value.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferFund)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_Transfer{msg}
		return true, err
	case 15: // value.commentProp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Comment)
		err := b.DecodeMessage(msg)
		m.Value = &AutonomyAction_CommentProp{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AutonomyAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AutonomyAction)
	// value
	switch x := m.Value.(type) {
	case *AutonomyAction_PropBoard:
		s := proto.Size(x.PropBoard)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_RvkPropBoard:
		s := proto.Size(x.RvkPropBoard)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_VotePropBoard:
		s := proto.Size(x.VotePropBoard)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_TmintPropBoard:
		s := proto.Size(x.TmintPropBoard)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_PropProject:
		s := proto.Size(x.PropProject)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_RvkPropProject:
		s := proto.Size(x.RvkPropProject)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_VotePropProject:
		s := proto.Size(x.VotePropProject)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_PubVotePropProject:
		s := proto.Size(x.PubVotePropProject)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_TmintPropProject:
		s := proto.Size(x.TmintPropProject)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_PropRule:
		s := proto.Size(x.PropRule)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_RvkPropRule:
		s := proto.Size(x.RvkPropRule)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_VotePropRule:
		s := proto.Size(x.VotePropRule)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_TmintPropRule:
		s := proto.Size(x.TmintPropRule)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_Transfer:
		s := proto.Size(x.Transfer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutonomyAction_CommentProp:
		s := proto.Size(x.CommentProp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AutonomyAction)(nil), "types.AutonomyAction")
}

func init() { proto.RegisterFile("autonomy.proto", fileDescriptor_0246b47df8434d60) }

var fileDescriptor_0246b47df8434d60 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0xaf, 0xd2, 0x40,
	0x14, 0x85, 0x01, 0xed, 0xe3, 0xbd, 0x5b, 0x5a, 0xc8, 0xd5, 0x68, 0x25, 0x1a, 0x89, 0x2b, 0x56,
	0x24, 0xa2, 0x2b, 0x13, 0x12, 0x40, 0x83, 0x6c, 0x8c, 0x4d, 0x43, 0xd8, 0x17, 0x1c, 0x13, 0xa4,
	0xed, 0x34, 0xc3, 0xb4, 0x49, 0xff, 0x71, 0xd7, 0xa6, 0xb7, 0xd3, 0x1f, 0x53, 0xea, 0x0e, 0x7a,
	0xcf, 0xf7, 0x25, 0xe7, 0x40, 0xc1, 0xf6, 0x13, 0xc9, 0x23, 0x1e, 0x66, 0x8b, 0x58, 0x70, 0xc9,
	0xd1, 0x90, 0x59, 0xcc, 0x6e, 0x53, 0xf3, 0xc4, 0x7d, 0xf1, 0xab, 0x78, 0x36, 0xb5, 0x62, 0xc1,
	0xff, 0xb0, 0xb3, 0x54, 0x5f, 0x41, 0x24, 0x01, 0x2b, 0x3e, 0x7f, 0xf8, 0x3b, 0x04, 0x7b, 0xa3,
	0x0c, 0x9b, 0xb3, 0xbc, 0xf0, 0x08, 0x3f, 0xc3, 0x53, 0x2c, 0x78, 0xbc, 0xcd, 0x05, 0x4e, 0x7f,
	0xd6, 0x9f, 0x9b, 0xcb, 0x97, 0x0b, 0xb2, 0x2e, 0x5c, 0xc1, 0x63, 0x7e, 0xf3, 0x03, 0xba, 0xed,
	0x7b, 0x5e, 0x1d, 0xc4, 0x35, 0x8c, 0x44, 0x7a, 0x75, 0x2b, 0x70, 0x40, 0xe0, 0x54, 0x81, 0x1e,
	0x4b, 0xf9, 0x95, 0xb5, 0x71, 0x8d, 0xc0, 0x35, 0x58, 0x29, 0x97, 0xac, 0x56, 0x3c, 0x23, 0x85,
	0xa3, 0x14, 0x47, 0x75, 0x6b, 0x0a, 0x74, 0x00, 0xbf, 0x83, 0x2d, 0xc3, 0x4b, 0x24, 0x6b, 0xc5,
	0x73, 0x52, 0xbc, 0x53, 0x8a, 0x03, 0x13, 0xe1, 0x25, 0xf2, 0xef, 0x3d, 0x2d, 0x0c, 0xbf, 0x80,
	0x99, 0x37, 0x73, 0x8b, 0xd9, 0x1c, 0x83, 0x2c, 0xaf, 0x5a, 0x23, 0xa8, 0xeb, 0xbe, 0xe7, 0x35,
	0xc3, 0xb8, 0x03, 0x5b, 0xd5, 0x2a, 0xf1, 0x07, 0xc2, 0xdf, 0x76, 0x4e, 0x51, 0x4b, 0x5a, 0x14,
	0xee, 0x60, 0x5c, 0xb6, 0x2b, 0x45, 0x43, 0x6d, 0xd3, 0xe6, 0x20, 0xb5, 0xa6, 0x0d, 0xe1, 0x4f,
	0xc0, 0x38, 0x39, 0x1d, 0x5b, 0xaa, 0x47, 0x6d, 0x18, 0xb7, 0x0e, 0xe8, 0xb6, 0x0e, 0x14, 0x7f,
	0xc0, 0xa4, 0x9a, 0xab, 0xd4, 0x3d, 0x91, 0xee, 0xfd, 0xff, 0x76, 0xae, 0x85, 0x77, 0x28, 0x7e,
	0x84, 0xc7, 0x7c, 0x3e, 0x2f, 0x09, 0x98, 0x03, 0xa4, 0x79, 0xd1, 0x1a, 0x3a, 0x3f, 0xed, 0x7b,
	0x5e, 0x15, 0xc3, 0x15, 0x98, 0x6a, 0x2c, 0xa2, 0x4c, 0xa2, 0xde, 0x74, 0xee, 0xab, 0xd8, 0x66,
	0x1e, 0x57, 0x30, 0x2a, 0x47, 0x22, 0x7e, 0x44, 0xfc, 0xeb, 0x8e, 0x59, 0x15, 0xad, 0xc5, 0xf1,
	0x1b, 0x58, 0x55, 0x09, 0xe2, 0x2d, 0xed, 0xf7, 0xbd, 0x2b, 0xaf, 0x24, 0x3a, 0x94, 0xd7, 0x96,
	0xc2, 0x8f, 0x6e, 0xbf, 0x99, 0x70, 0x6c, 0xad, 0xf6, 0x41, 0x3d, 0xde, 0x25, 0x51, 0xfe, 0xdf,
	0xac, 0x62, 0xb8, 0x04, 0xf3, 0xcc, 0xc3, 0x90, 0x15, 0x16, 0x67, 0x4c, 0x94, 0xad, 0xa8, 0xaf,
	0xc5, 0x25, 0xef, 0xda, 0x08, 0xa1, 0x0d, 0x03, 0x99, 0x39, 0x93, 0x59, 0x7f, 0x6e, 0x78, 0x03,
	0x99, 0x6d, 0x87, 0x60, 0xa4, 0x7e, 0x90, 0xb0, 0xd3, 0x03, 0xbd, 0xff, 0x9f, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xf5, 0x6a, 0x69, 0x9d, 0x40, 0x04, 0x00, 0x00,
}
