package errgroup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyGroup(t *testing.T) {
	g := New()
	assert.Nil(t, g.Wait())
	assert.Nil(t, g.Wait())
}

func TestSuccesfulCall(t *testing.T) {
	g := New()
	called1 := false
	g.Go(func() error {
		called1 = true
		return nil
	})
	assert.Nil(t, g.Wait())
	assert.True(t, called1)
	called2 := false
	g.Go(func() error {
		called2 = true
		return nil
	})
	assert.Nil(t, g.Wait())
	assert.True(t, called2)
}

func TestErrorCall(t *testing.T) {
	g := New()
	err := fmt.Errorf("error mcerrface")
	g.Go(func() error {
		return err
	})
	assert.Equal(t, err, g.Wait())
	g.Go(func() error {
		return nil
	})
	assert.Equal(t, err, g.Wait())
	g.Go(func() error {
		return fmt.Errorf("different error")
	})
	assert.Equal(t, err, g.Wait())
}

func TestCancel(t *testing.T) {
	g := New()
	err := fmt.Errorf("error mcerrface")
	g.Go(func() error {
		select{}  // never returns
		return nil
	})
	g.Cancel(err)
	assert.Equal(t, err, g.Wait())
}
