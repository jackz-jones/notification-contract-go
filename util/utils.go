package util

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jackz-jones/notification-contract-go/code"
	_const "github.com/jackz-jones/notification-contract-go/const"
	"github.com/jackz-jones/notification-contract-go/types"
	"github.com/jackz-jones/notification-contract-go/version"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

func SuccessNormal() protogo.Response {
	return sdk.Success([]byte("ok"))
}

func BoolToBytes(b bool) []byte {
	return []byte(strconv.FormatBool(b))
}

func BytesToBool(bs []byte) bool {
	res := false
	if len(bs) > 0 && string(bs) == "true" {
		res = true
	}

	return res
}

func StoreCurrentVersion() *protogo.Response {

	// 获取当前版本信息
	ver := version.VerInfo{Version: version.Version, CommitID: version.CommitID, BuildTime: version.BuildTime}
	verBytes, err := json.Marshal(&ver)
	if err != nil {
		errResp := code.ErrorResp(code.ErrInternalJsonMarshalFailed,
			fmt.Sprintf(code.InternalJsonMarshalFailedStrModifier, ver, err))
		return &errResp
	}

	// 记录合约版本
	if err = sdk.Instance.PutStateByte(_const.KeyNotifyContractVersion, "", verBytes); err != nil {
		errResp := code.ErrorResp(code.ErrInternalDataStoreFailed, fmt.Sprintf(code.InternalDataStoreFailedStrModifier,
			_const.KeyNotifyContractVersion, "", err))
		return &errResp
	}

	return nil
}

func StoreContractCreator() *protogo.Response {

	// 获取合约创建者
	origin, err := sdk.Instance.Origin()
	if err != nil {
		errFResp := code.ErrorResp(code.ErrGetOriginFailed, "")
		return &errFResp
	}

	// 记录合约创建者
	if err = sdk.Instance.PutStateByte(_const.KeyNotifyContractCreator, "", []byte(origin)); err != nil {
		errFResp := code.ErrorResp(code.ErrInternalDataStoreFailed, fmt.Sprintf(code.InternalDataStoreFailedStrModifier,
			_const.KeyNotifyContractCreator, "", err))
		return &errFResp
	}

	return nil
}

func StoreEnterpriseInfo(ei types.EnterpriseInfo) *protogo.Response {

	// marshal
	eiBytes, err := json.Marshal(ei)
	if err != nil {
		errFResp := code.ErrorResp(code.ErrInternalJsonMarshalFailed,
			fmt.Sprintf(code.InternalJsonMarshalFailedStrModifier, ei, err))
		return &errFResp
	}

	// put state
	err = sdk.Instance.PutStateByte(_const.KeyEnterpriseInfo, ei.ID, eiBytes)
	if err != nil {
		errFResp := code.ErrorResp(code.ErrInternalDataStoreFailed, fmt.Sprintf(code.InternalDataStoreFailedStrModifier,
			_const.KeyEnterpriseInfo, ei.ID, err))
		return &errFResp
	}

	return nil
}

func StoreFileInfo(fi types.FileInfo) *protogo.Response {

	// marshal
	fiBytes, err := json.Marshal(fi)
	if err != nil {
		errFResp := code.ErrorResp(code.ErrInternalJsonMarshalFailed,
			fmt.Sprintf(code.InternalJsonMarshalFailedStrModifier, fi, err))
		return &errFResp
	}

	// put state
	err = sdk.Instance.PutStateByte(_const.KeyFileInfo, fi.ID, fiBytes)
	if err != nil {
		errFResp := code.ErrorResp(code.ErrInternalDataStoreFailed, fmt.Sprintf(code.InternalDataStoreFailedStrModifier,
			_const.KeyFileInfo, fi.ID, err))
		return &errFResp
	}

	return nil
}

func StoreCallBackInfo(cbi types.CallBackInfo) *protogo.Response {

	// marshal
	cbiBytes, err := json.Marshal(cbi)
	if err != nil {
		errFResp := code.ErrorResp(code.ErrInternalJsonMarshalFailed,
			fmt.Sprintf(code.InternalJsonMarshalFailedStrModifier, cbi, err))
		return &errFResp
	}

	// put state
	err = sdk.Instance.PutStateByte(_const.KeyCallBackInfo, cbi.ID, cbiBytes)
	if err != nil {
		errFResp := code.ErrorResp(code.ErrInternalDataStoreFailed, fmt.Sprintf(code.InternalDataStoreFailedStrModifier,
			_const.KeyCallBackInfo, cbi.ID, err))
		return &errFResp
	}

	return nil
}
