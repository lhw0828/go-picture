package constants

// 用户角色
const (
	RoleAdmin  = "admin"  // 管理员
	RoleUser   = "user"   // 普通用户
	RoleMember = "member" // 团队成员
	RoleOwner  = "owner"  // 团队所有者
)

// 空间类型
const (
	SpaceTypePersonal = "personal" // 个人空间
	SpaceTypeTeam    = "team"     // 团队空间
	SpaceTypePublic  = "public"   // 公共空间
)

// 审核状态
const (
	ReviewStatusPending = 0 // 待审核
	ReviewStatusPass   = 1 // 通过
	ReviewStatusReject = 2 // 拒绝
)

// 图片相关常量
const (
	MaxImageSize    = 10 * 1024 * 1024 // 最大图片大小 10MB
	MaxThumbnailSize = 200              // 缩略图最大边长
)