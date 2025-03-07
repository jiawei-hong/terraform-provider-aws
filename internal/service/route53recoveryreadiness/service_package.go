package route53recoveryreadiness

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	endpoints_sdkv1 "github.com/aws/aws-sdk-go/aws/endpoints"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	route53recoveryreadiness_sdkv1 "github.com/aws/aws-sdk-go/service/route53recoveryreadiness"
)

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, m map[string]any) (*route53recoveryreadiness_sdkv1.Route53RecoveryReadiness, error) {
	sess := m["session"].(*session_sdkv1.Session)
	config := &aws_sdkv1.Config{Endpoint: aws_sdkv1.String(m["endpoint"].(string))}

	// Force "global" services to correct Regions.
	if m["partition"].(string) == endpoints_sdkv1.AwsPartitionID {
		config.Region = aws_sdkv1.String(endpoints_sdkv1.UsWest2RegionID)
	}

	return route53recoveryreadiness_sdkv1.New(sess.Copy(config)), nil
}
