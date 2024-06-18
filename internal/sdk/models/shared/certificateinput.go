// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CertificateInput struct {
	// PEM-encoded public certificate chain of the SSL key pair. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	Cert *string `json:"cert,omitempty"`
	// PEM-encoded public certificate chain of the alternate SSL key pair. This should only be set if you have both RSA and ECDSA types of certificate available and would like Kong to prefer serving using ECDSA certs when client advertises support for it. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	CertAlt *string `json:"cert_alt,omitempty"`
	// PEM-encoded private key of the SSL key pair. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	Key *string `json:"key,omitempty"`
	// PEM-encoded private key of the alternate SSL key pair. This should only be set if you have both RSA and ECDSA types of certificate available and would like Kong to prefer serving using ECDSA certs when client advertises support for it. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	KeyAlt *string `json:"key_alt,omitempty"`
	// An optional set of strings associated with the Certificate for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (o *CertificateInput) GetCert() *string {
	if o == nil {
		return nil
	}
	return o.Cert
}

func (o *CertificateInput) GetCertAlt() *string {
	if o == nil {
		return nil
	}
	return o.CertAlt
}

func (o *CertificateInput) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *CertificateInput) GetKeyAlt() *string {
	if o == nil {
		return nil
	}
	return o.KeyAlt
}

func (o *CertificateInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
