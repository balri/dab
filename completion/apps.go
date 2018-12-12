package main

import "github.com/posener/complete"

var apps = []string{
	"adminer",
	"ansible",
	"ballerina",
	"chronograf",
	"consul",
	"couchdb",
	"cyberchef",
	"dive",
	"elasticsearch",
	"fn",
	"grafana",
	"huginn",
	"influxdb",
	"kafka",
	"kafkacat",
	"kafka-rest",
	"kafka-topics-ui",
	"kapacitor",
	"logspout",
	"memcached",
	"minio",
	"mysql",
	"nats",
	"ngrok",
	"nikto",
	"ntopng",
	"pgadmin",
	"portainer",
	"postgres",
	"redis",
	"remote-syslog2",
	"selenium",
	"serveo",
	"sqliv",
	"sysdig",
	"telegraf",
	"traefik",
	"vault",
	"vaultbot",
	"vyne",
	"watchtower",
	"xsstrike",
	"zookeeper",
}

func predictApps(_ complete.Args) []string {
	return apps
}