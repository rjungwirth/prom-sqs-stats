package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Define metrics

	sqsApproxMessages = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sqs_approx_messages",
			Help: "Approximate number of messages in an SQS queue.",
		},
		[]string{"sqs_queue_name"},
	)

	sqsApproxMessagesNotVisible = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sqs_approx_messages_not_visible",
			Help: "Approximate number of not visible messages in an SQS queue.",
		},
		[]string{"sqs_queue_name"},
	)

	sqsApproxMessagesDelayed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sqs_approx_messages_delayed",
			Help: "Approximate number of delayed messages in an SQS queue.",
		},
		[]string{"sqs_queue_name"},
	)

	// Parameters

	awsAccountNumber string
	sqsQueueName     string
	awsRegion        string
)

// Initialize metrics
func init() {
	flag.StringVar(&awsAccountNumber, "account", "", "AWS acccount number")
	flag.StringVar(&sqsQueueName, "name", "", "Amazon SQS queue url")
	flag.StringVar(&awsRegion, "region", "eu-west-1", "AWS region name")

	prometheus.MustRegister(sqsApproxMessages)
	prometheus.MustRegister(sqsApproxMessagesNotVisible)
	prometheus.MustRegister(sqsApproxMessagesDelayed)
}

// Query SQS and update metrics
func queryAndUpdate(conn *sqs.SQS, queueUrl string, labels prometheus.Labels) error {
	params := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(queueUrl),
		AttributeNames: []*string{
			aws.String("ApproximateNumberOfMessages"),
			aws.String("ApproximateNumberOfMessagesNotVisible"),
			aws.String("ApproximateNumberOfMessagesDelayed"),
		},
	}

	resp, err := conn.GetQueueAttributes(params)
	if err != nil {
		return err
	}

	messages, err := strconv.ParseFloat(
		*resp.Attributes["ApproximateNumberOfMessages"], 64)
	if err != nil {
		return err
	}

	messagesNotVisible, err := strconv.ParseFloat(
		*resp.Attributes["ApproximateNumberOfMessagesNotVisible"], 64)
	if err != nil {
		return err
	}

	messagesDelayed, err := strconv.ParseFloat(
		*resp.Attributes["ApproximateNumberOfMessagesDelayed"], 64)
	if err != nil {
		return err
	}

	sqsApproxMessages.With(labels).Set(messages)
	sqsApproxMessagesNotVisible.With(labels).Set(messagesNotVisible)
	sqsApproxMessagesDelayed.With(labels).Set(messagesDelayed)

	return nil
}

func main() {
	flag.Parse()

	if sqsQueueName == "" {
		panic("-name is mandatory")
	}

	if awsAccountNumber == "" {
		panic("-account is mandatory")
	}

	sqsQueueUrl := fmt.Sprintf("https://%s.queue.amazonaws.com/%s/%s",
		awsRegion, awsAccountNumber, sqsQueueName)

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(awsRegion)})
	if err != nil {
		panic(err)
	}

	sqsConn := sqs.New(awsSession)

	go func() {
		for {
			err := queryAndUpdate(sqsConn, sqsQueueUrl, prometheus.Labels{"sqs_queue_name": sqsQueueName})
			if err != nil {
				panic(err)
			}

			time.Sleep(time.Second)
		}
	}()

	fmt.Println("Running on :8080")
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":8080", nil)
}
