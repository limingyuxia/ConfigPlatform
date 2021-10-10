package captcha

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	wraphh "github.com/turtlemonvh/gin-wraphh"
)

// @Tags 鉴权
// @Summary 获取或刷新图片验证码
// @Accept  json
// @Param RefeshCaptcha query model.RefeshCaptcha true "刷新验证码时的随机数"
// @Router /captcha/{{.captcha_id}}.png [get]
func RefeshCaptcha() gin.HandlerFunc {
	return wraphh.WrapHH(func(http.Handler) http.Handler {
		return captcha.Server(captcha.StdWidth, captcha.StdHeight)
	})
}
