build:
	sudo docker image build -t muhammadrivaldy05/basic-k8s .

push-image:
	sudo docker image push muhammadrivaldy05/basic-k8s

migration-user:
	@read -p "input the migration name: " name; \
	migrate create -ext sql -dir migrations $$name