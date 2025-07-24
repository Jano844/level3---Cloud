#!/bin/bash

# Build and Push Script for Go Backend
# Usage: ./build-and-push.sh

set -e  # Exit on any error

# Configuration
IMAGE_NAME="go-backend"
DOCKERHUB_REPO="jano844/level3-test"
TAG="v1.1.4"
FULL_IMAGE_NAME="${DOCKERHUB_REPO}:${TAG}"

echo "🚀 Starting build and push process..."
echo "📦 Image: ${FULL_IMAGE_NAME}"
echo ""

# Step 1: Build the Docker image
echo "🔨 Building Docker image..."
docker buildx build --platform linux/amd64 -t ${IMAGE_NAME} .

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

# Step 2: Tag the image for Docker Hub
echo "🏷️  Tagging image for Docker Hub..."
docker tag ${IMAGE_NAME} ${FULL_IMAGE_NAME}

if [ $? -eq 0 ]; then
    echo "✅ Tagging successful!"
else
    echo "❌ Tagging failed!"
    exit 1
fi

# Step 3: Push to Docker Hub
echo "📤 Pushing to Docker Hub..."
docker push ${FULL_IMAGE_NAME}

if [ $? -eq 0 ]; then
    echo "✅ Push successful!"
    echo "🎉 Image ${FULL_IMAGE_NAME} is now available on Docker Hub!"
else
    echo "❌ Push failed!"
    echo "💡 Make sure you're logged in to Docker Hub: docker login"
    exit 1
fi

echo ""
echo "🔗 Your image is available at: https://hub.docker.com/r/${DOCKERHUB_REPO}"
echo "📋 To use this image: docker pull ${FULL_IMAGE_NAME}"
