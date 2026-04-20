# Ratatoskr

_Embedded Python distribution for Go applications._

![Status](https://img.shields.io/badge/status-active-2A6E3F?style=for-the-badge)

## Attribution

Ratatoskr is a fork of
[kluctl/go-embed-python](https://github.com/kluctl/go-embed-python), originally
created by the kluctl team and licensed under Apache-2.0. The upstream project
is no longer actively maintained. This fork continues the work under the
Asgard EHS project, preserving the original design while updating the library
for continued compatibility with new Python releases and supported platforms.

All credit for the original design and implementation belongs to the kluctl
authors. Ratatoskr remains licensed under Apache-2.0, matching the upstream.

## Overview

Go applications that need Python capabilities — scientific libraries,
regulatory calculation engines, or ecosystem-specific tooling that only
exists in Python — face a hard choice: demand that users install Python
themselves, or use CGO-based bindings that are fragile across platforms.
Ratatoskr embeds a complete Python distribution inside the Go binary and
invokes it via subprocess, giving you Python's capabilities without either
cost. It is a library first — import it, call it, and it handles extraction,
path management, and the subprocess lifecycle for you — with utilities for
embedding pip packages alongside the interpreter.

## When Not to Use Ratatoskr

Ratatoskr is a subprocess-based bridge, not an in-process Python runtime. It
is:

- **Not a CGO binding.** Python runs in a separate process. If you need
  in-process Python (direct memory sharing, zero-overhead calls, embedding
  Python objects as Go values), Ratatoskr is the wrong tool.
- **Not for hot-path performance.** Every Python invocation has subprocess
  startup and IPC overhead. Use Ratatoskr when the work done per call
  outweighs the call itself, not for tight inner loops.
- **Not a sandbox.** The embedded Python interpreter runs with the full
  permissions of the host process. Do not use Ratatoskr to execute
  untrusted Python code.

## Quick Example

```go
package main

import (
    "os"

    "github.com/asgardehs/ratatoskr/python"
)

func main() {
    ep, err := python.NewEmbeddedPython("example")
    if err != nil {
        panic(err)
    }

    cmd, err := ep.PythonCmd("-c", "print('hello')")
    if err != nil {
        panic(err)
    }
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        panic(err)
    }
}
```

`NewEmbeddedPython` extracts the embedded distribution to a temporary
directory on first use and skips the extraction on subsequent runs if the
previously extracted copy is intact.

## Persistent Extraction for Desktop Apps

`NewEmbeddedPython` extracts into the system temp directory (`/tmp` on Linux,
`%TEMP%` on Windows, `/var/folders/...` on macOS). That location may be wiped
on reboot or cleaned by `systemd-tmpfiles`, forcing re-extraction on every
launch. For long-lived desktop applications, use
`NewEmbeddedPythonInCacheDir` instead:

```go
ep, err := python.NewEmbeddedPythonInCacheDir("myapp")
```

This places the extracted distribution under the user's OS cache directory —
`~/.cache/myapp/python-<hash>` on Linux (respecting `$XDG_CACHE_HOME`),
`~/Library/Caches/myapp/...` on macOS, and `%LOCALAPPDATA%\myapp\...` on
Windows. The extraction cost is paid once per installed version, and survives
reboots.

## Installation

```bash
go get github.com/asgardehs/ratatoskr
```

Requires Go 1.26 or later. No CGO — Ratatoskr uses subprocess IPC rather
than Python bindings, so cross-compilation is straightforward.

## Building from Source

```bash
git clone https://github.com/asgardehs/ratatoskr.git
cd ratatoskr
go build ./...
```

## Supported Platforms

| OS      | Architecture   |
| ------- | -------------- |
| Linux   | amd64, arm64   |
| macOS   | amd64, arm64   |
| Windows | amd64          |

Platform support follows the upstream
[python-build-standalone](https://github.com/astral-sh/python-build-standalone)
distributions that Ratatoskr embeds. Adding a platform means adding it to the
release workflow matrix; it is not a code change.

## How It Works

Ratatoskr uses the standalone Python distributions published by
[astral-sh/python-build-standalone](https://github.com/astral-sh/python-build-standalone).
At build time, the release workflow downloads, extracts, and packages the
supported distributions, which are then embedded into the Go binary using
`//go:embed`.

At runtime, `NewEmbeddedPython` extracts the embedded distribution into a
temporary folder the first time it is called. Subsequent calls reuse the
extracted distribution after verifying its integrity, so the extraction cost
is paid once per install, not once per invocation. The `EmbeddedPython`
object then exposes the interpreter as a helper for constructing
`exec.Cmd`-style Python invocations.

## Embedding Python Libraries

To bundle pip packages alongside the interpreter, create a generator under
your application (for example, `internal/pylibs/generate/main.go`):

```go
package main

import "github.com/asgardehs/ratatoskr/pip"

func main() {
    err := pip.CreateEmbeddedPipPackagesForKnownPlatforms(
        "requirements.txt",
        "./data/",
    )
    if err != nil {
        panic(err)
    }
}
```

Add a `go:generate` directive in a sibling file (`internal/pylibs/gen.go`):

```go
package pylibs

//go:generate go run -tags ratatoskr_embed ./generate
```

And a `requirements.txt` alongside it:

```
jinja2==3.1.2
```

Running `go generate -tags ratatoskr_embed ./...` populates
`internal/pylibs/data` with the platform-specific package archives. The
`ratatoskr_embed` build tag is required because the generator imports
Ratatoskr's embedded Python distribution, which is gated behind that tag;
without it the generator will fail to launch Python.

Load the generated archives at runtime via `embed_util.NewEmbeddedFiles()`
and attach the extracted path to your `EmbeddedPython` instance with
`AddPythonPath`. This is the same pattern the upstream
[go-jinja2](https://github.com/kluctl/go-jinja2) project uses.

## Releases

Ratatoskr follows [semantic versioning](https://semver.org/) on its Go API.
Each release bundles a single Python version, called out in the release
notes — a given Ratatoskr release advertises "bundles Python 3.13.x" in
its description, and upgrading Ratatoskr may upgrade Python alongside it.
Review release notes before upgrading if your application depends on
specific Python interpreter behavior.

Releases are cut manually via the `release` GitHub Actions workflow
(`workflow_dispatch`). The maintainer supplies the target version tag
(e.g. `v1.2.0`), the Python version, and the
[python-build-standalone](https://github.com/astral-sh/python-build-standalone)
release date. The workflow generates the embedded distribution, runs the
test suite on Linux, macOS, and Windows, and pushes the resulting tag.

## Project

- **License:** Apache-2.0 — see [LICENSE](LICENSE)
- **Code of Conduct:** see [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)
- **Contributing:** see [CONTRIBUTING.md](CONTRIBUTING.md)
- **Security:** report vulnerabilities to
  [muninn.developer@protonmail.com](mailto:muninn.developer@protonmail.com)

## Name

> _In Norse mythology, Ratatoskr is the squirrel who runs the length of
> Yggdrasil, the world tree, carrying messages between the eagle at its
> crown and the dragon Níðhöggr at its roots. He is the messenger who
> crosses between realms that would otherwise never meet. Here, Ratatoskr
> carries calls and replies between the Go process and the Python runtime
> embedded within it._

_Part of the [Asgard EHS family](https://asgardehs.github.io/)._
