mockgen -destination=src/mocks/mock_heartbeat.go -package=mocks -source=src/service/heartbeat.go . HealthReporter
mockgen -destination=src/mocks/mock_user.go -package=mocks -source=src/service/user.go . UserOperator
mockgen -destination=src/mocks/mock_messenger.go -package=mocks -source=src/service/messenger.go . Messenger