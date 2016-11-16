# prom-sqs-stats
A tiny service exposing some Amazon SQS metrics in Prometheus format.

```bash
go build .
```

## Build it
To build a Docker image:

```bash
make build
```

## Run it
```bash
docker run --rm -p 8080
    -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
    -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
    prom-sqs-stats:latest
        -name ${AMAZON_SQS_NAME}
        -account ${AWS_ACCOUNT_NUMBER}
        -region ${AWS_REGION}
```

Example:

```bash
docker run --rm -p 8080
    -e AWS_ACCESS_KEY_ID=AKAIZMCOZBQSMTQOLKXJ
    -e AWS_SECRET_ACCESS_KEY=ssywWrhtN6T9rCJMPiEAl1aNfw2WZCrKiPzyj0Mp
    prom-sqs-stats:latest
        -name gacgacgac-replicattion
        -account 123456789012
        -region eu-central-1
```
