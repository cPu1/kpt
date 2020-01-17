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
package pkgdocs

var READMEShort = `Fetch, update and sync packages using git`
var READMELong = `
Packages are collections of Resource configuration stored in git repositories.
They may be an entire repo, or a subdirectory within a repo.

- **Any git repository containing Resource configuration may be used as a package**,
  no additional structure or formatting is necessary.
- **Any package may be applied with ` + "`" + `kubectl apply -R -f` + "`" + `** or ` + "`" + `kpt svr apply -R -f` + "`" + `.
- Packages **may be customized in place either manually with (e.g. ` + "`" + `vi` + "`" + `) or
  programmatically** (e.g. [setters], [functions]).

Example imperative package workflow:

1. [kpt pkg get](get.md) to get a package
2. [kpt cfg set](../cfg/set.md), [kpt fn run](../fn/run.md) or ` + "`" + `vi` + "`" + ` to modify configuration
3. ` + "`" + `git add` + "`" + ` && ` + "`" + `git commit` + "`" + `
4. ` + "`" + `kubectl apply` + "`" + ` or [kpt svr apply](../svr/apply.md) to a cluster: 
5. [kpt pkg update](update.md) to pull in new changes
6. ` + "`" + `kubectl apply` + "`" + ` or [kpt svr apply](../svr/apply.md) to a cluster

Example declarative package workflow:

1. [kpt pkg init](init.md)
2. [kpt pkg sync set](sync-set.md) dev version of a package
3. [kpt pkg sync set](sync-set.md) prod version of a package
4. ` + "`" + `git add` + "`" + ` && ` + "`" + `git commit` + "`" + `
5. [kpt svr apply --context=dev](../svr/apply.md) or ` + "`" + `kubectl apply --context dev` + "`" + ` apply to dev
6. [kpt svr apply --context=prod](../svr/apply.md) or or ` + "`" + `kubectl apply --context prod` + "`" + ` apply to prod

#### Architecture

1. **Packages artifacts are Resource configuration** (rather than DSLs, templates, etc)
    * They may also contain supplemental non-Resource artifacts (e.g. README.md, Templates, etc).

2.  **Any existing git subdirectory containing Resource configuration** may be used as a package.
    * Nothing besides a git directory containing Resource configuration is required.
    * e.g. the [examples repo](https://github.com/kubernetes/examples/staging/cockroachdb) may
      be used as a package:

          kpt pkg get https://github.com/kubernetes/examples/staging/cockroachdb \
            my-cockroachdb
          kubectl apply -R -f my-cockroachdb

3. **Packages should use git references for versioning**.
    * Package authors should use semantic versioning when publishing packages.

          kpt pkg get https://github.com/kubernetes/examples/staging/cockroachdb@VERSION \
            my-cockroachdb
          kubectl apply -R -f my-cockroachdb

4. **Packages may be modified or customized in place**.
    * It is possible to directly modify the fetched package and merge upstream updates.

5. **The same package may be fetched multiple times** to separate locations.
    * Each instance may be modified and updated independently of the others.

          # fetch an instance of a java package
          kpt pkg get https://github.com/kubernetes/examples/staging/cockroachdb db1
          # make changes...

          # fetch a second instance of a java package
          kpt pkg get https://github.com/kubernetes/examples/staging/cockroachdb db2
          # make changes...

6. **Packages may pull upstream updates after they have been fetched and modified**.
    * Specify the target version to update to, and an (optional) update strategy for how to apply the
      upstream changes.

          kpt pkg get https://github.com/kubernetes/examples/staging/cockroachdb \
            my-cockroachdb
          # make changes...
          kpt pkg update my-cockroachdb@NEW_VERSION --strategy=resource-merge
`
var READMEExamples = `
    # get the package
    export SRC_REPO=git@github.com:GoogleContainerTools/kpt.git
    $ kpt pkg get $SRC_REPO/package-examples/helloworld-set@v0.1.0 helloworld

    # pull in upstream updates by merging Resources
    $ kpt pkg update helloworld@v0.2.0 --strategy=resource-merge

    # manage a collection of packages declaratively
    $ kpt pkg init ./ --description "my package"
    $ kpt pkg sync set $SRC_REPO/package-examples/helloworld-set@v0.1.0 \
        hello-world --strategy=resource-merge
    $ kpt pkg sync ./

    # update the package with sync
    $ kpt pkg sync set $SRC_REPO/package-examples/helloworld-set@v0.2.0 \
        hello-world --strategy=resource-merge
    $ kpt pkg sync ./
`

