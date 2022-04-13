// Code generated by genny. DO NOT EDIT.
// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package protoutils

import "github.com/stackrox/stackrox/generated/storage"

// *storage.Alert_Violation represents a generic type that we use in the function below.

// EqualStorageAlert_ViolationSlices returns whether the given two slices of proto objects (generically) have equal values.
func EqualStorageAlert_ViolationSlices(first, second []*storage.Alert_Violation) bool {
	if len(first) != len(second) {
		return false
	}
	for i, firstElem := range first {
		secondElem := second[i]
		if !protoEqualWrapper(firstElem, secondElem) {
			return false
		}
	}
	return true
}
