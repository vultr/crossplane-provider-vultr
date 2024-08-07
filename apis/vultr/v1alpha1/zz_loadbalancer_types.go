/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallRulesObservation struct {

	// The load balancer ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of ip this rule is - may be either v4 or v6.
	IPType *string `json:"ipType,omitempty" tf:"ip_type,omitempty"`

	// The assigned port (integer) on the attached instances that the load balancer should check against. Default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// IP address with subnet that is allowed through the firewall. You may also pass in cloudflare which will allow only CloudFlares IP range.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type FirewallRulesParameters struct {

	// The type of ip this rule is - may be either v4 or v6.
	// +kubebuilder:validation:Required
	IPType *string `json:"ipType" tf:"ip_type,omitempty"`

	// The assigned port (integer) on the attached instances that the load balancer should check against. Default value is 80.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// IP address with subnet that is allowed through the firewall. You may also pass in cloudflare which will allow only CloudFlares IP range.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

type ForwardingRulesObservation struct {

	// Port on instance side.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Protocol on instance side. Possible values: "http", "https", "tcp".
	BackendProtocol *string `json:"backendProtocol,omitempty" tf:"backend_protocol,omitempty"`

	// Port on load balancer side.
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// Protocol on load balancer side. Possible values: "http", "https", "tcp".
	FrontendProtocol *string `json:"frontendProtocol,omitempty" tf:"frontend_protocol,omitempty"`

	// The load balancer ID.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type ForwardingRulesParameters struct {

	// Port on instance side.
	// +kubebuilder:validation:Required
	BackendPort *float64 `json:"backendPort" tf:"backend_port,omitempty"`

	// Protocol on instance side. Possible values: "http", "https", "tcp".
	// +kubebuilder:validation:Required
	BackendProtocol *string `json:"backendProtocol" tf:"backend_protocol,omitempty"`

	// Port on load balancer side.
	// +kubebuilder:validation:Required
	FrontendPort *float64 `json:"frontendPort" tf:"frontend_port,omitempty"`

	// Protocol on load balancer side. Possible values: "http", "https", "tcp".
	// +kubebuilder:validation:Required
	FrontendProtocol *string `json:"frontendProtocol" tf:"frontend_protocol,omitempty"`
}

type HealthCheckObservation struct {

	// Time in seconds to perform health check. Default value is 15.
	CheckInterval *float64 `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`

	// Number of failed attempts encountered before failover. Default value is 5.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The path on the attached instances that the load balancer should check against. Default value is /
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The assigned port (integer) on the attached instances that the load balancer should check against. Default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol used to traffic requests to the load balancer. Possible values are http, or tcp. Default value is http.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Time in seconds to wait for a health check response. Default value is 5.
	ResponseTimeout *float64 `json:"responseTimeout,omitempty" tf:"response_timeout,omitempty"`

	// Number of failed attempts encountered before failover. Default value is 5.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckParameters struct {

	// Time in seconds to perform health check. Default value is 15.
	// +kubebuilder:validation:Optional
	CheckInterval *float64 `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`

	// Number of failed attempts encountered before failover. Default value is 5.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The path on the attached instances that the load balancer should check against. Default value is /
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The assigned port (integer) on the attached instances that the load balancer should check against. Default value is 80.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The protocol used to traffic requests to the load balancer. Possible values are http, or tcp. Default value is http.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Time in seconds to wait for a health check response. Default value is 5.
	// +kubebuilder:validation:Optional
	ResponseTimeout *float64 `json:"responseTimeout,omitempty" tf:"response_timeout,omitempty"`

	// Number of failed attempts encountered before failover. Default value is 5.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LoadBalancerObservation struct {

	// Array of instances that are currently attached to the load balancer.
	AttachedInstances []*string `json:"attachedInstances,omitempty" tf:"attached_instances,omitempty"`

	// The balancing algorithm for your load balancer. Options are roundrobin or leastconn. Default value is roundrobin
	BalancingAlgorithm *string `json:"balancingAlgorithm,omitempty" tf:"balancing_algorithm,omitempty"`

	// Name for your given sticky session.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Defines the firewall rules for a load balancer.
	FirewallRules []FirewallRulesObservation `json:"firewallRules,omitempty" tf:"firewall_rules,omitempty"`

	// List of forwarding rules for a load balancer. The configuration of a forwarding_rules is listened below.
	ForwardingRules []ForwardingRulesObservation `json:"forwardingRules,omitempty" tf:"forwarding_rules,omitempty"`

	// Boolean value that indicates if SSL is enabled.
	HasSSL *bool `json:"hasSsl,omitempty" tf:"has_ssl,omitempty"`

	// A block that defines the way load balancers should check for health. The configuration of a health_check is listed below.
	HealthCheck []HealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The load balancer ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IPv4 address for your load balancer.
	IPv4 *string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// IPv6 address for your load balancer.
	IPv6 *string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// The load balancer's label.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Boolean value that indicates if Proxy Protocol is enabled.
	ProxyProtocol *bool `json:"proxyProtocol,omitempty" tf:"proxy_protocol,omitempty"`

	// The region your load balancer is deployed in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a ssl is listed below.
	SSL []SSLObservation `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
	SSLRedirect *bool `json:"sslRedirect,omitempty" tf:"ssl_redirect,omitempty"`

	// Current status for the load balancer
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A VPC ID that the load balancer should be attached to.
	VPC *string `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type LoadBalancerParameters struct {

	// Array of instances that are currently attached to the load balancer.
	// +kubebuilder:validation:Optional
	AttachedInstances []*string `json:"attachedInstances,omitempty" tf:"attached_instances,omitempty"`

	// The balancing algorithm for your load balancer. Options are roundrobin or leastconn. Default value is roundrobin
	// +kubebuilder:validation:Optional
	BalancingAlgorithm *string `json:"balancingAlgorithm,omitempty" tf:"balancing_algorithm,omitempty"`

	// Name for your given sticky session.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Defines the firewall rules for a load balancer.
	// +kubebuilder:validation:Optional
	FirewallRules []FirewallRulesParameters `json:"firewallRules,omitempty" tf:"firewall_rules,omitempty"`

	// List of forwarding rules for a load balancer. The configuration of a forwarding_rules is listened below.
	// +kubebuilder:validation:Optional
	ForwardingRules []ForwardingRulesParameters `json:"forwardingRules,omitempty" tf:"forwarding_rules,omitempty"`

	// A block that defines the way load balancers should check for health. The configuration of a health_check is listed below.
	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The load balancer's label.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Boolean value that indicates if Proxy Protocol is enabled.
	// +kubebuilder:validation:Optional
	ProxyProtocol *bool `json:"proxyProtocol,omitempty" tf:"proxy_protocol,omitempty"`

	// The region your load balancer is deployed in.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a ssl is listed below.
	// +kubebuilder:validation:Optional
	SSL []SSLParameters `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
	// +kubebuilder:validation:Optional
	SSLRedirect *bool `json:"sslRedirect,omitempty" tf:"ssl_redirect,omitempty"`

	// A VPC ID that the load balancer should be attached to.
	// +kubebuilder:validation:Optional
	VPC *string `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type SSLObservation struct {

	// The SSL Certificate.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// The SSL certificate chain.
	Chain *string `json:"chain,omitempty" tf:"chain,omitempty"`
}

type SSLParameters struct {

	// The SSL Certificate.
	// +kubebuilder:validation:Required
	Certificate *string `json:"certificate" tf:"certificate,omitempty"`

	// The SSL certificate chain.
	// +kubebuilder:validation:Optional
	Chain *string `json:"chain,omitempty" tf:"chain,omitempty"`

	// The SSL certificates private key.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerParameters `json:"forProvider"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancer is the Schema for the LoadBalancers API. Get information about a Vultr Load Balancer.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vultr}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.forwardingRules)",message="forwardingRules is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.region)",message="region is a required parameter"
	Spec   LoadBalancerSpec   `json:"spec"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancer_Kind             = "LoadBalancer"
	LoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancer_Kind}.String()
	LoadBalancer_KindAPIVersion   = LoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
