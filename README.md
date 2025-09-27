# Chainpulse Platform

 Chainpulse is a hybrid commerce platform that combines traditional microservice architecture with blockchain-based features like escrow payments, loyalty tokens, and product NFTs. Built using Golang, Solidity, Apache Pulsar, and Kubernetes, it demonstrates real-world architecture with advanced backend and blockchain engineering.

---

## ✨ Key Features

- 🧱 Microservices — Clean separation of concerns using Go for services like account, product, order, payments, promotions, notifications, and loyalty.
- 🔐 Secure Auth — JWT-based authentication with role-based access control (RBAC).
- 📦 gRPC Communication — Efficient, scalable service-to-service communication.
- ⚡ Event-Driven Architecture — Apache Pulsar is used for asynchronous communication between services.
- 🛠 Smart Contracts — Ethereum-based contracts for:
  - Escrow payments (`Escrow.sol`)
  - Loyalty token system (`LoyaltyToken.sol`)
  - Product NFTs (`ProductNFT.sol`)
- 🌍 Kubernetes Ready — K8s manifests and Terraform configs included for deployment to AWS/GCP.
- 🔁 CI/CD Ready — Includes Makefile and Docker support for all services.
- 🧪 Testable — Each service and smart contract is designed for unit and integration testing.

---

## 🧩 Tech Stack

| Layer          | Tech                          |
| -------------- | ----------------------------- |
| Backend    | Golang (gRPC, REST)           |
| Messaging  | Apache Pulsar                 |
| Blockchain | Solidity + Hardhat            |
| Database   | PostgreSQL                    |
| Auth       | Firebase + JWT                |
| Deployment | Docker, Kubernetes, Terraform |
| CI/CD      | GitHub Actions (planned)      |

---

## 🗂️ Project Structure

