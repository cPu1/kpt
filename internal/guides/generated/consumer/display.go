// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package consumer

var DisplayGuide = `
*Tools can parse and render configuration so it is easier for humans to read.*

## Topics

[kpt cfg count], [kpt cfg tree],
[kpt cfg grep], [kpt cfg cat]

## Steps

1. [Fetch a remote package](#fetch-a-remote-package)
2. [Summarize resource counts](#summarize-resource-counts)
3. [Display resources as a tree](#display-resources-as-a-tree)
4. [Filter resources](#filter-resources)
5. [Dump all resources](#dump-all-resources)

## Fetch a remote package

Packages are fetched from remote git repository subdirectories with
[kpt pkg get].  In this guide we will use the [kubernetes examples] repository
as a public package catalogue.

### Fetch Command

  kpt pkg get https://github.com/kubernetes/examples/staging/ examples

Fetch the entire examples/staging directory as a kpt package under ` + "`" + `examples` + "`" + `.
This will contain many resources.

### List Command

  tree examples

### List Output

  examples/
  ├── Kptfile
  ├── cloud-controller-manager
  │   └── persistent-volume-label-initializer-config.yaml
  ├── cluster-dns
  │   ├── README.md
  ...
  
  79 directories, 329 files

The package is composed of 79 directories and, 329 files.  This is too many
to work with using tools such as ` + "`" + `less` + "`" + `.

## Summarize resource counts

### Count Example Command 1

  kpt cfg count examples/

The [` + "`" + `kpt cfg count` + "`" + `][kpt cfg count] command summarizes the resource counts to show the shape of a
package.

### Count Example Output 1

  ...
  Deployment: 10
  Endpoints: 1
  InitializerConfiguration: 1
  Namespace: 4
  Pod: 45
  ...

### Count Example Command 2

  kpt cfg count examples/cockroachdb/

Running [` + "`" + `count` + "`" + `][kpt cfg count] on a subdirectory will summarize that directory even if
it doesn't have a Kptfile.

### Count Example Output 2

  PodDisruptionBudget: 1
  Service: 2
  StatefulSet: 1

### Count Example Command 3

  kpt cfg count examples/ --kind=false

The total aggregate resource count can be shown with ` + "`" + `--kind=false` + "`" + `

### Count Example Output 3

  201

## Display resources as a tree

### Display Command

  kpt cfg tree examples/cockroachdb/ --image --replicas

Because the raw YAML configuration may be difficult for humans to easily
view and understand, kpt provides a command for rendering configuration
as a tree.  Flags may be provided to print specific fields under the resources.

### Display Output

  examples/cockroachdb
  ├── [cockroachdb-statefulset.yaml]  Service cockroachdb
  ├── [cockroachdb-statefulset.yaml]  StatefulSet cockroachdb
  │   ├── spec.replicas: 3
  │   └── spec.template.spec.containers
  │       └── 0
  │           └── image: cockroachdb/cockroach:v1.1.0
  ├── [cockroachdb-statefulset.yaml]  PodDisruptionBudget cockroachdb-budget
  └── [cockroachdb-statefulset.yaml]  Service cockroachdb-public

In addition to the built-in printable fields, [` + "`" + `kpt cfg tree` + "`" + `][kpt cfg tree] will print
arbitrary fields by providing the ` + "`" + `--field` + "`" + ` flag.

## Filter resources

### Filter Command

  kpt cfg grep "spec.replicas>3" examples | kpt cfg tree --replicas

Grep can be used to filter resources by field values.  The output of
[` + "`" + `kpt cfg grep` + "`" + `][kpt cfg grep] is the matching full resource configuration, which
may be piped to tree for rendering.

### Filter Output

  .
  ├── storage/minio
  │   └── [minio-distributed-statefulset.yaml]  StatefulSet minio
  │       └── spec.replicas: 4
  ├── sysdig-cloud
  │   └── [sysdig-rc.yaml]  ReplicationController sysdig-agent
  │       └── spec.replicas: 100
  └── volumes/vsphere
      └── [simple-statefulset.yaml]  StatefulSet web
          └── spec.replicas: 14

## Dump all resources

### Dump Command

  kpt cfg cat examples/cockroachdb

The raw YAML configuration may be dumped using [` + "`" + `kpt cfg cat` + "`" + `][kpt cfg cat].  This will
print only the YAML for Kubernetes resources.

### Dump Output

  apiVersion: v1
  kind: Service
  metadata:
    # This service is meant to be used by clients of the database. It exposes a ClusterIP that will
    # automatically load balance connections to the different database pods.
    name: cockroachdb-public
    labels:
      app: cockroachdb
  ...
`
