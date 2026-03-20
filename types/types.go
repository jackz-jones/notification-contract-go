package types

import "time"

// EnterpriseInfo 企业信息
type EnterpriseInfo struct {
	ID         string    `json:"id"`         // 链上唯一 id，企业信息文件 id
	OriginHash string    `json:"originHash"` // 企业信息文件原文 hash 值
	CreatedAt  time.Time `json:"createdAt"`  // 创建时间
	Address    string    `json:"address"`    // 企业链上地址
	Did        string    `json:"did"`        // 企业did
	Sender     string    `json:"sender"`     // 创建地址
}

// MessageType 消息类型
type MessageType int

const (
	MessageType_Enterprise_Create MessageType = 0 + iota // 企业注册
)

const (
	MessageType_EBL_CreateDraft  MessageType = 100 + iota // 创建电子提单草案 （涉及消息通知Msg1）
	MessageType_EBL_ApproveDraft                          // 进口企业审核同意电子提单草案 （涉及消息通知Msg3）
	MessageType_EBL_RejectDraft                           // 进口企业审核拒绝电子提单草案 （涉及消息通知Msg2）
	MessageType_EBL_Departure                             // 货物离港通知 （涉及消息通知Msg5）
	MessageType_EBL_Arrival                               // 货物到港通知 （涉及消息通知Msg8）
	MessageType_EBL_PickupReady                           // 待进口企业提货 （涉及消息通知Msg7）
)

const (
	MessageType_Audit_Create MessageType = 200 + iota // 创建审核通知
	MessageType_Audit_Update                          // 更新审核通知
)

const (
	MessageType_Logistics_Create MessageType = 300 + iota // 创建物流消息
)

const (
	MessageType_Voucher_Create  MessageType = 400 + iota // 创建支付凭证 （涉及消息通知Msg9）
	MessageType_Voucher_Approve                          // 支付凭证审核通过
	MessageType_Voucher_Reject                           // 支付凭证审核拒绝 （涉及消息通知Msg6）
)

// FileInfo 提单文件类上链信息
type FileInfo struct {
	ID         string      `json:"id"`         // 链上唯一 id，电子提单草案文件 id
	OriginHash string      `json:"originHash"` // 原文 hash 值
	MsgType    MessageType `json:"msgType"`    // 文件消息通知类型
	CreatedAt  time.Time   `json:"createdAt"`  // 创建时间
	Sender     string      `json:"sender"`     // 创建地址
}

// CallBackInfo 回调消息，告诉发起端跨链成功与否
type CallBackInfo struct {
	ID         string      `json:"id"`         // 链上唯一 id，文件 id
	OriginHash string      `json:"originHash"` // 文件原文 hash 值
	Msg        string      `json:"msg"`        // 如果失败，错误信息是啥
	Code       int         `json:"code"`       // 200000 表示成功，其他表示失败
	Sender     string      `json:"sender"`     // 创建地址
	MsgType    MessageType `json:"msgType"`    // 文件消息通知类型
}
