package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"

	"ConfigPlatform/conf/redis"
	"ConfigPlatform/model"
)

// @Tags 鉴权
// @Summary 获取图片验证码id
// @Success 200 {object} model.GetCaptchaResp
// @Router /captcha/get [get]
func GetCaptcha(c *gin.Context) {

	captchaId := model.GetCaptchaResp{
		CaptchaId: captcha.New(),
	}

	ResponseData(captchaId, c)
}

// @Tags 鉴权
// @Summary 校验图片验证码
// @Accept  json
// @Produce  json
// @Param ConfirmCaptcha body model.ConfirmCaptcha true "图片验证码的id和值"
// @Success 200 {object} model.CaptchaToken
// @Failure 400 {object} WebResponse
// @Router /captcha/confirm [post]
func ConfirmCaptcha(c *gin.Context) {

	var confirmCaptcha model.ConfirmCaptcha
	if err := c.ShouldBindJSON(&confirmCaptcha); err == nil {

		if !captcha.VerifyString(confirmCaptcha.CaptchaId, confirmCaptcha.CaptchaSolution) {
			ResponseError(CAPTCHA_ERROR, RETCODE_MSG[CAPTCHA_ERROR], c)
		} else {

			// 生成随机长度的token
			RandStr := func(length int) []byte {
				str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
				bytes := []byte(str)
				result := []byte{}
				rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
				for i := 0; i < length; i++ {
					result = append(result, bytes[rand.Intn(len(bytes))])
				}
				return result
			}

			// 对生成的token进行md5加密
			md5Token := md5.New()
			md5Token.Write(RandStr(8))
			captchaToken := hex.EncodeToString(md5Token.Sum(nil))

			// 存token，设置有效期为2分钟
			err = redis.RedisConn.Set(captchaToken, "", 2*time.Minute).Err()
			if err != nil {
				log.Print("set the captcha token failed: ", err)
				ResponseError(REDIS_ERROR, RETCODE_MSG[REDIS_ERROR], c)
				return
			}

			ResponseData(&model.CaptchaToken{
				CaptchaToken: captchaToken,
			}, c)
		}

	} else {
		log.Print("ConfirmCapcha param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}

}

// @Tags 鉴权
// @Summary 登录
// @Accept  json
// @Produce  json
// @Param Login body model.LoginReq true "登录的账号和密码"
// @Success 200 {object} model.LoginResp
// @Failure 400 {object} model.AuthError
// @Router /login [post]
func Login(c *gin.Context) {

}

// @Tags 鉴权
// @Summary 登出
// @Success 200 {string} string
// @Router /logout [get]
func Logout(c *gin.Context) {

}

// @Tags 鉴权
// @Summary 刷新token
// @Success 200 {object} model.LoginResp
// @Failure 400 {object} model.AuthError
// @Router /auth/refreshToken [get]
func RefeshToken(c *gin.Context) {

}

// @Tags 鉴权
// @Summary 发送邮箱验证码
// @Accept  json
// @Produce  json
// @Param SendEmailCode body model.SendEmailCode true "图片验证码返回的token和邮件地址"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /email/send [post]
func SendEmailCode(c *gin.Context) {
	var captchaToken model.SendEmailCode

	if err := c.ShouldBindJSON(&captchaToken); err == nil {

		// 校验图片验证码的token
		_, err := redis.RedisConn.Get(captchaToken.CaptchaToken).Result()
		if err != nil {
			ResponseError(CAPTCHA_TOKEN_ERROR, RETCODE_MSG[CAPTCHA_TOKEN_ERROR], c)
			return
		}

		// 获取6位随机验证码
		random := func() string {
			return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		}

		// 设置邮箱验证码有效期为2分钟
		emailCode := random()
		err = redis.RedisConn.Set(emailCode, captchaToken.CaptchaToken, 2*time.Minute).Err()
		if err != nil {
			log.Print("set the email code failed: ", err)
			ResponseError(REDIS_ERROR, RETCODE_MSG[REDIS_ERROR], c)
			return
		}

		// 设置收件人和发件人的信息
		e := &email.Email{
			To:      []string{captchaToken.EmailAddress},
			From:    "ConfigPlatform <2572597150@qq.com>",
			Subject: "注册验证码",
			HTML:    []byte("<h1>您的验证码为 " + emailCode + "</h1>"),
			Headers: textproto.MIMEHeader{},
		}

		// 获取发送邮件使用的服务器授权码
		secret, err := ioutil.ReadFile("./conf/email/email.conf")
		if err != nil {
			log.Print("read send email secret faile! ", err)
			ResponseError(READ_FAILED_ERROR, RETCODE_MSG[READ_FAILED_ERROR], c)
			return
		}

		// 发送邮件
		err = e.Send("smtp.qq.com:587", smtp.PlainAuth("", "2572597150@qq.com", string(secret), "smtp.qq.com"))
		if err != nil {
			log.Print("send email code failed: ", err)
			ResponseError(READ_FAILED_ERROR, RETCODE_MSG[READ_FAILED_ERROR], c)
			return
		}

		ResponseData("发送成功，请到邮箱查看。验证码2分钟有效", c)

	} else {
		log.Print("SendEmailCode param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}

// @Tags 鉴权
// @Summary 验证邮箱验证码
// @Accept  json
// @Produce  json
// @Param ConfirmEmailCode body model.ConfirmEmail true "邮箱验证码和图片验证码的token"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /email/confirm [post]
func ConfirmEmailCode(c *gin.Context) {
	var confirmEmail model.ConfirmEmail

	if err := c.ShouldBindJSON(&confirmEmail); err == nil {

		// 校验邮箱验证码
		captchaToken, err := redis.RedisConn.Get(confirmEmail.EmailCode).Result()
		if err != nil {
			log.Print("get email code error: ", err)
			ResponseError(EMAIL_CODE_ERROR, RETCODE_MSG[EMAIL_CODE_ERROR], c)
			return
		}

		// 验证邮箱验证码是不是和图片验证码匹配
		if captchaToken != confirmEmail.CaptchaToken {
			ResponseError(EMAIL_CODE_ERROR, RETCODE_MSG[EMAIL_CODE_ERROR], c)
			return
		}

		ResponseData("验证成功", c)
	} else {
		log.Print("ConfirmEmailCode param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}
}
