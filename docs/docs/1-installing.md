---
sidebar_position: 1
sidebar_label: Installation
---
# Installation Steps

To get Testkube up and running you need to:

1. Install the Testkube CLI.
2. Use HELM or the Testkube CLI to to install Testkube Server components in your cluster.
3. (optional) Configure Testkube's Dashboard UI Ingress for your ingress-controller, if needed.

Watch the full installation video from our product experts: [Testkube Installation Video](https://www.youtube.com/watch?v=bjQboi3Etys).

## **Installing the Testkube CLI**
Package dependencies:
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/)

Installing the Testkube CLI with Chocolatey and Homebrew will automatically install these dependencies if they are not present. For Linux-based systems please install them manually in advance.

### **From Scripts**
To install on Linux or MacOs, run
```bash
bash < <(curl -sSLf https://kubeshop.github.io/testkube/install.sh )
```
### **Through Package Managers**
#### **Homebrew (MacOS)**

You can install Testkube from [Homebrew](https://brew.sh/):
```bash
brew install testkube
```
Or directly from our tap. The Homebrew mantainers take a few days or a week to approve each one of our releases so you can use our tap to make sure you always have the most recent release.
```bash
brew tap kubeshop/homebrew-testkube
brew install kubeshop/testkube/testkube
```


#### **Chocolatey (Windows)**

**Using [Chocolatey](https://chocolatey.org/install):**

```bash
choco source add --name=testkube_repo --source=http://chocolatey.testkube.io/chocolatey
choco install testkube
```

#### **APT (Debian/Ubuntu)**

1. Download our public GPG key, and add it to the trusted keys:
```bash
wget -qO - https://repo.testkube.io/key.pub | sudo apt-key add -
```
2. Add our repository to your apt sources:
```bash
echo "deb https://repo.testkube.io/linux linux main" | sudo tee -a /etc/apt/sources.list
```
3. Make sure to get the updates:
```bash
sudo apt-get update
```

4. Install Testkube:
```bash
sudo apt-get install -y testkube
```

### **Manual Download**

If you don't want to use scripts or package managers you can always do a manual install:

1. Download the binary for the version and platform of your choice [here](https://github.com/kubeshop/testkube/releases)
2. Unpack it. For example, in Linux use (tar -zxvf testkube_1.5.1_Linux_arm64.tar.gz)
3. Move it to a location in the PATH. For example, `mv  testkube_0.6.5_Linux_arm64/kubectl-testkube /usr/local/bin/kubectl-testkube`.

For Windows, you will need to unpack the binary and add it to the `%PATH%` as well.

If you use a package manager that we don't support, please let us know here [#161](https://github.com/kubeshop/testkube/issues/161).


### **Testkube Server Components**
To deploy Testkube to your K8s cluster you will need the following packages installed:
- [Kubectl docs](https://kubernetes.io/docs/tasks/tools/) 
- [Helm docs](https://helm.sh/docs/intro/install/#through-package-managers)


### **Using Testkube's CLI to Deploy the Server Components**
The Testkube CLI provides a command to easly deploy the Testkube server components to your cluster.
Run:
```bash
testkube init
```
note: you must have your KUBECONFIG pointing to the desired location of the installation.

The above command will install the following components in your Kubernetes cluster:

1. Testkube API
2. `testkube` namespace
3. CRDs for Tests, TestSuites, Executors
4. MongoDB
5. Minio - default (can be disabled with `--no-minio` flag if you want to use S3 buckets)
6. Dashboard - default (can be disabled with `--no-dasboard` flag)


Confirm that Testkube is running:

```bash
kubectl get all -n testkube
```

Output:

```bash
NAME                                           READY   STATUS    RESTARTS   AGE
pod/cert-manager-847544bbd-fw2h8               1/1     Running   0          4m51s
pod/cert-manager-cainjector-5c747645bf-qgftx   1/1     Running   0          4m51s
pod/cert-manager-webhook-77b946cb6d-dl6gb      1/1     Running   0          4m51s
pod/testkube-dashboard-748cbcbb66-q8zzp        1/1     Running   0          4m51s
pod/testkube-api-server-546777c9f7-7g4kg       1/1     Running   0          4m51s
pod/testkube-mongodb-5d95f44fdd-cxqz6          1/1     Running   0          4m51s
pod/testkube-minio-testkube-64cd475b94-562hz   1/1     Running   0          4m51s

NAME                                      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                        AGE
service/cert-manager                      ClusterIP   10.106.81.214   <none>        9402/TCP                                       2d20h
service/cert-manager-webhook              ClusterIP   10.104.228.254  <none>        443/TCP                                        2d20h
service/testkube-minio-service-testkube   NodePort    10.43.121.107   <none>        9000:31222/TCP,9090:32002/TCP,9443:32586/TCP   4m51s
service/testkube-api-server               NodePort    10.43.66.13     <none>        8088:32203/TCP                                 4m51s
service/testkube-mongodb                  ClusterIP   10.43.126.230   <none>        27017/TCP                                      4m51s
service/testkube-dashboard                NodePort    10.43.136.34    <none>        80:31991/TCP                                   4m51s

NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/cert-manager              1/1     1            1           4m51s
deployment.apps/cert-manager-cainjector   1/1     1            1           4m51s
deployment.apps/cert-manager-webhook      1/1     1            1           4m51s
deployment.apps/testkube-dashboard        1/1     1            1           4m51s
deployment.apps/testkube-api-server       1/1     1            1           4m51s
deployment.apps/testkube-mongodb          1/1     1            1           4m51s
deployment.apps/testkube-minio-testkube   1/1     1            1           4m51s

NAME                                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/cert-manager-847544bbd               1         1         1       4m51s
replicaset.apps/cert-manager-cainjector-5c747645bf   1         1         1       4m51s
replicaset.apps/cert-manager-webhook-77b946cb6d      1         1         1       4m51s
replicaset.apps/testkube-dashboard-748cbcbb66        1         1         1       4m51s
replicaset.apps/testkube-api-server-546777c9f7       1         1         1       4m51s
replicaset.apps/testkube-mongodb-5d95f44fdd          1         1         1       4m51s
replicaset.apps/testkube-minio-testkube-64cd475b94   1         1         1       4m51s
```

By default Testkube is installed in the `testkube` namespace.

### **Using HELM to Deploy the Server Components**
1. Add the Kubeshop Helm repository as follows:
```bash
helm repo add testkube https://kubeshop.github.io/helm-charts
```

If this repo already exists, run `helm repo update` to retrieve
the `latest` versions of the packages.  You can then run `helm search repo
testkube` to see the charts.

2. To install the `testkube` chart:

```bash
helm install --create-namespace my-testkube testkube/testkube
```

Please note that, by default, the namespace for the intstallation will be `testkube`. If the `testkube` namespace does not exist, it will be created for you.

If you wish to install into a different namespace, please use following command:

```bash
helm install --namespace namespace_name my-testkube testkube/testkube
```

To uninstall the `testkube` chart if it was installed into default namespace:

```bash
helm delete my-testkube testkube/testkube
```

And from a namespace other than `testkube`:

```bash
helm delete --namespace namespace_name my-testkube testkube/testkube
```

#### **Helm Properties**

The following Helm defaults are used in the `testkube` chart:

| Parameter                            | Is optional | Default                              |
| ------------------------------------ | ----------- | ------------------------------------ |
| mongodb.auth.enabled                 | yes         | false                                |
| mongodb.service.port                 | yes         | "27017"                              |
| mongodb.service.portName             | yes         | "mongodb"                            |
| mongodb.service.nodePort             | yes         | true                                 |
| mongodb.service.clusterIP            | yes         | ""                                   |
| mongodb.nameOverride                 | yes         | "mongodb"                            |
| mongodb.fullnameOverride             | yes         | "testkube-mongodb"                   |
| testkube-api.image.repository        | yes         | "kubeshop/testkube-api-server"       |
| testkube-api.image.pullPolicy        | yes         | "Always"                             |
| testkube-api.image.tag               | yes         | "latest"                             |
| testkube-api.service.type            | yes         | "NodePort"                           |
| testkube-api.service.port            | yes         | 8088                                 |
| testkube-api.mongodb.dsn             | yes         | "mongodb://testkube-mongodb:27017"   |
| testkube-api.nats.uri                | yes         | "nats://testkube-nats"               |
| testkube-api.telemetryEnabled        | yes         | true                                 |
| testkube-api.storage.endpoint        | yes         | testkube-minio-service-testkube:9000 |
| testkube-api.storage.accessKeyId     | yes         | minio                                |
| testkube-api.storage.accessKey       | yes         | minio123                             |
| testkube-api.storage.scrapperEnabled | yes         | true                                 |
| testkube-api.slackToken              | yes         | ""                                   |
| testkube-api.slackChannelId          | yes         | ""                                   |

>For more configuration parameters of `MongoDB` chart please visit:
<https://github.com/bitnami/charts/tree/master/bitnami/mongodb#parameters>

>For more configuration parameters of `NATS` chart please visit:
<https://docs.nats.io/running-a-nats-service/nats-kubernetes/helm-charts>


## **Remove Testkube Server Components**
### **Using Helm:**
```bash
helm delete testkube
```
### **Using Testkube's CLI:**
```bash
testkube purge
```

## Intallation on OpenShift 

Because of upgrade issues from Mongo 11 to 13 Testkube can't work on root-less OpenShift environment by default. Fortunately we was able to install it manually. 

To do it we'll need empty OpenShift cluster and follow steps below: 

1. Save mongo chart values (let's name it `values.yaml`)

```yaml 
securityContext:
  enabled: true
  fsGroup: 1000650001
  runAsUser: 1000650001

podSecurityContext:
  enabled: false

containerSecurityContext:
  enabled: true
  runAsUser: 1000650001
  runAsNonRoot: true

volumePermissions:
  enabled: false

auth: 
   enabled: false
```

2. Install MongoDB 

```sh
helm install testkube-mongodb bitnami/mongodb --namespace=testkube --values values.yaml
```

3. Install Testkube configured to use our Custom MongoDB instance

```
helm install --create-namespace --namespace testkube testkube testkube/testkube --set mongodb.enabled=false --set testkube-dashboard.service.port=8080
```

Please notice that we've just installed MongoDB with `testkube-mongodb` helm release name it'll allow us to not reconfigure Testkube API MongoDB connection URI. If you've intalled with different name / namespace please adjust `--set testkube-api.mongodb.dsn: "mongodb://testkube-mongodb:27017"` to your MongoDB service.

That's all your Testkube instance should be able to run on Openshift now. 