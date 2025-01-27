package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func getParameter(name string) string {
	sess := session.Must(session.NewSession())
	ssmSvc := ssm.New(sess)

	param, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		log.Fatalf("Unable to get parameter %v: %v", name, err)
	}

	return *param.Parameter.Value
}
