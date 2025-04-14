// Copyright (c) 2025 Renorm Labs. All rights reserved.

package spy_test

import (
	"testing"

	"github.com/renormlabs/assert"
	"github.com/renormlabs/assert/internal/spy"
)

func TestSpyExpectFailure(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	s.Errorf("test failed")
	assert.True(t, s.SpiedOnAFailure())
}

func TestSpyExpectSuccess(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectSuccess()
	assert.False(t, s.SpiedOnAFailure())
}

func TestSpyFail(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	s.Fail()
}

func TestSpyFailNow(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	s.FailNow()
}

func TestSpyError(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	s.Error("test error")
}

func TestSpyErrorf(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	s.Errorf("test error %s", "foo")
}

func TestSpyFatalf(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	s.Fatalf("test fatal %s", "foo")
}

func TestSpyFatal(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	s.Fatal("test fatal")
}
