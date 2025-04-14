// Copyright (c) 2025 Renorm Labs. All rights reserved.

// Package assert provides simple, readable assertions for testing in Go. It is designed to be used
// with the testing package.
package assert

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// [Equalf] asserts that two values are equal. If they are not, it reports an error with the given formatting.
func Equalf[T comparable](t testing.TB, expected, actual T, format string, args ...any) {
	t.Helper()
	if expected != actual {
		t.Errorf(format, args...)
	}
}

// [Equal] asserts that two values are equal using [Equalf] with a default message.
func Equal[T comparable](t testing.TB, expected, actual T) {
	t.Helper()
	Equalf(t, expected, actual, "expected %v to equal %v", actual, expected)
}

// [NotEqualf] asserts that two values are not equal. If they are, it reports an error with the given formatting.
func NotEqualf[T comparable](t testing.TB, expected, actual T, format string, args ...any) {
	t.Helper()
	if expected == actual {
		t.Errorf(format, args...)
	}
}

// [NotEqual] asserts that two values are not equal using [NotEqualf] with a default message.
func NotEqual[T comparable](t testing.TB, expected, actual T) {
	t.Helper()
	NotEqualf(t, expected, actual, "expected %v to not equal %v", actual, expected)
}

// [Truef] asserts that a boolean value is true. If it is not, it reports an error with the given formatting.
func Truef(t testing.TB, actual bool, format string, args ...any) {
	t.Helper()
	Equalf(t, true, actual, format, args...)
}

// [True] asserts that a boolean value is true using [Truef] with a default message.
func True(t testing.TB, actual bool) {
	t.Helper()
	Truef(t, actual, "expected %v to be true", actual)
}

// [Falsef] asserts that a boolean value is false. If it is not, it reports an error with the given formatting.
func Falsef(t testing.TB, actual bool, format string, args ...any) {
	t.Helper()
	Equalf(t, false, actual, format, args...)
}

// [False] asserts that a boolean value is false using [Falsef] with a default message.
func False(t testing.TB, actual bool) {
	t.Helper()
	Falsef(t, actual, "expected %v to be false", actual)
}

// [Nilf] asserts that a value is nil. If it is not, it reports an error with the given formatting.
func Nilf(t testing.TB, actual any, format string, args ...any) {
	t.Helper()
	if actual != nil {
		t.Errorf(format, args...)
	}
}

// [Nil] asserts that a value is nil using [Nilf] with a default message.
func Nil(t testing.TB, actual any) {
	t.Helper()
	Nilf(t, actual, "expected %v to be nil", actual)
}

// [NotNilf] asserts that a value is not nil. If it is, it reports an error with the given formatting.
func NotNilf(t testing.TB, actual any, format string, args ...any) {
	t.Helper()
	if actual == nil {
		t.Errorf(format, args...)
	}
}

// [NotNil] asserts that a value is not nil using [NotNilf] with a default message.
func NotNil(t testing.TB, actual any) {
	t.Helper()
	NotNilf(t, actual, "expected %v to not be nil", actual)
}

// [StringContainsf] asserts that a string contains a substring. If it does not, it reports an error with the given formatting.
func StringContainsf(t testing.TB, str, substr string, format string, args ...any) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf(format, args...)
	}
}

// [StringContains] asserts that a string contains a substring using [StringContainsf] with a default message.
func StringContains(t testing.TB, str, substr string) {
	t.Helper()
	StringContainsf(t, str, substr, "expected %v to contain the substring %v", str, substr)
}

// [StringNotContainsf] asserts that a string does not contain a substring. If it does, it reports an error with the given formatting.
func StringDoesNotContainf(t testing.TB, str, substr string, format string, args ...any) {
	t.Helper()
	if strings.Contains(str, substr) {
		t.Errorf(format, args...)
	}
}

// [StringNotContains] asserts that a string does not contain a substring using [StringNotContainsf] with a default message.
func StringDoesNotContain(t testing.TB, str, substr string) {
	t.Helper()
	StringDoesNotContainf(t, str, substr, "expected %v to not contain the substring %v", str, substr)
}

// [Panicsf] asserts that a function panics. If it does not, it reports an error with the given formatting.
func Panicsf(t testing.TB, f func(), format string, args ...any) (recovery any) {
	t.Helper()
	defer func() {
		recovery = recover()
		if recovery == nil {
			t.Errorf(format, args...)
		}
	}()
	f()
	return
}

// [Panics] asserts that a function panics using [Panicsf] with a default message.
func Panics(t testing.TB, f func()) any {
	t.Helper()
	return Panicsf(t, f, "expected function to panic, did not panic")
}

// [DoesNotPanicf] asserts that a function does not panic. If it does, it reports an error with the given formatting.
func DoesNotPanicf(t testing.TB, f func(), format string, args ...any) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf(format, args...)
		}
	}()
	f()
}

// [DoesNotPanic] asserts that a function does not panic using [DoesNotPanicf] with a default message.
func DoesNotPanic(t testing.TB, f func()) {
	t.Helper()
	DoesNotPanicf(t, f, "expected function to not panic, did panic")
}

// [ErrorIsf] asserts that an error is of a specific type. If it is not, it reports an error with the given formatting.
func ErrorIsf(t testing.TB, err, target error, format string, args ...any) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Errorf(format, args...)
	}
}

// [ErrorIs] asserts that an error is of a specific type using [ErrorIsf] with a default message.
func ErrorIs(t testing.TB, err, target error) {
	t.Helper()
	ErrorIsf(t, err, target, "expected error %v to be %v", err, target)
}

