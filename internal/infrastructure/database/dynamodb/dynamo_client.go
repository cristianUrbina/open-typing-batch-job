package dynamodb

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoClient() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"), // specify your region
	})
	if err != nil {
		log.Fatalf("Failed to create DynamoDB session: %v", err)
	}
	return dynamodb.New(sess)
}
