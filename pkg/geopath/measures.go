package geopath

import (
	"context"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	serverUrl string = "https://direct.cco.io"
)

func NewFrameClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Frame], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Frame](ctx, svr, "/v1/frames", oauth2)
}

func NewFrameConstructionClassificationClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[ConstructionClassification], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[ConstructionClassification](ctx, svr, "/v1/construction/classifications", oauth2)
}

func NewFrameConstructionPlacementClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[ConstructionPlacement], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[ConstructionPlacement](ctx, svr, "/v1/construction/placements", oauth2)
}

func NewFrameConstructionTypeClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[ConstructionType], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[ConstructionType](ctx, svr, "/v1/construction/types", oauth2)
}

func NewFrameMeasuresClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.ChildClient[Measure], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewChildClient[Measure](ctx, svr, "/v1/frames", "/measures", oauth2)
}

func NewMeasuresClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Measure], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Measure](ctx, svr, "/v1/measures", oauth2)
}
