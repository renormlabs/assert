// Copyright (c) 2025 Renorm Labs. All rights reserved.

package spy

import (
	"testing"

	"github.com/renormlabs/assert"
)

type TBSpy struct {
	testing.TB
	failed bool
}

func SpyOn(t testing.TB) *TBSpy {
	return &TBSpy{TB: t}
}

func (s *TBSpy) SpiedOnAFailure() bool {
	return s.failed
}

func (s *TBSpy) Fail() {
	s.TB.Helper()
	s.failed = true
}

func (s *TBSpy) FailNow() {
	s.TB.Helper()
	s.failed = true
}

func (s *TBSpy) Error(args ...interface{}) {
	s.TB.Helper()
	s.failed = true
	s.TB.Log(args...)
}

func (s *TBSpy) Errorf(format string, args ...interface{}) {
	s.TB.Helper()
	s.failed = true
	s.TB.Logf(format, args...)
}

func (s *TBSpy) Fatalf(format string, args ...interface{}) {
	s.TB.Helper()
	s.failed = true
	s.TB.Logf(format, args...)
}

func (s *TBSpy) Fatal(args ...interface{}) {
	s.TB.Helper()
	s.failed = true
	s.TB.Log(args...)
}

func (s *TBSpy) ExpectFailure() {
	s.TB.Helper()
	assert.Truef(s.TB, s.failed, "[spy] Expected test failure but got success")
}

func (s *TBSpy) ExpectSuccess() {
	s.TB.Helper()
	assert.Falsef(s.TB, s.failed, "[spy] Expected test success but got failure")
}
