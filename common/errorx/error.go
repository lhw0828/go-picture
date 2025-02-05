package errorx

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
	SpaceNotEnough             = 40300 // 空间容量不足
	SpaceNotExist              = 40301 // 空间不存在
	SpaceNameNotNull           = 40302 // 空间名称不能为空
	CreateSpaceFailed          = 40303 // 创建空间失败
	InvalidSpaceLevel          = 40304 // 无效的空间等级
	GetSpaceFailed             = 40305 // 获取空间失败
	OnlyCreateOneSpaceEachType = 40306 // 每个类型只能创建一个空间
	DeleteSpaceFailed          = 40307 // 删除空间失败
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
	UserNotExist       = 40500 // 用户不存在
	PasswordWrong      = 40501 // 密码错误
	UserExist          = 40502 // 用户已存在
	UserNotLogin       = 40503 // 用户未登录
	UserNotActive      = 40504 // 用户未激活
	RegisterFail       = 40505 // 注册失败
	LoginFail          = 40506 // 登录失败
	GenerateTokenFail  = 40507 // 生成token失败
	PasswordNotMatch   = 40508 // 密码不匹配
	PasswordIsNull     = 40509 // 密码为空
	PasswordTooShort   = 40510 // 密码太短
	PasswordTooLong    = 40511 // 密码太长
	PasswordNotSet     = 40512 // 密码未设置
	UserAccountNotNull = 40513 // 用户账号不能为空
	UserRoleNotNull    = 40514 // 用户角色不能为空
	CreateUserFailed   = 40515 // 创建用户失败

)

// 常用参数校正错误码 (406xx)
const (
	LengthLess4 = 40600 // 长度小于4
	LengthLess8 = 40601 // 长度小于8
)

const (
	DBError       = 1002 // 数据库错误
	InternalError = 1003 // 内部错误
	NotFoundError = 1004 // 资源不存在
	QuotaError    = 1005 // 配额不足
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
	SpaceNotEnoughMsg             = "空间容量不足"
	SpaceNotExistMsg              = "空间不存在"
	SpaceNameNotNullMsg           = "空间名称不能为空"
	CreateSpaceFailedMsg          = "创建空间失败"
	InvalidSpaceLevelMsg          = "无效的空间等级"
	GetSpaceFailedMsg             = "获取空间失败"
	OnlyCreateOneSpaceEachTypeMsg = "每个类型只能创建一个空间"
	DeleteSpaceFailedMsg          = "删除空间失败"

	// 团队相关错误信息
	NotTeamMemberMsg     = "不是团队成员"
	NotTeamOwnerMsg      = "不是团队所有者"
	AlreadyTeamMemberMsg = "已经是团队成员"
	AddTeamMemberFailMsg = "添加团队成员失败"

	// 用户相关错误信息
	UserNotExistMsg       = "用户不存在"
	PasswordWrongMsg      = "密码错误"
	UserExistMsg          = "用户已存在"
	UserNotLoginMsg       = "用户未登录"
	UserNotActiveMsg      = "用户未激活"
	RegisterFailMsg       = "注册失败"
	LoginFailMsg          = "登录失败"
	GenerateTokenFailMsg  = "生成token失败"
	PasswordNotMatchMsg   = "密码不匹配"
	PasswordIsNullMsg     = "密码为空"
	PasswordTooShortMsg   = "密码太短"
	PasswordTooLongMsg    = "密码太长"
	PasswordNotSetMsg     = "密码未设置"
	UserAccountNotNullMsg = "用户账号不能为空"
	UserRoleNotNullMsg    = "用户角色不能为空"
	CreateUserFailedMsg   = "创建用户失败"

	// 常用参数校正错误信息
	LengthLess4Msg = "长度小于4"
	LengthLess8Msg = "长度小于8"
)

// 错误码和消息映射
var errorMap = map[int]string{
	// 系统级别错误
	SystemErr:        SystemErrMsg,
	ParamError:       ParamErrorMsg,
	InvalidTimeRange: InvalidTimeRangeMsg,

	// 认证授权相关错误
	UnauthorizedErr: UnauthorizedErrMsg,
	ForbiddenErr:    ForbiddenErrMsg,

	// 资源相关错误
	NotFoundErr: NotFoundErrMsg,

	// 空间相关错误
	SpaceNotEnough:    SpaceNotEnoughMsg,
	SpaceNotExist:     SpaceNotExistMsg,
	SpaceNameNotNull:  SpaceNameNotNullMsg,
	CreateSpaceFailed: CreateSpaceFailedMsg,
	InvalidSpaceLevel: InvalidSpaceLevelMsg,
	GetSpaceFailed:    GetSpaceFailedMsg,

	// 团队相关错误
	NotTeamMember:     NotTeamMemberMsg,
	NotTeamOwner:      NotTeamOwnerMsg,
	AlreadyTeamMember: AlreadyTeamMemberMsg,
	AddTeamMemberFail: AddTeamMemberFailMsg,

	// 用户相关错误
	UserNotExist:      UserNotExistMsg,
	PasswordWrong:     PasswordWrongMsg,
	UserExist:         UserExistMsg,
	UserNotLogin:      UserNotLoginMsg,
	UserNotActive:     UserNotActiveMsg,
	RegisterFail:      RegisterFailMsg,
	LoginFail:         LoginFailMsg,
	GenerateTokenFail: GenerateTokenFailMsg,
	PasswordNotMatch:  PasswordNotMatchMsg,

	// 参数校验错误
	LengthLess4: LengthLess4Msg,
	LengthLess8: LengthLess8Msg,
}

// 获取错误信息
func GetMessage(code int) string {
	if msg, ok := errorMap[code]; ok {
		return msg
	}
	return SystemErrMsg
}

// 创建错误
func NewError(code int) error {
	return &CodeError{
		Code:    code,
		Message: GetMessage(code),
	}
}

// 创建带自定义消息的错误
func NewErrorWithMsg(code int, msg string) error {
	return &CodeError{
		Code:    code,
		Message: msg,
	}
}

// 错误结构体
type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 修改 Error 方法的实现
func (e *CodeError) Error() string {
	return e.Message
}

func NewCodeError(code int, message string) error {
	return &CodeError{
		Code:    code,
		Message: message,
	}
}

func NewSystemError(message string) error {
	return NewCodeError(SystemErr, message)
}

func NewParamError(message string) error {
	return NewCodeError(ParamError, message)
}
