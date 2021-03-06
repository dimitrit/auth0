package management

import "encoding/json"

type Connection struct {
	// A generated string identifying the connection.
	ID *string `json:"id,omitempty"`

	// The name of the connection. Must start and end with an alphanumeric
	// character and can only contain alphanumeric characters and '-'. Max
	// length 128.
	Name *string `json:"name,omitempty"`

	// The identity provider identifier for the connection. Can be any of the
	// following:
	//
	// "ad", "adfs", "amazon", "dropbox", "bitbucket", "aol", "auth0-adldap",
	// "auth0-oidc", "auth0", "baidu", "bitly", "box", "custom", "daccount",
	// "dwolla", "email", "evernote-sandbox", "evernote", "exact", "facebook",
	// "fitbit", "flickr", "github", "google-apps", "google-oauth2", "guardian",
	//  "instagram", "ip", "linkedin", "miicard", "oauth1", "oauth2",
	// "office365", "paypal", "paypal-sandbox", "pingfederate",
	// "planningcenter", "renren", "salesforce-community", "salesforce-sandbox",
	//  "salesforce", "samlp", "sharepoint", "shopify", "sms", "soundcloud",
	// "thecity-sandbox", "thecity", "thirtysevensignals", "twitter", "untappd",
	//  "vkontakte", "waad", "weibo", "windowslive", "wordpress", "yahoo",
	// "yammer" or "yandex".
	Strategy *string `json:"strategy,omitempty"`

	// True if the connection is domain level
	IsDomainConnection *bool `json:"is_domain_connection,omitempty"`

	// Options for validation.
	Options *ConnectionOptions `json:"options,omitempty"`

	// The identifiers of the clients for which the connection is to be
	// enabled. If the array is empty or the property is not specified, no
	// clients are enabled.
	EnabledClients []interface{} `json:"enabled_clients,omitempty"`

	// Defines the realms for which the connection will be used (ie: email
	// domains). If the array is empty or the property is not specified, the
	// connection name will be added as realm.
	Realms []interface{} `json:"realms,omitempty"`

	Metadata *interface{} `json:"metadata,omitempty"`
}

// ConnectionOptions general options
type ConnectionOptions struct {
	// Options for validation.
	Validation map[string]interface{} `json:"validation,omitempty"`

	// Password strength level, can be one of:
	// "none", "low", "fair", "good", "excellent" or null.
	PasswordPolicy *string `json:"passwordPolicy,omitempty"`

	// Options for password history policy.
	PasswordHistory map[string]interface{} `json:"password_history,omitempty"`

	// Options for password expiration policy.
	PasswordNoPersonalInfo map[string]interface{} `json:"password_no_personal_info,omitempty"`

	// Options for password dictionary policy.
	PasswordDictionary map[string]interface{} `json:"password_dictionary,omitempty"`

	// Options for password complexity options.
	PasswordComplexityOptions map[string]interface{} `json:"password_complexity_options,omitempty"`

	APIEnableUsers               *bool `json:"api_enable_users,omitempty"`
	BasicProfile                 *bool `json:"basic_profile,omitempty"`
	ExtAdmin                     *bool `json:"ext_admin,omitempty"`
	ExtIsSuspended               *bool `json:"ext_is_suspended,omitempty"`
	ExtAgreedTerms               *bool `json:"ext_agreed_terms,omitempty"`
	ExtGroups                    *bool `json:"ext_groups,omitempty"`
	ExtNestedGroups              *bool `json:"ext_nested_groups,omitempty"`
	ExtAssignedPlans             *bool `json:"ext_assigned_plans,omitempty"`
	ExtProfile                   *bool `json:"ext_profile,omitempty"`
	EnabledDatabaseCustomization *bool `json:"enabledDatabaseCustomization,omitempty"`
	BruteForceProtection         *bool `json:"brute_force_protection,omitempty"`
	ImportMode                   *bool `json:"import_mode,omitempty"`
	DisableSignup                *bool `json:"disable_signup,omitempty"`
	RequiresUsername             *bool `json:"requires_username,omitempty"`

	// Options for adding parameters in the request to the upstream IdP.
	UpstreamParams *interface{} `json:"upstream_params,omitempty"`

	ClientID            *string       `json:"client_id,omitempty"`
	ClientSecret        *string       `json:"client_secret,omitempty"`
	TenantDomain        *string       `json:"tenant_domain,omitempty"`
	DomainAliases       []interface{} `json:"domain_aliases,omitempty"`
	UseWsfed            *bool         `json:"use_wsfed,omitempty"`
	WaadProtocol        *string       `json:"waad_protocol,omitempty"`
	WaadCommonEndpoint  *bool         `json:"waad_common_endpoint,omitempty"`
	AppID               *string       `json:"app_id,omitempty"`
	AppDomain           *string       `json:"app_domain,omitempty"`
	MaxGroupsToRetrieve *string       `json:"max_groups_to_retrieve,omitempty"`

	// Scripts for the connection
	// Allowed keys are: "get_user", "login", "create", "verify", "change_password", "delete" or "change_email".
	CustomScripts map[string]interface{} `json:"customScripts,omitempty"`
	// configuration variables that can be used in custom scripts
	Configuration map[string]interface{} `json:"configuration,omitempty"`

	// Options to add integration with Twilio
	// https://community.auth0.com/t/using-management-api-to-create-a-twilio-connection/23576/3
	Totp                *ConnectionOptionsTotp `json:"totp,omitempty"`
	Name                *string                `json:"name,omitempty"`
	TwilioSid           *string                `json:"twilio_sid,omitempty"`
	TwilioToken         *string                `json:"twilio_token,omitempty"`
	From                *string                `json:"from,omitempty"`
	Syntax              *string                `json:"syntax,omitempty"`
	Template            *string                `json:"template,omitempty"`
	MessagingServiceSid *string                `json:"messaging_service_sid,omitempty"`

	// Adfs
	AdfsServer *string `json:"adfs_server,omitempty"`

	// Salesforce community
	CommunityBaseURL *string `json:"community_base_url"`

	// Passwordless email or Google-Oauth2
	Email *ConnectionOptionsEmail `json:"email,omitempty"`

	// Windowslive strategy version:
	//   1 => Live Connect (DEPRECATED)
	//   2 => Azure AD (personal accounts)
	StrategyVersion *int `json:"strategy_version"`
}

