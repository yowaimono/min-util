package validation

import (
	"github.com/dlclark/regexp2"
)

// IsValidIDCard 验证身份证号码
func IsValidIDCard(idCard string) bool {
	reg := regexp2.MustCompile(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`, 0)
	match, _ := reg.MatchString(idCard)
	return match
}

// IsValidMobile 验证手机号码
func IsValidMobile(mobile string) bool {
	reg := regexp2.MustCompile(`^1[3456789]\d{9}$`, 0)
	match, _ := reg.MatchString(mobile)
	return match
}

// IsValidChineseName 验证中文人名
func IsValidChineseName(name string) bool {
	reg := regexp2.MustCompile(`^[\p{Han}]{2,10}$`, 0)
	match, _ := reg.MatchString(name)
	return match
}

// IsValidPassword 验证密码强度
func IsValidPassword(password string) bool {
	reg := regexp2.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`, 0)
	match, _ := reg.MatchString(password)
	return match
}

// IsValidAccount 验证账号
func IsValidAccount(account string) bool {
	reg := regexp2.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{4,15}$`, 0)
	match, _ := reg.MatchString(account)
	return match
}

// IsValidEmail 验证邮箱
func IsValidEmail(email string) bool {
	reg := regexp2.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, 0)
	match, _ := reg.MatchString(email)
	return match
}

// IsValidURL 验证URL
func IsValidURL(url string) bool {
	reg := regexp2.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`, 0)
	match, _ := reg.MatchString(url)
	return match
}

// IsValidIP 验证IP地址
func IsValidIP(ip string) bool {
	reg := regexp2.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`, 0)
	match, _ := reg.MatchString(ip)
	return match
}

// IsValidIPv4 验证IPv4地址
func IsValidIPv4(ip string) bool {
	reg := regexp2.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`, 0)
	match, _ := reg.MatchString(ip)
	return match
}

// IsValidIPv6 验证IPv6地址
func IsValidIPv6(ip string) bool {
	reg := regexp2.MustCompile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`, 0)
	match, _ := reg.MatchString(ip)
	return match
}

// IsValidPostalCode 验证邮政编码
func IsValidPostalCode(postalCode string) bool {
	reg := regexp2.MustCompile(`^[1-9]\d{5}$`, 0)
	match, _ := reg.MatchString(postalCode)
	return match
}

// IsValidDate 验证日期格式（YYYY-MM-DD）
func IsValidDate(date string) bool {
	reg := regexp2.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`, 0)
	match, _ := reg.MatchString(date)
	return match
}


// IsValidTime 验证时间格式（HH:MM:SS）
func IsValidTime(timeStr string) bool {
	reg := regexp2.MustCompile(`^([01]\d|2[0-3]):([0-5]\d):([0-5]\d)$`, 0)
	match, _ := reg.MatchString(timeStr)
	return match
}

// IsValidCreditCard 验证信用卡号
func IsValidCreditCard(cardNumber string) bool {
	reg := regexp2.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|7[0-9]{15}|(2131|1800|35\d{3})\d{11})$`, 0)
	match, _ := reg.MatchString(cardNumber)
	return match
}

// IsValidUsername 验证用户名
func IsValidUsername(username string) bool {
	reg := regexp2.MustCompile(`^[a-zA-Z0-9_]{4,20}$`, 0)
	match, _ := reg.MatchString(username)
	return match
}

// IsPositiveInteger 验证正整数
func IsPositiveInteger(num string) bool {
	reg := regexp2.MustCompile(`^[1-9]\d*$`, 0)
	match, _ := reg.MatchString(num)
	return match
}