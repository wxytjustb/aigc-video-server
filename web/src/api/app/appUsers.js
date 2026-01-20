import service from '@/utils/request'
// @Tags AppUsers
// @Summary 创建appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "创建appUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /appUsers/createAppUsers [post]
export const createAppUsers = (data) => {
  return service({
    url: '/appUsers/createAppUsers',
    method: 'post',
    data
  })
}

// @Tags AppUsers
// @Summary 删除appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "删除appUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUsers/deleteAppUsers [delete]
export const deleteAppUsers = (params) => {
  return service({
    url: '/appUsers/deleteAppUsers',
    method: 'delete',
    params
  })
}

// @Tags AppUsers
// @Summary 批量删除appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除appUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUsers/deleteAppUsers [delete]
export const deleteAppUsersByIds = (params) => {
  return service({
    url: '/appUsers/deleteAppUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags AppUsers
// @Summary 更新appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "更新appUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUsers/updateAppUsers [put]
export const updateAppUsers = (data) => {
  return service({
    url: '/appUsers/updateAppUsers',
    method: 'put',
    data
  })
}

// @Tags AppUsers
// @Summary 用id查询appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AppUsers true "用id查询appUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUsers/findAppUsers [get]
export const findAppUsers = (params) => {
  return service({
    url: '/appUsers/findAppUsers',
    method: 'get',
    params
  })
}

// @Tags AppUsers
// @Summary 分页获取appUsers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取appUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUsers/getAppUsersList [get]
export const getAppUsersList = (params) => {
  return service({
    url: '/appUsers/getAppUsersList',
    method: 'get',
    params
  })
}

// @Tags AppUsers
// @Summary 不需要鉴权的appUsers表接口
// @Accept application/json
// @Produce application/json
// @Param data query appReq.AppUsersSearch true "分页获取appUsers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appUsers/getAppUsersPublic [get]
export const getAppUsersPublic = () => {
  return service({
    url: '/appUsers/getAppUsersPublic',
    method: 'get',
  })
}
