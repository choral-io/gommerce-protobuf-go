// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gommerce/iam/v1beta/users.proto

package iam_v1beta

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

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Realm

	// no validation rules for Disabled

	// no validation rules for Approved

	// no validation rules for Verified

	// no validation rules for Immutable

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDeletedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "DeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExpiresAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiresAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "ExpiresAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFirstLoginTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "FirstLoginTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "FirstLoginTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFirstLoginTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "FirstLoginTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastActiveTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "LastActiveTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "LastActiveTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastActiveTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "LastActiveTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Flags

	// no validation rules for Attributes

	if all {
		switch v := interface{}(m.GetDisplayName()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "DisplayName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "DisplayName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDisplayName()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "DisplayName",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDescription()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "Description",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "Description",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDescription()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "Description",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMaskedPhoneNumber()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "MaskedPhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "MaskedPhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaskedPhoneNumber()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "MaskedPhoneNumber",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMaskedEmailAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "MaskedEmailAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "MaskedEmailAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaskedEmailAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "MaskedEmailAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "Creator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersRequestMultiError, or nil if none found.
func (m *ListUsersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Size

	// no validation rules for Sort

	if len(errors) > 0 {
		return ListUsersRequestMultiError(errors)
	}

	return nil
}

// ListUsersRequestMultiError is an error wrapping multiple validation errors
// returned by ListUsersRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUsersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersRequestMultiError) AllErrors() []error { return m }

// ListUsersRequestValidationError is the validation error returned by
// ListUsersRequest.Validate if the designated constraints aren't met.
type ListUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersRequestValidationError) ErrorName() string { return "ListUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersRequestValidationError{}

// Validate checks the field values on ListUsersResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersResponseMultiError, or nil if none found.
func (m *ListUsersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUsersResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Page

	// no validation rules for Size

	// no validation rules for Total

	if len(errors) > 0 {
		return ListUsersResponseMultiError(errors)
	}

	return nil
}

// ListUsersResponseMultiError is an error wrapping multiple validation errors
// returned by ListUsersResponse.ValidateAll() if the designated constraints
// aren't met.
type ListUsersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersResponseMultiError) AllErrors() []error { return m }

// ListUsersResponseValidationError is the validation error returned by
// ListUsersResponse.Validate if the designated constraints aren't met.
type ListUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersResponseValidationError) ErrorName() string {
	return "ListUsersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersResponseValidationError{}

// Validate checks the field values on GetIdentityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetIdentityRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetIdentityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetIdentityRequestMultiError, or nil if none found.
func (m *GetIdentityRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetIdentityRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetIdentityRequestMultiError(errors)
	}

	return nil
}

// GetIdentityRequestMultiError is an error wrapping multiple validation errors
// returned by GetIdentityRequest.ValidateAll() if the designated constraints
// aren't met.
type GetIdentityRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetIdentityRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetIdentityRequestMultiError) AllErrors() []error { return m }

// GetIdentityRequestValidationError is the validation error returned by
// GetIdentityRequest.Validate if the designated constraints aren't met.
type GetIdentityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetIdentityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetIdentityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetIdentityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetIdentityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetIdentityRequestValidationError) ErrorName() string {
	return "GetIdentityRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetIdentityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetIdentityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetIdentityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetIdentityRequestValidationError{}

// Validate checks the field values on GetIdentityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetIdentityResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetIdentityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetIdentityResponseMultiError, or nil if none found.
func (m *GetIdentityResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetIdentityResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetIdentityResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetIdentityResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetIdentityResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetIdentityResponseMultiError(errors)
	}

	return nil
}

// GetIdentityResponseMultiError is an error wrapping multiple validation
// errors returned by GetIdentityResponse.ValidateAll() if the designated
// constraints aren't met.
type GetIdentityResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetIdentityResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetIdentityResponseMultiError) AllErrors() []error { return m }

// GetIdentityResponseValidationError is the validation error returned by
// GetIdentityResponse.Validate if the designated constraints aren't met.
type GetIdentityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetIdentityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetIdentityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetIdentityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetIdentityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetIdentityResponseValidationError) ErrorName() string {
	return "GetIdentityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetIdentityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetIdentityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetIdentityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetIdentityResponseValidationError{}
