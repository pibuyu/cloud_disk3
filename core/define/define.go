package define

import "github.com/golang-jwt/jwt/v4"

// token加密盐值
const SECRET_KEY = "huhaifeng_key"

// redis连接配置
const REDIS_CONN_IP = "localhost:6379"
const REDIS_CONN_PWD = ""

// 验证码有效时间（秒）
const CODE_EXPIRE = 300

type UserClaim struct {
	Id       int64
	Identity string
	Name     string
	jwt.StandardClaims
}