```bash
chainpulse-platform/
├── account-service/           # Auth, roles
├── product-service/           # Product listing, reviews, wishlist 
├── order-payment-service/     # Cart, order, escrow payment logic
├── notification-service/      #  notifications sms,email and push
├── api-gateway/               # Single entry point for all services
├── contracts/                 # Smart contracts written in Solidity
├── deployments/               # Kubernetes and Terraform configs
├── scripts/                   # DB seeding, contract deployment
└── pulsar/                    # Pulsar setup for event streaming

# Directory Structure

├── account-service
│   ├── cmd
│   │   ├── main.go
│   ├── internal
│   │   ├── infrastructure
│   │   │   ├── firebase
│   │   │   │   ├── firebase.go
│   │   │   │   ├── firebase_impl.go
│   │   │   └── postgres_repo
│   │   │       ├── db.go
│   │   │       ├── db_impl.go
│   │   ├── models
│   │   │   ├── errors.go
│   │   │   ├── factory.go
│   │   │   ├── type.go
│   │   ├── pkgs
│   │   │   ├── config
│   │   │   │   ├── config.go
│   │   │   ├── db
│   │   │   │   ├── db.go
│   │   │   ├── logger
│   │   │   │   ├── logger.go
│   │   │   ├── middleware
│   │   │   │   ├── middleware.go
│   │   │   │   ├── roleBasedAuth.go
│   │   │   └── validator
│   │   │       ├── validator.go
│   │   ├── routes
│   │   │   └── handlers
│   │   │       └── dto
│   │   │           ├── req.go
│   │   │           ├── res.go
│   │   │       ├── auth.go
│   │   │       ├── grpc-handler.go
│   │   │       ├── role-related.go
│   │   │   ├── app.go
│   │   │   ├── route-init.go
│   │   └── services
│   │       ├── auth_impl.go
│   │       ├── helpers.go
│   │       ├── memo-cache.go
│   │       ├── service.go
│   │       ├── service_impl.go
│   └── proto
│       ├── account.pb.go
│       ├── account.proto
│       ├── account_grpc.pb.go
│   ├── .env
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
├── api-gateway
│   ├── cmd
│   │   ├── main.go
│   ├── grpcclients
│   │   ├── account.go
│   │   ├── checkout.go
│   │   ├── product.go
│   ├── internal
│   │   ├── infrasracture
│   │   └── routes
│   │       └── handlers
│   │           └── dto
│   │               ├── req.go
│   │               ├── res.go
│   │           ├── account.go
│   │           ├── checkout.go
│   │           ├── product.go
│   │           ├── route-init.go
│   │           ├── server.go
│   │       ├── app.go
│   ├── pkg
│   │   ├── config
│   │   │   ├── config.go
│   │   ├── logger
│   │   │   ├── logger.go
│   │   └── middleware
│   │       ├── auth.go
│   │       ├── ratelimit.go
│   └── proto
│       └── gen
│           └── github.com
│           ├── account.pb.go
│           ├── account_grpc.pb.go
│       ├── account.proto
│   ├── .env
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
├── cmd
│   ├── account-service
│   │   ├── main.go
│   ├── api-gateway
│   │   ├── main.go
│   ├── loyalty-service
│   │   ├── main.go
│   ├── notification-service
│   │   ├── main.go
│   ├── order-payment-service
│   │   ├── main.go
│   ├── product-service
│   │   ├── main.go
│   └── promotion-service
│       ├── main.go
├── contracts
│   └── migrations
│       ├── 1_deploy.js
│       ├── 2_seed.js
│   ├── Escrow.sol
│   ├── LoyaltyToken.sol
│   ├── ProductNFT.sol
├── deployments
│   ├── k8s
│   │   ├── account.yaml
│   │   ├── loyalty.yaml
│   │   ├── notification.yaml
│   │   ├── order-payment.yaml
│   │   ├── product.yaml
│   │   ├── promotion.yaml
│   │   ├── pulsar.yaml
│   └── terraform
│       ├── aws
│       └── gcp
├── loyalty-service
│   ├── internal
│   │   ├── infrastructure
│   │   │   └── db
│   │   │       ├── postgres.go
│   │   ├── models
│   │   │   ├── coin.go
│   │   │   ├── reward.go
│   │   ├── routes
│   │   │   └── handlers
│   │   │       ├── handler.go
│   │   └── services
│   │       ├── loyalty_service.go
│   └── proto
│       ├── loyalty.proto
│   ├── Dockerfile
├── nest-api-gateway
│   ├── .vscode
│   │   ├── settings.json
│   ├── src
│   │   ├── decorators
│   │   │   ├── admin.decorator.ts
│   │   │   ├── buyer.decorator.ts
│   │   │   ├── permission.decorator.ts
│   │   │   ├── public.decorator.ts
│   │   │   ├── seller.decorator.ts
│   │   ├── events
│   │   │   ├── sms.event.ts
│   │   ├── guards
│   │   │   ├── auth.guard.ts
│   │   │   ├── jwt.guard.ts
│   │   │   ├── local-auth.guard.ts
│   │   ├── listeners
│   │   │   ├── sms.listener.ts
│   │   ├── middlewares
│   │   │   ├── logging.middleware.ts
│   │   ├── modules
│   │   │   ├── account
│   │   │   │   ├── dto
│   │   │   │   │   ├── create-account.dto.ts
│   │   │   │   │   ├── update-account.dto.ts
│   │   │   │   └── entities
│   │   │   │       ├── account.entity.ts
│   │   │   │   ├── account.controller.ts
│   │   │   │   ├── account.module.ts
│   │   │   │   ├── account.service.ts
│   │   │   ├── auth
│   │   │   │   ├── dto
│   │   │   │   │   ├── Login.dto.ts
│   │   │   │   ├── enum
│   │   │   │   │   ├── Auth.enum.ts
│   │   │   │   ├── interface
│   │   │   │   │   ├── JWTPayload.ts
│   │   │   │   ├── service
│   │   │   │   │   ├── admin-auth.service.ts
│   │   │   │   │   ├── buyer-auth.service.ts
│   │   │   │   │   ├── seller-auth.service.ts
│   │   │   │   └── strategies
│   │   │   │       ├── jwt.strategy.ts
│   │   │   │       ├── local.strategy.ts
│   │   │   │   ├── auth.controller.ts
│   │   │   │   ├── auth.module.ts
│   │   │   ├── database
│   │   │   │   ├── database.config.ts
│   │   │   │   ├── database.module.ts
│   │   │   │   ├── database.providers.ts
│   │   │   └── shared
│   │   │       ├── Service
│   │   │       │   ├── interceptor.service.ts
│   │   │       │   ├── notification.service.ts
│   │   │       │   ├── sms.service.ts
│   │   │       │   ├── storage.service.ts
│   │   │       │   ├── tools.service.ts
│   │   │       ├── enums
│   │   │       └── utils
│   │   │           ├── random.util.ts
│   │   │       ├── common.controller.ts
│   │   │       ├── regex.ts
│   │   │       ├── shared.module.ts
│   │   └── proto
│   │       ├── account.proto
│   │   ├── app.module.ts
│   │   ├── main.ts
│   └── test
│       ├── app.e2e-spec.ts
│       ├── jest-e2e.json
│   ├── .env
│   ├── .env_example
│   ├── .eslintrc.js
│   ├── .gitignore
│   ├── .prettierrc
│   ├── README.md
│   ├── eslint.config.mjs
│   ├── nest-cli.json
│   ├── ormlogs.log
│   ├── package-lock.json
│   ├── package.json
│   ├── tsconfig.build.json
│   ├── tsconfig.json
├── notification-service
│   ├── internal
│   │   ├── infrastructure
│   │   │   ├── db
│   │   │   │   ├── postgres.go
│   │   │   └── email
│   │   │       ├── smtp.go
│   │   ├── models
│   │   │   ├── notification.go
│   │   ├── routes
│   │   │   └── handlers
│   │   │       ├── handler.go
│   │   └── services
│   │       ├── notification_service.go
│   └── proto
│       ├── notification.proto
│   ├── Dockerfile
├── order-payment-service
│   ├── internal
│   │   ├── infrastructure
│   │   │   ├── db
│   │   │   │   ├── postgres.go
│   │   │   ├── eth
│   │   │   │   ├── escrow.go
│   │   │   ├── grpcclient
│   │   │   │   ├── product_client.go
│   │   │   └── kafka
│   │   │       ├── publisher.go
│   │   ├── models
│   │   │   ├── cart.go
│   │   │   ├── order.go
│   │   │   ├── payment.go
│   │   ├── routes
│   │   │   └── handlers
│   │   │       ├── handler.go
│   │   └── services
│   │       ├── order_service.go
│   │       ├── payment_service.go
│   └── proto
│       ├── order_payment.proto
│   ├── Dockerfile
├── product-service
│   ├── internal
│   │   ├── infrastructure
│   │   │   ├── db
│   │   │   │   ├── postgres.go
│   │   │   └── grpcclient
│   │   │       ├── order_client.go
│   │   ├── models
│   │   │   ├── product.go
│   │   │   ├── review.go
│   │   │   ├── wishlist.go
│   │   ├── routes
│   │   │   └── handlers
│   │   │       ├── handler.go
│   │   └── services
│   │       ├── product_service.go
│   └── proto
│       ├── product.proto
│   ├── Dockerfile
├── promotion-service
│   ├── internal
│   │   ├── infrastructure
│   │   │   └── db
│   │   │       ├── postgres.go
│   │   ├── models
│   │   │   ├── banner.go
│   │   │   ├── promo.go
│   │   ├── routes
│   │   │   └── handlers
│   │   │       ├── handler.go
│   │   └── services
│   │       ├── promo_service.go
│   └── proto
│       ├── promotion.proto
│   ├── Dockerfile
├── pulsar
│   ├── docker-compose.yml
│   ├── topics.sh
└── scripts
    ├── deploy-smart-contracts.js
    ├── seed-db.go
├── .env
├── .gitignore
├── Makefile
├── README.md
├── firebase-adminsdk.json
├── go.mod
├── go.sum
├── private_key.pem
├── public_key.pem
├── rbac_model.conf

# End Directory Structure