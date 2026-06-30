# Security Policy

AlpineJudge is designed to safely execute untrusted code using strong isolation and resource constraints.

## Reporting a Vulnerability

If you discover a security issue, please report it responsibly by contacting:

<security@nsci.dev>

Please avoid publicly disclosing vulnerabilities until they have been addressed.

## Security Model

AlpineJudge enforces multiple layers of security including:

- container-based isolation with containerd integration
- strict CPU, memory, and execution limits
- no network access in execution environments
- ephemeral execution containers

