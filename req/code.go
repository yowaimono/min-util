package req


// 定义业务相关的枚举
type ErrorCode int

const (
	// 通用错误
	ErrBadRequest ErrorCode = iota + 400
	ErrUnauthorized
	ErrForbidden
	ErrNotFound
	ErrInternalServerError

	// 用户相关错误 (1001-1099)
	ErrUserNotExist ErrorCode = iota + 1000
	ErrUserAlreadyExists
	ErrInvalidCredentials
	ErrUserLocked

	// 权限相关错误 (1101-1199)
	ErrPermissionDenied ErrorCode = iota + 1100
	ErrTokenExpired
	ErrTokenInvalid

	// 支付相关错误 (1201-1299)
	ErrPayError ErrorCode = iota + 1200
	ErrPaymentFailed
	ErrPaymentTimeout

	// 会员相关错误 (1301-1399)
	ErrMembershipExpired ErrorCode = iota + 1300
	ErrMembershipNotActive
	ErrMembershipUpgradeFailed

	// 注册登录相关错误 (1401-1499)
	ErrRegistrationFailed ErrorCode = iota + 1400
	ErrLoginFailed
	ErrAccountDisabled

	// 其他业务相关错误 (1501-1599)
	ErrInvalidInput ErrorCode = iota + 1500
	ErrResourceAlreadyExists
	ErrResourceNotFound
	ErrOperationFailed
)

// 定义枚举对应的错误信息
var errorMessages = map[ErrorCode]string{
	// 通用错误
	ErrBadRequest:          "Bad Request",
	ErrUnauthorized:        "Unauthorized",
	ErrForbidden:           "Forbidden",
	ErrNotFound:            "Not Found",
	ErrInternalServerError: "Internal Server Error",

	// 用户相关错误
	ErrUserNotExist:        "User not exist!",
	ErrUserAlreadyExists:   "User already exists!",
	ErrInvalidCredentials:  "Invalid credentials!",
	ErrUserLocked:          "User is locked!",

	// 权限相关错误
	ErrPermissionDenied:    "Permission denied!",
	ErrTokenExpired:        "Token expired!",
	ErrTokenInvalid:        "Token invalid!",

	// 支付相关错误
	ErrPayError:            "Pay error! Please try again.",
	ErrPaymentFailed:       "Payment failed!",
	ErrPaymentTimeout:      "Payment timeout!",

	// 会员相关错误
	ErrMembershipExpired:   "Membership expired!",
	ErrMembershipNotActive: "Membership not active!",
	ErrMembershipUpgradeFailed: "Membership upgrade failed!",

	// 注册登录相关错误
	ErrRegistrationFailed:  "Registration failed!",
	ErrLoginFailed:         "Login failed!",
	ErrAccountDisabled:     "Account disabled!",

	// 其他业务相关错误
	ErrInvalidInput:        "Invalid input!",
	ErrResourceAlreadyExists: "Resource already exists!",
	ErrResourceNotFound:    "Resource not found!",
	ErrOperationFailed:     "Operation failed!",
}
