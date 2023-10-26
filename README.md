# Cloud Engineering: Implementing and testing microservices in Go
![User Service](https://github.com/flohansen/hsfl-master-ai-cloud-engineering-01/actions/workflows/user-service.yml/badge.svg)
![Product Service](https://github.com/flohansen/hsfl-master-ai-cloud-engineering-01/actions/workflows/product-service.yml/badge.svg)
[![codecov](https://codecov.io/gh/flohansen/hsfl-master-ai-cloud-engineering-01/graph/badge.svg?token=2SLAN65JV3)](https://codecov.io/gh/flohansen/hsfl-master-ai-cloud-engineering-01)

In this example we implement microservices of a webshop. This includes
the following services:

* [User Service](src/user-service/): Authentication features like registration and login.
* [Product Service](src/product-service/): Holds detailed information about products like prices, sellers, etc.

## Developing

### Unit Testing
Mocks are generated using the [gomock](https://github.com/uber-go/mock)
framework and located at `_mocks` in every microservice source folder.
