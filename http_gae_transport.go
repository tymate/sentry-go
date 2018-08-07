package sentry

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/appengine/urlfetch"
)

const (
	gaeContextClass = "sentry-go.gae-context"
)

type gaeTransport struct {
	*httpTransport
}

func NewGAETransport() Transport {
	t := newHTTPTransport()
	httpT := t.(*httpTransport)
	return &gaeTransport{httpT}
}

func (t *gaeTransport) Send(dsn string, p Packet) error {
	pp := *p.(*packet)
	ctxOpt, ok := pp[gaeContextClass]
	if !ok {
		return errors.New("missing Google AppEngine context")
	}

	opt, ok := ctxOpt.(*gaeContextOption)
	if !ok {
		return errors.New("invalid GAEContext option")
	}

	delete(pp, gaeContextClass)

	t.client = urlfetch.Client(opt.ctx)
	return t.Send(dsn, &pp)
}

type gaeContextOption struct {
	ctx context.Context
}

func (o *gaeContextOption) Class() string {
	return gaeContextClass
}

func (o *gaeContextOption) Omit() bool {
	return false
}

func GAEContext(ctx context.Context) Option {
	return &gaeContextOption{ctx}
}
