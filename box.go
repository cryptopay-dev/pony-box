package box

import (
	"errors"

	"go.uber.org/dig"
)

type (
	Box struct {
		dig *dig.Container
	}

	Provider struct {
		fn   interface{}
		opts []dig.ProvideOption
	}
)

var (
	ErrEmptyProviders = errors.New("empty providers")
)

func New() Box {
	c := dig.New()

	return Box{
		dig: c,
	}
}

func NewProvider(fn interface{}, opts ...dig.ProvideOption) Provider {
	return Provider{
		fn:   fn,
		opts: opts,
	}
}

func (b Box) Provide(providers ...Provider) error {
	if len(providers) == 0 {
		return ErrEmptyProviders
	}

	for _, provider := range providers {
		if err := b.dig.Provide(provider.fn, provider.opts...); err != nil {
			return dig.RootCause(err)
		}
	}

	return nil
}

func (b Box) Invoke(fn interface{}, opts ...dig.InvokeOption) error {
	return b.dig.Invoke(fn, opts...)
}
