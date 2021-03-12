run:
	go run cmd/app/main.go

install:
	go mod tidy

	sudo apt-get install tor tor-geoipdb privoxy -y
	sudo mv /etc/privoxy/config /etc/privoxy/config.backup
	sudo cp privoxy.config /etc/privoxy/config
	sudo service privoxy restart 