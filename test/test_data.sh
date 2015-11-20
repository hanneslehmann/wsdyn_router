#!/bin/bash
echo "Preparing Test Data within etcd"
curl -L -XPUT localhost:2379/v2/keys/directory/helloservice/ -d dir=true
curl -L -XPUT localhost:2379/v2/keys/directory/helloservice/host -d value="localhost:5553"
curl -L -XPUT localhost:2379/v2/keys/directory/helloservice/url -d value="/hello"
curl -L -XPUT localhost:2379/v2/keys/directory/movieservice/ -d dir=true
curl -L -XPUT localhost:2379/v2/keys/directory/movieservice/host -d value="localhost:5552"
curl -L -XPUT localhost:2379/v2/keys/directory/movieservice/url -d value="/movie"