// [ErrorIsNotf] asserts that an error is not of a specific type. If it is, it reports an error with the given formatting.
func ErrorIsNotf(t testing.TB, err, target error, format string, args ...any) {
	t.Helper()
	if errors.Is(err, target) {
		t.Errorf(format, args...)
	}
}

// [ErrorIsNot] asserts that an error is not of a specific type using [ErrorIsNotf] with a default message.
func ErrorIsNot(t testing.TB, err, target error) {
	t.Helper()
	ErrorIsNotf(t, err, target, "expected error %v to not be %v", err, target)
}

// [MapContainsKeyf] asserts that a map contains a specific key. If it does not, it reports an error with the given formatting.
func MapContainsKeyf[K comparable, V any](t testing.TB, m map[K]V, key K, format string, args ...any) {
	t.Helper()
	if _, ok := m[key]; !ok {
		t.Errorf(format, args...)
	}
}

// [MapContainsKey] asserts that a map contains a specific key using [MapContainsKeyf] with a default message.
func MapContainsKey[K comparable, V any](t testing.TB, m map[K]V, key K) {
	t.Helper()
	MapContainsKeyf(t, m, key, "expected map to contain key %v", key)
}

// [MapDoesNotContainKeyf] asserts that a map does not contain a specific key. If it does, it reports an error with the given formatting.
func MapDoesNotContainKeyf[K comparable, V any](t testing.TB, m map[K]V, key K, format string, args ...any) {
	t.Helper()
	if _, ok := m[key]; ok {
		t.Errorf(format, args...)
	}
}

// [MapDoesNotContainKey] asserts that a map does not contain a specific key using [MapDoesNotContainKeyf] with a default message.
func MapDoesNotContainKey[K comparable, V any](t testing.TB, m map[K]V, key K) {
	t.Helper()
	MapDoesNotContainKeyf(t, m, key, "expected map to not contain key %v", key)
}

// [EmptyMapf] asserts that a map is empty. If it is not, it reports an error with the given formatting.
func EmptyMapf[K comparable, V any](t testing.TB, m map[K]V, format string, args ...any) {
	t.Helper()
	if len(m) != 0 {
		t.Errorf(format, args...)
	}
}

// [EmptyMap] asserts that a map is empty using [EmptyMapf] with a default message.
func EmptyMap[K comparable, V any](t testing.TB, m map[K]V) {
	t.Helper()
	EmptyMapf(t, m, "expected map %v to be empty", m)
}

// [NotEmptyMapf] asserts that a map is not empty. If it is, it reports an error with the given formatting.
func NotEmptyMapf[K comparable, V any](t testing.TB, m map[K]V, format string, args ...any) {
	t.Helper()
	if len(m) == 0 {
		t.Errorf(format, args...)
	}
}

// [NotEmptyMap] asserts that a map is not empty using [NotEmptyMapf] with a default message.
func NotEmptyMap[K comparable, V any](t testing.TB, m map[K]V) {
	t.Helper()
	NotEmptyMapf(t, m, "expected map %v to not be empty", m)
}

// [EmptySlicef] asserts that a slice is empty. If it is not, it reports an error with the given formatting.
func EmptySlicef[T any](t testing.TB, slice []T, format string, args ...any) {
	t.Helper()
	if len(slice) != 0 {
		t.Errorf(format, args...)
	}
}

// [EmptySlice] asserts that a slice is empty using [EmptySlicef] with a default message.
func EmptySlice[T any](t testing.TB, slice []T) {
	t.Helper()
	EmptySlicef(t, slice, "expected slice %v to be empty", slice)
}

// [NotEmptySlicef] asserts that a slice is not empty. If it is, it reports an error with the given formatting.
func NotEmptySlicef[T any](t testing.TB, slice []T, format string, args ...any) {
	t.Helper()
	if len(slice) == 0 {
		t.Errorf(format, args...)
	}
}

// [NotEmptySlice] asserts that a slice is not empty using [NotEmptySlicef] with a default message.
func NotEmptySlice[T any](t testing.TB, slice []T) {
	t.Helper()
	NotEmptySlicef(t, slice, "expected slice %v to not be empty", slice)
}

// [DeepEqualf] asserts that two values are equal per [reflect.DeepEqual]. If they are not, it reports an error with the given formatting.
func DeepEqualf[T any](t testing.TB, expected, actual T, format string, args ...any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(format, args...)
	}
}

// [DeepEqual] asserts that two values are equal per [reflect.DeepEqual] using [DeepEqualf] with a default message.
func DeepEqual[T any](t testing.TB, expected, actual T) {
	t.Helper()
	DeepEqualf(t, expected, actual, "expected %v to reflect.DeepEqual %v", actual, expected)
}

// [NotDeepEqualf] asserts that two values are not equal per [reflect.DeepEqual]. If they are, it reports an error with the given formatting.
func NotDeepEqualf[T any](t testing.TB, expected, actual T, format string, args ...any) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		t.Errorf(format, args...)
	}
}

// [NotDeepEqual] asserts that two values are not equal per [reflect.DeepEqual] using [NotDeepEqualf] with a default message.
func NotDeepEqual[T any](t testing.TB, expected, actual T) {
	t.Helper()
	NotDeepEqualf(t, expected, actual, "expected %v to not reflect.DeepEqual %v", actual, expected)
}
