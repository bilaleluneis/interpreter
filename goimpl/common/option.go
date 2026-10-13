// Package common provides core shared constraints and monadic types.
package common

import (
	"errors"
	"fmt"
)

var (
	// ErrNone represents an empty Option without an explicit underlying error.
	ErrNone = errors.New("option is empty")
	// ErrPredicateFailed represents a failure during Filter when predicate returns false.
	ErrPredicateFailed = errors.New("predicate check failed")
)

// Option represents an optional value of type T, which either contains
// a value of type T (Some/Ok) or an error implementing Go's error interface (None/Err).
// It enables clean, fluent method chaining as an alternative to verbose "if err != nil" checks.
type Option[T any] struct {
	value    T
	err      error
	hasValue bool
}

// Some creates an Option containing a successful value of type T.
// The value is stored directly without additional cloning.
func Some[T any](value T) Option[T] {
	return Option[T]{
		value:    value,
		hasValue: true,
	}
}

// Ok is an alias for Some to align with Result / Ok conventions.
func Ok[T any](value T) Option[T] {
	return Some(value)
}

// None creates an Option representing a failure with an optional error.
// If no error is supplied or err is nil, ErrNone is used by default.
func None[T any](err ...error) Option[T] {
	var e error
	if len(err) > 0 && err[0] != nil {
		e = err[0]
	} else {
		e = ErrNone
	}
	return Option[T]{
		err:      e,
		hasValue: false,
	}
}

// Err creates an Option representing a failure with the given error.
// If err is nil, ErrNone is used.
func Err[T any](err error) Option[T] {
	if err == nil {
		err = ErrNone
	}
	return Option[T]{
		err:      err,
		hasValue: false,
	}
}

// From constructs an Option from standard Go (value, error) pairs.
// If err is non-nil, an Err Option is returned; otherwise Some(value) is returned.
func From[T any](value T, err error) Option[T] {
	if err != nil {
		return Err[T](err)
	}
	return Some(value)
}

// FromOk constructs an Option from standard Go (value, bool) comma-ok pairs.
// If ok is false, an Err Option with the given error (or ErrNone) is returned.
func FromOk[T any](value T, ok bool, err ...error) Option[T] {
	if !ok {
		if len(err) > 0 && err[0] != nil {
			return Err[T](err[0])
		}
		return Err[T](ErrNone)
	}
	return Some(value)
}

// FromPtr constructs an Option from a pointer.
// If ptr is nil, an Err Option with the given error (or ErrNone) is returned.
// If ptr is non-nil, Some(*ptr) is returned.
func FromPtr[T any](ptr *T, err ...error) Option[T] {
	if ptr == nil {
		if len(err) > 0 && err[0] != nil {
			return Err[T](err[0])
		}
		return Err[T](ErrNone)
	}
	return Some(*ptr)
}

// Try executes fn and captures both returned (T, error) and any panics as an Option[T].
func Try[T any](fn func() (T, error)) (opt Option[T]) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				opt = Err[T](e)
			} else {
				opt = Err[T](fmt.Errorf("panic: %v", r))
			}
		}
	}()
	val, err := fn()
	return From(val, err)
}

// TryValue executes fn and captures any panic as an Option[T].
func TryValue[T any](fn func() T) (opt Option[T]) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				opt = Err[T](e)
			} else {
				opt = Err[T](fmt.Errorf("panic: %v", r))
			}
		}
	}()
	return Some(fn())
}

// IsSome returns true if the Option contains a value.
func (o Option[T]) IsSome() bool {
	return o.hasValue
}

// IsNone returns true if the Option represents a failure / empty state.
func (o Option[T]) IsNone() bool {
	return !o.hasValue
}

// IsOk is an alias for IsSome.
func (o Option[T]) IsOk() bool {
	return o.hasValue
}

// IsErr is an alias for IsNone.
func (o Option[T]) IsErr() bool {
	return !o.hasValue
}

// Value returns the contained value and true if present, or zero value of T and false.
// The returned value is a direct copy of the stored value.
func (o Option[T]) Value() (T, bool) {
	if !o.hasValue {
		var zero T
		return zero, false
	}
	return o.value, true
}

