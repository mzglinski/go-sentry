package sentry

import (
	"context"
	"fmt"
)

// Owner represents the owner of a resource in Sentry.
type Owner struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// UptimeMonitor represents an uptime monitor in Sentry.
type UptimeMonitor struct {
	ID              *string     `json:"id,omitempty"`
	Name            *string     `json:"name,omitempty"`
	URL             *string     `json:"url,omitempty"`
	IntervalSeconds *int        `json:"intervalSeconds,omitempty"`
	TimeoutMs       *int        `json:"timeoutMs,omitempty"`
	Status          *string     `json:"status,omitempty"`
	Owner           *Owner      `json:"owner,omitempty"`
	Environment     *string     `json:"environment,omitempty"`
	Method          *string     `json:"method,omitempty"`
	Headers         interface{} `json:"headers,omitempty"`
	Body            *string     `json:"body,omitempty"`
	TraceSampling   *bool       `json:"traceSampling,omitempty"`
}

// UptimeMonitorParams represents the parameters for creating or updating an uptime monitor.
type UptimeMonitorParams struct {
	Name            *string     `json:"name,omitempty"`
	URL             *string     `json:"url,omitempty"`
	IntervalSeconds *int        `json:"interval_seconds,omitempty"`
	TimeoutMs       *int        `json:"timeout_ms,omitempty"`
	Status          *string     `json:"status,omitempty"`
	Owner           *string     `json:"owner,omitempty"`
	Environment     *string     `json:"environment,omitempty"`
	Method          *string     `json:"method,omitempty"`
	Headers         interface{} `json:"headers,omitempty"`
	Body            *string     `json:"body,omitempty"`
	TraceSampling   *bool       `json:"trace_sampling,omitempty"`
}

// UptimeService provides methods for accessing Sentry uptime monitoring API endpoints.
type UptimeService service

// Get retrieves an uptime monitor.
func (s *UptimeService) Get(ctx context.Context, organizationSlug, projectSlug, monitorID string) (*UptimeMonitor, *Response, error) {
	u := fmt.Sprintf("0/projects/%s/%s/uptime/%s/", organizationSlug, projectSlug, monitorID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	monitor := new(UptimeMonitor)
	resp, err := s.client.Do(ctx, req, monitor)
	if err != nil {
		return nil, resp, err
	}
	return monitor, resp, nil
}

// Create creates a new uptime monitor.
func (s *UptimeService) Create(ctx context.Context, organizationSlug, projectSlug string, params *UptimeMonitorParams) (*UptimeMonitor, *Response, error) {
	u := fmt.Sprintf("0/projects/%s/%s/uptime/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	monitor := new(UptimeMonitor)
	resp, err := s.client.Do(ctx, req, monitor)
	if err != nil {
		return nil, resp, err
	}
	return monitor, resp, nil
}

// Update updates an existing uptime monitor.
func (s *UptimeService) Update(ctx context.Context, organizationSlug, projectSlug, monitorID string, params *UptimeMonitorParams) (*UptimeMonitor, *Response, error) {
	u := fmt.Sprintf("0/projects/%s/%s/uptime/%s/", organizationSlug, projectSlug, monitorID)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	monitor := new(UptimeMonitor)
	resp, err := s.client.Do(ctx, req, monitor)
	if err != nil {
		return nil, resp, err
	}
	return monitor, resp, nil
}

// Delete deletes an uptime monitor.
func (s *UptimeService) Delete(ctx context.Context, organizationSlug, projectSlug, monitorID string) (*Response, error) {
	u := fmt.Sprintf("0/projects/%s/%s/uptime/%s/", organizationSlug, projectSlug, monitorID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}
