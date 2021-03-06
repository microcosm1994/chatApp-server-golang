package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"chatAppServer/utils"
)

/*TokenController Token控制器 */
type TokenController struct {
	utils.Redis
}

/*Config Token配置项 */
type Config struct {
	PrivateKey    []byte
	PrivateKeyID  string
	Subject       string
	Scopes        []string
	TokenURL      string
	Expires       time.Duration
	Audience      string
	PrivateClaims map[string]interface{}
	UseIDToken    bool
}

/*SigningMethod 签名 */
type SigningMethod interface {
    Verify(signingString, signature string, key interface{}) error 
    Sign(signingString string, key interface{}) (string, error)    
    Alg() string
}

/*Claims 确定令牌是否由于任何受支持的原因而无效 */
type Claims interface {
    Valid() error
}

/*Token JWT Token */
type Token struct {
	Raw       string                 
	Method    SigningMethod          
	Header    map[string]interface{} 
	Claims    Claims                 
	Signature string                 
	Valid     bool                  
}

/*CreateToken 生成token*/
func (c *TokenController) CreateToken(phone string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
		"nbf": time.Now().Add(2*time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("11111111"))
	// 保存token到redis
	c.Redis.Set(tokenString, phone, 60 * 60 * 24)
	return tokenString
}


