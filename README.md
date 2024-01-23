# Thumbnail Generator Worker

## Overview

`thumbnail-generator-worker` is a Golang application designed to consume messages from AWS SNS (Simple Notification Service) and generate thumbnails from images stored in an S3 bucket. This microservice is an integral component of my exploration into the `claim check` event architecture pattern.

## Features

- Consumes messages from AWS SNS.
- Retrieves images from an S3 bucket.
- Generates thumbnails for the retrieved images.

## Prerequisites

Before running the application, make sure you have:

- Golang installed on your machine.
- AWS credentials properly configured.
- Necessary dependencies installed. (List them if any)

## Installation

Clone the repository:

```bash
git clone https://github.com/your_username/thumbnail-generator-worker.git
cd thumbnail-generator-worker
```