type ConnectionManager struct {
	*Management
}

// PasswordlessEmail for one-time password authentication via email messages
// See https://auth0.com/docs/connections/passwordless/guides/email-otp
type PasswordlessEmail struct {
	// allowed key for syntax: "liquid"
	Syntax  *string `json:"syntax,omitempty"`
	From    *string `json:"from,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Body    *string `json:"body,omitempty"`
}

// ConnectionOptionsEmail holds a pointer to either a bool or to a PasswordlessEmail
type ConnectionOptionsEmail struct {
	v interface{}
}

// MarshalJSON marshals a ConnectionOptionsEmail
func (e *ConnectionOptionsEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.v)
}

// UnmarshalJSON unmarshals a ConnectionOptionsEmail
func (e *ConnectionOptionsEmail) UnmarshalJSON(b []byte) (err error) {
	// check if this is an passwordless email option
	var p PasswordlessEmail
	if err = json.Unmarshal(b, &p); err == nil {
		e.v = &p
		return
	}

	// nope, so maybe it is a google-oauth2 email option
	var q bool
	if err = json.Unmarshal(b, &q); err != nil {
		// it's neither :(
		return err
	}

	e.v = &q
	return
}

// PasswordlessEmail returns the value of a ConnectionOptionsEmail as a
// passwordless email, or false if ConnectionOptionsEmail is not a pointer
// to a passwordless email value
func (e *ConnectionOptionsEmail) PasswordlessEmail() (*PasswordlessEmail, bool) {
	if e == nil || e.v == nil {
		return nil, false
	}

	if email, ok := e.v.(*PasswordlessEmail); ok {
		return email, ok
	}

	return nil, false
}

// Bool returns the value of a ConnectionOptionsEmail as a bool, or false if
// ConnectionOptionsEmail is not a pointer to a boolean value
func (e *ConnectionOptionsEmail) Bool() (*bool, bool) {
	if e == nil || e.v == nil {
		return nil, false
	}

	if email, ok := e.v.(*bool); ok {
		return email, ok
	}

	return nil, false
}

// SetPasswordlessEmail sets the value of ConnectionOptionsEmail to a passwordless email
func (e *ConnectionOptionsEmail) SetPasswordlessEmail(v *PasswordlessEmail) {
	e.v = v
}

// SetBool sets the value of ConnectionOptionsEmail to a bool
func (e *ConnectionOptionsEmail) SetBool(v *bool) {
	e.v = v
}

type ConnectionOptionsTotp struct {
	TimeStep *int `json:"time_step,omitempty"`
	Length   *int `json:"length,omitempty"`
}

type ConnectionList struct {
	List
	Connections []*Connection `json:"connections"`
}

func newConnectionManager(m *Management) *ConnectionManager {
	return &ConnectionManager{m}
}

// Create a new connection.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/post_connections
func (m *ConnectionManager) Create(c *Connection) error {
	return m.post(m.uri("connections"), c)
}

// Read retrieves a connection by its id.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/get_connections_by_id
func (m *ConnectionManager) Read(id string) (c *Connection, err error) {
	err = m.get(m.uri("connections", id), &c)
	return
}

// List all connections.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/get_connections
func (m *ConnectionManager) List(opts ...ListOption) (c *ConnectionList, err error) {
	opts = m.defaults(opts)
	err = m.get(m.uri("connections")+m.q(opts), &c)
	return
}

// Update a connection.
//
// Note: if you use the options parameter, the whole options object will be
// overridden, so ensure that all parameters are present.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/patch_connections_by_id
func (m *ConnectionManager) Update(id string, c *Connection) (err error) {
	return m.patch(m.uri("connections", id), c)
}

// Delete a connection and all its users.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/delete_connections_by_id
func (m *ConnectionManager) Delete(id string) (err error) {
	return m.delete(m.uri("connections", id))
}

// ReadByName retrieves a connection by its name. This is a helper method when a
// connection id is not readily available.
func (m *ConnectionManager) ReadByName(name string) (*Connection, error) {
	c, err := m.List(Parameter("name", name))
	if err != nil {
		return nil, err
	}
	if len(c.Connections) > 0 {
		return c.Connections[0], nil
	}
	return nil, &managementError{404, "Not Found", "Connection not found"}
}
