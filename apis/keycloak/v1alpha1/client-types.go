package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClientSpec defines the desired state of Client
// +k8s:openapi-gen=true
type ClientSpec struct {
	// Selector for looking up KeycloakRealm Custom Resources.
	// +kubebuilder:validation:Required
	RealmSelector *metav1.LabelSelector `json:"realmSelector"`
	// Keycloak Client REST object.
	// +kubebuilder:validation:Required
	Client *KeycloakAPIClient `json:"client"`
}

type KeycloakAPIClient struct {
	// Client ID. If not specified, automatically generated.
	// +optional
	ID string `json:"id,omitempty"`
	// Client ID.
	// +kubebuilder:validation:Required
	ClientID string `json:"clientId"`
	// Client name.
	// +optional
	Name string `json:"name,omitempty"`
	// Surrogate Authentication Required option.
	// +optional
	SurrogateAuthRequired bool `json:"surrogateAuthRequired,omitempty"`
	// Client enabled flag.
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// What Client authentication type to use.
	// +optional
	ClientAuthenticatorType string `json:"clientAuthenticatorType,omitempty"`
	// Client Secret. The Operator will automatically create a Secret based on this value.
	// +optional
	Secret string `json:"secret,omitempty"`
	// Application base URL.
	// +optional
	BaseURL string `json:"baseUrl,omitempty"`
	// Application Admin URL.
	// +optional
	AdminURL string `json:"adminUrl,omitempty"`
	// Application root URL.
	// +optional
	RootURL string `json:"rootUrl,omitempty"`
	// Client description.
	// +optional
	Description string `json:"description,omitempty"`
	// Default Client roles.
	// +optional
	DefaultRoles []string `json:"defaultRoles,omitempty"`
	// A list of valid Redirection URLs.
	// +optional
	RedirectUris []string `json:"redirectUris,omitempty"`
	// A list of valid Web Origins.
	// +optional
	WebOrigins []string `json:"webOrigins,omitempty"`
	// Not Before setting.
	// +optional
	NotBefore int `json:"notBefore,omitempty"`
	// True if a client supports only Bearer Tokens.
	// +optional
	BearerOnly bool `json:"bearerOnly,omitempty"`
	// True if Consent Screen is required.
	// +optional
	ConsentRequired bool `json:"consentRequired,omitempty"`
	// True if Standard flow is enabled.
	// +optional
	StandardFlowEnabled bool `json:"standardFlowEnabled,omitempty"`
	// True if Implicit flow is enabled.
	// +optional
	ImplicitFlowEnabled bool `json:"implicitFlowEnabled,omitempty"`
	// True if Direct Grant is enabled.
	// +optional
	DirectAccessGrantsEnabled bool `json:"directAccessGrantsEnabled,omitempty"`
	// True if Service Accounts are enabled.
	// +optional
	ServiceAccountsEnabled bool `json:"serviceAccountsEnabled,omitempty"`
	// True if this is a public Client.
	// +optional
	PublicClient bool `json:"publicClient,omitempty"`
	// True if this client supports Front Channel logout.
	// +optional
	FrontchannelLogout bool `json:"frontchannelLogout,omitempty"`
	// Protocol used for this Client.
	// +optional
	Protocol string `json:"protocol,omitempty"`
	// Client Attributes.
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`
	// True if Full Scope is allowed.
	// +optional
	FullScopeAllowed bool `json:"fullScopeAllowed,omitempty"`
	// Node registration timeout.
	// +optional
	NodeReRegistrationTimeout int `json:"nodeReRegistrationTimeout,omitempty"`
	// Protocol Mappers.
	// +optional
	ProtocolMappers []KeycloakProtocolMapper `json:"protocolMappers,omitempty"`
	// True to use a Template Config.
	// +optional
	UseTemplateConfig bool `json:"useTemplateConfig,omitempty"`
	// True to use Template Scope.
	// +optional
	UseTemplateScope bool `json:"useTemplateScope,omitempty"`
	// True to use Template Mappers.
	// +optional
	UseTemplateMappers bool `json:"useTemplateMappers,omitempty"`
	// Access options.
	// +optional
	Access map[string]bool `json:"access,omitempty"`
}


// ClientStatus defines the observed state of Client
// +k8s:openapi-gen=true
type ClientStatus struct {
	// Current phase of the operator.
	Phase StatusPhase `json:"phase"`
	// Human-readable message indicating details about current operator phase or error.
	Message string `json:"message"`
	// True if all resources are in a ready state and all work is done.
	Ready bool `json:"ready"`
	// A map of all the secondary resources types and names created for this CR. e.g "Deployment": [ "DeploymentName1", "DeploymentName2" ]
	SecondaryResources map[string][]string `json:"secondaryResources,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Client is the Schema for the keycloakclients API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=keycloakclients,scope=Namespaced
type Client struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClientSpec   `json:"spec,omitempty"`
	Status ClientStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClientList contains a list of Client
type ClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Client `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Client{}, &ClientList{})
}

func (i *Client) UpdateStatusSecondaryResources(kind string, resourceName string) {
	i.Status.SecondaryResources = UpdateStatusSecondaryResources(i.Status.SecondaryResources, kind, resourceName)
}