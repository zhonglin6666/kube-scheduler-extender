module kube-scheduler-extender

go 1.13

require (
	github.com/emicklei/go-restful v2.11.2+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/cli/v2 v2.1.1

	k8s.io/api v0.0.0-20191016225839-816a9b7df678
	k8s.io/apiextensions-apiserver v0.0.0-20191015221719-7d47edc353ef
	k8s.io/apimachinery v0.0.0-20191016225534-b1267f8c42b4
	k8s.io/apiserver v0.0.0-20191015220424-a5d070e3855f // indirect
	k8s.io/client-go v11.0.1-0.20191004102930-01520b8320fc+incompatible
	k8s.io/cloud-provider v0.0.0-20191004111010-9775d7be8494
	k8s.io/kube-openapi v0.0.0-20190918143330-0270cf2f1c1d // indirect
	k8s.io/kubernetes v1.14.8
	k8s.io/utils v0.0.0-20191010214722-8d271d903fe4
)
