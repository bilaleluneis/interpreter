// Package common provides core shared constraints, immutability helpers, and monadic types.
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

// Option represents an immutable optional value of type T, which either contains
// a value of type T (Some/Ok) or an error implementing Go's error interface (None/Err).
// It enables clean, fluent method chaining as an alternative to verbose "if err != nil" checks.
type Option[T any] struct {
	value    T
	err      error
	hasValue bool
}

// Some creates an Option containing a successful value of type T.
// A deep copy of value is made when possible to guarantee value isolation and immutability.
// If DeepCopy fails, an Err Option with the failure error is returned.
func Some[T any](value T) Option[T] {
	copied, err := DeepCopy(value)
	if err != nil {
		return Err[T](err)
	}
	return Option[T]{
		value:    copied,
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
// Returns a defensive deep copy when possible to preserve value semantics.
func (o Option[T]) Value() (T, bool) {
	if !o.hasValue {
		var zero T
		return zero, false
	}
	copied, err := DeepCopy(o.value)
	if err != nil {
		var zero T
		return zero, false
	}
	return copied, true
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
// Returns a defensive deep copy when possible to prevent external mutation of internal state.
func (o Option[T]) Result() (T, error) {
	if o.hasValue {
		copied, err := DeepCopy(o.value)
		if err != nil {
			var zero T
			return zero, err
		}
		return copied, nil
	}
	var zero T
	return zero, o.Error()
}

// Unwrap returns the value if present, or panics with the error.
// Returns a defensive deep copy when possible to preserve immutability.
func (o Option[T]) Unwrap() T {
	if !o.hasValue {
		panic(fmt.Sprintf("called Unwrap() on None Option: %v", o.Error()))
	}
	copied, err := DeepCopy(o.value)
	if err != nil {
		panic(fmt.Sprintf("called Unwrap() failed to copy value: %v", err))
	}
	return copied
}

// UnwrapOr returns the value if present, or defaultValue if None or if copying fails.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if !o.hasValue {
		return defaultValue
	}
	copied, err := DeepCopy(o.value)
	if err != nil {
		return defaultValue
	}
	return copied
}

// UnwrapOrElse returns the value if present, or evaluates fn(err) if None or if copying fails.
func (o Option[T]) UnwrapOrElse(fn func(error) T) T {
	if !o.hasValue {
		return fn(o.Error())
	}
	copied, err := DeepCopy(o.value)
	if err != nil {
		return fn(err)
	}
	return copied
}

// UnwrapOrDefault returns the contained value if present,
// or the zero value of T if None or if copying fails.
func (o Option[T]) UnwrapOrDefault() T {
	if !o.hasValue {
		var zero T
		return zero
	}
	copied, err := DeepCopy(o.value)
	if err != nil {
		var zero T
		return zero
	}
	return copied
}

// Ptr returns a pointer to a copy of the value if present, or nil if None or if copying fails.
// It returns an independent copy so mutating the referenced value does not affect the Option.
func (o Option[T]) Ptr() *T {
	if !o.hasValue {
		return nil
	}
	copied, err := DeepCopy(o.value)
	if err != nil {
		return nil
	}
	return &copied
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

// Then chains a fallible operation returning (T, error).
// If the Option is Some, fn is invoked with the contained value.
// If fn returns a non-nil error, an Err Option is returned.
// If the Option is None, the failure is propagated without invoking fn.
// This replaces verbose "if err != nil" checks with a clean fluent call.
func (o Option[T]) Then(fn func(T) (T, error)) Option[T] {
	if !o.hasValue {
		return o
	}
	val, err := fn(o.value)
	if err != nil {
		return Err[T](err)
	}
	return Some(val)
}

// ThenOk chains an operation returning (T, bool), such as a map lookup or lookup function.
// If fn returns false, it returns an Err with the provided error (or ErrNone).
func (o Option[T]) ThenOk(fn func(T) (T, bool), err ...error) Option[T] {
	if !o.hasValue {
		return o
	}
	val, ok := fn(o.value)
	if !ok {
		if len(err) > 0 && err[0] != nil {
			return Err[T](err[0])
		}
		return Err[T](ErrNone)
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

// AndThen chains operations on the success path.
// If the Option is Some, fn is invoked with the value.
// If the Option is None, the failure is propagated unchanged.
func (o Option[T]) AndThen(fn func(T) Option[T]) Option[T] {
	if !o.hasValue {
		return o
	}
	return fn(o.value)
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
func (o Option[T]) Map(fn func(T) T) Option[T] {
	if !o.hasValue {
		return o
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
// A defensive copy of the value is provided to action
// so side effects cannot mutate internal state.
func (o Option[T]) Tap(action func(T)) Option[T] {
	if o.hasValue {
		if copied, err := DeepCopy(o.value); err == nil {
			action(copied)
		}
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
// A defensive copy of the value is provided to onSuccess
// so external callers cannot mutate internal state.
func (o Option[T]) Fold(onSuccess func(T), onFailure func(error)) {
	if o.hasValue {
		copied, err := DeepCopy(o.value)
		if err != nil {
			onFailure(err)
		} else {
			onSuccess(copied)
		}
	} else {
		onFailure(o.Error())
	}
}

// Pair holds two values of types T and U.
type Pair[T any, U any] struct {
	First  T
	Second U
}

// Then chains a fallible operation transforming T into U via (U, error).
// If opt is Some, fn is invoked with the contained value.
// If opt is None, the failure is propagated.
func Then[T any, U any](opt Option[T], fn func(T) (U, error)) Option[U] {
	if !opt.hasValue {
		return Err[U](opt.Error())
	}
	val, err := fn(opt.value)
	if err != nil {
		return Err[U](err)
	}
	return Some(val)
}

// Map transforms an Option[T] into Option[U] by applying fn to the contained value.
func Map[T any, U any](opt Option[T], fn func(T) U) Option[U] {
	if !opt.hasValue {
		return Err[U](opt.Error())
	}
	return Some(fn(opt.value))
}

// FlatMap transforms an Option[T] into Option[U] by applying fn to the contained value.
func FlatMap[T any, U any](opt Option[T], fn func(T) Option[U]) Option[U] {
	if !opt.hasValue {
		return Err[U](opt.Error())
	}
	return fn(opt.value)
}

// AndThen is an alias for FlatMap.
func AndThen[T any, U any](opt Option[T], fn func(T) Option[U]) Option[U] {
	return FlatMap(opt, fn)
}

// MapOption is an alias for Map.
func MapOption[T any, U any](opt Option[T], fn func(T) U) Option[U] {
	return Map(opt, fn)
}

// FlatMapOption is an alias for FlatMap.
func FlatMapOption[T any, U any](opt Option[T], fn func(T) Option[U]) Option[U] {
	return FlatMap(opt, fn)
}

// Match executes onSuccess or onFailure and returns the computed result R.
// A defensive copy of the value is provided to onSuccess to preserve immutability.
func Match[T any, R any](opt Option[T], onSuccess func(T) R, onFailure func(error) R) R {
	if opt.hasValue {
		copied, err := DeepCopy(opt.value)
		if err != nil {
			return onFailure(err)
		}
		return onSuccess(copied)
	}
	return onFailure(opt.Error())
}

// Zip combines two options into an Option[Pair[T, U]].
// If either option is None, the first encountered error is returned.
func Zip[T any, U any](optA Option[T], optB Option[U]) Option[Pair[T, U]] {
	return ZipWith(optA, optB, func(a T, b U) Pair[T, U] {
		return Pair[T, U]{First: a, Second: b}
	})
}

// ZipWith combines two options using fn into an Option[R].
// If optA is None, optA's error is returned; otherwise if optB is None, optB's error is returned.
func ZipWith[T any, U any, R any](optA Option[T], optB Option[U], fn func(T, U) R) Option[R] {
	if !optA.hasValue {
		return Err[R](optA.Error())
	}
	if !optB.hasValue {
		return Err[R](optB.Error())
	}
	return Some(fn(optA.value, optB.value))
}

// ErrorAs extracts a concrete error type E from the Option's error.
// Returns the concrete error and true if matched, or the zero value of E and false otherwise.
func ErrorAs[E error, T any](opt Option[T]) (E, bool) {
	var target E
	if opt.IsSome() {
		return target, false
	}
	err := opt.Error()
	if err == nil {
		return target, false
	}
	if concrete, ok := err.(E); ok {
		return concrete, true
	}
	if errors.As(err, &target) {
		return target, true
	}
	return target, false
}
