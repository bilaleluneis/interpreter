package common

import (
	"errors"
	"fmt"
	"testing"
)

type customTestError struct {
	code    int
	message string
}

func (e *customTestError) Error() string {
	return fmt.Sprintf("code=%d: %s", e.code, e.message)
}

func TestSomeAndOk(t *testing.T) {
	opt1 := Some(42)
	if !opt1.IsSome() || opt1.IsNone() {
		t.Fatal("expected opt1 to be Some")
	}
	if !opt1.IsOk() || opt1.IsErr() {
		t.Fatal("expected opt1 to be Ok")
	}
	val, ok := opt1.Value()
	if !ok || val != 42 {
		t.Fatalf("expected value 42, got %d (ok=%v)", val, ok)
	}
	if opt1.Error() != nil {
		t.Fatalf("expected nil error on Some, got %v", opt1.Error())
	}

	opt2 := Ok("hello")
	if !opt2.IsSome() {
		t.Fatal("expected opt2 to be Some")
	}
	if v, _ := opt2.Value(); v != "hello" {
		t.Fatalf("expected hello, got %s", v)
	}
}

func TestNoneAndErr(t *testing.T) {
	optNone := None[int]()
	if !optNone.IsNone() || optNone.IsSome() {
		t.Fatal("expected optNone to be None")
	}
	if !optNone.IsErr() || optNone.IsOk() {
		t.Fatal("expected optNone to be Err")
	}
	if _, ok := optNone.Value(); ok {
		t.Fatal("expected Value() ok to be false on None")
	}
	if !errors.Is(optNone.Error(), ErrNone) {
		t.Fatalf("expected ErrNone, got %v", optNone.Error())
	}

	myErr := errors.New("custom error")
	optErr := Err[string](myErr)
	if !optErr.IsNone() {
		t.Fatal("expected optErr to be None")
	}
	if !errors.Is(optErr.Error(), myErr) {
		t.Fatalf("expected myErr, got %v", optErr.Error())
	}

	optErrNil := Err[int](nil)
	if !errors.Is(optErrNil.Error(), ErrNone) {
		t.Fatalf("expected ErrNone when nil error passed, got %v", optErrNil.Error())
	}
}

func TestZeroValueOption(t *testing.T) {
	var opt Option[int]
	if opt.IsSome() {
		t.Fatal("zero value Option should not be Some")
	}
	if !opt.IsNone() {
		t.Fatal("zero value Option should be None")
	}
	if !errors.Is(opt.Error(), ErrNone) {
		t.Fatalf("expected ErrNone on zero value, got %v", opt.Error())
	}
	val, ok := opt.Value()
	if ok || val != 0 {
		t.Fatalf("expected zero value with ok=false, got %d, %v", val, ok)
	}
}

func TestFrom(t *testing.T) {
	res1 := From(100, nil)
	if !res1.IsSome() {
		t.Fatal("expected Some for nil error")
	}
	if v, _ := res1.Value(); v != 100 {
		t.Fatalf("expected 100, got %d", v)
	}

	res2 := From(0, errors.New("fail"))
	if !res2.IsNone() {
		t.Fatal("expected None for non-nil error")
	}
	if res2.Error().Error() != "fail" {
		t.Fatalf("expected fail error, got %v", res2.Error())
	}
}

func TestResult(t *testing.T) {
	val, err := Some(25).Result()
	if err != nil || val != 25 {
		t.Fatalf("expected (25, nil), got (%d, %v)", val, err)
	}

	myErr := errors.New("failed computation")
	val2, err2 := Err[int](myErr).Result()
	if !errors.Is(err2, myErr) || val2 != 0 {
		t.Fatalf("expected (0, myErr), got (%d, %v)", val2, err2)
	}
}

func TestUnwrapAndVariants(t *testing.T) {
	if Some(10).Unwrap() != 10 {
		t.Fatal("expected Unwrap to return 10")
	}
	if Some(10).UnwrapOr(20) != 10 {
		t.Fatal("expected UnwrapOr to return 10")
	}
	if Err[int](errors.New("err")).UnwrapOr(20) != 20 {
		t.Fatal("expected UnwrapOr to return default 20")
	}

	called := false
	val := Err[int](errors.New("test")).UnwrapOrElse(func(err error) int {
		called = true
		return 99
	})
	if !called || val != 99 {
		t.Fatalf("expected UnwrapOrElse to evaluate fallback to 99, got %d", val)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected Unwrap to panic on None")
		}
	}()
	Err[int](errors.New("boom")).Unwrap()
}

