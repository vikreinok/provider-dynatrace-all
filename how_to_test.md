# How to Test the Dynatrace Crossplane Provider

This guide explains how to run manual and automated integration tests for the generated Dynatrace Crossplane provider using your trial credentials. Testing is done by running the provider locally on your machine while utilizing a lightweight local Kubernetes cluster (`kind`) exclusively as the API control plane.

## Prerequisites

1.  **Docker**: Ensure Docker Desktop is active and running.
2.  **Kind**: Kubernetes in Docker ([install guide](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)).
3.  **Go**: Go version `>=1.24` is required.
4.  **Kubectl**: Kubernetes command-line tool.
5.  **Helm**: Kubernetes package manager ([install guide](https://helm.sh/docs/intro/install/)).

---

---

## Dynamic SafeStart & Selective CRD Loading (Memory Optimization)

Because the complete Dynatrace provider contains all **431 resources** (resulting in 862 CRDs across cluster-scoped and namespaced variants), applying all of them simultaneously can overwhelm and crash the API server of constrained local Kubernetes clusters (like `kind` inside Docker Desktop), causing timeout/disconnect errors.

To address this, we leverage Upjet's built-in **SafeStart** feature:
1. The provider binary and schema are fully generated with all **431 resources** included.
2. When starting the provider locally via `make run`, it dynamically queries the Kubernetes API server to check which CRDs are registered.
3. It automatically **gates** (disables) all controllers whose CRDs are missing, and only activates reconciliation loops for the CRDs that are actually present in the cluster!

Therefore, we can compile the full provider, but selectively apply only the **12 relevant CRDs** to the local `kind` cluster. The manager will start up cleanly, instantly, and comfortably within memory limit!

---

## Step 1: Create a Local Control Plane

Create a minimal Kubernetes cluster using `kind`:

```sh
kind create cluster --name dynatrace-test
```

Install Crossplane into this local cluster:

```sh
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm upgrade --install crossplane crossplane-stable/crossplane \
    --namespace crossplane-system \
    --create-namespace \
    --wait
```

---

## Step 2: Apply the Custom Resource Definitions (CRDs) Selectively

Instead of applying the entire `package/crds/` folder (which would attempt to load all 862 CRDs and crash the control plane), selectively apply **only the CRDs for the 12 target resources** you are testing. 

Since these CRDs are large, we use **Server-Side Apply** (`--server-side`):

```sh
# 1. Apply Cluster-Scoped CRDs
kubectl apply --server-side -f package/crds/iam.dynatrace.crossplane.io_groups.yaml
kubectl apply --server-side -f package/crds/iam.dynatrace.crossplane.io_policies.yaml
kubectl apply --server-side -f package/crds/iam.dynatrace.crossplane.io_policyboundaries.yaml
kubectl apply --server-side -f package/crds/iam.dynatrace.crossplane.io_policybindingsv2s.yaml
kubectl apply --server-side -f package/crds/management.dynatrace.crossplane.io_zonev2s.yaml

kubectl apply --server-side -f package/crds/alerting.dynatrace.crossplane.io_alertings.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.crossplane.io_v2logspipelines.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.crossplane.io_v2spanspipelines.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.crossplane.io_v2logspipelinegroups.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.crossplane.io_v2spanspipelinegroups.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.crossplane.io_v2logsroutings.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.crossplane.io_v2spansroutings.yaml

# 2. Apply Namespaced-Scoped CRDs
kubectl apply --server-side -f package/crds/iam.dynatrace.m.crossplane.io_groups.yaml
kubectl apply --server-side -f package/crds/iam.dynatrace.m.crossplane.io_policies.yaml
kubectl apply --server-side -f package/crds/iam.dynatrace.m.crossplane.io_policyboundaries.yaml
kubectl apply --server-side -f package/crds/iam.dynatrace.m.crossplane.io_policybindingsv2s.yaml
kubectl apply --server-side -f package/crds/management.dynatrace.m.crossplane.io_zonev2s.yaml
kubectl apply --server-side -f package/crds/alerting.dynatrace.m.crossplane.io_alertings.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.m.crossplane.io_v2logspipelines.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.m.crossplane.io_v2spanspipelines.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.m.crossplane.io_v2logspipelinegroups.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.m.crossplane.io_v2spanspipelinegroups.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.m.crossplane.io_v2logsroutings.yaml
kubectl apply --server-side -f package/crds/openpipeline.dynatrace.m.crossplane.io_v2spansroutings.yaml

# 3. Apply core ProviderConfig CRDs
kubectl apply --server-side -f package/crds/dynatrace.crossplane.io_providerconfigs.yaml
kubectl apply --server-side -f package/crds/dynatrace.crossplane.io_providerconfigusages.yaml
kubectl apply --server-side -f package/crds/dynatrace.m.crossplane.io_clusterproviderconfigs.yaml
kubectl apply --server-side -f package/crds/dynatrace.m.crossplane.io_providerconfigs.yaml
kubectl apply --server-side -f package/crds/dynatrace.m.crossplane.io_providerconfigusages.yaml
```

---

## Step 3: Configure the Credentials & ProviderConfig

We will configure the provider to utilize the provided Dynatrace trial account credentials:

1. **Apply the credentials secret**:
   We have already pre-populated the `secret.yaml` file with the client ID, secret, and account ID. Apply it to the cluster:
   ```sh
   kubectl apply -f secret.yaml
   ```

2. **Apply the ProviderConfig**:
   The `providerconfig.yaml` points the provider to use the `dynatrace-creds` secret:
   ```sh
   kubectl apply -f providerconfig.yaml
   ```

---

## Step 4: Run the Provider Locally

Since we are pair-programming and running out-of-cluster, you can boot up the controller engine natively from the repository root:

```sh
export PATH=$PATH:$HOME/go/bin
make run
```

You should see standard logs indicating that the provider has successfully connected to your Kind API server and is ready to reconcile resources.

---

## Step 5: Deploy the Test Suite

We have created 9 logically ordered test manifests under `examples/cluster/iam/` for the target 12 resources. 

Deploy them all at once:

```sh
kubectl apply -f examples/namespaced/iam/
```

Watch the resources converge to `READY: True`:

```sh
watch kubectl get managed
```

Crossplane will automatically reconcile and resolve resource cross-references (such as groups, policy boundaries, and pipelines) in the correct topological order.

---

## Step 6: Verify Drift Correction

1. Go to your **Dynatrace Environment Console** (`https://iok04871.live.dynatrace.com`).
2. Manually change a value of a synchronized resource (e.g. modify the description of the `Uptest Policy`).
3. Wait for the next reconciliation cycle (defaults to 1 minute locally).
4. Verify the manual change is automatically reverted to match the definition in your Kubernetes manifest.

---

## Step 7: Complete Teardown & Cleanup

To ensure that the environment is completely clean and no resources or credentials remain active:

1. **Delete all managed resources** (this will trigger the provider to delete the corresponding configurations in Dynatrace):
   ```sh
   kubectl delete -f examples/namespaced/iam/
   ```
   Wait until all managed resources are successfully deleted from the cluster.

2. **Delete the secret and config**:
   ```sh
   kubectl delete -f secret.yaml
   ```

3. **Stop the local provider**:
   Press `Ctrl+C` in the terminal running `make run`.

4. **Delete the Kind cluster**:
   ```sh
   kind delete cluster --name dynatrace-test
   ```
