package e

type ErrorCode struct {
	Code int
	Msg  string
}

/**
错误码组成：应用标识+错误类型+错误编码

错误码位数：5位
错误码示例：40102
使用规范：只增不改，避免混乱

应用标识: (1 位数字)
*  用户调用  U: 4
*  系统内部: S: 5
*  外部依赖: T: 6
*  中间件  : M: 7
错误类型: (1 位数字)
*  参数错误：P    : 1
*  业务错误：B    : 2
*  系统错误：S    : 3
*  网络错误：N    : 4
*  数据库错误：D  : 5
*  缓存错误：C    : 6
*  RPC错误：R     : 7
*  文件IO错误：F  : 8
*  其他错误：O    : 9
错误编码: (3位数字)
*  001自增
*/

var (
	SUCCESS = &ErrorCode{200, "成功"}

	FAIL = &ErrorCode{500, "系统错误"}

	/**
	用户错误
	*/

	//	用户调用	  	参数错误
	UserParamInvalid = &ErrorCode{41001, "参数错误"}

	//	用户调用	  	业务错误
	UserBusinessUserNotBind = &ErrorCode{42001, "用户未绑定"}

	UserBusinessUserPwdError = &ErrorCode{42002, "密码错误"}

	//	用户调用	  	系统错误
	UserSysUserNotExists = &ErrorCode{43001, "用户不存在"}

	//	用户调用	  	网络错误
	UserNetInvalid = &ErrorCode{44001, "网络错误"}

	//	用户调用	  	数据库错误
	UserDataInvalid = &ErrorCode{45001, "数据库错误"}

	//	用户调用	  	缓存错误
	UserCacheInvalid = &ErrorCode{46001, "缓存错误"}

	//	用户调用	  	RPC错误
	UserRpcInvalid = &ErrorCode{47001, "RPC错误"}

	//	用户调用	  	文件IO错误
	UserFileInvalid = &ErrorCode{48001, "文件IO错误"}

	//	用户调用	  	其他错误
	UserOtherInvalid = &ErrorCode{49001, "用户其他错误"}

	/**
	系统错误
	*/

	//	系统内部	  	参数错误
	SysParamInvalid = &ErrorCode{51001, "参数错误"}

	//	系统内部	  	业务错误
	SysBusinessUserNotBind = &ErrorCode{52001, "用户未绑定"}

	//	系统内部	  	系统错误
	SysUserNotExists = &ErrorCode{53001, "用户不存在"}

	//	系统内部	  	网络错误
	SysNetInvalid = &ErrorCode{54001, "网络错误"}

	//	系统内部	  	数据库错误
	SysDataInvalid = &ErrorCode{55001, "数据库错误"}

	//	系统内部	  	缓存错误
	SysCacheInvalid = &ErrorCode{56001, "缓存错误"}

	//	系统内部	  	RPC错误
	SysRpcInvalid = &ErrorCode{57001, "RPC错误"}

	//	系统内部	  	文件IO错误
	SysFileInvalid = &ErrorCode{58001, "文件IO错误"}

	//	系统内部	  	其他错误
	SysOtherInvalid = &ErrorCode{59001, "系统其他错误"}

	/**
	外部依赖
	*/

	//	外部依赖	  	参数错误
	ThirdParamInvalid = &ErrorCode{61001, "参数错误"}

	//	外部依赖	  	业务错误
	ThirdBusinessUserNotBind = &ErrorCode{62001, "用户未绑定"}

	//	外部依赖	  	系统错误
	ThirdUserNotExists = &ErrorCode{63001, "用户不存在"}

	//	外部依赖	  	网络错误
	ThirdNetInvalid = &ErrorCode{64001, "网络错误"}

	//	外部依赖	  	数据库错误
	ThirdDataInvalid = &ErrorCode{65001, "数据库错误"}

	//	外部依赖	  	缓存错误
	ThirdCacheInvalid = &ErrorCode{66001, "缓存错误"}

	//	外部依赖	  	RPC错误
	ThirdRpcInvalid = &ErrorCode{67001, "RPC错误"}

	//	外部依赖	  	文件IO错误
	ThirdFileInvalid = &ErrorCode{68001, "文件IO错误"}

	//	外部依赖	  	其他错误
	ThirdOtherInvalid = &ErrorCode{69001, "系统其他错误"}
)
