# [Kata](https://github.com/dbtedman/kata) // [Scrutinize](https://github.com/dbtedman/kata-scrutinize)

> ⚠️ WARNING: Not production ready code.

[![CI GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/kata-scrutinize/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-scrutinize/actions/workflows/ci.yml)

Tool for developers to scrutinise web application information architectures.

-   [Getting Started](#getting-started)
-   [Design](#design)
-   [License](#license)

## Getting Started

```shell
nvm use && make && ./scrutinise -url https://danieltedman.com
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

## Design

### Domain Entities

| Entity | Notes |
| :----- | :---- |
| ` `    |       |

### Domain Use Cases

| Use Case | Notes |
| :------- | :---- |
| ` `      |       |

### Gateways

| Gateway | Notes |
| :------ | :---- |
| ` `     |       |

### Security Mitigations

> Initially based on the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/).

| Security Risk                                                                                                                       | Mitigation |
| :---------------------------------------------------------------------------------------------------------------------------------- | :--------- |
| [A01:2021-Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)                                           |            |
| [A02:2021-Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)                                         |            |
| [A03:2021-Injection](https://owasp.org/Top10/A03_2021-Injection/)                                                                   |            |
| [A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)                                                       |            |
| [A05:2021-Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)                                   |            |
| [A06:2021-Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)                 |            |
| [A07:2021-Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/) |            |
| [A08:2021-Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)             |            |
| [A09:2021-Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)     |            |
| [A10:2021-Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)                    |            |

## License

The [MIT License](./LICENSE.md) is used by this project.
