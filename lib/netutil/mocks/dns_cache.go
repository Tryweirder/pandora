// Code generated by mockery v1.0.0
package netmock

import "github.com/stretchr/testify/mock"

// DNSCache is an autogenerated mock type for the DNSCache type
type DNSCache struct {
	mock.Mock
}

// Add provides a mock function with given fields: addr, resolved
func (_m *DNSCache) Add(addr string, resolved string) {
	_m.Called(addr, resolved)
}

// Get provides a mock function with given fields: addr
func (_m *DNSCache) Get(addr string) (string, bool) {
	ret := _m.Called(addr)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}