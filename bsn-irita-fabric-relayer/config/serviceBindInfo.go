package config

import "github.com/spf13/viper"

const (
	service_Name        = "service.service_name"
	service_Description = "service.service_description"
	service_Schemas     = "service.service_schemas"
	service_Provider    = "service.service_provider"
	service_Fee         = "service.service_fee"
	service_Qos         = "service.service_qos"
)

type ServiceBindInfo struct {
	Name        string
	Description string
	Schemas     string
	Provider    string
	Fee         string
	Qos         uint64
}

func GetServiceBindInfo(v *viper.Viper) *ServiceBindInfo {

	service := &ServiceBindInfo{
		Name:        v.GetString(service_Name),
		Description: v.GetString(service_Description),
		Schemas:     v.GetString(service_Schemas),
		Provider:    v.GetString(service_Provider),
		Fee:         v.GetString(service_Fee),
		Qos:         v.GetUint64(service_Qos),
	}
	return service
}
