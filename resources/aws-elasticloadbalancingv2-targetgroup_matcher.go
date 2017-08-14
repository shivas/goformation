package resources

// AWS::ElasticLoadBalancingV2::TargetGroup.Matcher AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html
type AWSElasticLoadBalancingV2TargetGroupMatcher struct {

	// HttpCode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html#cfn-elasticloadbalancingv2-targetgroup-matcher-httpcode
	HttpCode string `json:"HttpCode"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2TargetGroupMatcher) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::TargetGroup.Matcher"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2TargetGroupMatcher) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}