var DescShort = `Display package descriptions`
var DescLong = `
    kpt pkg desc [DIR...]

` + "`" + `desc` + "`" + ` reads package information in given DIRs and displays it in tabular format.
Input can be a list of package directories (defaults to the current directory if not specifed).
Any directory with a Kptfile is considered to be a package.

    kpt pkg desc [DIR]...
`
var DescExamples = `
    # display description for package in current directory
    kpt pkg desc

    # display description for packages in directories with 'prod-' prefix
    kpt pkg desc prod-*
`

var DiffShort = `Show changes between local and upstream source package`
var DiffLong = `
    kpt pkg diff [LOCAL_PKG_DIR@VERSION]

Diff commands lets you answer the following questions:
  - What have I changed in my package relative to the upstream source package
  - What has changed in the upstream source package between the original source version and target version

You can specify a diffing tool with options to show the changes. By default, it
uses 'diff' commandline tool.

Args:

  LOCAL_PKG_DIR:
    Local package to compare. Command will fail if the directory doesn't exist, or does not
    contain a Kptfile.  Defaults to the current working directory.

  VERSION:
    A git tag, branch, ref or commit. Specified after the local_package with @ -- pkg_dir@version.
    Defaults to the local package version that was last fetched.

Envs:

  KPT_EXTERNAL_DIFF:
   Commandline diffing tool (diff by default) that will be used to show changes. For ex.
   # Use meld to show changes
   KPT_EXTERNAL_DIFF=meld kpt pkg diff

  KPT_EXTERNAL_DIFF_OPTS:
   Commandline options to use for the diffing tool. For ex.
   # Using "-a" diff option
   KPT_EXTERNAL_DIFF_OPTS="-a" kpt pkg diff --diff-tool meld

Flags:
  diff-type:
    The type of changes to view (local by default). Following types are supported:
	 local: shows changes in local package relative to upstream source package at original version
	 remote: shows changes in upstream source package at target version relative to original version
	 combined: shows changes in local package relative to upstream source package at target version
	 3way: shows changes in local package and source package at target version relative to original version side by side

  diff-tool:
    Commandline tool (diff by default) for showing the changes.
	# Show changes using 'meld' commandline tool
	kpt pkg diff @master --diff-tool meld

	Note that it overrides the KPT_EXTERNAL_DIFF environment variable.

  diff-opts:
    Commandline options to use with the diffing tool.
	# Show changes using "diff" with recurive options
	kpt pkg diff @master --diff-tool meld --diff-opts "-r"

	Note that it overrides the KPT_EXTERNAL_DIFF_OPTS environment variable.
`
var DiffExamples = `
    # Show changes in current package relative to upstream source package
    kpt pkg diff

    # Show changes in current package relative to upstream source package using meld tool with auto compare option.
    kpt pkg diff --diff-tool meld --diff-tool-opts "-a"

    # Show changes in upstream source package between current version and target version
    kpt pkg diff @v4.0.0 --diff-type remote

    # Show changes in current package relative to target version
    kpt pkg diff @v4.0.0 --diff-type combined

    # Show 3way changes between the local package, upstream package at original version and upstream package at target version using meld
    kpt pkg diff @v4.0.0 --diff-type 3way --diff-tool meld --diff-tool-opts "-a"`

