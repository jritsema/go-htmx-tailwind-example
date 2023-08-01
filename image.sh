#!/bin/bash
set -e

# login to ecr
ACCOUNT=$(aws sts get-caller-identity --query Account --output text)
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com

# create ecr repo (if not exists)
REPO_NAME=$1
aws ecr describe-repositories --repository-names ${REPO_NAME} || aws ecr create-repository --repository-name ${REPO_NAME}

# write repo to file for later use
URI=$(aws ecr describe-repositories --repository-names ${REPO_NAME} | jq '.repositories[0].repositoryUri' -r)
echo ${URI} > ecr-repo

# docker buildx 
docker buildx build --push --platform linux/amd64 -t ${URI} .
