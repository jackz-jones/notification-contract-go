package event

import (
	"encoding/json"

	"github.com/jackz-jones/notification-contract-go/types"
	"github.com/jackz-jones/notification-contract-go/util"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

const (
	EnterpriseNotifiedEvent = "EnterpriseNotifiedEvent" // 企业身份通知事件
	FileNotifiedEvent       = "FileNotifiedEvent"       // 提单、运输、支付凭证类的文件通知事件
	CallbackEvent           = "CallbackEvent"           // 回调通知事件
)

// EmitEnterpriseNotifiedEvent 企业身份创建事件
func EmitEnterpriseNotifiedEvent(ei types.EnterpriseInfo, needCrossChain bool) error {
	enterpriseInfoBytes, err := json.Marshal(ei)
	if err != nil {
		return err
	}

	sdk.Instance.EmitEvent(EnterpriseNotifiedEvent, []string{string(enterpriseInfoBytes), string(util.BoolToBytes(needCrossChain))})
	return nil
}

// EmitFileNotifiedEvent 提单、运输、支付凭证类的文件通知事件
func EmitFileNotifiedEvent(fi types.FileInfo, needCrossChain bool) error {
	fileInfoBytes, err := json.Marshal(fi)
	if err != nil {
		return err
	}

	sdk.Instance.EmitEvent(FileNotifiedEvent, []string{string(fileInfoBytes), string(util.BoolToBytes(needCrossChain))})
	return nil
}

// EmitCallbackEvent 回调通知事件
func EmitCallbackEvent(ci types.CallBackInfo) error {
	callbackInfoBytes, err := json.Marshal(ci)
	if err != nil {
		return err
	}

	sdk.Instance.EmitEvent(CallbackEvent, []string{string(callbackInfoBytes)})
	return nil
}
