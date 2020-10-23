package core

import (
	log "github.com/sirupsen/logrus"
)

// Relayer is a relayer implementing the interchain service invocation between the application chain and Irita-Hub
type Relayer struct {
	IritaHubChain IritaHubChainI
	AppChain      AppChainI
	Logger        *log.Logger
}

// NewRelayer constructs a new Relayer instance
func NewRelayer(iritaHubChain IritaHubChainI, appChain AppChainI, logger *log.Logger) Relayer {
	return Relayer{
		IritaHubChain: iritaHubChain,
		AppChain:      appChain,
		Logger:        logger,
	}
}

// BuildInterchainRequest builds an interchain request from an interchain event
func (r Relayer) BuildInterchainRequest(icEvent InterchainEventI) (InterchainRequestI, error) {
	serviceBinding, err := r.AppChain.GetServiceBinding(icEvent.GetServiceName())
	if err != nil {
		return nil, err
	}

	return InterchainRequestAdaptor{
		ServiceName:   icEvent.GetServiceName(),
		Provider:      serviceBinding.GetProvider(),
		Input:         icEvent.GetInput(),
		Timeout:       icEvent.GetTimeout(),
		ServiceFeeCap: serviceBinding.GetServiceFee(),
	}, nil
}

// HandleInterchainEvent handles the interchain event
func (r Relayer) HandleInterchainEvent(icEvent InterchainEventI) {
	r.Logger.Infof("got the interchain event on %s: %+v", r.AppChain.GetChainID(), icEvent)

	icRequest, err := r.BuildInterchainRequest(icEvent)
	if err != nil {
		r.Logger.Errorf(
			"failed to build the interchain request: %s",
			err,
		)

		return
	}

	callback := func(icRequestID string, response ResponseI) {
		r.Logger.Infof(
			"got the response of the interchain request on %s: %+v",
			r.IritaHubChain.GetChainID(),
			response,
		)

		err := r.AppChain.SendResponse(icEvent.GetRequestID(), response)
		if err != nil {
			r.Logger.Errorf(
				"failed to send the response to %s: %s",
				r.AppChain.GetChainID(),
				err,
			)

			return
		}

		r.Logger.Infof(
			"response sent to %s successfully",
			r.AppChain.GetChainID(),
		)
	}

	err = r.IritaHubChain.SendInterchainRequest(icRequest, callback)
	if err != nil {
		r.Logger.Errorf(
			"failed to handle the interchain request %+v on %s: %s",
			icRequest,
			r.IritaHubChain.GetChainID(),
			err,
		)
	}
}

// Start starts the relayer process
func (r Relayer) Start() {
	r.Logger.Infof("relayer started")
	r.Logger.Infof("watching the interchain events on %s", r.AppChain.GetChainID())

	err := r.AppChain.InterchainEventListener(r.HandleInterchainEvent)
	if err != nil {
		r.Logger.Errorf(
			"failed to listen to the interchain events on %s: %s",
			r.AppChain.GetChainID(),
			err,
		)
	}
}
