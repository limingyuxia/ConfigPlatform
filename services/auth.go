package services

import (
	"log"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"

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
// @Param ConfirmCaptcha body model.ConfirmCaptcha true "校验图片验证码"
// @Success 200 {string} string
// @Failure 400 {object} WebResponse
// @Router /captcha/confirm [post]
func ConfirmCaptcha(c *gin.Context) {

	var confirmCaptcha model.ConfirmCaptcha
	if err := c.ShouldBindJSON(&confirmCaptcha); err == nil {

		if !captcha.VerifyString(confirmCaptcha.CaptchaId, confirmCaptcha.CaptchaSolution) {
			ResponseError(CAPTCHA_ERROR, RETCODE_MSG[CAPTCHA_ERROR], c)
		} else {
			ResponseData("验证成功", c)
		}

	} else {
		log.Print("ConfirmCapcha param error: ", err)
		ResponseError(PARAMS_ERROR, RETCODE_MSG[PARAMS_ERROR], c)
	}

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
