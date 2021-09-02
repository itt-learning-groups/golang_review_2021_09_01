# Tentative/Draft Topic List for "Learn&Kount/DevOps"

This is a working draft of the topic list for the upcoming learning-group that will begin probably Sept. 29 after Golang Review finishes up.

This draft needs refinement: It's not at all in its final state yet.

Topics:

* gRPC at Kount
  * ...
* Brief anatomy of a Kount Go backend service and its deployment
  * Brief intro to Docker & K8s
  * Brief intro to k8s app config & how Kount middleware is used for app config and server/logger/etc setup
  * Brief intro to CI/CD options and how Kount's choices fit into the larger CI/CD ecosystem
* Brief anatomy of a Kount backend orchestrator
* Brief anatomy of Kount's Request Router
  * Brief intro to API gateway concept
  * Overview of Request Router
  * Overview of AWS API Gateway and similar cloud api-gateway alternatives
* Brief anatomy of a Kount JS frontend service and its deployment
  * Brief intro to options for deploying a frontend/Angular app
  * How Kount uses S3 for frontend apps
  * Brief intro to Cloudfront and how Kount uses it
* Cloud infra at Kount
  * Brief intro to how Kount uses AWS
  * Brief intro to important AWS services that Kount uses
  * Intro to IAM and access policies
  * Intro to Okta at Kount for internal user-access control
    * Intro to OAuth, OIDC, SAML, and SSO
    * Brief intro to Okta
    * Brief intro to how Kount uses Okta
    * Details on how Kount uses Okta and AWS for internal user-access control
  * Intro to AWS CLI and how Kount configures its cloud access for users
    * Ex. 1: Use AWS CLI to set up two buckets in separate AWS accounts with shared cross-account access
  * Level-set on networking concepts
  * Intro to subnets and security groups
  * Intro to AWS accounts, VPC/networking, and cross-account access config (including transit-gateway) at Kount
    * Ex. 2: Use kount-shared-infra's transit-gateway config to debug failed cross-account access
  * Intro to DNS and Route53 at Kount
    * Ex. 3: Create/manipulate Route53 records for a Kount service/environment
  * Deep dive on EC2
    * Ex. 4: Create a subnet, security group, and "emergency/POC workstation" EC2 in kount-dev
    * Ex. 5: Create a highly-available, fault-tolerant set of EC2s in private subnets across several AZs in kount-dev
  * Deep dive on load balancing and autoscaling
    * Ex. 6: Create load-balanced, autoscaled set of EC2s in private subnets across several AZs in kount-dev
  * Deep dive on containers/Docker
  * Dockerfiles at Kount: selected special topics and debugging case-studies
    * Ex. 7: Creating or debugging a dockerfile locally
  * Intro to container orchestration concepts
  * Intro to ECR
    * Ex. 8: Create an ECR repo and use it
  * Intro to ECS and EKS
    * Ex. 9: Create an EKS cluster and use it
  * Getting AWS certified
* CI/CD at Kount
  * Gitlab intro
    * Ex. 10: Create a free Gitlab account and create a project + CI/CD pipeline
    * Ex. 11: Install and run a Gitlab runner for your project locally
  * Gitlab server at Kount
  * Gitlab runners & their config at Kount
    * Ex. 12: Case-study on debugging + updating Gitlab runner IAM policy
  * Job images, ECR, and Platform Dockerfile projects
    * Ex. 13: Create and deploy a Platform-style Dockerfile project
  * Environments at Kount
  * Project setup and config at Kount
  * CI/CD variables (both Gitlab-defined and custom) and secrets
  * Your personal space in Kount GitLab; forking projects
    * Ex. 14: Fork a Kount repo and retrofit it to deploy an exerimental/POC pipeline to an ad-hoc environment in kount-dev
  * Brief intro to the GitLab API and GitLab webhooks/plugins
  * Some useful git tactics for Kount projects
  * Garden overview and topics
    * Ex. 15: Selected case-studies in Garden-related pipeline debugging
* IaC at Kount
  * Intro to IaC concepts and tooling options
  * Intro to common IaC options for AWS
    * Ex. 16: Automate a deployment using the AWS CLI
    * Ex. 17: Automate a deployment using a Boto Python script
    * Ex. 18: Automate a deployment using the AWS Go SDK
  * Intro to Cloudformation
    * Ex. 19: Automate a deployment using Cloudformation
  * Cloudformation at Kount
  * Intro to Terraform
    * Ex. 20: Automate a deployment using Terraform
  * The Terraform platform at Kount
    * Ex. 21: Create and use a custom Terraform module
* Intro to K8s, helm, and Istio
  * ...
* Deep dive on K8s at Kount
  * ...
* Deep dive on Terraform and creating selected AWS resources for Kount
  * ...
* Metrics/logging/monitoring/alerting at Kount
  * ...
* Data warehouse at Kount
  * ...
* Misc infra & CI/CD topics at Kount
  * Intro to TLS and certificates
  * TLS and certificates at Kount
  * AWS KMS at Kount
    * Ex. 22: Create a KMS secret + alias
  * Sonarqube
    * Ex. 23: Customize the sonarqube scan for a project's deployment
  * wget & checksum
    * Ex. 24: Create a Dockerfile with a checksum verification
  * VPC endpoints
  * Cloud bad practices to avoid
  * Cloud costs/billing at Kount
  * AWS Lambda deep dive
    * Ex. 25: Create a lambda function
  * AWS DynamoDB deep dive
    * Ex. 26: Create and set up a DynamoDB table
  * AWS RDS deep dive
    * Ex. 27: Create, set up, and initialize an RDS DB cluster
  * AWS Elasticsearch at Kount
  * Sigma and data visualization at Kount
  * Kafka and AWS Kinesis at Kount
  * AWS Athena & S3
  * AWS Glue at Kount
