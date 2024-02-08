# Credit Card Validator with Luhn Algorithm

This project is a web-enabled microservice built in Go that validates credit card numbers using the Luhn algorithm. It exposes an HTTP endpoint to accept credit card numbers via a POST request and returns a JSON response indicating whether the provided number is valid or not.

## Getting Started

These instructions will help you set up and run the credit card validator microservice in a Docker container.

### Prerequisites

- Docker installed on your machine.

### Build and Run

1. Clone the repository:
   ```bash
   git clone https://github.com/shinusAlokin/Project1_Shinus_credit_card_validator.git


2. Navigate to project directory:
   ```bash
   cd Project1_shinus_credit_card_validator

3. Build the Docker image:
   ```bash
   docker build -t creditcardvalidator .

4. Run the Docker image:
  ```bash
 docker run -p 8080:8080 creditcardvalidator



