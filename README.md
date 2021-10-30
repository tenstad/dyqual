# dyqual
[Gomega Equal](https://onsi.github.io/gomega/) with [dyff](https://github.com/homeport/dyff)

## Usage

```golang
import (
    . "github.com/onsi/gomega"
    . "github.com/tenstad/dyqual"
)
```

```golang
Expect(ACTUAL).To(Dyqual(EXPECTED))
```

## Example error message

Comparisons of two `Container`s from `k8s.io/api/core/v1`

### `Dyqual`

`Dyqual` uses [dyff](https://github.com/homeport/dyff)
to display the changes between `EXPECTED` and `ACTUAL`.

```txt
  v1.Container not as expected
    
    name
      ± value change
        - ubuntu
        + alpine
    
    ports
      + one list entry added:
        - name: http
          hostport: 0
          containerport: 80
          protocol:
          hostip:
```

### Gomega `Equal`

[Gomega](https://onsi.github.io/gomega/) `Equal` displays all struct fields,
regardless of the actual differences.

```txt
  Expected
      <v1.Container>: {
          Name: "alpine",
          Image: "",
          Command: nil,
          Args: nil,
          WorkingDir: "",
          Ports: [
              {Name: "http", HostPort: 0, ContainerPort: 80, Protocol: "", HostIP: ""},
          ],
          EnvFrom: nil,
          Env: nil,
          Resources: {Limits: nil, Requests: nil},
          VolumeMounts: nil,
          VolumeDevices: nil,
          LivenessProbe: nil,
          ReadinessProbe: nil,
          StartupProbe: nil,
          Lifecycle: nil,
          TerminationMessagePath: "",
          TerminationMessagePolicy: "",
          ImagePullPolicy: "",
          SecurityContext: nil,
          Stdin: false,
          StdinOnce: false,
          TTY: false,
      }
  to equal
      <v1.Container>: {
          Name: "ubuntu",
          Image: "",
          Command: nil,
          Args: nil,
          WorkingDir: "",
          Ports: nil,
          EnvFrom: nil,
          Env: nil,
          Resources: {Limits: nil, Requests: nil},
          VolumeMounts: nil,
          VolumeDevices: nil,
          LivenessProbe: nil,
          ReadinessProbe: nil,
          StartupProbe: nil,
          Lifecycle: nil,
          TerminationMessagePath: "",
          TerminationMessagePolicy: "",
          ImagePullPolicy: "",
          SecurityContext: nil,
          Stdin: false,
          StdinOnce: false,
          TTY: false,
      }
```
