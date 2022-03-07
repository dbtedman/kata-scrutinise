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

> 🚧 Placeholder

## License

The [MIT License](./LICENSE.md) is used by this project.
