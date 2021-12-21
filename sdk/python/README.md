# Vultr Resource Provider

The Gandi Resource Provider lets you manage
[Gandi](https://www.gandi.com) resources. This is not an official
provider and I have only very limited time to work on it. It works for
me and that's the main goal.

## Installing

This package was not uploaded to language-specific registries. While
the SDK builds for Node.JS and .NET, I don't know how to install the
resulting SDK without pushing it to NPM and NuGet. Feel free to open a
pull request if you know how to do that in a way similar to Python.

### Python

To use from Python, install using `pip`:

```bash
pip install git+https://github.com/vincentbernat/pulumi-gandi.git#egg=pulumi_gandi&subdirectory=sdk/python
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/vincentbernat/pulumi-gandi/sdk/go/...
```

## Configuration

There is configuration points. However, it needs the `GANDI_KEY`
environment variable set to work.

## Reference

The detailed reference documentation is not published.
