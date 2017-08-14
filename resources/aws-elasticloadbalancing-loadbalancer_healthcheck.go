package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.HealthCheck AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
type AWSElasticLoadBalancingLoadBalancer_HealthCheck struct {

	// HealthyThreshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-healthythreshold

	HealthyThreshold string `json:"HealthyThreshold"`

	// Interval AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-interval

	Interval string `json:"Interval"`

	// Target AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-target

	Target string `json:"Target"`

	// Timeout AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-timeout

	Timeout string `json:"Timeout"`

	// UnhealthyThreshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-unhealthythreshold

	UnhealthyThreshold string `json:"UnhealthyThreshold"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_HealthCheck) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.HealthCheck"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_HealthCheck) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_HealthCheckResources retrieves all AWSElasticLoadBalancingLoadBalancer_HealthCheck items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_HealthCheck(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_HealthCheck {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_HealthCheck{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_HealthCheck{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_HealthCheckWithName retrieves all AWSElasticLoadBalancingLoadBalancer_HealthCheck items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_HealthCheck(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_HealthCheck, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_HealthCheck{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_HealthCheck{}, errors.New("resource not found")

}
