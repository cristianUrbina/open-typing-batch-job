package dynamodb

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type CodeSnippetRepository struct {
	svc       dynamodbiface.DynamoDBAPI
	tableName string
}

func NewCodeSnippetRepository(db *dynamodb.DynamoDB) *CodeSnippetRepository {
	return &CodeSnippetRepository{
		svc:       db,
		tableName: *aws.String("code_snippets"),
	}
}

func (r *CodeSnippetRepository) Save(snippet domain.CodeSnippet) error {
	_, err := r.svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(snippet.Name),
			},
			"Repository": {
				S: aws.String(snippet.Repository),
			},
			"RepoDir": {
				S: aws.String(snippet.RepoDir),
			},
			"Content": {
				S: aws.String(snippet.Content),
			},
			"Language": {
				S: aws.String(snippet.Language),
			},
		},
	})
	return err
}

func (r *CodeSnippetRepository) GetByRepository(repoName string) ([]domain.CodeSnippet, error) {
	result, err := r.svc.Scan(&dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("Repository = :repoName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":repoName": {
				S: aws.String(repoName),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var snippets []domain.CodeSnippet
	for _, item := range result.Items {
		snippets = append(snippets, domain.CodeSnippet{
			Name:       *item["Name"].S,
			Content:    *item["Content"].S,
			Language:   *item["Language"].S,
			Repository: *item["Repository"].S,
			RepoDir:    *item["RepoDir"].S,
		})
	}

	return snippets, nil
}

func (r *CodeSnippetRepository) GetByFileName(fileName string) ([]domain.CodeSnippet, error) {
	result, err := r.svc.Scan(&dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("Name = :fileName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":fileName": {
				S: aws.String(fileName),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var snippets []domain.CodeSnippet
	for _, item := range result.Items {
		snippets = append(snippets, domain.CodeSnippet{
			Name:       *item["Name"].S,
			Content:    *item["Content"].S,
			Language:   *item["Language"].S,
			Repository: *item["Repository"].S,
			RepoDir:    *item["RepoDir"].S,
		})
	}

	return snippets, nil
}

func (r *CodeSnippetRepository) GetByLanguage(language string) ([]domain.CodeSnippet, error) {
	result, err := r.svc.Scan(&dynamodb.ScanInput{
		TableName:                aws.String(r.tableName),
		FilterExpression:         aws.String("#lang = :language"),
		ExpressionAttributeNames: map[string]*string{"#lang": aws.String("Language")}, // 👈 Avoid reserved keyword
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":language": {S: aws.String(language)},
		},
	})
	if err != nil {
		return nil, err
	}

	var snippets []domain.CodeSnippet
	for _, item := range result.Items {
		snippets = append(snippets, domain.CodeSnippet{
			Name:       *item["Name"].S,
			Content:    *item["Content"].S,
			Language:   *item["Language"].S,
			Repository: *item["Repository"].S,
			RepoDir:    *item["RepoDir"].S,
		})
	}

	return snippets, nil
}

func (r *CodeSnippetRepository) Delete(snippetID string) error {
	_, err := r.svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(snippetID),
			},
		},
	})
	return err
}
