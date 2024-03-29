#!/bin/bash

export LOG_LEVEL=error


SERVER_PORT=8081
ADDRESS="localhost:${SERVER_PORT}"
TEMP_FILE=/tmp/practicumTestTempFile

echo "Run test 1"
./cmd/metricstest -test.v -test.run=^TestIteration1$ \
            -binary-path=cmd/server/server

 echo "Run test 2"
./cmd/metricstest -test.v -test.run=^TestIteration2[AB]*$ \
  -source-path=. \
  -agent-binary-path=cmd/agent/agent

echo "Run test 3"
./cmd/metricstest -test.v -test.run=^TestIteration3[AB]*$ \
  -source-path=. \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server

echo "Run test 4"
./cmd/metricstest -test.v -test.run=^TestIteration4$ \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server \
  -server-port=$SERVER_PORT \
  -source-path=.

echo "Run test 5"
./cmd/metricstest -test.v -test.run=^TestIteration5$ \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server \
  -server-port=$SERVER_PORT \
  -source-path=.

echo "Run test 6"
TEMP_FILE=/tmp/practicumTestTempFile
./cmd/metricstest -test.v -test.run=^TestIteration6$ \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server \
  -server-port=$SERVER_PORT \
  -source-path=.

echo "Run test 7"
./cmd/metricstest -test.v -test.run=^TestIteration7$ \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server \
  -server-port=$SERVER_PORT \
  -source-path=.


echo "Run test 8"
./cmd/metricstest -test.v -test.run=^TestIteration8$ \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -server-port=$SERVER_PORT \
            -source-path=.

echo "Run test 9"
./cmd/metricstest -test.v -test.run=^TestIteration9$ \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -file-storage-path=$TEMP_FILE \
            -server-port=$SERVER_PORT \
            -source-path=.

 echo "Run test 10"
./cmd/metricstest -test.v -test.run=^TestIteration10[AB]$ \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -database-dsn='postgres://mcollector:supersecretpassword@localhost:5432/metrics?sslmode=disable' \
            -server-port=$SERVER_PORT \
            -source-path=.
 echo "Run test 11"
./cmd/metricstest -test.v -test.run=^TestIteration11$ \
            -agent-binary-path=cmd/agent/agent \
            -binary-path=cmd/server/server \
            -database-dsn='postgres://mcollector:supersecretpassword@localhost:5432/metrics?sslmode=disable' \
            -server-port=$SERVER_PORT \
            -source-path=.

echo "Run test 12"
./cmd/metricstest -test.v -test.run=^TestIteration12$ \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server \
  -database-dsn='postgres://mcollector:supersecretpassword@localhost:5432/metrics?sslmode=disable' \
  -server-port=$SERVER_PORT \
  -source-path=.

echo "Run test 13"
./cmd/metricstest -test.v -test.run=^TestIteration13$ \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server \
  -database-dsn='postgres://mcollector:supersecretpassword@localhost:5432/metrics?sslmode=disable' \
  -server-port=$SERVER_PORT \
  -source-path=.

echo "Run test 14"
./cmd/metricstest -test.v -test.run=^TestIteration14$ \
  -agent-binary-path=cmd/agent/agent \
  -binary-path=cmd/server/server \
  -database-dsn='postgres://mcollector:supersecretpassword@localhost:5432/metrics?sslmode=disable' \
  -key="${TEMP_FILE}" \
  -server-port=$SERVER_PORT \
  -source-path=.
