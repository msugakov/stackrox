package predicate

import (
	"reflect"

	"github.com/stackrox/stackrox/pkg/search"
)

// internal predicates

type alwaysTrueIntType struct{}

func (alwaysTrueIntType) Evaluate(value reflect.Value) (*search.Result, bool) {
	return &search.Result{}, true
}

type alwaysFalseIntType struct{}

func (alwaysFalseIntType) Evaluate(value reflect.Value) (*search.Result, bool) {
	return nil, false
}

var (
	alwaysTrue  internalPredicate = alwaysTrueIntType{}
	alwaysFalse internalPredicate = alwaysFalseIntType{}
)

// external predicates

type alwaysTrueType struct{}

func (alwaysTrueType) Evaluate(interface{}) (*search.Result, bool) {
	return &search.Result{}, true
}
func (alwaysTrueType) Matches(interface{}) bool {
	return true
}

type alwaysFalseType struct{}

func (alwaysFalseType) Evaluate(interface{}) (*search.Result, bool) {
	return nil, false
}
func (alwaysFalseType) Matches(interface{}) bool {
	return false
}

var (
	// AlwaysTrue is a predicate that always evaluates to true.
	AlwaysTrue Predicate = alwaysTrueType{}

	// AlwaysFalse is a predicate that always evaluates to false.
	AlwaysFalse Predicate = alwaysFalseType{}
)
