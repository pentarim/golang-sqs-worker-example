package worker

import (
	"os"

	"github.com/nabeken/aws-go-sqs/queue"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Getenv(name, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
}

func NewSQSQueue(s *sqs.SQS, name string) (*queue.Queue, error) {
	stackName := Getenv("AWS_STACK_NAME", defaultStackName)
	if stackName != "" {
		stackName += "-"
	}
	return queue.New(s, stackName+name)
}
