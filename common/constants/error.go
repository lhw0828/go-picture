package constants

// 系统级别错误码
const (
	SystemErr        = 50000 // 系统错误
	ParamError       = 40000 // 参数错误
	InvalidTimeRange = 40001 // 无效的时间范围
)

// 认证授权相关错误码 (401xx)
const (
	UnauthorizedErr = 40100 // 未授权
	ForbiddenErr    = 40101 // 无权限
)

// 资源相关错误码 (402xx)
const (
	NotFoundErr = 40200 // 资源不存在
)

// 空间相关错误码 (403xx)
const (
	SpaceNotEnough    = 40300 // 空间容量不足
	SpaceNotExist     = 40301 // 空间不存在
	SpaceNameNotNull  = 40302 // 空间名称不能为空
	CreateSpaceFailed = 40303 // 创建空间失败
	InvalidSpaceLevel = 40304 // 无效的空间等级
	GetSpaceFailed    = 40305 // 获取空间失败
)

// 团队相关错误码 (404xx)
const (
	NotTeamMember     = 40400 // 不是团队成员
	NotTeamOwner      = 40401 // 不是团队所有者
	AlreadyTeamMember = 40402 // 已经是团队成员
	AddTeamMemberFail = 40403 // 添加团队成员失败
)

// 用户相关错误码 (405xx)
const (
	UserNotExist      = 40500 // 用户不存在
	PasswordWrong     = 40501 // 密码错误
	UserExist         = 40502 // 用户已存在
	UserNotLogin      = 40503 // 用户未登录
	UserNotActive     = 40504 // 用户未激活
	RegisterFail      = 40505 // 注册失败
	LoginFail         = 40506 // 登录失败
	GenerateTokenFail = 40507 // 生成token失败
	PasswordNotMatch  = 40508 // 密码不匹配
)

// 常用参数校正错误码 (406xx)
const (
	LengthLess4 = 40600 // 长度小于4
	LengthLess8 = 40601 // 长度小于8
)

// 错误信息
const (
	// 系统级别错误信息
	SuccessMsg          = "success"
	SystemErrMsg        = "系统错误"
	ParamErrorMsg       = "参数错误"
	InvalidTimeRangeMsg = "无效的时间范围"

	// 认证授权相关错误信息
	UnauthorizedErrMsg = "未授权"
	ForbiddenErrMsg    = "无权限"

	// 资源相关错误信息
	NotFoundErrMsg = "资源不存在"

	// 空间相关错误信息
	SpaceNotEnoughMsg    = "空间容量不足"
	SpaceNotExistMsg     = "空间不存在"
	SpaceNameNotNullMsg  = "空间名称不能为空"
	CreateSpaceFailedMsg = "创建空间失败"
	InvalidSpaceLevelMsg = "无效的空间等级"
	GetSpaceFailedMsg    = "获取空间失败"

	// 团队相关错误信息
	NotTeamMemberMsg     = "不是团队成员"
	NotTeamOwnerMsg      = "不是团队所有者"
	AlreadyTeamMemberMsg = "已经是团队成员"
	AddTeamMemberFailMsg = "添加团队成员失败"

	// 用户相关错误信息
	UserNotExistMsg      = "用户不存在"
	PasswordWrongMsg     = "密码错误"
	UserExistMsg         = "用户已存在"
	UserNotLoginMsg      = "用户未登录"
	UserNotActiveMsg     = "用户未激活"
	RegisterFailMsg      = "注册失败"
	LoginFailMsg         = "登录失败"
	GenerateTokenFailMsg = "生成token失败"
	PasswordNotMatchMsg  = "密码不匹配"

	// 常用参数校正错误信息
	LengthLess4Msg = "长度小于4"
	LengthLess8Msg = "长度小于8"
)
