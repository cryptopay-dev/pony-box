package box

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestNewProvider(t *testing.T) {
	t.Run("without options", func(t *testing.T) {
		p := NewProvider(func() error {
			return nil
		})

		assert.NotNil(t, p)
		assert.NotNil(t, p.fn)
		assert.Nil(t, p.opts)
	})

	t.Run("with options", func(t *testing.T) {
		p := NewProvider(func() error {
			return nil
		}, dig.Name("testing"))

		assert.NotNil(t, p)
		assert.NotNil(t, p.fn)
		assert.NotNil(t, p.opts)
	})
}

type (
	A struct{}
	B struct{}
)

func (a A) hello() string {
	return "hello, a"
}

func (b B) hello() string {
	return "hello, b"
}

func TestNew(t *testing.T) {
	dep1 := NewProvider(func() A {
		return A{}
	})

	dep2 := NewProvider(func() B {
		return B{}
	})

	dep3 := NewProvider(func() (B, error) {
		return B{}, errors.New("unknown error")
	})

	t.Run("with dependencies", func(t *testing.T) {
		b := New()
		{
			err := b.Provide(dep1, dep2)
			assert.NoError(t, err)
		}

		{
			err := b.Invoke(func(a A, b B) {
				assert.Equal(t, "hello, a", a.hello())
				assert.Equal(t, "hello, b", b.hello())
			})
			assert.NoError(t, err)
		}
	})

	t.Run("missing dependencies", func(t *testing.T) {
		b := New()
		{
			err := b.Provide(dep1)
			assert.NoError(t, err)
		}

		{
			err := b.Invoke(func(a A, b B) {
			})
			assert.Error(t, err)
			assert.True(t, strings.Contains(err.Error(), "type box.B is not in the container"))
		}
	})

	t.Run("error in dependencies", func(t *testing.T) {
		b := New()
		{
			err := b.Provide(dep1, dep3)
			assert.NoError(t, err)
		}

		{
			err := b.Invoke(func(a A, b B) {
			})
			assert.Error(t, err)
			assert.True(t, strings.Contains(err.Error(), "unknown error"))
		}
	})

	t.Run("providing not function as a dependency", func(t *testing.T) {
		b := New()
		{
			err := b.Provide(dep1, NewProvider(A{}))
			assert.Error(t, err)
		}
	})

	t.Run("empty providers", func(t *testing.T) {
		b := New()
		err := b.Provide()
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyProviders, err)
	})
}
