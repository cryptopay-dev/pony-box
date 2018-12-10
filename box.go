package box

import (
	"errors"

	"go.uber.org/dig"
)

type (
	// Box is handling dig container
	Box struct {
		dig *dig.Container
	}

	// Provider handles constructor fn and options provided
	Provider struct {
		fn   interface{}
		opts []dig.ProvideOption
	}
)

var (
	// ErrEmptyProviders thrown when empty providers provided
	ErrEmptyProviders = errors.New("empty providers")
)

// New creates new Box
func New() Box {
	c := dig.New()

	return Box{
		dig: c,
	}
}

// NewProvider creates new provider that can be added to box
func NewProvider(fn interface{}, opts ...dig.ProvideOption) Provider {
	return Provider{
		fn:   fn,
		opts: opts,
	}
}

// Provide setting up a list of providers inside a box, can throw ErrEmptyProviders when providers list is empty.
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

// Invoke run fn function with all it's dependencies trying to resolve them through box
func (b Box) Invoke(fn interface{}, opts ...dig.InvokeOption) error {
	return b.dig.Invoke(fn, opts...)
}
