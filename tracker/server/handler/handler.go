package handler

import (
	"context"

	"github.com/darren-west/marathon/tracker"
	//log "github.com/sirupsen/logrus"
)

type LocationHandler struct{}

func (LocationHandler) Create(ctx context.Context, location *tracker.CreateRequest) (*tracker.CreateResponse, error) {

	return new(tracker.CreateResponse), nil
}
