package singleton_test

import (
	singleton "design-pattern/01_singleton"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}