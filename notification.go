package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackz-jones/notification-contract-go/code"
	_const "github.com/jackz-jones/notification-contract-go/const"
	"github.com/jackz-jones/notification-contract-go/event"
	"github.com/jackz-jones/notification-contract-go/types"
	"github.com/jackz-jones/notification-contract-go/util"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type EblNotify struct{}

func (e *EblNotify) InitContract() protogo.Response {

	// 记录合约创建者
	errResp := util.StoreContractCreator()
	if errResp != nil {
		return *errResp
	}

	// 记录合约版本
	errResp = util.StoreCurrentVersion()
	if errResp != nil {
		return *errResp
	}

	return sdk.Success([]byte("Init success"))
}

func (e *EblNotify) UpgradeContract() protogo.Response {

	// 更新当前合约版本
	errResp := util.StoreCurrentVersion()
	if errResp != nil {
		return *errResp
	}

	return sdk.Success([]byte("Upgrade success"))
}

func (e *EblNotify) InvokeContract(method string) (resp protogo.Response) {
	args := sdk.Instance.GetArgs()
	if len(method) == 0 {
		return sdk.Error("method of param should not be empty")
	}

	defer func() {
		if resp.Status != sdk.OK {
			sdk.Instance.Warnf(resp.Message)
		}
	}()

	switch method {
	case _const.MethodNotifyEnterpriseInfo:
		needCrossChainBytes := args[_const.ParamNeedCrossChain]
		needCrossChain := util.BytesToBool(needCrossChainBytes)

		var ei types.EnterpriseInfo
		enterpriseInfoBytes := args[_const.ParamEnterpriseInfo]
		err := json.Unmarshal(enterpriseInfoBytes, &ei)
		if err != nil {
			return code.ErrorResp(code.ErrInternalJsonUnmarshalFailed,
				fmt.Sprintf(code.InternalJsonUnmarshalFailedStrModifier, string(enterpriseInfoBytes), err))
		}

		return e.NotifyEnterpriseInfo(ei, needCrossChain)

	case _const.MethodNotifyFileInfo:
		needCrossChainBytes := args[_const.ParamNeedCrossChain]
		needCrossChain := util.BytesToBool(needCrossChainBytes)

		var fi types.FileInfo
		fileInfoBytes := args[_const.ParamFileInfo]
		err := json.Unmarshal(fileInfoBytes, &fi)
		if err != nil {
			return code.ErrorResp(code.ErrInternalJsonUnmarshalFailed,
				fmt.Sprintf(code.InternalJsonUnmarshalFailedStrModifier, string(fileInfoBytes), err))
		}

		return e.NotifyFileInfo(fi, needCrossChain)

	case _const.MethodCallback:
		var ci types.CallBackInfo
		callbackInfoBytes := args[_const.ParamCallbackInfo]
		err := json.Unmarshal(callbackInfoBytes, &ci)
		if err != nil {
			return code.ErrorResp(code.ErrInternalJsonUnmarshalFailed,
				fmt.Sprintf(code.InternalJsonUnmarshalFailedStrModifier, string(callbackInfoBytes), err))
		}

		return e.Callback(ci)

	case _const.MethodQueryContractVersion:
		return e.QueryContractVersion()

	default:
		return code.ErrorResp(code.ErrInvalidMethod, fmt.Sprintf("unknown method %s", method))
	}
}

func (e *EblNotify) NotifyEnterpriseInfo(ei types.EnterpriseInfo, needCrossChain bool) protogo.Response {
	sender, err := sdk.Instance.Sender()
	if err != nil {
		return code.ErrorResp(code.ErrInternalGetSenderFailed, "")
	}

	// set sender
	ei.Sender = sender

	// store enterprise info
	errResp := util.StoreEnterpriseInfo(ei)
	if errResp != nil {
		return *errResp
	}

	// emit event
	err = event.EmitEnterpriseNotifiedEvent(ei, needCrossChain)
	if err != nil {
		return code.ErrorResp(code.ErrInternalEmitEvent, fmt.Sprintf(code.InternalEmitEventFailedStrModifier,
			event.EnterpriseNotifiedEvent, err))
	}

	return util.SuccessNormal()
}

func (e *EblNotify) NotifyFileInfo(fi types.FileInfo, needCrossChain bool) protogo.Response {
	sender, err := sdk.Instance.Sender()
	if err != nil {
		return code.ErrorResp(code.ErrInternalGetSenderFailed, "")
	}

	// set sender
	fi.Sender = sender

	// store file info
	errResp := util.StoreFileInfo(fi)
	if errResp != nil {
		return *errResp
	}

	// emit event
	err = event.EmitFileNotifiedEvent(fi, needCrossChain)
	if err != nil {
		return code.ErrorResp(code.ErrInternalEmitEvent, fmt.Sprintf(code.InternalEmitEventFailedStrModifier,
			event.FileNotifiedEvent, err))
	}

	return util.SuccessNormal()
}

func (e *EblNotify) Callback(cbi types.CallBackInfo) protogo.Response {
	sender, err := sdk.Instance.Sender()
	if err != nil {
		return code.ErrorResp(code.ErrInternalGetSenderFailed, "")
	}

	// set sender
	cbi.Sender = sender

	// store file info
	errResp := util.StoreCallBackInfo(cbi)
	if errResp != nil {
		return *errResp
	}

	// emit event
	err = event.EmitCallbackEvent(cbi)
	if err != nil {
		return code.ErrorResp(code.ErrInternalEmitEvent, fmt.Sprintf(code.InternalEmitEventFailedStrModifier,
			event.CallbackEvent, err))
	}

	return util.SuccessNormal()
}

func (e *EblNotify) QueryContractVersion() protogo.Response {
	verBytes, err := sdk.Instance.GetStateByte(_const.KeyNotifyContractVersion, "")
	if err != nil {
		return code.ErrorResp(code.ErrInternalDataReadFailed, fmt.Sprintf(code.InternalDataReadFailedStrModifier,
			_const.KeyNotifyContractVersion, "", err))
	}

	return sdk.Success(verBytes)
}

func main() {
	err := sandbox.Start(new(EblNotify))
	if err != nil {
		log.Fatal(err)
	}
}
