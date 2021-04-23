package service

import (
	"github.com/spf13/viper"
	cfg "relayer/config"
)

const (
	Prefix      = "service"
	ServiceName = "service_name"
	Schemas     = "schemas"
	Provider    = "provider"
	ServiceFee  = "service_fee"
	QoS         = "qos"
)

// ServiceBinding defines an iService binding
type ServiceBindInfo struct {
	ServiceName string   // service name
	Schemas     string   // input and output schemas
	Provider    string   // service provider
	ServiceFee  string   // service fee
	QoS         uint64    // quality of service, in terms of the minimum response time
}

// ServiceBindInfo constructs a new ServiceBindInfo from viper
func NewServiceBindInfo(v *viper.Viper) ServiceBindInfo {
	return ServiceBindInfo{
		ServiceName: v.GetString(cfg.GetConfigKey(Prefix, ServiceName)),
		Schemas:     v.GetString(cfg.GetConfigKey(Prefix, Schemas)),
		Provider:    v.GetString(cfg.GetConfigKey(Prefix, Provider)),
		ServiceFee:  v.GetString(cfg.GetConfigKey(Prefix, ServiceFee)),
		QoS:         v.GetUint64(cfg.GetConfigKey(Prefix, QoS)),
	}
}
