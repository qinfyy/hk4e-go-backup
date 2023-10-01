package controller

import (
	"net/http"

	"hk4e/node/api"
	"hk4e/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (c *Controller) serverStopState(ctx *gin.Context) {
	stopServerInfo, err := c.discoveryClient.GetStopServerInfo(ctx.Request.Context(), &api.GetStopServerInfoReq{ClientIpAddr: ""})
	if err != nil {
		logger.Error("get stop server info error: %v", err)
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: stopServerInfo})
}

type ServerStopChangeReq struct {
	StopServer bool   `json:"stop_server"`
	StartTime  uint32 `json:"start_time"`
	EndTime    uint32 `json:"end_time"`
}

func (c *Controller) serverStopChange(ctx *gin.Context) {
	req := new(ServerStopChangeReq)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	_, err = c.discoveryClient.SetStopServerInfo(ctx.Request.Context(), &api.StopServerInfo{
		StopServer: req.StopServer,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	})
	if err != nil {
		logger.Error("set stop server info error: %v", err)
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: nil})
}

func (c *Controller) serverWhiteList(ctx *gin.Context) {
	whiteList, err := c.discoveryClient.GetWhiteList(ctx.Request.Context(), &api.NullMsg{})
	if err != nil {
		logger.Error("get white list error: %v", err)
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: whiteList.IpAddrList})
}

type ServerWhiteAdd struct {
	IpAddr string `json:"ip_addr"`
}

func (c *Controller) serverWhiteAdd(ctx *gin.Context) {
	req := new(ServerWhiteAdd)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	_, err = c.discoveryClient.SetWhiteList(ctx.Request.Context(), &api.SetWhiteListReq{
		IsAdd:  true,
		IpAddr: req.IpAddr,
	})
	if err != nil {
		logger.Error("set white list error: %v", err)
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: nil})
}

type ServerWhiteDel struct {
	IpAddr string `json:"ip_addr"`
}

func (c *Controller) serverWhiteDel(ctx *gin.Context) {
	req := new(ServerWhiteDel)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	_, err = c.discoveryClient.SetWhiteList(ctx.Request.Context(), &api.SetWhiteListReq{
		IsAdd:  false,
		IpAddr: req.IpAddr,
	})
	if err != nil {
		logger.Error("set white list error: %v", err)
		ctx.JSON(http.StatusOK, &CommonRsp{Code: -1, Msg: "", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, &CommonRsp{Code: 0, Msg: "", Data: nil})
}