package core

// ResponseAdaptor is the wrapped response struct of Irita-Hub
type ResponseAdaptor struct {
	StatusCode  int
	Result      string
	Output      string
}

// GetErrMsg implements ResponseI
func (r ResponseAdaptor) GetErrMsg() string {
	switch r.StatusCode {
	case 200:
		return ""

	case 400, 500:
		return r.Result

	default:
		return ""
	}
}

// GetOutput implements ResponseI
func (r ResponseAdaptor) GetOutput() string {
	switch r.StatusCode {
	case 200:
		return r.Output

	case 400, 500:
		return r.Result

	default:
		return ""
	}
}
