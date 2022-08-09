# Golang microservice

This repo has the source code for the golang microservice exposing the REST API for the `/DevOps` path.

## Authorization
As per the problem statement, the request requires to have an API_KEY and JWT set. To obtain a JWT send a GET request to the `/token` path and it return a signed JWT issued by the same service

```shell
curl https://devsu.amacedop.xyz/token

{"accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjAwMTYzMzB9.0MrzueSl4NDyYhBO-NTZ9w8P1q_l0xLmvIvzxTzDz0U"}
```

## Deployment
This microservice has been deployed in AKS and is publicly exposed in the domain [devsu.amacedop.xyz](https://devsu.amacedop.xyz/token). It uses Terraform for IaC and is deployed as a Helm Chart using automated pipelines.