var GetShort = `Fetch a package from a git repository`
var GetLong = `
    kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] LOCAL_DEST_DIRECTORY [flags]

  REPO_URI:

    URI of a git repository containing 1 or more packages as subdirectories.
    In most cases the .git suffix should be specified to delimit the REPO_URI from the PKG_PATH,
    but this is not required for widely recognized repo prefixes.  If get cannot parse the repo
    for the directory and version, then it will print an error asking for '.git' to be specified
    as part of the argument.
    e.g. https://github.com/kubernetes/examples.git
    Specify - to read Resources from stdin and write to a LOCAL_DEST_DIRECTORY.

  PKG_PATH:

    Path to remote subdirectory containing Kubernetes Resource configuration files or directories.
    Defaults to the root directory.
    Uses '/' as the path separator (regardless of OS).
    e.g. staging/cockroachdb

  VERSION:

    A git tag, branch, ref or commit for the remote version of the package to fetch.
    Defaults to the repository master branch.
    e.g. @master

  LOCAL_DEST_DIRECTORY:

    The local directory to write the package to.
    e.g. ./my-cockroachdb-copy

    * If the directory does NOT exist: create the specified directory and write
      the package contents to it
    * If the directory DOES exist: create a NEW directory under the specified one,
      defaulting the name to the Base of REPO/PKG_PATH
    * If the directory DOES exist and already contains a directory with the same name
      of the one that would be created: fail

  --pattern string
  
    Pattern to use for writing files.  May contain the following formatting verbs
    %n: metadata.name, %s: metadata.namespace, %k: kind (default "%n_%k.yaml")
`
var GetExamples = `
    # fetch package cockroachdb from github.com/kubernetes/examples/staging/cockroachdb
    # creates directory ./cockroachdb/ containing the package contents
    kpt pkg get https://github.com/kubernetes/examples.git/staging/cockroachdb@master ./

    # fetch a cockroachdb
    # if ./my-package doesn't exist, creates directory ./my-package/ containing the package contents
    kpt pkg get https://github.com/kubernetes/examples.git/staging/cockroachdb@master ./my-package/

    # fetch package examples from github.com/kubernetes/examples
    # creates directory ./examples fetched from the provided commit
    kpt pkg get https://github.com/kubernetes/examples.git/@8186bef8e5c0621bf80fa8106bd595aae8b62884 ./
`

var InitShort = `Initialize suggested package meta for a local config directory`
var InitLong = `
Any directory containing Kubernetes Resource Configuration may be treated as
remote package without the existence of additional packaging metadata.

* Resource Configuration may be placed anywhere under DIR as *.yaml files.
* DIR may contain additional non-Resource Configuration files.
* DIR must be pushed to a git repo or repo subdirectory.

Init will augment an existing local directory with packaging metadata to help
with discovery.

Init will:

* Create a Kptfile with package name and metadata if it doesn't exist
* Create a README.md for package documentation if it doesn't exist.


    kpt pkg init DIR [flags]

  DIR:

    Defaults to '.'. Init fails if DIR does not exist

  --description string

    short description of the package. (default "sample description")

  --name string

    package name.  defaults to the directory base name.

  --tag strings

    list of tags for the package.

  --url string

    link to page with information about the package.
`
var InitExamples = `
    # writes Kptfile package meta if not found
    kpt pkg init ./ --tag kpt.dev/app=cockroachdb --description "my cockroachdb implementation"`

var ManShort = `Format and display package documentation if it exists`
var ManLong = `
    kpt pkg man LOCAL_PKG_DIR [flags]

  LOCAL_PKG_DIR:

    local path to a package.

If package documentation is missing from the package or 'man' is not installed,
the command will fail.
`
var ManExamples = `
    # display package documentation
    kpt pkg man my-package/

    # display subpackage documentation
    kpt pkg man my-package/sub-package/
`