func TestMethodChainingSuccessAndFailure(t *testing.T) {
	// Chain happy path
	var loggedSuccess int
	res := Some(5).
		Map(func(v int) int { return v * 2 }).
		AndThen(func(v int) Option[int] {
			return Some(v + 10)
		}).
		Filter(func(v int) bool {
			return v > 15
		}, errors.New("too small")).
		OnSuccess(func(v int) {
			loggedSuccess = v
		}).
		OrElse(func(err error) Option[int] {
			t.Fatal("OrElse should not run on success")
			return Some(0)
		})

	if !res.IsSome() || loggedSuccess != 20 {
		t.Fatalf("expected success with 20, got res=%v, logged=%d", res, loggedSuccess)
	}

	// Chain failure path
	var loggedFailure string
	resErr := Some(5).
		Map(func(v int) int { return v * 2 }).
		AndThen(func(v int) Option[int] {
			return Err[int](&customTestError{code: 404, message: "not found"})
		}).
		Filter(func(v int) bool {
			t.Fatal("Filter should not run on error")
			return true
		}).
		OnSuccess(func(v int) {
			t.Fatal("OnSuccess should not run on failure")
		}).
		OnFailure(func(err error) {
			loggedFailure = err.Error()
		}).
		MapError(func(err error) error {
			return fmt.Errorf("wrapped: %w", err)
		})

	if !resErr.IsNone() {
		t.Fatal("expected failure result")
	}
	if loggedFailure != "code=404: not found" {
		t.Fatalf("expected logged failure, got %s", loggedFailure)
	}

	// Recover path
	recovered := resErr.Recover(func(err error) int {
		return 999
	})
	if !recovered.IsSome() || recovered.Unwrap() != 999 {
		t.Fatalf("expected recovered 999, got %v", recovered)
	}
}

func TestErrorAsGenericAndMethod(t *testing.T) {
	customErr := &customTestError{code: 403, message: "forbidden"}
	opt := Err[int](fmt.Errorf("request failed: %w", customErr))

	// Generic function ErrorAs[E]
	extracted, ok := ErrorAs[*customTestError](opt)
	if !ok {
		t.Fatal("expected ErrorAs to extract *customTestError")
	}
	if extracted.code != 403 || extracted.message != "forbidden" {
		t.Fatalf("unexpected extracted error fields: %+v", extracted)
	}

	// Method AsError
	var target *customTestError
	if !opt.AsError(&target) {
		t.Fatal("expected opt.AsError to succeed")
	}
	if target.code != 403 {
		t.Fatalf("expected code 403, got %d", target.code)
	}

	// Failure case on Some
	someOpt := Some(123)
	if _, ok := ErrorAs[*customTestError](someOpt); ok {
		t.Fatal("ErrorAs on Some should return false")
	}
	if someOpt.AsError(&target) {
		t.Fatal("AsError on Some should return false")
	}
}

func TestTransformFunctions(t *testing.T) {
	// MapOption transforms int to string
	optInt := Some(42)
	optStr := MapOption(optInt, func(v int) string {
		return fmt.Sprintf("val=%d", v)
	})
	if !optStr.IsSome() || optStr.Unwrap() != "val=42" {
		t.Fatalf("expected 'val=42', got %v", optStr)
	}

	// FlatMapOption int to string
	optFlat := FlatMapOption(optInt, func(v int) Option[string] {
		return Some(fmt.Sprintf("flat=%d", v))
	})
	if optFlat.Unwrap() != "flat=42" {
		t.Fatalf("expected 'flat=42', got %v", optFlat)
	}

	// FlatMapOption propagation of error
	errOpt := Err[int](errors.New("upstream err"))
	optFlatErr := FlatMapOption(errOpt, func(v int) Option[string] {
		return Some("won't run")
	})
	if !optFlatErr.IsNone() || optFlatErr.Error().Error() != "upstream err" {
		t.Fatalf("expected error propagated, got %v", optFlatErr)
	}

	// Match
	matchRes := Match(optInt,
		func(v int) string { return fmt.Sprintf("success %d", v) },
		func(err error) string { return "failure" },
	)
	if matchRes != "success 42" {
		t.Fatalf("expected 'success 42', got %s", matchRes)
	}

	matchErrRes := Match(errOpt,
		func(v int) string { return "success" },
		func(err error) string { return fmt.Sprintf("err: %s", err) },
	)
	if matchErrRes != "err: upstream err" {
		t.Fatalf("expected 'err: upstream err', got %s", matchErrRes)
	}
}

