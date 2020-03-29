module github.com/fizx/iron-man

go 1.14

replace github.com/fizx/jarvis => ../
replace github.com/fizx/iron-man => ./

require (
	github.com/apache/thrift v0.13.1-0.20200118205551-397645ac2487
	github.com/fizx/jarvis v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis v6.15.7+incompatible // indirect
	github.com/reddit/baseplate.go v0.0.0-20200327170315-03f12f4745fb // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
