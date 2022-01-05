cd /tmp
git clone https://github.com/exexute/govwa.git
cd /tmp/govwa

# build govwa
go mod tidy

# run govwa
nohup go run -gcflags "all=-N -l" app.go &

sleep 60

ps aux
curl localhost:8888