func TestFold(t *testing.T) {
	succCalled := false
	Some(1).Fold(
		func(v int) { succCalled = true },
		func(err error) { t.Fatal("onFailure called on Some") },
	)
	if !succCalled {
		t.Fatal("expected onSuccess to be called")
	}

	failCalled := false
	Err[int](errors.New("err")).Fold(
		func(v int) { t.Fatal("onSuccess called on Err") },
		func(err error) { failCalled = true },
	)
	if !failCalled {
		t.Fatal("expected onFailure to be called")
	}
}

func TestUnwrapOrDefaultAndPtr(t *testing.T) {
	optSome := Some(42)
	if optSome.UnwrapOrDefault() != 42 {
		t.Fatalf("expected 42, got %d", optSome.UnwrapOrDefault())
	}
	ptr := optSome.Ptr()
	if ptr == nil || *ptr != 42 {
		t.Fatalf("expected pointer to 42, got %v", ptr)
	}

	optNone := None[int]()
	if optNone.UnwrapOrDefault() != 0 {
		t.Fatalf("expected zero value 0, got %d", optNone.UnwrapOrDefault())
	}
	if optNone.Ptr() != nil {
		t.Fatal("expected nil pointer on None")
	}
}

func TestIsErrorMethod(t *testing.T) {
	errTarget := errors.New("target error")
	optErr := Err[int](fmt.Errorf("wrapped: %w", errTarget))
	if !optErr.Is(errTarget) {
		t.Fatal("expected optErr.Is(errTarget) to be true")
	}
	if optErr.Is(errors.New("other")) {
		t.Fatal("expected false for unmatched error")
	}

	optSome := Some(10)
	if optSome.Is(errTarget) {
		t.Fatal("expected false for Some")
	}
}

func TestFromOkAndFromPtr(t *testing.T) {
	// FromOk
	m := map[string]int{"a": 1}
	val, ok := m["a"]
	opt1 := FromOk(val, ok)
	if !opt1.IsSome() || opt1.Unwrap() != 1 {
		t.Fatalf("expected Some(1), got %v", opt1)
	}

	valMissing, okMissing := m["b"]
	opt2 := FromOk(valMissing, okMissing, errors.New("key missing"))
	if !opt2.IsNone() || opt2.Error().Error() != "key missing" {
		t.Fatalf("expected Err('key missing'), got %v", opt2)
	}

	opt2Default := FromOk(valMissing, okMissing)
	if !opt2Default.IsNone() || !errors.Is(opt2Default.Error(), ErrNone) {
		t.Fatalf("expected ErrNone, got %v", opt2Default.Error())
	}

	// FromPtr
	n := 99
	optPtr := FromPtr(&n)
	if !optPtr.IsSome() || optPtr.Unwrap() != 99 {
		t.Fatalf("expected Some(99), got %v", optPtr)
	}

	var nilPtr *int
	optNilPtr := FromPtr(nilPtr)
	if !optNilPtr.IsNone() || !errors.Is(optNilPtr.Error(), ErrNone) {
		t.Fatalf("expected ErrNone for nil pointer, got %v", optNilPtr.Error())
	}

	customPtrErr := errors.New("null pointer")
	optNilPtrCustom := FromPtr(nilPtr, customPtrErr)
	if !errors.Is(optNilPtrCustom.Error(), customPtrErr) {
		t.Fatalf("expected customPtrErr, got %v", optNilPtrCustom.Error())
	}
}

