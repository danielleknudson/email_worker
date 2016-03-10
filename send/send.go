package send

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ses"
)

type Email struct {
  Recipient string `json:"recipient"`
  Sender    string `json:"sender"`
  Subject   string `json:"subject"`
  Body      string `json:"body"`
}

func SendEmail(email *Email) {
  // the file location and load default profile
  credentials.NewSharedCredentials("/Users/dani/.aws/credentials", "default")

  svc := ses.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

  params := &ses.SendEmailInput{
    Destination: &ses.Destination{
      ToAddresses: []*string{
        aws.String(email.Recipient),
        // More values...
      },
    },
    Message: &ses.Message{
      Body: &ses.Body{
        Text: &ses.Content{
          Data:    aws.String(email.Body),
          Charset: aws.String("utf-8"),
        },
      },
      Subject: &ses.Content{
        Data:    aws.String(email.Subject),
        Charset: aws.String("utf-8"),
      },
    },
    Source: aws.String(email.Sender),
    ReplyToAddresses: []*string{
      aws.String(email.Sender),
    },
  }

  resp, err := svc.SendEmail(params)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println(resp)
}
