package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Endpoint:         aws.String("http://localhost:4566"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.AnonymousCredentials,
	})

	// Crie um novo cliente S3
	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo"

	// Crie um novo bucket
	_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(s3Bucket),
	})
	if err != nil {
		fmt.Println("Erro ao criar o bucket", err)
		return
	}

	fmt.Println("Bucket criado com sucesso!")
}

func main() {
	files, err := os.ReadDir("./tmp")
	if err != nil {
		log.Printf("Error reading directory: %s\n", err)
		panic(err)
	}
	for _, file := range files {
		uploadFile(file.Name())
	}
}

func uploadFile(fileName string) {
	filePath := fmt.Sprintf("./tmp/%s", fileName)
	log.Printf("Uploading file %s to bucket %s\n", filePath, s3Bucket)
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", filePath)
		return
	}
	defer f.Close()
	s3Input := &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	}
	_, err = s3Client.PutObject(s3Input)
	if err != nil {
		log.Printf("Error uploading file: %s\n", filePath)
		return
	}
	log.Printf("File %s uploaded successfully\n", filePath)
}