func TestTryAndTryValue(t *testing.T) {
	// Try successful
	opt1 := Try(func() (int, error) {
		return 123, nil
	})
	if !opt1.IsSome() || opt1.Unwrap() != 123 {
		t.Fatalf("expected Some(123), got %v", opt1)
	}

	// Try with error
	opt2 := Try(func() (int, error) {
		return 0, errors.New("failed call")
	})
	if !opt2.IsNone() || opt2.Error().Error() != "failed call" {
		t.Fatalf("expected error 'failed call', got %v", opt2)
	}

	// Try catching panic with error
	opt3 := Try(func() (int, error) {
		panic(errors.New("panic error"))
	})
	if !opt3.IsNone() || opt3.Error().Error() != "panic error" {
		t.Fatalf("expected caught panic error, got %v", opt3)
	}

	// Try catching string panic
	opt4 := Try(func() (int, error) {
		panic("boom")
	})
	if !opt4.IsNone() || opt4.Error().Error() != "panic: boom" {
		t.Fatalf("expected caught panic 'panic: boom', got %v", opt4)
	}

	// TryValue
	optVal := TryValue(func() string {
		return "hello"
	})
	if !optVal.IsSome() || optVal.Unwrap() != "hello" {
		t.Fatalf("expected 'hello', got %v", optVal)
	}

	optValPanic := TryValue(func() string {
		panic("value panic")
	})
	if !optValPanic.IsNone() || optValPanic.Error().Error() != "panic: value panic" {
		t.Fatalf("expected caught panic, got %v", optValPanic)
	}
}

func TestThenAndCheckChaining(t *testing.T) {
	// Simulate multi-step pipeline replacing "if err != nil"
	stepDouble := func(n int) (int, error) {
		return n * 2, nil
	}
	stepAddTen := func(n int) (int, error) {
		return n + 10, nil
	}
	checkPositive := func(n int) error {
		if n <= 0 {
			return errors.New("must be positive")
		}
		return nil
	}
	stepFail := func(n int) (int, error) {
		return 0, errors.New("step error")
	}

	// Success chain in a single clean one-liner
	res, err := Ok(5).
		Then(stepDouble).
		Check(checkPositive).
		Then(stepAddTen).
		Result()

	if err != nil || res != 20 {
		t.Fatalf("expected 20 and nil error, got res=%d, err=%v", res, err)
	}

	// Failure at Check short-circuits subsequent steps
	stepAfterCheckCalled := false
	resCheckErr, errCheck := Ok(-5).
		Check(checkPositive).
		Then(func(n int) (int, error) {
			stepAfterCheckCalled = true
			return n * 2, nil
		}).
		Result()

	if errCheck == nil || errCheck.Error() != "must be positive" {
		t.Fatalf("expected 'must be positive', got %v", errCheck)
	}
	if stepAfterCheckCalled {
		t.Fatal("step after failed Check should not have been called")
	}
	if resCheckErr != 0 {
		t.Fatalf("expected 0 for failed result, got %d", resCheckErr)
	}

	// Failure at Then step propagates error
	_, errThen := Ok(5).
		Then(stepFail).
		Then(stepAddTen).
		Result()

	if errThen == nil || errThen.Error() != "step error" {
		t.Fatalf("expected 'step error', got %v", errThen)
	}
}

func TestThenOk(t *testing.T) {
	m := map[int]int{2: 200, 4: 400}
	lookup := func(k int) (int, bool) {
		v, ok := m[k]
		return v, ok
	}

	// Success
	res := Ok(2).ThenOk(lookup).Unwrap()
	if res != 200 {
		t.Fatalf("expected 200, got %d", res)
	}

	// Failure with custom error
	notFoundErr := errors.New("id not found")
	resErr := Ok(99).ThenOk(lookup, notFoundErr)
	if !errors.Is(resErr.Error(), notFoundErr) {
		t.Fatalf("expected notFoundErr, got %v", resErr.Error())
	}
}

func TestTapAndTapError(t *testing.T) {
	var tappedVal int
	var tappedErr string

	Ok(42).
		Tap(func(v int) { tappedVal = v }).
		TapError(func(err error) { t.Fatal("TapError should not run on Ok") })

	if tappedVal != 42 {
		t.Fatalf("expected tappedVal 42, got %d", tappedVal)
	}

	Err[int](errors.New("err tapped")).
		Tap(func(v int) { t.Fatal("Tap should not run on Err") }).
		TapError(func(err error) { tappedErr = err.Error() })

	if tappedErr != "err tapped" {
		t.Fatalf("expected 'err tapped', got %s", tappedErr)
	}
}

func TestOrElseValueAndRecoverWith(t *testing.T) {
	// OrElseValue
	val := None[int]().OrElseValue(100).Unwrap()
	if val != 100 {
		t.Fatalf("expected 100, got %d", val)
	}

	valSome := Some(50).OrElseValue(100).Unwrap()
	if valSome != 50 {
		t.Fatalf("expected 50, got %d", valSome)
	}

	// RecoverWith success
	recovered := Err[int](errors.New("fail")).RecoverWith(func(err error) (int, error) {
		return 777, nil
	})
	if recovered.Unwrap() != 777 {
		t.Fatalf("expected 777, got %d", recovered.Unwrap())
	}

	// RecoverWith failure
	recovFailErr := errors.New("recovery failed too")
	recovFail := Err[int](errors.New("fail")).RecoverWith(func(err error) (int, error) {
		return 0, recovFailErr
	})
	if !errors.Is(recovFail.Error(), recovFailErr) {
		t.Fatalf("expected recovFailErr, got %v", recovFail.Error())
	}
}

