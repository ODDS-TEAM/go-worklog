package worklog_test

import "github.com/stretchr/testify/mock"

type MockedHomeDir struct {
	mock.Mock
}

func (m *MockedHomeDir) Path() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockedHomeDir) With(path string) string {
	args := m.Called(path)
	return args.String(0)
}
