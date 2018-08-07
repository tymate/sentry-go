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
	ctxOpt, ok := (*p.(*packet))[gaeContextClass]
	if !ok {
		return errors.New("missing Google AppEngine context")
	}

	opt, ok := ctxOpt.(*gaeContextOption)
	if !ok {
		return errors.New("invalid GAEContext option")
	}

	t.client = urlfetch.Client(opt.ctx)
	return t.Send(dsn, p)
}

type gaeContextOption struct {
	ctx context.Context
}

func (o *gaeContextOption) Class() string {
	return gaeContextClass
}

func (o *gaeContextOption) Omit() bool {
	return true
}

func GAEContext(ctx context.Context) Option {
	return &gaeContextOption{ctx}
}
