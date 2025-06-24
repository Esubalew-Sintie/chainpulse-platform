
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
│   ├── dto
│   │   ├── req.go
│   │   ├── res.go
│   ├── grpcclients
│   │   ├── account.go
│   │   ├── checkout.go
│   │   ├── product.go
│   ├── middleware
│   │   ├── auth.go
│   │   ├── ratelimit.go
│   ├── proto
│   │   └── gen
│   │       └── github.com
│   │       ├── account.pb.go
│   │       ├── account_grpc.pb.go
│   │   ├── account.proto
│   └── routes
│       └── handlers
│           ├── account.go
│           ├── checkout.go
│           ├── product.go
│       ├── grpc.go
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
├── Makefile
├── README.md
├── firebase-adminsdk.json
├── go.mod
├── go.sum
├── private_key.pem
├── public_key.pem
├── rbac_model.conf

# End Directory Structure