// Error returns the underlying error if None/Err, or nil if Some/Ok.
// If the Option was default-initialized as zero-value, ErrNone is returned.
func (o Option[T]) Error() error {
	if o.hasValue {
		return nil
	}
	if o.err == nil {
		return ErrNone
	}
	return o.err
}

// Result unpacks the Option into a standard Go (T, error) tuple.
// The returned value is a direct copy of the stored value.
func (o Option[T]) Result() (T, error) {
	if o.hasValue {
		return o.value, nil
	}
	var zero T
	return zero, o.Error()
}

// Unwrap returns the value if present, or panics with the error.
// The returned value is a direct copy of the stored value.
func (o Option[T]) Unwrap() T {
	if !o.hasValue {
		panic(fmt.Sprintf("called Unwrap() on None Option: %v", o.Error()))
	}
	return o.value
}

// UnwrapOr returns the value if present, or defaultValue if None.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if !o.hasValue {
		return defaultValue
	}
	return o.value
}

// UnwrapOrElse returns the value if present, or evaluates fn(err) if None.
func (o Option[T]) UnwrapOrElse(fn func(error) T) T {
	if !o.hasValue {
		return fn(o.Error())
	}
	return o.value
}

// UnwrapOrDefault returns the contained value if present,
// or the zero value of T if None.
func (o Option[T]) UnwrapOrDefault() T {
	if !o.hasValue {
		var zero T
		return zero
	}
	return o.value
}

// Ptr returns a pointer to the contained value if present, or nil if None.
// The pointer references the receiver copy's stored value.
func (o Option[T]) Ptr() *T {
	if !o.hasValue {
		return nil
	}
	return &o.value
}

// Is reports whether any error in the Option's error tree matches target using errors.Is.
// If the Option is Some, it returns false.
func (o Option[T]) Is(target error) bool {
	if o.hasValue {
		return false
	}
	return errors.Is(o.Error(), target)
}

// AsError checks whether the Option's error can be cast to target using errors.As.
// target must be a non-nil pointer to a concrete error type or error interface.
func (o Option[T]) AsError(target any) bool {
	if o.hasValue {
		return false
	}
	return errors.As(o.Error(), target)
}

// Then chains a fallible operation returning (U, error).
// If the Option is Some, fn is invoked with the contained value.
// If fn returns a non-nil error, an Err Option is returned.
// If the Option is None, the failure is propagated without invoking fn.
// This replaces verbose "if err != nil" checks with a clean fluent call.
func (o Option[T]) Then[U any](fn func(T) (U, error)) Option[U] {
	if !o.hasValue {
		return Err[U](o.Error())
	}
	val, err := fn(o.value)
	if err != nil {
		return Err[U](err)
	}
	return Some(val)
}

// ThenOk chains an operation returning (U, bool), such as a map lookup or lookup function.
// If fn returns false, it returns an Err with the provided error (or ErrNone).
func (o Option[T]) ThenOk[U any](fn func(T) (U, bool), err ...error) Option[U] {
	if !o.hasValue {
		return Err[U](o.Error())
	}
	val, ok := fn(o.value)
	if !ok {
		if len(err) > 0 && err[0] != nil {
			return Err[U](err[0])
		}
		return Err[U](ErrNone)
	}
	return Some(val)
}

// Check validates the contained value with fn.
// If the Option is Some and fn returns a non-nil error, an Err Option is returned.
// If fn returns nil, the Option is returned unchanged.
// If the Option is None, fn is not called and the failure is propagated.
func (o Option[T]) Check(fn func(T) error) Option[T] {
	if !o.hasValue {
		return o
	}
	if err := fn(o.value); err != nil {
		return Err[T](err)
	}
	return o
}

// FlatMap chains operations on the success path.
// If the Option is Some, fn is invoked with the value.
// If the Option is None, the failure is propagated unchanged.
func (o Option[T]) FlatMap[U any](fn func(T) Option[U]) Option[U] {
	if !o.hasValue {
		return Err[U](o.Error())
	}
	return fn(o.value)
}

