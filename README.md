# Template Wizard CLI project

The Template Wizard CLI is a simple command-line tool written in Go that helps you quickly generate messages from predefined templates. It allows you to select sections from a JSON template file, replace placeholders with user-provided data, and copy the final message to your clipboard for easy pasting.

### Dockerized Deployment

This project is containerized using a multi-stage Docker build. The first stage uses the official Go image to build the binary, and the second stage packages it into a minimal Alpine Linux image. To build and run the container, simply use:

```docker build -t template-wizard-cli .```
```docker run --rm -ti template-wizard-cli```