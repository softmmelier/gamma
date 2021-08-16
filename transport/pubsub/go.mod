module github.com/softmmelier/gamma/transport/redis

go 1.16

replace github.com/softmmelier/gamma/app => /home/tibi/projects/gamma/app

require (
	github.com/go-redis/redis/v8 v8.11.3
	github.com/softmmelier/gamma/app v0.0.0-00010101000000-000000000000
)
