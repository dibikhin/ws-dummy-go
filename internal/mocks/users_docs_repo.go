// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ws-dummy-go/internal/dummy/domain"

	mock "github.com/stretchr/testify/mock"
)

// UsersDocsRepo is an autogenerated mock type for the UsersDocsRepo type
type UsersDocsRepo struct {
	mock.Mock
}

type UsersDocsRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *UsersDocsRepo) EXPECT() *UsersDocsRepo_Expecter {
	return &UsersDocsRepo_Expecter{mock: &_m.Mock}
}

// Insert provides a mock function with given fields: ctx, name
func (_m *UsersDocsRepo) Insert(ctx context.Context, name string) (domain.UserID, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 domain.UserID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.UserID, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.UserID); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(domain.UserID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersDocsRepo_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type UsersDocsRepo_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *UsersDocsRepo_Expecter) Insert(ctx interface{}, name interface{}) *UsersDocsRepo_Insert_Call {
	return &UsersDocsRepo_Insert_Call{Call: _e.mock.On("Insert", ctx, name)}
}

func (_c *UsersDocsRepo_Insert_Call) Run(run func(ctx context.Context, name string)) *UsersDocsRepo_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UsersDocsRepo_Insert_Call) Return(_a0 domain.UserID, _a1 error) *UsersDocsRepo_Insert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UsersDocsRepo_Insert_Call) RunAndReturn(run func(context.Context, string) (domain.UserID, error)) *UsersDocsRepo_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// NewUsersDocsRepo creates a new instance of UsersDocsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersDocsRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersDocsRepo {
	mock := &UsersDocsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
