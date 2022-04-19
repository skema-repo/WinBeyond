// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: RRXX.proto

package xxx

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on HealthcheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthcheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthcheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthcheckRequestMultiError, or nil if none found.
func (m *HealthcheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthcheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return HealthcheckRequestMultiError(errors)
	}
	return nil
}

// HealthcheckRequestMultiError is an error wrapping multiple validation errors
// returned by HealthcheckRequest.ValidateAll() if the designated constraints
// aren't met.
type HealthcheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthcheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthcheckRequestMultiError) AllErrors() []error { return m }

// HealthcheckRequestValidationError is the validation error returned by
// HealthcheckRequest.Validate if the designated constraints aren't met.
type HealthcheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthcheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthcheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthcheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthcheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthcheckRequestValidationError) ErrorName() string {
	return "HealthcheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HealthcheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthcheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthcheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthcheckRequestValidationError{}

// Validate checks the field values on HealthcheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthcheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthcheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthcheckResponseMultiError, or nil if none found.
func (m *HealthcheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthcheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return HealthcheckResponseMultiError(errors)
	}
	return nil
}

// HealthcheckResponseMultiError is an error wrapping multiple validation
// errors returned by HealthcheckResponse.ValidateAll() if the designated
// constraints aren't met.
type HealthcheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthcheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthcheckResponseMultiError) AllErrors() []error { return m }

// HealthcheckResponseValidationError is the validation error returned by
// HealthcheckResponse.Validate if the designated constraints aren't met.
type HealthcheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthcheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthcheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthcheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthcheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthcheckResponseValidationError) ErrorName() string {
	return "HealthcheckResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthcheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthcheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthcheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthcheckResponseValidationError{}

// Validate checks the field values on HelloRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloRequestMultiError, or
// nil if none found.
func (m *HelloRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMsg()) < 3 {
		err := HelloRequestValidationError{
			field:  "Msg",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HelloRequestMultiError(errors)
	}
	return nil
}

// HelloRequestMultiError is an error wrapping multiple validation errors
// returned by HelloRequest.ValidateAll() if the designated constraints aren't met.
type HelloRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloRequestMultiError) AllErrors() []error { return m }

// HelloRequestValidationError is the validation error returned by
// HelloRequest.Validate if the designated constraints aren't met.
type HelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloRequestValidationError) ErrorName() string { return "HelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e HelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloRequestValidationError{}

// Validate checks the field values on HelloReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloReplyMultiError, or
// nil if none found.
func (m *HelloReply) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	// no validation rules for Code

	if len(errors) > 0 {
		return HelloReplyMultiError(errors)
	}
	return nil
}

// HelloReplyMultiError is an error wrapping multiple validation errors
// returned by HelloReply.ValidateAll() if the designated constraints aren't met.
type HelloReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloReplyMultiError) AllErrors() []error { return m }

// HelloReplyValidationError is the validation error returned by
// HelloReply.Validate if the designated constraints aren't met.
type HelloReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloReplyValidationError) ErrorName() string { return "HelloReplyValidationError" }

// Error satisfies the builtin error interface
func (e HelloReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloReplyValidationError{}
