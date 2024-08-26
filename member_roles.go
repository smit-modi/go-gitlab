package gitlab

import (
	"fmt"
	"net/http"
)

// MemberRolesService handles communication with the member roles related
// methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/member_roles.html
type MemberRolesService struct {
	client *Client
}

// MemberRole represents a GitLab member role.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/member_roles.html
type MemberRole struct {
	ID                         int              `json:"id"`
	Name                       string           `json:"name"`
	Description                string           `json:"description,omitempty"`
	GroupId                    int              `json:"group_id"`
	BaseAccessLevel            AccessLevelValue `json:"base_access_level"`
	AdminCICDVariables         bool             `json:"admin_cicd_variables,omitempty"`
	AdminComplianceFramework   bool             `json:"admin_compliance_framework,omitempty"`
	AdminGroupMembers          bool             `json:"admin_group_member,omitempty"`
	AdminMergeRequests         bool             `json:"admin_merge_request,omitempty"`
	AdminPushRules             bool             `json:"admin_push_rules,omitempty"`
	AdminTerraformState        bool             `json:"admin_terraform_state,omitempty"`
	AdminVulnerability         bool             `json:"admin_vulnerability,omitempty"`
	AdminWebHook               bool             `json:"admin_web_hook,omitempty"`
	ArchiveProject             bool             `json:"archive_project,omitempty"`
	ManageDeployTokens         bool             `json:"manage_deploy_tokens,omitempty"`
	ManageGroupAccesToken      bool             `json:"manage_group_access_tokens,omitempty"`
	ManageMergeRequestSettings bool             `json:"manage_merge_request_settings,omitempty"`
	ManageProjectAccessToken   bool             `json:"manage_project_access_tokens,omitempty"`
	ManageSecurityPolicyLink   bool             `json:"manage_security_policy_link,omitempty"`
	ReadCode                   bool             `json:"read_code,omitempty"`
	ReadRunners                bool             `json:"read_runners,omitempty"`
	ReadDependency             bool             `json:"read_dependency,omitempty"`
	ReadVulnerability          bool             `json:"read_vulnerability,omitempty"`
	RemoveGroup                bool             `json:"remove_group,omitempty"`
	RemoveProject              bool             `json:"remove_project,omitempty"`
}

// ListMemberRoles gets a list of member roles for a specified group.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/member_roles.html#list-all-member-roles-of-a-group
func (s *MemberRolesService) ListMemberRoles(gid interface{}, options ...RequestOptionFunc) ([]*MemberRole, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/member_roles", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mrs []*MemberRole
	resp, err := s.client.Do(req, &mrs)
	if err != nil {
		return nil, resp, err
	}

	return mrs, resp, nil
}

// CreateMemberRoleOptions represents the available CreateMemberRole() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/member_roles.html#add-a-member-role-to-a-group
type CreateMemberRoleOptions struct {
	Name                       *string           `url:"name,omitempty" json:"name,omitempty"`
	BaseAccessLevel            *AccessLevelValue `url:"base_access_level,omitempty" json:"base_access_level,omitempty"`
	Description                *string           `url:"description,omitempty" json:"description,omitempty"`
	AdminCICDVariables         *bool             `url:"admin_cicd_variables" json:"admin_cicd_variables,omitempty"`
	AdminComplianceFramework   *bool             `url:"admin_compliance_framework" json:"admin_compliance_framework,omitempty"`
	AdminGroupMembers          *bool             `url:"admin_group_member" json:"admin_group_member,omitempty"`
	AdminMergeRequest          *bool             `url:"admin_merge_request,omitempty" json:"admin_merge_request,omitempty"`
	AdminPushRules             *bool             `url:"admin_push_rules" json:"admin_push_rules,omitempty"`
	AdminTerraformState        *bool             `url:"admin_terraform_state" json:"admin_terraform_state,omitempty"`
	AdminVulnerability         *bool             `url:"admin_vulnerability,omitempty" json:"admin_vulnerability,omitempty"`
	AdminWebHook               *bool             `url:"admin_web_hook" json:"admin_web_hook,omitempty"`
	ArchiveProject             *bool             `url:"archive_project" json:"archive_project,omitempty"`
	ManageDeployTokens         *bool             `url:"manage_deploy_tokens" json:"manage_deploy_tokens,omitempty"`
	ManageGroupAccesToken      *bool             `url:"manage_group_access_tokens" json:"manage_group_access_tokens,omitempty"`
	ManageMergeRequestSettings *bool             `url:"manage_merge_request_settings" json:"manage_merge_request_settings,omitempty"`
	ManageProjectAccessToken   *bool             `url:"manage_project_access_tokens" json:"manage_project_access_tokens,omitempty"`
	ManageSecurityPolicyLink   *bool             `url:"manage_security_policy_link" json:"manage_security_policy_link,omitempty"`
	ReadCode                   *bool             `url:"read_code,omitempty" json:"read_code,omitempty"`
	ReadRunners                *bool             `url:"read_runners" json:"read_runners,omitempty"`
	ReadDependency             *bool             `url:"read_dependency,omitempty" json:"read_dependency,omitempty"`
	ReadVulnerability          *bool             `url:"read_vulnerability,omitempty" json:"read_vulnerability,omitempty"`
	RemoveGroup                *bool             `url:"remove_group" json:"remove_group,omitempty"`
	RemoveProject              *bool             `url:"remove_project" json:"remove_project,omitempty"`
}

// CreateMemberRole creates a new member role for a specified group.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/member_roles.html#add-a-member-role-to-a-group
func (s *MemberRolesService) CreateMemberRole(gid interface{}, opt *CreateMemberRoleOptions, options ...RequestOptionFunc) (*MemberRole, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/member_roles", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	mr := new(MemberRole)
	resp, err := s.client.Do(req, mr)
	if err != nil {
		return nil, resp, err
	}

	return mr, resp, nil
}

// DeleteMemberRole deletes a member role from a specified group.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/member_roles.html#remove-member-role-of-a-group
func (s *MemberRolesService) DeleteMemberRole(gid interface{}, memberRole int, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/member_roles/%d", PathEscape(group), memberRole)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
