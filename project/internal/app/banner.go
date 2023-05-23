package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Добавить баннер в слот
// @Description
// @Tags     Banners
// @Produce  json
// @Router   /api/add-banner-to-slot/{banner_id}/{slot_id} [post]
// @Success  200 {string} "" "OK"
// @Failure  500 {object} ErrorStruct "StatusInternalServerError"
func (a *App) addBannerToSlot(ctx *gin.Context) {
	bannerId, err := strconv.ParseInt(ctx.Param("banner_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	slotId, err := strconv.ParseInt(ctx.Param("slot_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	err = a.rep.AddBannerToSlot(int32(bannerId), int32(slotId))
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка получения добавления баннера в слот")
		return
	}
	ctx.JSON(http.StatusOK, "")
}

// @Summary Удаленить баннер из слота
// @Description
// @Tags     Banners
// @Produce  json
// @Router   /api/remove-banner-from-slot/{banner_id}/{slot_id} [post]
// @Success  200 {string} "" "OK"
// @Failure  500 {object} ErrorStruct "StatusInternalServerError"
func (a *App) removeBannerFromSlot(ctx *gin.Context) {
	bannerId, err := strconv.ParseInt(ctx.Param("banner_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	slotId, err := strconv.ParseInt(ctx.Param("slot_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	err = a.rep.RemoveBannerFromSlot(int32(bannerId), int32(slotId))
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка удаления баннера из слота")
		return
	}
	ctx.JSON(http.StatusOK, "")
}

// @Summary Переход по беннеру
// @Description
// @Tags     Banners
// @Produce  json
// @Router   /api/click-on-banner/{banner_id}/{slot_id}/{soc_group_id} [post]
// @Success  200 {string} "" "OK"
// @Failure  500 {object} ErrorStruct "StatusInternalServerError"
func (a *App) clickOnBanner(ctx *gin.Context) {
	bannerId, err := strconv.ParseInt(ctx.Param("banner_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	slotId, err := strconv.ParseInt(ctx.Param("slot_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	socGroupId, err := strconv.ParseInt(ctx.Param("soc_group_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	stat, err := a.rep.ClickOnBanner(int32(bannerId), int32(slotId), int32(socGroupId))
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка записи статистики при переходе по беннеру")
		return
	}
	if err := a.queue.SendEvent(ctx, a.exchangeName, stat); err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка отправки сообщения в очередь")
	}
	ctx.JSON(http.StatusOK, "")
}

// @Summary Выбрать баннер для слота
// @Description
// @Tags     Banners
// @Produce  json
// @Router   /api/choose-banner-for-slot/{soc_group_id}/{slot_id} [post]
// @Success  200 {integer} 0 "Выбранный баннер"
// @Failure  500 {object} ErrorStruct "StatusInternalServerError"
func (a *App) chooseBannerForSlot(ctx *gin.Context) {
	socGroupId, err := strconv.ParseInt(ctx.Param("soc_group_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	slotId, err := strconv.ParseInt(ctx.Param("slot_id"), 10, 32)
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка сервера")
		return
	}
	bannerId, stat, err := a.rep.ChooseBannerForSlot(int32(socGroupId), int32(slotId))
	if err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка выбора баннера для слота")
		return
	}
	if err := a.queue.SendEvent(ctx, a.exchangeName, stat); err != nil {
		ServeError(ctx, http.StatusInternalServerError, "Ошибка отправки сообщения в очередь")
	}
	ctx.JSON(http.StatusOK, bannerId)
}