func TestCrossTypeChainingFunctions(t *testing.T) {
	// Then[T, U]
	parseInt := func(s string) (int, error) {
		if s == "invalid" {
			return 0, errors.New("not an int")
		}
		var n int
		_, err := fmt.Sscanf(s, "%d", &n)
		return n, err
	}

	optStr := Ok("123")
	optInt := Then(optStr, parseInt)
	if !optInt.IsSome() || optInt.Unwrap() != 123 {
		t.Fatalf("expected 123, got %v", optInt)
	}

	optErr := Then(Ok("invalid"), parseInt)
	if !optErr.IsNone() {
		t.Fatal("expected failure on invalid")
	}

	// Zip and ZipWith
	optA := Ok(10)
	optB := Ok("stars")
	zipped := Zip(optA, optB)
	if !zipped.IsSome() || zipped.Unwrap().First != 10 || zipped.Unwrap().Second != "stars" {
		t.Fatalf("unexpected zipped result: %+v", zipped)
	}

	zipFormatted := ZipWith(optA, optB, func(a int, b string) string {
		return fmt.Sprintf("%d %s", a, b)
	})
	if zipFormatted.Unwrap() != "10 stars" {
		t.Fatalf("expected '10 stars', got %s", zipFormatted.Unwrap())
	}

	zipFail := Zip(Err[int](errors.New("failA")), optB)
	if !zipFail.IsNone() || zipFail.Error().Error() != "failA" {
		t.Fatalf("expected failA, got %v", zipFail)
	}
}

type mutableData struct {
	Items []string
	Meta  map[string]int
}

func TestImmutabilityProtectionWithPointersAndReferences(t *testing.T) {
	// 1. Mutating input slice/map after Some does not affect Option
	origSlice := []string{"one", "two"}
	origMap := map[string]int{"count": 1}
	d := mutableData{
		Items: origSlice,
		Meta:  origMap,
	}

	opt := Some(d)

	// Mutate originals
	origSlice[0] = "mutated"
	origMap["count"] = 999

	stored := opt.Unwrap()
	if stored.Items[0] != "one" {
		t.Fatalf("expected stored slice to be isolated from input mutation, got %s", stored.Items[0])
	}
	if stored.Meta["count"] != 1 {
		t.Fatalf("expected stored map to be isolated from input mutation, got %d", stored.Meta["count"])
	}

	// 2. Mutating returned value from Unwrap/Value/Result does not affect Option
	unwrapped := opt.Unwrap()
	unwrapped.Items[0] = "tampered"
	unwrapped.Meta["count"] = 888

	storedAgain, _ := opt.Value()
	if storedAgain.Items[0] != "one" {
		t.Fatalf("expected stored value to remain unchanged after modifying unwrapped copy, got %s", storedAgain.Items[0])
	}
	if storedAgain.Meta["count"] != 1 {
		t.Fatalf("expected stored map to remain unchanged after modifying unwrapped copy, got %d", storedAgain.Meta["count"])
	}

	// 3. Mutating value from Ptr does not affect Option
	p := opt.Ptr()
	p.Items[0] = "ptr_tampered"

	storedAfterPtr := opt.Unwrap()
	if storedAfterPtr.Items[0] != "one" {
		t.Fatalf("expected stored value to be unaffected by mutating *Ptr(), got %s", storedAfterPtr.Items[0])
	}

	// 4. Tap receives isolated copy
	opt.Tap(func(v mutableData) {
		v.Items[0] = "tapped_mutation"
	})
	if opt.Unwrap().Items[0] != "one" {
		t.Fatalf("expected stored value to be unaffected by mutations inside Tap")
	}

	// 5. Fold receives isolated copy
	opt.Fold(func(v mutableData) {
		v.Items[0] = "fold_mutation"
	}, func(err error) {})
	if opt.Unwrap().Items[0] != "one" {
		t.Fatalf("expected stored value to be unaffected by mutations inside Fold")
	}
}
