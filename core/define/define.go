package define

import (
	"github.com/golang-jwt/jwt/v4"
)

// token加密盐值
const SECRET_KEY = "huhaifeng_key"

// redis连接配置
const REDIS_CONN_IP = "localhost:6379"
const REDIS_CONN_PWD = ""

// 验证码有效时间（秒）
const CODE_EXPIRE = 300

var TencentSecretID = "AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF"
var TencentSecretKey = "wooiR3meTGwiV9g5Qix6Tn3EbvLi0Kcl"
var TencentCloudURL = "https://hhf-1317635862.cos.ap-guangzhou.myqcloud.com"

// 分页默认参数
var PAGE_SIZE = 10
var PAGE_NUM = 1

type UserClaim struct {
	Id       int64
	Identity string
	Name     string
	jwt.StandardClaims
}
