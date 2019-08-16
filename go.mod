module github.com/kanisterio/kanister

go 1.12

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.1.1-0.20160913182117-3b1ae45394a2
	github.com/graymeta/stow => github.com/kastenhq/stow v0.1.1-kasten
	github.com/rook/operator-kit => github.com/kastenhq/operator-kit v0.0.0-20180316185208-859e831cc18d
)

require (
	cloud.google.com/go v0.37.2 // indirect
	github.com/Azure/azure-sdk-for-go v31.1.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.9.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.2.0 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/IBM-Cloud/ibm-cloud-cli-sdk v0.3.0 // indirect
	github.com/IBM/ibmcloud-storage-volume-lib v0.0.0-20190705114640-e7a4b11c618c
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/Masterminds/sprig v2.15.0+incompatible
	github.com/aokoli/goutils v1.1.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/aws/aws-sdk-go v1.20.12
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/elazarl/goproxy v0.0.0-20190711103511-473e67f1d7d2 // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20190711103511-473e67f1d7d2 // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-openapi/errors v0.19.2 // indirect
	github.com/go-openapi/strfmt v0.18.0
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/googleapis/gnostic v0.3.0 // indirect
	github.com/graymeta/stow v0.0.0-00010101000000-000000000000
	github.com/gregjones/httpcache v0.0.0-20181110185634-c63ab54fda8f // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jarcoal/httpmock v1.0.4 // indirect
	github.com/jpillora/backoff v0.0.0-20170918002102-8eab2debe79d
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kubernetes-csi/external-snapshotter v1.1.0
	github.com/luci/go-render v0.0.0-20160219211803-9a04cc21af0f
	github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/renier/xmlrpc v0.0.0-20170708154548-ce4a1a486c03 // indirect
	github.com/rook/operator-kit v0.0.0-00010101000000-000000000000
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.1
	github.com/softlayer/softlayer-go v0.0.0-20190615201252-ba6e7f295217 // indirect
	github.com/spf13/cobra v0.0.3
	go.opencensus.io v0.20.2 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	google.golang.org/api v0.3.1
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/genproto v0.0.0-20190404172233-64821d5d2107 // indirect
	google.golang.org/grpc v1.19.1 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
	k8s.io/api v0.0.0-20190708174958-539a33f6e817
	k8s.io/apiextensions-apiserver v0.0.0-20190708181606-527eacf2d4b7
	k8s.io/apimachinery v0.0.0-20190612125636-6a5db36e93ad
	k8s.io/client-go v11.0.1-0.20190708175433-62e1c231c5dc+incompatible
	k8s.io/code-generator v0.0.0-20190814100514-5654a2f3dfb8 // indirect
	sigs.k8s.io/controller-runtime v0.0.0-00010101000000-000000000000
	sigs.k8s.io/yaml v1.1.0 // indirect
)

// Pinned to kubernetes-1.13.4
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190222213804-5cb15d344471
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190228180357-d002e88f6236
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221213512-86fb29eff628
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190228174230-b40b2a5939e4
)

replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.29.0
	k8s.io/kube-state-metrics => k8s.io/kube-state-metrics v1.6.0
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)
