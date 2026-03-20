package types

import (
	"chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
)

// Notify 消息通知接口
type Notify interface {

	// NotifyEnterpriseInfo 企业身份类消息通知
	NotifyEnterpriseInfo(ei EnterpriseInfo, needCrossChain bool) protogo.Response

	// NotifyFileInfo 提单文件类消息通知
	NotifyFileInfo(fi FileInfo, needCrossChain bool) protogo.Response

	// Callback 跨链消息通知发出，不管成功还是失败需要告诉发起端，只会跨链服务调用
	Callback(cbi CallBackInfo) protogo.Response

	// QueryContractVersion 查询合约版本
	QueryContractVersion() protogo.Response
}
