module github.com/quay/quay-operator

go 1.1

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/quay/clair/v4 v4.0.0-alpha.7.0.20200717155243-2238b246a10b
	github.com/quay/config-tool v0.1.2-0.20200803212319-e26fe9918b18
	github.com/stretchr/testify v1.6.1
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/controller-runtime v0.5.0
	sigs.k8s.io/kustomize/api v0.5.0
	sigs.k8s.io/yaml v1.2.0
)
