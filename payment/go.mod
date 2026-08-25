module github.com/MoMentalochka/HomeWork/payment

require (
	github.com/MoMentalochka/HomeWork/platform v0.0.0-00010101000000-000000000000
	github.com/MoMentalochka/HomeWork/shared v0.0.0-00010101000000-000000000000
	github.com/brianvoe/gofakeit/v7 v7.3.0
	github.com/caarlos0/env/v11 v11.4.1
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.12.1
	go.uber.org/zap v1.28.0
	google.golang.org/grpc v1.83.0
)

require (
	github.com/stretchr/objx v0.5.3 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260630182238-925bb5da69e7 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

go 1.26.6

replace github.com/MoMentalochka/HomeWork/shared => ../shared

replace github.com/MoMentalochka/HomeWork/platform => ../platform
