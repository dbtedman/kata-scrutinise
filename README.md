# [Scrutinize](https://github.com/dbtedman/kata-scrutinize)

> **⚠️ WARNING:** Not production ready code, instead a [Code Kata](https://github.com/dbtedman#code-kata) intended to
> hone my programming skills through practice and repetition.

[![CI GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/kata-scrutinize/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-scrutinize/actions/workflows/ci.yml)
[![sast workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-scrutinize/sast?style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/kata-scrutinize/actions/workflows/sast.yml)
![language: go](https://img.shields.io/badge/language-go-blue.svg?style=for-the-badge&logo=go)

Tool for developers to scrutinise web application information architectures.

-   [Getting Started](#getting-started)
-   [Verification](#verification)
-   [Design](#design)
-   [References](#references)
-   [License](#license)

## Getting Started

### Prepare

Begin by [installing Go](https://go.dev/doc/install) if you have not done so already.

You can test your install by calling the following command:

```shell
go version
```

Your version must be greater than or equal to the version defined in `./go.mod` file.

### Install, Verify, and Build

Install, verify, and build `./scrutinise` binary.

```shell
nvm use && make
```

### Scrutinise a URL

```shell
./scrutinise -url https://danieltedman.com
```

You will see output like the following:

```
{"level":"debug","msg":"checking: https://danieltedman.com//","time":"2022-01-17T21:52:15+10:00"}
{"level":"debug","msg":"checking: https://danieltedman.com/mailto:dbtedman@gmail.com","time":"2022-01-17T21:52:15+10:00"}
{"level":"debug","msg":"checking: http://gohugo.io","time":"2022-01-17T21:52:16+10:00"}
{"level":"debug","msg":"checking: https://danieltedman.com/my-work","time":"2022-01-17T21:52:16+10:00"}
{"level":"debug","msg":"checking: https://themes.gohugo.io/hugo-theme-hello-friend-ng/","time":"2022-01-17T21:52:16+10:00"}
{"level":"debug","msg":"checking: https://www.linkedin.com/in/dbtedman","time":"2022-01-17T21:52:16+10:00"}
{"level":"debug","msg":"checking: https://github.com/dbtedman","time":"2022-01-17T21:52:16+10:00"}
{"level":"info","msg":"Results written to file \"results.csv\"","time":"2022-01-17T21:52:17+10:00"}
```

The results are written to the file `results.csv`.

### Help

Learn about the available commands in the help menu.

```shell
./scrutinise --help
```

## Verification

### Linting

-   [Prettier](https://prettier.io)
-   [gofmt](https://pkg.go.dev/cmd/gofmt)

```shell
make lint
```

These rules can then be automatically applied:

```shell
make format
```

### Unit Testing

```shell
make test
```

## Design

### Domain Entities

_Placeholder_

### Domain Use Cases

_Placeholder_

### Gateways

_Placeholder_

### Security Mitigations

> Initially based on the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/).

#### [A01:2021-Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)

[Github Security](https://github.com/features/security) detects secrets incorrectly committed into the repository.

#### [A02:2021-Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)

_Placeholder_

#### [A03:2021-Injection](https://owasp.org/Top10/A03_2021-Injection/)

_Placeholder_

#### [A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)

_Placeholder_

#### [A05:2021-Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)

_Placeholder_

#### [A06:2021-Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)

[Snyk](https://snyk.io) and [Github Security](https://github.com/features/security) scan Gradle and NPM dependencies for know vulnerabilities and create pull requests to resolve the vulnerabilities when available.

#### [A07:2021-Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/)

_Placeholder_

#### [A08:2021-Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)

_Placeholder_

#### [A09:2021-Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)

_Placeholder_

#### [A10:2021-Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

_Placeholder_

## References

_Placeholder_

## License

The [MIT License](./LICENSE.md) is used by this project.
