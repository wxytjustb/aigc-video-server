package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppUsersApi struct{}

// CreateAppUsers 创建appUsers表
// @Tags AppUsers
// @Summary 创建appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.AppUsers true "创建appUsers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /appUsers/createAppUsers [post]
func (appUsersApi *AppUsersApi) CreateAppUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var appUsers app.AppUsers
	err := c.ShouldBindJSON(&appUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appUsersService.CreateAppUsers(ctx, &appUsers)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAppUsers 删除appUsers表
// @Tags AppUsers
// @Summary 删除appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.AppUsers true "删除appUsers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /appUsers/deleteAppUsers [delete]
func (appUsersApi *AppUsersApi) DeleteAppUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := appUsersService.DeleteAppUsers(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAppUsersByIds 批量删除appUsers表
// @Tags AppUsers
// @Summary 批量删除appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /appUsers/deleteAppUsersByIds [delete]
func (appUsersApi *AppUsersApi) DeleteAppUsersByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := appUsersService.DeleteAppUsersByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAppUsers 更新appUsers表
// @Tags AppUsers
// @Summary 更新appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.AppUsers true "更新appUsers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /appUsers/updateAppUsers [put]
func (appUsersApi *AppUsersApi) UpdateAppUsers(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var appUsers app.AppUsers
	err := c.ShouldBindJSON(&appUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appUsersService.UpdateAppUsers(ctx, appUsers)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAppUsers 用id查询appUsers表
// @Tags AppUsers
// @Summary 用id查询appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询appUsers表"
// @Success 200 {object} response.Response{data=app.AppUsers,msg=string} "查询成功"
// @Router /appUsers/findAppUsers [get]
func (appUsersApi *AppUsersApi) FindAppUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reappUsers, err := appUsersService.GetAppUsers(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reappUsers, c)
}

// GetAppUsersList 分页获取appUsers表列表
// @Tags AppUsers
// @Summary 分页获取appUsers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query appReq.AppUsersSearch true "分页获取appUsers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /appUsers/getAppUsersList [get]
func (appUsersApi *AppUsersApi) GetAppUsersList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo appReq.AppUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := appUsersService.GetAppUsersInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetAppUsersPublic 不需要鉴权的appUsers表接口
// @Tags AppUsers
// @Summary 不需要鉴权的appUsers表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appUsers/getAppUsersPublic [get]
func (appUsersApi *AppUsersApi) GetAppUsersPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	appUsersService.GetAppUsersPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的appUsers表接口信息",
	}, "获取成功", c)
}
