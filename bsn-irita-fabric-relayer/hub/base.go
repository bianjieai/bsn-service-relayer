package hub

// ServiceInput defines the service input
type ServiceInput struct {
	Header map[string]interface{} `json:"header"`
	Body   map[string]interface{} `json:"body"`
}

// AddHeader adds the given kv to the header
func (s *ServiceInput) AddHeader(key string, value interface{}) {
	s.Header[key] = value
}
