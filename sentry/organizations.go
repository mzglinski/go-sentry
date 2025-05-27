package sentry

import (
	"context"
	"fmt"
	"time"
)

// OrganizationStatus represents a Sentry organization's status.
type OrganizationStatus struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

// OrganizationQuota represents a Sentry organization's quota.
type OrganizationQuota struct {
	MaxRate         *int `json:"maxRate"`
	MaxRateInterval *int `json:"maxRateInterval"`
	AccountLimit    *int `json:"accountLimit"`
	ProjectLimit    *int `json:"projectLimit"`
}

// OrganizationAvailableRole represents a Sentry organization's available role.
type OrganizationAvailableRole struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

// Organization represents detailed information about a Sentry organization.
// Based on https://github.com/getsentry/sentry/blob/22.5.0/src/sentry/api/serializers/models/organization.py#L263-L288
type Organization struct {
	// Basic
	ID                       *string             `json:"id,omitempty"`
	Slug                     *string             `json:"slug,omitempty"`
	Status                   *OrganizationStatus `json:"status,omitempty"`
	Name                     *string             `json:"name,omitempty"`
	DateCreated              *time.Time          `json:"dateCreated,omitempty"`
	IsEarlyAdopter           *bool               `json:"isEarlyAdopter,omitempty"`
	Require2FA               *bool               `json:"require2FA,omitempty"`
	RequireEmailVerification *bool               `json:"requireEmailVerification,omitempty"`
	Avatar                   *OrganizationAvatar `json:"avatar,omitempty"`
	Features                 []string            `json:"features,omitempty"`

	// Detailed
	// TODO: experiments
	Quota                *OrganizationQuota          `json:"quota,omitempty"`
	IsDefault            *bool                       `json:"isDefault,omitempty"`
	DefaultRole          *string                     `json:"defaultRole,omitempty"`
	AvailableRoles       []OrganizationAvailableRole `json:"availableRoles,omitempty"`
	OrgRoleList          []OrganizationRoleListItem  `json:"orgRoleList,omitempty"`
	TeamRoleList         []TeamRoleListItem          `json:"teamRoleList,omitempty"`
	OpenMembership       *bool                       `json:"openMembership,omitempty"`
	AllowSharedIssues    *bool                       `json:"allowSharedIssues,omitempty"`
	EnhancedPrivacy      *bool                       `json:"enhancedPrivacy,omitempty"`
	DataScrubber         *bool                       `json:"dataScrubber,omitempty"`
	DataScrubberDefaults *bool                       `json:"dataScrubberDefaults,omitempty"`
	SensitiveFields      []string                    `json:"sensitiveFields,omitempty"`
	SafeFields           []string                    `json:"safeFields,omitempty"`
	StoreCrashReports    *int                        `json:"storeCrashReports,omitempty"`
	AttachmentsRole      *string                     `json:"attachmentsRole,omitempty"`
	DebugFilesRole       *string                     `json:"debugFilesRole,omitempty"`
	EventsMemberAdmin    *bool                       `json:"eventsMemberAdmin,omitempty"`
	AlertsMemberWrite    *bool                       `json:"alertsMemberWrite,omitempty"`
	ScrubIPAddresses     *bool                       `json:"scrubIPAddresses,omitempty"`
	ScrapeJavaScript     *bool                       `json:"scrapeJavaScript,omitempty"`
	AllowJoinRequests    *bool                       `json:"allowJoinRequests,omitempty"`
	RelayPiiConfig       *string                     `json:"relayPiiConfig,omitempty"`
	TrustedRelays        []TrustedRelay              `json:"trustedRelays,omitempty"`
	Access               []string                    `json:"access,omitempty"`
	Role                 *string                     `json:"role,omitempty"`
	PendingAccessRequests *int                        `json:"pendingAccessRequests,omitempty"`
	// TODO: onboardingTasks

	// Fields from Update API, not necessarily in GET response structure but useful for consistency
	HideAIFeatures      *bool `json:"hideAiFeatures,omitempty"`
	CodecovAccess       *bool `json:"codecovAccess,omitempty"`
	GithubPRBot         *bool `json:"githubPRBot,omitempty"`
	GithubOpenPRBot     *bool `json:"githubOpenPRBot,omitempty"`
	GithubNudgeInvite   *bool `json:"githubNudgeInvite,omitempty"`
	GitlabPRBot         *bool `json:"gitlabPRBot,omitempty"`
	AllowMemberProjectCreation *bool `json:"allowMemberProjectCreation,omitempty"`
}

// OrganizationAvatar represents a Sentry organization's avatar.
// This is distinct from the User's Avatar struct.
type OrganizationAvatar struct {
	AvatarType *string `json:"avatarType,omitempty"`
	AvatarUUID *string `json:"avatarUuid,omitempty"`
	// Avatar (base64 content) is not part of the GET response, it is for uploads via UpdateOrganizationParams
}

