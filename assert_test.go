// Copyright (c) 2025 Renorm Labs. All rights reserved.

package assert_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/renormlabs/assert"
	"github.com/renormlabs/assert/internal/spy"
)

var errToCompare = errors.New("test error")

// Tests expected to PASS:
func TestEqual(t *testing.T) {
	assert.Equal(t, 42, 42)
}
func TestNotEqual(t *testing.T) {
	assert.NotEqual(t, 1, 2)
}
func TestTrue(t *testing.T) {
	assert.True(t, true)
}
func TestFalse(t *testing.T) {
	assert.False(t, false)
}
func TestNil(t *testing.T) {
	var v testing.TB
	assert.Nil(t, v)
}
func TestNotNil(t *testing.T) {
	assert.NotNil(t, "hello")
}
func TestStringContains(t *testing.T) {
	assert.StringContains(t, "golang testing helpers", "testing")
}
func TestStringNotContains(t *testing.T) {
	assert.StringDoesNotContain(t, "golang testing helpers", "foo")
}
func TestPanics(t *testing.T) {
	f := func() {
		panic("test")
	}
	assert.Panics(t, f)
}
func TestDoesNotPanic(t *testing.T) {
	f := func() {}
	assert.DoesNotPanic(t, f)
}
func TestErrorIs(t *testing.T) {
	f := func() error {
		return fmt.Errorf("foo: %w", errToCompare)
	}
	assert.ErrorIs(t, f(), errToCompare)
}
func TestErrorIsNot(t *testing.T) {
	f := func() error {
		return fmt.Errorf("test error")
	}
	assert.ErrorIsNot(t, f(), errToCompare)
	assert.Equal(t, f().Error(), errToCompare.Error())
}

func TestMapContainsKey(t *testing.T) {
	m := map[string]string{
		"foo": "bar",
	}
	assert.MapContainsKey(t, m, "foo")
}

func TestMapDoesNotContainKey(t *testing.T) {
	m := map[string]string{
		"foo": "bar",
	}
	assert.MapDoesNotContainKey(t, m, "baz")
}

func TestEmptyMap(t *testing.T) {
	m := map[string]string{}
	assert.EmptyMap(t, m)
}

func TestNotEmptyMap(t *testing.T) {
	m := map[string]string{
		"foo": "bar",
	}
	assert.NotEmptyMap(t, m)
}

func TestEmptySlice(t *testing.T) {
	s := []string{}
	assert.EmptySlice(t, s)
}

func TestNotEmptySlice(t *testing.T) {
	s := []string{"foo", "bar"}
	assert.NotEmptySlice(t, s)
}

func TestDeepEqual(t *testing.T) {
	type person struct {
		Name    string
		Age     int
		Friends []person
	}
	p1 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Bob", Age: 25}},
	}
	p2 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Bob", Age: 25}},
	}
	assert.DeepEqual(t, p1, p2)
	assert.NotEqual(t, &p1, &p2)
}

func TestNotDeepEqual(t *testing.T) {
	type person struct {
		Name    string
		Age     int
		Friends []person
	}
	p1 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Bob", Age: 25}},
	}
	p2 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Bob", Age: 26}},
	}
	assert.NotDeepEqual(t, p1, p2)
}

// Tests expected to FAIL:
// The following tests are designed to fail. To check that they do indeed fail, we wrap the
// testing.T in a custom struct that tracks whether an error was reported. Test authors can use the
// following pattern:
//
//	func TestExample(t *testing.T) {
//		s := assert.SpyOn(t)
//		defer s.ExpectFailure()
//		assert.True(s, false)
//	}

func TestTrueFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.True(s, false)
}

func TestFalseFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.False(s, true)
}

func TestEqualFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.Equal(s, 1, 2)
}

func TestNotEqualFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.NotEqual(s, "same", "same")
}

func TestNilFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.Nil(s, 123)
}

func TestNotNilFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.NotNil(s, nil)
}

func TestStringContainsFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.StringContains(s, "golang testing helpers", "python")
}

func TestStringDoesNotContainFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	assert.StringDoesNotContain(s, "golang testing helpers", "helpers")
}

func TestPanicsFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	f := func() {}
	assert.Panics(s, f)
}

func TestDoesNotPanicFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	f := func() {
		panic("test")
	}
	assert.DoesNotPanic(s, f)
}

func TestErrorIsFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	f := func() error {
		return fmt.Errorf("foo: %w", errToCompare)
	}
	assert.ErrorIs(s, f(), errors.New("not the same"))
}

func TestErrorIsNotFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	f := func() error {
		return fmt.Errorf("foo: %w", errToCompare)
	}
	assert.ErrorIsNot(s, f(), errToCompare)
	assert.Equal(s, f().Error(), errToCompare.Error())
}

func TestMapContainsKeyFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	m := map[string]string{
		"foo": "bar",
	}
	assert.MapContainsKey(s, m, "baz")
}

func TestMapDoesNotContainKeyFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	m := map[string]string{
		"foo": "bar",
	}
	assert.MapDoesNotContainKey(s, m, "foo")
}

func TestEmptyMapFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	m := map[string]string{
		"foo": "bar",
	}
	assert.EmptyMap(s, m)
}

func TestNotEmptyMapFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	m := map[string]string{}
	assert.NotEmptyMap(s, m)
}

func TestEmptySliceFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	slice := []string{"foo", "bar"}
	assert.EmptySlice(s, slice)
}

func TestNotEmptySliceFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	slice := []string{}
	assert.NotEmptySlice(s, slice)
}

func TestDeepEqualFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	type person struct {
		Name    string
		Age     int
		Friends []person
	}
	p1 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Bob", Age: 25}},
	}
	p2 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Charlie", Age: 25}},
	}
	assert.DeepEqual(s, p1, p2)
}
func TestNotDeepEqualFails(t *testing.T) {
	s := spy.SpyOn(t)
	defer s.ExpectFailure()
	type person struct {
		Name    string
		Age     int
		Friends []person
	}
	p1 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Bob", Age: 25}},
	}
	p2 := person{
		Name:    "Alice",
		Age:     30,
		Friends: []person{{Name: "Bob", Age: 25}},
	}
	assert.NotDeepEqual(s, p1, p2)
}
