module github.com/cuijxin/redis-operator

go 1.13

require (
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/operator-framework/operator-sdk v0.17.0
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kubectl v0.17.4
	sigs.k8s.io/controller-runtime v0.5.2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
)
