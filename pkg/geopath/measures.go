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

func NewConstructionClassificationClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[ConstructionClassification], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[ConstructionClassification](ctx, svr, "/v1/construction/classifications", oauth2)
}

func NewConstructionPlacementClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[ConstructionPlacement], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[ConstructionPlacement](ctx, svr, "/v1/construction/placements", oauth2)
}

func NewConstructionTypeClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[ConstructionType], error) {
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

func NewFrameHistoryClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.ChildClient[Measure], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewChildClient[Measure](ctx, svr, "/v1/frames", "/history", oauth2)
}

func NewIlluminationTypeClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[IlluminationType], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[IlluminationType](ctx, svr, "/v1/illumination/types", oauth2)
}

func NewLocationTypeClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[LocationType], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[LocationType](ctx, svr, "/v1/locations/types", oauth2)
}

func NewMeasuresClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Measure], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Measure](ctx, svr, "/v1/measures", oauth2)
}

func NewMeasuresStatusClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[MeasuresStatus], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[MeasuresStatus](ctx, svr, "/v1/measures/status", oauth2)
}

// TODO: /v1/measures/status
// Support for this client requires customization

func NewMediaTypeClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[MediaType], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[MediaType](ctx, svr, "/v1/media/types", oauth2)
}

func NewSegmentIDClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[SegmentID], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}
	return clients.NewClient[SegmentID](ctx, svr, "/v1/segments/ids", oauth2)
}

func NewSegmentNameClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[SegmentName], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[SegmentName](ctx, svr, "/v1/segments/names", oauth2)
}
