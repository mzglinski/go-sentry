package sentry

import (
	"context"
	"fmt"
)

// ExternalUserService provides methods for accessing Sentry external users API endpoints.
type ExternalUserService service

// CreateExternalUserParams represents the parameters for creating an external user.
type CreateExternalUserParams struct {
	UserID        int    `json:"user_id"`
	ExternalName  string `json:"external_name"`
	ExternalProvider string `json:"provider"`
	IntegrationID int    `json:"integration_id"`
	ID            int    `json:"id"`
	ExternalID    string `json:"external_id,omitempty"`
}

// ExternalUser represents a Sentry external user.
type ExternalUser struct {
	ID            string `json:"id"`
	UserID        string `json:"userId"`
	ExternalName  string `json:"externalName"`
	Provider      string `json:"provider"`
	IntegrationID string `json:"integrationId"`
}

// Create a new external user.
// POST /api/0/organizations/{organization_id_or_slug}/external-users/
func (s *ExternalUserService) Create(ctx context.Context, organizationIDOrSlug string, params *CreateExternalUserParams) (*ExternalUser, *Response, error) {
	u := fmt.Sprintf("0/organizations/%s/external-users/", organizationIDOrSlug)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	externalUser := new(ExternalUser)
	resp, err := s.client.Do(ctx, req, externalUser)
	if err != nil {
		return nil, resp, err
	}
	return externalUser, resp, nil
}

// UpdateExternalUserParams represents the parameters for updating an external user.
// It's identical to CreateExternalUserParams as Sentry's Update API expects the full object.
type UpdateExternalUserParams = CreateExternalUserParams

// Update an existing external user mapping.
// PUT /api/0/organizations/{organization_id_or_slug}/external-users/{external_user_id}/
func (s *ExternalUserService) Update(ctx context.Context, organizationIDOrSlug string, externalUserMappingID string, params *UpdateExternalUserParams) (*ExternalUser, *Response, error) {
	u := fmt.Sprintf("0/organizations/%s/external-users/%s/", organizationIDOrSlug, externalUserMappingID)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	updatedExternalUser := new(ExternalUser)
	resp, err := s.client.Do(ctx, req, updatedExternalUser)
	if err != nil {
		return nil, resp, err
	}
	return updatedExternalUser, resp, nil
}

// Delete an external user.
// DELETE /api/0/organizations/{organization_id_or_slug}/external-users/{external_user_id}/
func (s *ExternalUserService) Delete(ctx context.Context, organizationIDOrSlug string, externalUserID string) (*Response, error) {
	u := fmt.Sprintf("0/organizations/%s/external-users/%s/", organizationIDOrSlug, externalUserID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
} 