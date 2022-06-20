package minio_learning

import (
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

/*
NOTE: Some Instructions to use minio
-------------------------------------
The MinIO Go Client SDK provides simple APIs to access any Amazon S3 compatible object storage.

Download from Github
go get github.com/minio/minio-go/v7

Initialize MinIO Client
MinIO client requires the following four parameters specified to connect to an Amazon S3 compatible object storage.

RESOURCE LINK: https://docs.min.io/docs/golang-client-quickstart-guide.html


LOGIN LINK: https://play.min.io:9443/login

USERNAME: Q3AM3UQ867SPQQA43P2F
PASSWORD: zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG

*/

func InitializeMinIOClient() {
    endpoint := "play.min.io"
    accessKeyID := "Q3AM3UQ867SPQQA43P2F"
    secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
    useSSL := true

    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("%#v\n", minioClient) // minioClient is now set up
}
