//go:build !external
// +build !external

package main

import (
	"time"

	"code.cloudfoundry.org/healthcheck"
)

func newHealthCheck(
	network, uri, port string,
	timeout time.Duration,
	useHTTP2 bool,
) healthcheck.HealthCheck {
	return healthcheck.NewHealthCheck(network, uri, port, timeout, useHTTP2)
}
