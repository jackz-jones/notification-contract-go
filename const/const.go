package _const

// 一些参数名称定义
const (
	ParamNeedCrossChain = "needCrossChain" // 是否需要跨链
	ParamEnterpriseInfo = "enterpriseInfo" // 企业参数
	ParamFileInfo       = "fileInfo"       // 文件参数
	ParamCallbackInfo   = "callbackInfo"   // 回调参数
)

// 一些写状态的 key 定义
const (
	KeyNotifyContractCreator = "NotifyContractCreator" // 合约创建者
	KeyNotifyContractVersion = "NotifyContractVersion" // 合约版本
	KeyEnterpriseInfo        = "EnterpriseInfo"        // 企业信息
	KeyFileInfo              = "FileInfo"              // 文件信息
	KeyCallBackInfo          = "CallBackInfo"          // 回调信息
)

// 一些合约方法名定义
const (
	MethodNotifyEnterpriseInfo = "NotifyEnterpriseInfo" // 企业信息通知
	MethodNotifyFileInfo       = "NotifyFileInfo"       // 文件信息通知
	MethodCallback             = "Callback"             // 回调通知
	MethodQueryContractVersion = "QueryContractVersion" // 查询合约版本
)
