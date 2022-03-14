#!/bin/sh
set -xeuo pipefail

mkdir -p ~/.kube
mkdir -p ~/.aws

if [ -z ${K8S_CONFIG+x} ]
then
  echo "K8S_CONFIG env is not set"; false
fi

if [ -z ${AWS_CONFIG+x} ]
then
  echo "AWS_CONFIG env is not set"; false
fi

echo $K8S_CONFIG | base64 -d > ~/.kube/config

echo $AWS_CONFIG | base64 -d > ~/.aws/credentials

./app $@
