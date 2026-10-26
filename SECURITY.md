# Cherryctl Security Policy

## Disclosing security bugs
Cherry Servers takes security of its open source projects seriously.

This project uses GitHub's private vulnerability reporting. Before submitting vulnerability reports we kindly ask you to familiarize yourself with [GitHub's documentation](https://docs.github.com/en/code-security/how-tos/report-and-fix-vulnerabilities/report-privately). 

We ask you to follow responsible disclosure principles. If your findings are deemed valid and relevant, we will do our best to acknowledge your work.

Please do not report vulnerabilities relating to dependencies of this project. Instead direct them to appropriate channels.

Your report should include the following details, at minimum:
- a detailed description
- precise versions that are vulnerable (commit hashes)
- possible impact
- suggested mitigations, if any
- reproduction procedure

We will try to acknowledge your report as soon as possible. We may ask you to provide more information regarding your report.

The maintainers will try to keep you informed on the report status as handling of the issue progresses.

## Vulnerability Management
Vulnerabilities of dependencies are managed using Dependabot.

## Build Environment Security
### Actions
Permissions in GHA workflows are scoped to least privilege. If you find an overpermissioning issue, we kindly ask you to open an issue or a pull request.

We lock our GHA dependencies using [workflow lockfiles](https://github.blog/changelog/2026-10-14-github-actions-workflow-lockfiles-in-public-preview/).

### Releasing
Release artifacts are checksummed with SHA256. Checksums files are signed with [Sigstore keyless signing](https://docs.sigstore.dev/cosign/signing/overview/). In addition, this repository uses [GitHub Artifact Attestation](https://docs.github.com/en/actions/how-tos/secure-your-work/use-artifact-attestations/use-artifact-attestations). 

Immutable releases are enabled for this repository.

## Binary files
No release binaries are available in this repository. No other binaries are available in this repositroy (images, test fixtures, etc).

## Changes
Please open an issue or a pull request if you have any suggestions regarding this policy.
