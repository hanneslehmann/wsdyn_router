#!/bin/bash
echo "Preparing Test Data within etcd"
curl -L -XPUT localhost:2379/v2/keys/directory/helloservice/ -d dir=true
curl -L -XPUT localhost:2379/v2/keys/directory/helloservice/url -d value="localhost:5552"
curl -L -XPUT localhost:2379/v2/keys/directory/helloservice/host -d value="/hello"
curl -L -XPUT localhost:2379/v2/keys/directory/movieservice/ -d dir=true
curl -L -XPUT localhost:2379/v2/keys/directory/movieservice/url -d value="localhost:5552"
curl -L -XPUT localhost:2379/v2/keys/directory/movieservice/host -d value="/movie"
