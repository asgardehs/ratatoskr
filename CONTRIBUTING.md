# Contributing to Ratatoskr

Ratatoskr is a small, focused library with a narrow purpose: embed a working
Python distribution in a Go binary and invoke it without CGO. Contributions
that serve that purpose are welcome. Contributions that expand it — new APIs,
additional runtimes, broader scope — will be discussed before they are
written, not after.

This document is direct about what does and does not work because the
alternative is wasted effort. Your time is worth respecting, and so is the
project's. Clear expectations serve both.

## Before You Start

A few things to know about how this project is run, so you can decide whether
contributing here is a good use of your time.

Ratatoskr is solo-maintained. Review cadence depends on the maintainer's
availability, and not every contribution will be accepted. The project has a
defined scope and a defined shape, and changes that don't fit either will be
declined — sometimes after discussion, sometimes on sight. Declining a
contribution is not a comment on its quality; it means it doesn't belong in
this particular library.

The project is a fork of an unmaintained upstream. Part of what the fork
exists to do is keep the library stable and predictable for its downstream
users. That orientation shapes what kinds of changes are welcome: bug fixes,
platform support, Python version upgrades, and documentation improvements
are always appropriate topics. Large refactors, architectural changes, and
new API surface are not, unless they have been discussed first and the
maintainer has agreed the direction is right.

## Reporting Issues

A good bug report for Ratatoskr includes, at minimum:

- Your operating system and architecture (e.g., `linux-amd64`, `darwin-arm64`)
- Your Go version (`go version`)
- The Python version Ratatoskr is bundling (from the release notes)
- The version or commit of Ratatoskr you are using
- A minimal reproduction — ideally a short Go file that demonstrates the
  problem, along with the commands you ran and the output you saw

Platform-specific surface is a real source of bugs in a library like this,
so the OS/architecture information is not optional. "It doesn't work on my
machine" without those details cannot be investigated.

If you are not sure whether something is a bug or expected behavior, open an
issue anyway and label it as a question. An unclear behavior that needs
documentation is itself a kind of bug.

## Before You Submit a PR

There are two kinds of pull requests: ones that should just be sent, and
ones that should be discussed first. Knowing which is which saves everyone's
time.

**Just send it.** Typo fixes, broken-link corrections, obvious-bug fixes
with a clear root cause, small documentation improvements, and test
additions for existing behavior don't need prior discussion. Open a PR.

**Open an issue first.** Anything that fits one of the following categories
should be discussed before code is written:

- Changes to the public API (`EmbeddedPython`, `EmbeddedFiles`, the `pip`
  package, or any exported symbol downstream code depends on)
- Adding or removing platforms from the supported matrix
- Python version upgrades (see below for the specific process)
- New features that add surface area rather than fix behavior
- Refactors that touch more than a handful of files
- Anything you're not sure about

Pull requests in these categories that arrive without prior discussion will
usually be closed without a detailed review. This is not personal, and it is
not a comment on code quality. It is a consequence of scarce review time
being spent on changes whose direction has already been agreed on. A
ten-minute issue conversation before you start can save you hours of work
on a PR that does not fit the project's direction.

### Python Version Upgrades

Python version upgrades are a specific class of change that affects every
downstream user. They are welcome, but they require a short coordination
step first. Before starting work on an upgrade:

1. Open an issue proposing the target Python version.
2. Wait for confirmation that the version is the right target (usually
   aligning with the latest
   [python-build-standalone](https://github.com/astral-sh/python-build-standalone)
   release, but not always).
3. Confirm you have the ability to test the upgrade across the full platform
   matrix, or coordinate with the maintainer on how testing will be handled.

Once the direction is agreed on, the actual upgrade is usually a workflow
change to `.github/workflows/release.yml` plus any compatibility fixes the
new Python version requires. Release notes for the upgrade should explicitly
call out the Python version change.

## Development Setup

```bash
git clone https://github.com/asgardehs/ratatoskr.git
cd ratatoskr
go build ./...
go test ./...
```

The example program under `example/` is the simplest end-to-end check that
the library is working on your machine:

```bash
go run ./example
```

### Regenerating Embedded Distributions

If you are working on code that touches the embedded Python or pip package
generation, the relevant generator entry point lives under
`./hack/build-tag.sh`. The full regeneration is network-heavy — it downloads
the full set of `python-build-standalone` archives for every supported
platform — so run it only when necessary. For most contributions, the
committed embedded data is sufficient and does not need to be regenerated.

## What Makes a Good PR

A pull request is more likely to be accepted quickly if it:

- Has a clear, narrow scope — one change per PR, not five
- Preserves the no-CGO guarantee (this is non-negotiable)
- Does not break any currently supported platform
- Keeps the public API backwards-compatible, unless the change is part of
  an agreed major version bump
- Includes tests for behavior changes, not just code changes
- Updates documentation when it changes public surface
- Has commits that are readable on their own (see conventions below)

A pull request that bundles a bug fix with a refactor with a new feature
will be asked to split before it is reviewed. Not because the work is
unwelcome, but because reviewing bundled changes is harder and slower, and
each piece deserves its own decision.

## Commits and PRs

Commit messages follow the Asgard EHS project conventions, which are also
documented in the
[brand guidelines](https://asgardehs.github.io/docs/brand/#voice-and-tone):

- Imperative mood: "Add arm64 build target," not "Added" or "Adds"
- Present tense
- Summary line under 72 characters
- If the change is not self-explanatory, a blank line followed by a
  paragraph that explains why — not what, since the diff shows what

Pull request descriptions should link to the issue they address (if one
exists), summarize the change in one or two sentences, and call out anything
a reviewer should pay particular attention to. PRs that are simply labeled
"fix bug" or "update code" will be asked for more context before review.

## Code of Conduct

Participation in this project is governed by the
[Asgard EHS Code of Conduct](CODE_OF_CONDUCT.md). By contributing, you agree
to uphold its expectations.

## License

Ratatoskr is licensed under Apache-2.0. Contributions are accepted under the
same license. By submitting a pull request, you confirm that you have the
right to submit the code and that you agree to license it under Apache-2.0.

The original upstream,
[kluctl/go-embed-python](https://github.com/kluctl/go-embed-python), is also
Apache-2.0 licensed, and continues to hold copyright on the code it
contributed. See the project [README](README.md#attribution) for full
attribution details.
