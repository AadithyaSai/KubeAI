# KubeAI

KubeAI is a command-line tool that provides an AI-powered assistant for Kubernetes. It helps you with generating YAML manifests, kubectl commands, and other Kubernetes-related tasks.

## Installation

To install KubeAI, you need to have Go installed on your system. You can then install the tool using the following command:

```bash
go install github.com/aadithyasai/kubeai/cmd/kubeai@latest
```

## Usage

To use KubeAI, you can run the following command:

```bash
kubeai [QUESTION]
```

For example:

```bash
kubeai "create a deployment with 3 replicas of nginx"
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue on the [GitHub repository](https://github.com/aadithyasai/kubeai).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
