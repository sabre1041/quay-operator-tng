package securityscanner

import (
	"errors"

	"github.com/creasty/defaults"
)

// SecurityScannerFieldGroup represents the SecurityScannerFieldGroup config fields
type SecurityScannerFieldGroup struct {
	FeatureSecurityScanner              bool     `default:"false" validate:"" json:"FEATURE_SECURITY_SCANNER" yaml:"FEATURE_SECURITY_SCANNER"`
	SecurityScannerEndpoint             string   `default:"" validate:"" json:"SECURITY_SCANNER_ENDPOINT" yaml:"SECURITY_SCANNER_ENDPOINT"`
	SecurityScannerIndexingInterval     int      `default:"30" validate:"" json:"SECURITY_SCANNER_INDEXING_INTERVAL" yaml:"SECURITY_SCANNER_INDEXING_INTERVAL"`
	SecurityScannerNotifications        bool     `default:"false" validate:"" json:"SECURITY_SCANNER_NOTIFICATIONS" yaml:"SECURITY_SCANNER_NOTIFICATIONS"`
	SecurityScannerV4Endpoint           string   `default:"" validate:"" json:"SECURITY_SCANNER_V4_ENDPOINT" yaml:"SECURITY_SCANNER_V4_ENDPOINT"`
	SecurityScannerV4NamespaceWhitelist []string `default:"[]" validate:"" json:"SECURITY_SCANNER_V4_NAMESPACE_WHITELIST" yaml:"SECURITY_SCANNER_V4_NAMESPACE_WHITELIST"`
}

// NewSecurityScannerFieldGroup creates a new SecurityScannerFieldGroup
func NewSecurityScannerFieldGroup(fullConfig map[string]interface{}) (*SecurityScannerFieldGroup, error) {
	newSecurityScannerFieldGroup := &SecurityScannerFieldGroup{}
	defaults.Set(newSecurityScannerFieldGroup)

	if value, ok := fullConfig["FEATURE_SECURITY_SCANNER"]; ok {
		newSecurityScannerFieldGroup.FeatureSecurityScanner, ok = value.(bool)
		if !ok {
			return newSecurityScannerFieldGroup, errors.New("FEATURE_SECURITY_SCANNER must be of type bool")
		}
	}
	if value, ok := fullConfig["SECURITY_SCANNER_ENDPOINT"]; ok {
		newSecurityScannerFieldGroup.SecurityScannerEndpoint, ok = value.(string)
		if !ok {
			return newSecurityScannerFieldGroup, errors.New("SECURITY_SCANNER_ENDPOINT must be of type string")
		}
	}
	if value, ok := fullConfig["SECURITY_SCANNER_INDEXING_INTERVAL"]; ok {
		newSecurityScannerFieldGroup.SecurityScannerIndexingInterval, ok = value.(int)
		if !ok {
			return newSecurityScannerFieldGroup, errors.New("SECURITY_SCANNER_INDEXING_INTERVAL must be of type int")
		}
	}
	if value, ok := fullConfig["SECURITY_SCANNER_NOTIFICATIONS"]; ok {
		newSecurityScannerFieldGroup.SecurityScannerNotifications, ok = value.(bool)
		if !ok {
			return newSecurityScannerFieldGroup, errors.New("SECURITY_SCANNER_NOTIFICATIONS must be of type bool")
		}
	}
	if value, ok := fullConfig["SECURITY_SCANNER_V4_ENDPOINT"]; ok {
		newSecurityScannerFieldGroup.SecurityScannerV4Endpoint, ok = value.(string)
		if !ok {
			return newSecurityScannerFieldGroup, errors.New("SECURITY_SCANNER_V4_ENDPOINT must be of type string")
		}
	}
	if value, ok := fullConfig["SECURITY_SCANNER_V4_NAMESPACE_WHITELIST"]; ok {
		newSecurityScannerFieldGroup.SecurityScannerV4NamespaceWhitelist, ok = value.([]string)
		if !ok {
			return newSecurityScannerFieldGroup, errors.New("SECURITY_SCANNER_V4_NAMESPACE_WHITELIST must be of type []string")
		}
	}

	return newSecurityScannerFieldGroup, nil
}