var SyncSetShort = `Add a sync dependency to a Kptfile`
var SyncSetLong = `
Add a sync dependency to a Kptfile.

While is it possible to directly edit the Kptfile, ` + "`" + `set` + "`" + ` can be used to add or update
Kptfile dependencies.

    kpt pkg  get REPO_URI[.git]/PKG_PATH[@VERSION] LOCAL_DEST_DIRECTORY [flags]

  REPO_URI:

    URI of a git repository containing 1 or more packages as subdirectories.
    In most cases the .git suffix should be specified to delimit the REPO_URI from the PKG_PATH,
    but this is not required for widely recognized repo prefixes.  If get cannot parse the repo
    for the directory and version, then it will print an error asking for '.git' to be specified
    as part of the argument.
    e.g. https://github.com/kubernetes/examples.git
    Specify - to read Resources from stdin and write to a LOCAL_DEST_DIRECTORY.

  PKG_PATH:

    Path to remote subdirectory containing Kubernetes Resource configuration files or directories.
    Defaults to the root directory.
    Uses '/' as the path separator (regardless of OS).
    e.g. staging/cockroachdb

  VERSION:

    A git tag, branch, ref or commit for the remote version of the package to fetch.
    Defaults to the repository master branch.
    e.g. @master

  LOCAL_DEST_DIRECTORY:

    The local directory to write the package to.
    e.g. ./my-cockroachdb-copy

    * If the directory does NOT exist: create the specified directory and write
      the package contents to it
    * If the directory DOES exist: create a NEW directory under the specified one,
      defaulting the name to the Base of REPO/PKG_PATH
    * If the directory DOES exist and already contains a directory with the same name
      of the one that would be created: fail

  --strategy:

    Controls how changes to the local package are handled.  Defaults to fast-forward.

    * resource-merge: perform a structural comparison of the original / updated Resources, and merge
	  the changes into the local package.  See ` + "`" + `kpt help apis merge3` + "`" + ` for details on merge.
    * fast-forward: fail without updating if the local package was modified since it was fetched.
    * alpha-git-patch: use 'git format-patch' and 'git am' to apply a patch of the
      changes between the source version and destination version.
      **REQUIRES THE LOCAL PACKAGE TO HAVE BEEN COMMITTED TO A LOCAL GIT REPO.**
    * force-delete-replace: THIS WILL WIPE ALL LOCAL CHANGES TO
      THE PACKAGE.  DELETE the local package at local_pkg_dir/ and replace it
      with the remote version.
`
var SyncSetExamples = `
  Create a new package and add a dependency to it

    # init a package so it can be synced
    kpt pkg init

    # add a dependency to the package
    kpt pkg sync set https://github.com/GoogleContainerTools/kpt.git/package-examples/helloworld-set \
        hello-world

    # sync the dependencies
    kpt pkg sync .

  Update an existing package dependency

    # add a dependency to an existing package
    kpt pkg sync set https://github.com/GoogleContainerTools/kpt.git/package-examples/helloworld-set@v0.2.0 \
        hello-world --strategy=resource-merge

[demo]: https://storage.googleapis.com/kpt-dev/docs/pkg-sync.gif "kpt pkg sync"`

