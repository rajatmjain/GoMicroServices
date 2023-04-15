# GoMicroServices

GoMicroServices is a collection of microservices written in Golang. This project serves as a starting point for building microservices-based applications using Golang and other complementary technologies.

## Getting Started

### Prerequisites
To run this project, you need to have the following installed on your machine:

1. Docker
2. Minikube
3. Helm

### Installation
1. Clone this repository to your local machine using `git clone https://github.com/rajatmjain/GoMicroServices.git`.
2. Build the Docker image for the project using `docker build -t gomicroservices:latest .` in the project directory.
3. Start Minikube using minikube start.
4. Deploy the application using the Kubernetes deployment and service configurations found in the k8s directory:
`kubectl apply -f kubernetes/`

5. View the application in your browser by running minikube service gomicroservices-service.
Architecture

The project is composed of the following microservices:

Health Service
Details Service
Each service has its own REST API.

## Technologies Used

1. Golang
2. Docker
3. Minikube
4. Kubernetes
5. Helm

## Contributing

Contributions are always welcome! If you'd like to contribute, please follow the contribution guidelines.