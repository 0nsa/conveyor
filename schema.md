## <a name="resource-artifact"></a>Artifact

An artifact is the result of a successful build. It represents a built Docker image and will tell what what you need to pull to obtain the image.

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **build:id** | *uuid* | unique identifier of build | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **id** | *uuid* | unique identifier of artifact | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **[image](#resource-build)** | *string* | the name of the Docker image. This can be pulled with `docker pull` | `"remind101/acme-inc:139759bd61e98faeec619c45b1060b4288952164"` |

### Artifact Info



```
GET /artifacts/{artifact_id_or_build_identity}
```


#### Curl Example

```bash
$ curl -n http://localhost:8080/artifacts/$ARTIFACT_ID_OR_BUILD_IDENTITY
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "image": "remind101/acme-inc:139759bd61e98faeec619c45b1060b4288952164",
  "build": {
    "id": "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```


## <a name="resource-build"></a>Build

A build represents a request to build a git commit for a repo.

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **branch** | *string* | the branch within the GitHub repository that the build was triggered from | `"master"` |
| **completed_at** | *nullable date-time* | when the build moved to the `"succeeded"` or `"failed"` state | `null` |
| **created_at** | *date-time* | when the build was created | `"2015-01-01T12:00:00Z"` |
| **id** | *uuid* | unique identifier of build | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **repository** | *string* | the GitHub repository that this build is for | `"remind101/acme-inc"` |
| **sha** | *string* | the git commit to build | `"139759bd61e98faeec619c45b1060b4288952164"` |
| **started_at** | *nullable date-time* | when the build moved to the `"building"` state | `null` |
| **state** | *string* | the current state of the build<br/> **one of:**`"pending"` or `"building"` or `"succeeded"` or `"failed"` | `"building"` |

### Build Create

Create a new build and start it. Note that you cannot start a new build for a sha that is already in a "pending" or "building" state. You should cancel the existing build first, or wait for it to complete. You must specify either a `branch` OR a `sha`. If you provide a `branch` but no `sha`, Conveyor will use the GitHub API to resolve the HEAD commit on that branch to a sha. If you provide a `sha` but no `branch`, branch caching will be disabled.

```
POST /builds
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **repository** | *string* | the GitHub repository that this build is for | `"remind101/acme-inc"` |


#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **branch** | *string* | the branch within the GitHub repository that the build was triggered from | `"master"` |
| **sha** | *string* | the git commit to build | `"139759bd61e98faeec619c45b1060b4288952164"` |


#### Curl Example

```bash
$ curl -n -X POST http://localhost:8080/builds \
  -d '{
  "repository": "remind101/acme-inc",
  "branch": "master",
  "sha": "139759bd61e98faeec619c45b1060b4288952164"
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "repository": "remind101/acme-inc",
  "branch": "master",
  "sha": "139759bd61e98faeec619c45b1060b4288952164",
  "state": "building",
  "created_at": "2015-01-01T12:00:00Z",
  "started_at": "2015-01-01T12:00:00Z",
  "completed_at": null
}
```

### Build Info

Info for existing build.

```
GET /builds/{build_id_or_repo_sha}
```


#### Curl Example

```bash
$ curl -n http://localhost:8080/builds/$BUILD_ID_OR_REPO_SHA
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "repository": "remind101/acme-inc",
  "branch": "master",
  "sha": "139759bd61e98faeec619c45b1060b4288952164",
  "state": "building",
  "created_at": "2015-01-01T12:00:00Z",
  "started_at": "2015-01-01T12:00:00Z",
  "completed_at": null
}
```


## <a name="resource-error"></a>Error

Defines the format that errors are returned in

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **id** | *string* | unique identifier of error | `"example"` |
| **message** | *string* | human readable message | `"example"` |