var SyncShort = `Sync package dependencies using a manifest`
var SyncLong = `
Sync uses a manifest to manage a collection of dependencies.

The manifest declares *all* direct dependencies of a package in a Kptfile.
When ` + "`" + `sync` + "`" + ` is run, it will ensure each dependency has been fetched at the
specified ref.

This is an alternative to managing package dependencies individually using
the ` + "`" + `get` + "`" + ` and ` + "`" + `update` + "`" + ` commands.

    kpt pkg sync LOCAL_PKG_DIR [flags]

  LOCAL_PKG_DIR:
  
    Local package with dependencies to sync.  Directory must exist and contain a Kptfile.

#### Env Vars

  KPT_CACHE_DIR:
  
    Controls where to cache remote packages during updates.
    Defaults to ~/.kpt/repos/
    
#### Dependencies

For each dependency in the Kptfile, ` + "`" + `sync` + "`" + ` will ensure that it exists locally with the
matching repo and ref.

Dependencies are specified in the ` + "`" + `Kptfile` + "`" + ` ` + "`" + `dependencies` + "`" + ` field and can be added or updated
with ` + "`" + `kpt pkg sync set` + "`" + `.  e.g.

    kpt pkg sync set https://github.com/GoogleContainerTools/kpt.git/package-examples/helloworld-set \
        hello-world

Or edit the Kptfile directly:

    apiVersion: kpt.dev/v1alpha1
    kind: Kptfile
    dependencies:
    - name: hello-world
      git:
        repo: "https://github.com/GoogleContainerTools/kpt.git"
        directory: "/package-examples/helloworld-set"
        ref: "master"

Dependencies have following schema:

    name: <local path (relative to the Kptfile) to fetch the dependency to>
    git:
      repo: <git repository>
      directory: <sub-directory under the git repository>
      ref: <git reference -- e.g. tag, branch, commit, etc>
    updateStrategy: <strategy to use when updating the dependency -- see kpt help update for more details>
    ensureNotExists: <remove the dependency, mutually exclusive with git>

Dependencies maybe be updated by updating their ` + "`" + `git.ref` + "`" + ` field and running ` + "`" + `kpt pkg sync` + "`" + `
against the directory.
`
var SyncExamples = `
  Example Kptfile to sync:

    # file: my-package-dir/Kptfile

    apiVersion: kpt.dev/v1alpha1
    kind: Kptfile
    # list of dependencies to sync
    dependencies:
    - name: local/destination/dir
      git:
        # repo is the git respository
        repo: "https://github.com/pwittrock/examples"
        # directory is the git subdirectory
        directory: "staging/cockroachdb"
        # ref is the ref to fetch
        ref: "v1.0.0"
    - name: local/destination/dir1
      git:
        repo: "https://github.com/pwittrock/examples"
        directory: "staging/javaee"
        ref: "v1.0.0"
      # set the strategy for applying package updates
      updateStrategy: "resource-merge"
    - name: app2
      path: local/destination/dir2
      # declaratively delete this dependency
      ensureNotExists: true

  Example invocation:

    # print the dependencies that would be modified
    kpt pkg sync my-package-dir/ --dry-run

    # sync the dependencies
    kpt pkg sync my-package-dir/

[demo]: https://storage.googleapis.com/kpt-dev/docs/pkg-sync.gif "kpt pkg sync"`

var UpdateShort = `Update a local package with changes from a remote source repo`
var UpdateLong = `
    kpt pkg update LOCAL_PKG_DIR[@VERSION] [flags]

  LOCAL_PKG_DIR:

    Local package to update.  Directory must exist and contain a Kptfile to be updated.

  VERSION:

  	A git tag, branch, ref or commit.  Specified after the local_package with @ -- pkg@version.
    Defaults the local package version that was last fetched.

	Version types:

    * branch: update the local contents to the tip of the remote branch
    * tag: update the local contents to the remote tag
    * commit: update the local contents to the remote commit

  --strategy:

    Controls how changes to the local package are handled.  Defaults to fast-forward.

    * resource-merge: perform a structural comparison of the original / updated Resources, and merge
	  the changes into the local package.  See ` + "`" + `kpt help apis merge3` + "`" + ` for details on merge.
    * fast-forward: fail without updating if the local package was modified since it was fetched.
    * alpha-git-patch: use 'git format-patch' and 'git am' to apply a patch of the
      changes between the source version and destination version.
      **REQUIRES THE LOCAL PACKAGE TO HAVE BEEN COMMITTED TO A LOCAL GIT REPO.**
    * force-delete-replace: THIS WILL WIPE ALL LOCAL CHANGES TO
      THE PACKAGE.  DELETE the local package at local_pkg_dir/ and replace it
      with the remote version.

  -r, --repo string

    Git repo url for updating contents.  Defaults to the repo the package was fetched from.

  --dry-run

    Print the 'alpha-git-patch' strategy patch rather than merging it.

#### Env Vars

  KPT_CACHE_DIR:

    Controls where to cache remote packages when fetching them to update local packages.
    Defaults to ~/.kpt/repos/
`
var UpdateExamples = `
    # update my-package-dir/
    kpt pkg update my-package-dir/

    # update my-package-dir/ to match the v1.3 branch or tag
    kpt pkg update my-package-dir/@v1.3

    # update applying a git patch
    git add my-package-dir/
    git commit -m "package updates"
    kpt pkg  update my-package-dir/@master --strategy alpha-git-patch
`
