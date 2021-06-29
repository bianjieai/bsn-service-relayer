package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"relayer/appchains"
	"strconv"

	"github.com/gin-gonic/gin"

	"relayer/logging"
)

// HTTPService represents an HTTP service
type HTTPService struct {
	Router   *gin.Engine
	AppChain appchains.AppChainHandlerI
}

// NewHTTPService constructs a new HTTPService instance
func NewHTTPService(
	appChain appchains.AppChainHandlerI,
) *HTTPService {
	srv := HTTPService{
		Router:   gin.Default(),
		AppChain: appChain,
	}

	srv.createRouter()

	return &srv
}

func (srv *HTTPService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Router.ServeHTTP(w, r)
}

func (srv *HTTPService) createRouter() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	api := r.Group("/api/v0")
	fabric := api.Group("/fabric")
	{
		fabric.POST("/regSideChain", srv.AddChain)
		fabric.POST("/removeAppChain", srv.DeleteChain)
		fabric.POST("/updateAppChain", srv.UpdateChain)
	}

	ser := api.Group("/service")
	{
		ser.POST("/bindings/:chainid", srv.AddServiceBinding)
		ser.PUT("/bindings/:chainid/:svcname", srv.UpdateServiceBinding)
		ser.GET("/bindings/:chainid/:svcname", srv.GetServiceBinding)
	}

	//r.POST("/chains", srv.AddChain)
	//r.POST("/chains/:chainid/delete", srv.DeleteChain)

	//r.POST("/chains/:chainid/start", srv.StartChain)
	//r.POST("/chains/:chainid/stop", srv.StopChain)
	//
	//r.GET("/chains", srv.GetChains)
	//r.GET("/chains/:chainid/status", srv.GetChainStatus)

	r.GET("/health", srv.ShowHealth)

	srv.Router = r
}

//由于默认http.Request.Body类型为io.ReadCloser类型,即只能读一次，读完后直接close掉,后续流程无法继续读取。
func ReadBody(body io.ReadCloser) error {
	var bodyBytes []byte // 我们需要的body内容

	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("Invalid request body")
	}

	// 新建缓冲区并替换原有Request.body
	body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// 当前函数可以使用body内容
	//_ := bodyBytes
	return err
}

func (srv *HTTPService) AddChain(c *gin.Context) {

	var bodyBytes []byte
	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	logging.Logger.Infof("AddChain data is %s", string(bodyBytes))

	chainID, err := srv.AppChain.RegisterChain(bodyBytes) //.AddChain([]byte(req.ChainParams))
	if err != nil {
		logging.Logger.Errorf(err.Error())
		onError(c, err)
		return
	}

	onSuccess(c, AddChainResult{ChainID: strconv.FormatUint(chainID, 10)})
}

func (srv *HTTPService) DeleteChain(c *gin.Context) {
	var bodyBytes []byte // 我们需要的body内容

	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	logging.Logger.Infof("DeleteChain data is %s", string(bodyBytes))

	err = srv.AppChain.DeleteChain(bodyBytes)
	if err != nil {
		onError(c, err)
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) UpdateChain(c *gin.Context) {
	var bodyBytes []byte // 我们需要的body内容

	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	logging.Logger.Infof("UpdateChain data is %s", string(bodyBytes))

	err = srv.AppChain.UpdateChain(bodyBytes)
	if err != nil {
		onError(c, err)
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) AddServiceBinding(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	logging.Logger.Infof("AddServiceBinding data is %s", string(bodyBytes))

	chainID := c.Param("chainid")

	err = srv.AppChain.AddServiceBinding(chainID, bodyBytes)
	if err != nil {
		onError(c, err)
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) UpdateServiceBinding(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	logging.Logger.Infof("UpdateServiceBinding data is %s", string(bodyBytes))

	chainID := c.Param("chainid")

	err = srv.AppChain.UpdateServiceBinding(chainID, bodyBytes)
	if err != nil {
		onError(c, err)
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) GetServiceBinding(c *gin.Context) {
	chainID := c.Param("chainid")
	serviceName := c.Param("servicename")

	binding, err := srv.AppChain.GetServiceBinding(chainID, serviceName)
	if err != nil {
		onError(c, err)
		return
	}

	onSuccess(c, binding)
}

//func (srv *HTTPService) StartChain(c *gin.Context) {
//	chainID := c.Param("chainid")
//
//	err := srv.ChainManager.StartChain(chainID)
//	if err != nil {
//		onError(c, err)
//		return
//	}
//
//	onSuccess(c, nil)
//}
//
//func (srv *HTTPService) StopChain(c *gin.Context) {
//	chainID := c.Param("chainid")
//
//	err := srv.ChainManager.StopChain(chainID)
//	if err != nil {
//		onError(c, err)
//		return
//	}
//
//	onSuccess(c, nil)
//}

//func (srv *HTTPService) GetChains(c *gin.Context) {
//	chains := srv.ChainManager.GetChains()
//
//	onSuccess(c, chains)
//}

//func (srv *HTTPService) GetChainStatus(c *gin.Context) {
//	chainID := c.Param("chainid")
//
//	state, height, err := srv.ChainManager.GetChainStatus(chainID)
//	if err != nil {
//		onError(c, err)
//		return
//	}
//
//	onSuccess(c, ChainStatus{State: state, Height: height})
//}

// ShowHealth returns the health state
func (srv *HTTPService) ShowHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": true})
}

func onError(c *gin.Context, err error) {
	logging.Logger.Errorf(err.Error())

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:  CODE_ERROR,
		Error: err.Error(),
	})
}

func onSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code:   CODE_SUCCESS,
		Msg:    "success",
		Result: result,
	})
}