// TrustedRelay represents a Sentry organization's trusted relay.
type TrustedRelay struct {
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	PublicKey   *string    `json:"publicKey,omitempty"`
	Description *string    `json:"description,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastSeen    *time.Time `json:"lastSeen,omitempty"`
}

// OrganizationsService provides methods for accessing Sentry organization API endpoints.
// https://docs.sentry.io/api/organizations/
type OrganizationsService service

// List organizations available to the authenticated session.
// https://docs.sentry.io/api/organizations/list-your-organizations/
func (s *OrganizationsService) List(ctx context.Context, params *ListCursorParams) ([]*Organization, *Response, error) {
	u, err := addQuery("0/organizations/", params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	orgs := []*Organization{}
	resp, err := s.client.Do(ctx, req, &orgs)
	if err != nil {
		return nil, resp, err
	}
	return orgs, resp, nil
}

// Get a Sentry organization.
// https://docs.sentry.io/api/organizations/retrieve-an-organization/
func (s *OrganizationsService) Get(ctx context.Context, slug string) (*Organization, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/", slug)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}
	return org, resp, nil
}

// CreateOrganizationParams are the parameters for OrganizationService.Create.
type CreateOrganizationParams struct {
	Name       *string `json:"name,omitempty"`
	Slug       *string `json:"slug,omitempty"`
	AgreeTerms *bool   `json:"agreeTerms,omitempty"`
}

// Create a new Sentry organization.
func (s *OrganizationsService) Create(ctx context.Context, params *CreateOrganizationParams) (*Organization, *Response, error) {
	u := "0/organizations/"
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}
	return org, resp, nil
}

// UpdateOrganizationParams are the parameters for OrganizationService.Update.
type UpdateOrganizationParams struct {
	Name                 *string                      `json:"name,omitempty"`
	Slug                 *string                      `json:"slug,omitempty"`
	IsEarlyAdopter       *bool                        `json:"isEarlyAdopter,omitempty"`
	HideAIFeatures       *bool                        `json:"hideAiFeatures,omitempty"`
	CodecovAccess        *bool                        `json:"codecovAccess,omitempty"`
	DefaultRole          *string                      `json:"defaultRole,omitempty"`
	OpenMembership       *bool                        `json:"openMembership,omitempty"`
	EventsMemberAdmin    *bool                        `json:"eventsMemberAdmin,omitempty"`
	AlertsMemberWrite    *bool                        `json:"alertsMemberWrite,omitempty"`
	AttachmentsRole      *string                      `json:"attachmentsRole,omitempty"`
	DebugFilesRole       *string                      `json:"debugFilesRole,omitempty"`
	AvatarType           *string                      `json:"avatarType,omitempty"`
	Avatar               *string                      `json:"avatar,omitempty"` // base64 encoded image
	Require2FA           *bool                        `json:"require2FA,omitempty"`
	AllowSharedIssues    *bool                        `json:"allowSharedIssues,omitempty"`
	EnhancedPrivacy      *bool                        `json:"enhancedPrivacy,omitempty"`
	ScrapeJavaScript     *bool                        `json:"scrapeJavaScript,omitempty"`
	StoreCrashReports    *int                         `json:"storeCrashReports,omitempty"`
	AllowJoinRequests    *bool                        `json:"allowJoinRequests,omitempty"`
	DataScrubber         *bool                        `json:"dataScrubber,omitempty"`
	DataScrubberDefaults *bool                        `json:"dataScrubberDefaults,omitempty"`
	SensitiveFields      []string                     `json:"sensitiveFields,omitempty"`
	SafeFields           []string                     `json:"safeFields,omitempty"`
	ScrubIPAddresses     *bool                        `json:"scrubIPAddresses,omitempty"`
	RelayPiiConfig       *string                      `json:"relayPiiConfig,omitempty"`
	TrustedRelays        []TrustedRelayUpdateParams `json:"trustedRelays,omitempty"`
	GithubPRBot          *bool                        `json:"githubPRBot,omitempty"`
	GithubOpenPRBot      *bool                        `json:"githubOpenPRBot,omitempty"`
	GithubNudgeInvite    *bool                        `json:"githubNudgeInvite,omitempty"`
	GitlabPRBot          *bool                        `json:"gitlabPRBot,omitempty"`
	AllowMemberProjectCreation *bool                  `json:"allowMemberProjectCreation,omitempty"`
}

// TrustedRelayUpdateParams are the parameters for updating trusted relays.
type TrustedRelayUpdateParams struct {
	Name        *string `json:"name,omitempty"`
	PublicKey   *string `json:"publicKey,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Update a Sentry organization.
// https://docs.sentry.io/api/organizations/update-an-organization/
func (s *OrganizationsService) Update(ctx context.Context, slug string, params *UpdateOrganizationParams) (*Organization, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/", slug)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}
	return org, resp, nil
}

// Delete a Sentry organization.
func (s *OrganizationsService) Delete(ctx context.Context, slug string) (*Response, error) {
	u := fmt.Sprintf("0/organizations/%v/", slug)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
