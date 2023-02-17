package errors_test

import (
	"errors"
	"testing"

	httpcerrors "github.com/ns-jsattler/go-httpc/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMulti(t *testing.T) {
	t.Run("handles nil old error", func(t *testing.T) {
		err := httpcerrors.Append(nil, errors.New("something wrong"))
		require.Error(t, err)
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("handles nil new error", func(t *testing.T) {
		err := httpcerrors.Append(errors.New("something wrong"), nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("handles nil old and new errors", func(t *testing.T) {
		err := httpcerrors.Append(nil, nil)
		require.NoError(t, err)
	})

	t.Run("existing non-multi conflict error", func(t *testing.T) {
		old := &testConflictErr{}
		err := httpcerrors.Append(old, errors.New("something wrong"))
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Conflicter)(nil), err)
		assert.True(t, err.(httpcerrors.Conflicter).Conflict())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("existing non-multi exists error", func(t *testing.T) {
		old := &testExistsErr{}
		err := httpcerrors.Append(old, errors.New("something wrong"))
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Exister)(nil), err)
		assert.True(t, err.(httpcerrors.Exister).Exists())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("existing non-multi notFound error", func(t *testing.T) {
		old := &testNotFoundErr{}
		err := httpcerrors.Append(old, errors.New("something wrong"))
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.NotFounder)(nil), err)
		assert.True(t, err.(httpcerrors.NotFounder).NotFound())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("existing non-multi retry error", func(t *testing.T) {
		old := &testRetryErr{}
		err := httpcerrors.Append(old, errors.New("something wrong"))
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Retrier)(nil), err)
		assert.True(t, err.(httpcerrors.Retrier).Retry())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("existing non-multi temporary error", func(t *testing.T) {
		old := &testTemporaryErr{}
		err := httpcerrors.Append(old, errors.New("something wrong"))
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Temporarier)(nil), err)
		assert.True(t, err.(httpcerrors.Temporarier).Temporary())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("append conflict error", func(t *testing.T) {
		new := &testConflictErr{}
		err := httpcerrors.Append(errors.New("something wrong"), new)
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Conflicter)(nil), err)
		assert.True(t, err.(httpcerrors.Conflicter).Conflict())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("append exists error", func(t *testing.T) {
		new := &testExistsErr{}
		err := httpcerrors.Append(errors.New("something wrong"), new)
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Exister)(nil), err)
		assert.True(t, err.(httpcerrors.Exister).Exists())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("append notFound error", func(t *testing.T) {
		new := &testNotFoundErr{}
		err := httpcerrors.Append(errors.New("something wrong"), new)
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.NotFounder)(nil), err)
		assert.True(t, err.(httpcerrors.NotFounder).NotFound())
		assert.Contains(t, err.Error(), "something wrong")
	})

	t.Run("append retry error", func(t *testing.T) {
		new := &testRetryErr{}
		err := httpcerrors.Append(errors.New("something wrong"), new)
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Retrier)(nil), err)
		assert.True(t, err.(httpcerrors.Retrier).Retry())
	})

	t.Run("append temporary error", func(t *testing.T) {
		new := &testTemporaryErr{}
		err := httpcerrors.Append(errors.New("something wrong"), new)
		require.Error(t, err)
		require.Implements(t, (*httpcerrors.Temporarier)(nil), err)
		assert.True(t, err.(httpcerrors.Temporarier).Temporary())
		assert.Contains(t, err.Error(), "something wrong")
	})
}

type testConflictErr struct{}

func (err *testConflictErr) Error() string {
	return "conflict"
}

func (err *testConflictErr) Conflict() bool {
	return true
}

type testExistsErr struct{}

func (err *testExistsErr) Error() string {
	return "exists"
}

func (err *testExistsErr) Exists() bool {
	return true
}

type testNotFoundErr struct{}

func (err *testNotFoundErr) Error() string {
	return "not found"
}

func (err *testNotFoundErr) NotFound() bool {
	return true
}

type testRetryErr struct{}

func (err *testRetryErr) Error() string {
	return "retry me!"
}

func (err *testRetryErr) Retry() bool {
	return true
}

type testTemporaryErr struct{}

func (err *testTemporaryErr) Error() string {
	return "retry me!"
}

func (err *testTemporaryErr) Temporary() bool {
	return true
}
