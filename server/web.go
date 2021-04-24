package server

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"relayer/logging"
)

// HTTPService represents an HTTP service
type HTTPService struct {
	Router        *gin.Engine
	ChainManager  *ChainManager
	MarketManager *MarketManager
}

// NewHTTPService constructs a new HTTPService instance
func NewHTTPService(
	chainManager *ChainManager,
	marketManager *MarketManager,
) *HTTPService {
	srv := HTTPService{
		Router:        gin.Default(),
		ChainManager:  chainManager,
		MarketManager: marketManager,
	}

	srv.createRouter()

	return &srv
}

func (srv *HTTPService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Router.ServeHTTP(w, r)
}

func (srv *HTTPService) createRouter() {
	r := gin.Default()

	api := r.Group("/api/v0")
	fiscobcos := api.Group("/fiscobcos")
	{
		fiscobcos.POST("/chains", srv.AddChain)
		fiscobcos.POST("/chains/:chainid/update", srv.UpdateChain)
		fiscobcos.POST("/chains/:chainid/delete", srv.DeleteChain)
		fiscobcos.POST("/chains/:chainid/start", srv.StartChain)
		fiscobcos.POST("/chains/:chainid/stop", srv.StopChain)
		fiscobcos.GET("/chains", srv.GetChains)
		fiscobcos.GET("/chains/:chainid/status", srv.GetChainStatus)
	}

	ser := api.Group("/service")
	{
		ser.POST("/bindings/:chainid", srv.AddServiceBinding)
		ser.PUT("/bindings/:chainid/:svcname", srv.UpdateServiceBinding)
		ser.GET("/bindings/:chainid/:svcname", srv.GetServiceBinding)
	}

	r.GET("/health", srv.ShowHealth)

	srv.Router = r
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

	chainID, err := srv.ChainManager.AddChain(bodyBytes)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, AddChainResult{ChainID: chainID})
}

func (srv *HTTPService) DeleteChain(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.ChainManager.DeleteChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) UpdateChain(c *gin.Context) {
	var req AddChainRequest
	if err := c.BindJSON(&req); err != nil {
		onError(c, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.ChainManager.DeleteChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	chainID, err = srv.ChainManager.AddChain([]byte(req.ChainParams))
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, AddChainResult{ChainID: chainID})
}

func (srv *HTTPService) StartChain(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.ChainManager.StartChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) StopChain(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.ChainManager.StopChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) GetChains(c *gin.Context) {
	chains := srv.ChainManager.GetChains()

	onSuccess(c, chains)
}

func (srv *HTTPService) GetChainStatus(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	state, height, err := srv.ChainManager.GetChainStatus(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, ChainStatus{State: state, Height: height})
}

func (srv *HTTPService) AddServiceBinding(c *gin.Context) {
	var req AddServiceBindingRequest
	if err := c.BindJSON(&req); err != nil {
		onError(c, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.MarketManager.AddServiceBinding(chainID, req)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) UpdateServiceBinding(c *gin.Context) {
	var req UpdateServiceBindingRequest
	if err := c.BindJSON(&req); err != nil {
		onError(c, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	serviceName := c.Param("svcname")
	if err := ValidateServiceName(serviceName); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.MarketManager.UpdateServiceBinding(chainID, serviceName, req)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) GetServiceBinding(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	serviceName := c.Param("svcname")
	if err := ValidateServiceName(serviceName); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	binding, err := srv.MarketManager.GetServiceBinding(chainID, serviceName)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, binding)
}

// ShowHealth returns the health state
func (srv *HTTPService) ShowHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": true})
}

func onError(c *gin.Context, code int, msg string) {
	logging.Logger.Errorf(msg)

	c.JSON(code, ErrorResponse{
		Code:  CODE_ERROR,
		Error: msg,
	})
}

func onSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code:   CODE_SUCCESS,
		Result: result,
	})
}
