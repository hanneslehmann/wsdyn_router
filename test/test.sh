#!/bin/bash
echo "Running Tests"
curl -X GET localhost:6667/hello/hannes
echo '-> hello hannes!'
curl -X GET localhost:6667/service/bla/bla
echo
echo '-> Requested: "/service/bla/bla"'
curl -X POST localhost:6667/test/test  -d "{\"host\": \"that\", \"url\": \"1.2.3.4\"}"