// AndThen is an alias for FlatMap.
func (o Option[T]) AndThen[U any](fn func(T) Option[U]) Option[U] {
	return o.FlatMap(fn)
}

// OrElse chains operations on the failure path.
// If the Option is None, fn is invoked with the error.
// If the Option is Some, it is returned unchanged.
func (o Option[T]) OrElse(fn func(error) Option[T]) Option[T] {
	if o.hasValue {
		return o
	}
	return fn(o.Error())
}

// OrElseValue provides an alternative value if the Option is None.
func (o Option[T]) OrElseValue(fallback T) Option[T] {
	if !o.hasValue {
		return Some(fallback)
	}
	return o
}

// Map applies fn to the contained value if Some, returning Some(fn(value)).
// If None, the failure is propagated unchanged.
func (o Option[T]) Map[U any](fn func(T) U) Option[U] {
	if !o.hasValue {
		return Err[U](o.Error())
	}
	return Some(fn(o.value))
}

// MapError applies fn to the underlying error if None.
// If Some, it is returned unchanged.
func (o Option[T]) MapError(fn func(error) error) Option[T] {
	if o.hasValue {
		return o
	}
	return Err[T](fn(o.Error()))
}

// Filter evaluates predicate on the value if Some.
// If predicate returns false,
// it returns an Err with the provided error (or ErrPredicateFailed).
func (o Option[T]) Filter(predicate func(T) bool, err ...error) Option[T] {
	if !o.hasValue {
		return o
	}
	if predicate(o.value) {
		return o
	}
	if len(err) > 0 && err[0] != nil {
		return Err[T](err[0])
	}
	return Err[T](ErrPredicateFailed)
}

// Tap executes action if Some, and returns self for method chaining.
// The stored value is passed directly to action.
func (o Option[T]) Tap(action func(T)) Option[T] {
	if o.hasValue {
		action(o.value)
	}
	return o
}

// TapError executes action if None, and returns self for method chaining.
func (o Option[T]) TapError(action func(error)) Option[T] {
	if !o.hasValue {
		action(o.Error())
	}
	return o
}

// OnSuccess is an alias for Tap.
func (o Option[T]) OnSuccess(action func(T)) Option[T] {
	return o.Tap(action)
}

// OnFailure is an alias for TapError.
func (o Option[T]) OnFailure(action func(error)) Option[T] {
	return o.TapError(action)
}

// Recover recovers from a failure by producing a value of type T from the error.
func (o Option[T]) Recover(fn func(error) T) Option[T] {
	if !o.hasValue {
		return Some(fn(o.Error()))
	}
	return o
}

// RecoverWith recovers from a failure by invoking fn(err), which returns (T, error).
// If fn succeeds, Some(val) is returned; if fn fails, Err(err) is returned.
// If the Option is Some, it is returned unchanged.
func (o Option[T]) RecoverWith(fn func(error) (T, error)) Option[T] {
	if !o.hasValue {
		val, err := fn(o.Error())
		if err != nil {
			return Err[T](err)
		}
		return Some(val)
	}
	return o
}

// Fold executes onSuccess if Some, or onFailure if None.
// The stored value is passed directly to onSuccess.
func (o Option[T]) Fold(onSuccess func(T), onFailure func(error)) {
	if o.hasValue {
		onSuccess(o.value)
	} else {
		onFailure(o.Error())
	}
}

// Match executes onSuccess or onFailure and returns the computed result R.
// The stored value is passed directly to onSuccess.
func (o Option[T]) Match[R any](onSuccess func(T) R, onFailure func(error) R) R {
	if o.hasValue {
		return onSuccess(o.value)
	}
	return onFailure(o.Error())
}

// ZipWith combines two options using fn into an Option[R].
// If the receiver is None, its error is returned; otherwise if other is None, other's error is returned.
func (o Option[T]) ZipWith[U any, R any](other Option[U], fn func(T, U) R) Option[R] {
	if !o.hasValue {
		return Err[R](o.Error())
	}
	if !other.hasValue {
		return Err[R](other.Error())
	}
	return Some(fn(o.value, other.value))
}
