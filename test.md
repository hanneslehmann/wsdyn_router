etcd
curl -X GET localhost:2379/v2/keys/directory/helloservice/url
{"action":"get","node":{"key":"/directory/helloservice/url","value":"/hello","modifiedIndex":24,"createdIndex":24}}

testserver1
curl -X GET localhost:5551/hello/chad
Hello: chad

httprouterd
curl -X GET localhost:6667/services/helloservice/chad
Requested: "/services/helloservice/chad"
