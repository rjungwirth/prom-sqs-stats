# prom-sqs-stats
A tiny service exposing some Amazon SQS metrics in Prometheus format.

```bash
go build .
```

## Usage

### Build it

```bash
make build
```

### Run it
```bash
docker run --rm -p 8080
    -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
    -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}
    wmgaca/prom-sqs-stats:latest
        -name ${AMAZON_SQS_NAME}
        -account ${AWS_ACCOUNT_NUMBER}
        -region ${AWS_REGION}
```

Example:

```bash
docker run --rm -p 8080
    -e AWS_ACCESS_KEY_ID=AKAIZMCOZBQSMTQOLKXJ
    -e AWS_SECRET_ACCESS_KEY=ssywWrhtN6T9rCJMPiEAl1aNfw2WZCrKiPzyj0Mp
    wmgaca/prom-sqs-stats:latest
        -name wmgaca-repl-fra
        -account 123456789012
        -region eu-central-1
```

### Query the metrics
Metrics are available under the `/metrics` endpoint:

```bash
~ curl -Ss localhost:32755/metrics | grep sqs
# HELP sqs_approx_messages Approximate number of messages in an SQS queue.
# TYPE sqs_approx_messages gauge
sqs_approx_messages{sqs_queue_name="wmgaca-repl-fra"} 0
# HELP sqs_approx_messages_delayed Approximate number of delayed messages in an SQS queue.
# TYPE sqs_approx_messages_delayed gauge
sqs_approx_messages_delayed{sqs_queue_name="wmgaca-repl-fra"} 0
# HELP sqs_approx_messages_not_visible Approximate number of not visible messages in an SQS queue.
# TYPE sqs_approx_messages_not_visible gauge
sqs_approx_messages_not_visible{sqs_queue_name="wmgaca-repl-fra"} 0
```
