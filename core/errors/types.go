// 公共错误库
package errors

// 0-199	预留
// 200-299  参数错误
// 300-399	类型错误
// 400-499  中间件问题
// 500-599  内部错误
// 600-699  操作系统错误
// 700-699  协议错误

var (
	// ErrUnsupportedContentType indicates unacceptable or lack of Content-Type
	ErrUnsupportedContentType = New("unsupported content type", 300)

	// ErrInvalidQueryParams indicates invalid query parameters
	ErrInvalidQueryParams = New("invalid query parameters", 200)

	// ErrNotFoundParam indicates that the parameter was not found in the query
	ErrNotFoundParam = New("parameter not found in the query", 201)

	// ErrMalformedEntity indicates a malformed entity specification
	ErrMalformedEntity = New("malformed entity specification", 500)
	ErrCompressData    = New("数据压缩错误", 501)
)

var (
	ErrRedisConnection    = New("redis连接错误", 400)
	ErrRedisConfig        = New("Redis配置错误", 401)
	ErrDatabaseConnect    = New("数据库连接错误", 403)
	ErrDatabaseSessionNil = New("数据库会话为空", 405)
	ErrDatabaseCommit     = New("数据库提交错误", 406)
	ErrDataBindError      = New("数据绑定错误", 407)
	ErrDataMarshal        = New("数据序列化错误", 408)
	ErrDataUnmarshal      = New("数据反序列化错误", 409)
)

var (
	ErrAuthConnection = New("Auth Grpc 连接错误", 404)
	ErrCloseTracer    = New("关闭Tracer错误", 407)
	ErrTimeParse      = New("时间解析错误", 408)
)

var (
	ErrorSystemBusy      = New("系统繁忙,请稍后再试!", 10000)
	ErrorParam           = New("参数错误", 10001)
	ErrorCode            = New("验证码错误", 10002)
	ErrorUserCreate      = New("用户创建失败", 10003)
	ErrorUserNotFound    = New("用户不存在", 10004)
	ErrorPassword        = New("密码错误", 10005)
	ErrorRole            = New("角色获取失败", 10006)
	ErrorValue           = New("该值不可用，系统中已存在!", 10007)
	ErrorRoleEdit        = New("修改角色信息失败", 10008)
	ErrorRoleDel         = New("删除角色信息失败", 10009)
	ErrorRoleDelAuth     = New("删除角色权限失败", 10010)
	ErrorRoleBindUser    = New("角色绑定用户失败", 10011)
	ErrorRoleAuthMenu    = New("角色授权菜单权限失败", 10012)
	ErrorUserNameExist   = New("用户名已存在", 10013)
	ErrorUserEdit        = New("用户信息修改失败", 10014)
	ErrorUserDel         = New("删除用户失败", 10015)
	ErrorUserDelRole     = New("删除用户绑定角色信息失败", 10016)
	ErrorMenuAdd         = New("添加菜单失败", 10017)
	ErrorMenuGet         = New("菜单信息获取失败", 10018)
	ErrorMenuEdit        = New("编辑菜单失败", 10019)
	ErrorMenuDel         = New("删除菜单失败", 10020)
	ErrorRoleCreate      = New("创建角色失败", 10021)
	ErrorRoleBind        = New("用户已经绑定角色", 10022)
	ErrorOperate         = New("操作失败,请稍等后再试!", 10023)
	ErrorMenuParent      = New("上级菜单不可以是自己!", 10024)
	ErrorOldPassword     = New("旧密码验证失败", 10025)
	ErrorOldPassword2    = New("旧密码错误", 10026)
	ErrorAgentAdd        = New("添加代理失败", 10027)
	ErrorAgentDel        = New("该用户不是代理", 10028)
	ErrorAgentGet        = New("代理获取失败", 10029)
	ErrorAccountMoney    = New("账户余额不足", 10030)
	ErrorWithdrawGet     = New("提现信息获取失败", 10031)
	ErrorWithdrawCheck   = New("请勿重复审核", 10032)
	ErrorAccountFreeze   = New("账户已冻结", 10033)
	ErrorAccountDel      = New("账户已删除", 10034)
	ErrorWithdrawCheck2  = New("请勿重复审核", 10035)
	ErrorAccountPayType  = New("请先申请收款方式", 10036)
	ErrorDataNotFound    = New("数据没找到", 10037)
	ErrorDataValidate    = New("数据验证错误", 10038)
	ErrorAccountRegister = New("账号注册失败", 10039)
	ErrorAccountType     = New("账号类型错误", 10040)
	ErrorAccountNotfound = New("账号不存在", 10041)
	ErrorStatus          = New("状态错误", 10042)
	ErrorNotThatAccount  = New("不是该账户数据", 10043)
)
