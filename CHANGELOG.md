<!-- markdownlint-disable MD024 -->
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add skip validate flag for nightlies by @nicholas-fedor in [#81](https://github.com/nicholas-fedor/todo/pull/81)
- Add sbom attestation subject-checksums path by @nicholas-fedor in [#68](https://github.com/nicholas-fedor/todo/pull/68)

### Changed

- Improve conditional by @nicholas-fedor in [#106](https://github.com/nicholas-fedor/todo/pull/106)
- Improve last commit query by @nicholas-fedor in [#105](https://github.com/nicholas-fedor/todo/pull/105)
- Correct release config header by @nicholas-fedor in [#103](https://github.com/nicholas-fedor/todo/pull/103)
- Use snapshot flag instead of skip for dry runs by @nicholas-fedor in [#101](https://github.com/nicholas-fedor/todo/pull/101)
- Use boolean coercion directly by @nicholas-fedor in [#99](https://github.com/nicholas-fedor/todo/pull/99)
- Use proper GH actions glob pattern for tags by @nicholas-fedor in [#97](https://github.com/nicholas-fedor/todo/pull/97)
- Refactor clean cache by @nicholas-fedor in [#89](https://github.com/nicholas-fedor/todo/pull/89)
- Change nightly image tag by @nicholas-fedor in [#87](https://github.com/nicholas-fedor/todo/pull/87)
- Update nightly build config by @nicholas-fedor in [#85](https://github.com/nicholas-fedor/todo/pull/85)
- Correct quick start example by @nicholas-fedor in [#83](https://github.com/nicholas-fedor/todo/pull/83)
- Omit verifiable build for nightlies by @nicholas-fedor in [#79](https://github.com/nicholas-fedor/todo/pull/79)
- Restore snapshot argument for nightly builds by @nicholas-fedor in [#75](https://github.com/nicholas-fedor/todo/pull/75)
- Tighten up release and changelog configurations by @nicholas-fedor in [#65](https://github.com/nicholas-fedor/todo/pull/65)

### Chores

- Update step-security/harden-runner action to v2.19.1 by @renovate[bot] in [#112](https://github.com/nicholas-fedor/todo/pull/112)
- Update github/codeql-action digest to e46ed2c by @renovate[bot] in [#110](https://github.com/nicholas-fedor/todo/pull/110)
- Update module github.com/klauspost/compress to v1.18.6 by @renovate[bot] in [#108](https://github.com/nicholas-fedor/todo/pull/108)
- Update securego/gosec action to v2.26.1 by @renovate[bot] in [#92](https://github.com/nicholas-fedor/todo/pull/92)
- Update dependency @commitlint/cli to v20.5.2 by @renovate[bot] in [#60](https://github.com/nicholas-fedor/todo/pull/60)
- Update module github.com/mattn/go-isatty to v0.0.22 by @renovate[bot] in [#74](https://github.com/nicholas-fedor/todo/pull/74)
- Update goreleaser/goreleaser-action digest to 1a80836 by @renovate[bot] in [#72](https://github.com/nicholas-fedor/todo/pull/72)
- Add github-actions[bot] for CHANGELOG.md by @nicholas-fedor in [#69](https://github.com/nicholas-fedor/todo/pull/69)
- Add codeowners file by @nicholas-fedor in [#67](https://github.com/nicholas-fedor/todo/pull/67)
- Update orhun/git-cliff-action digest to f50e115 by @renovate[bot] in [#63](https://github.com/nicholas-fedor/todo/pull/63)

### Fixed

- Fix nightly build config and steps by @nicholas-fedor in [#77](https://github.com/nicholas-fedor/todo/pull/77)
- Fix source sbom attestation path by @nicholas-fedor in [#70](https://github.com/nicholas-fedor/todo/pull/70)

### Removed

- Remove unsupported flag option by @nicholas-fedor in [#96](https://github.com/nicholas-fedor/todo/pull/96)
- Remove invalid skip checksum by @nicholas-fedor in [#94](https://github.com/nicholas-fedor/todo/pull/94)
- Remove redundant workflow by @nicholas-fedor in [#91](https://github.com/nicholas-fedor/todo/pull/91)

## [0.1.2] - 2026-04-26

### Changed

- Disable cache by @nicholas-fedor in [#57](https://github.com/nicholas-fedor/todo/pull/57)
- Change nonroot user and document by @nicholas-fedor in [#28](https://github.com/nicholas-fedor/todo/pull/28)
- Correct goreleaser nightly configuration by @nicholas-fedor in [#26](https://github.com/nicholas-fedor/todo/pull/26)

### Chores

- Lock file maintenance by @renovate[bot] in [#61](https://github.com/nicholas-fedor/todo/pull/61)
- Update actions/setup-go action to v6 by @renovate[bot] in [#50](https://github.com/nicholas-fedor/todo/pull/50)
- Update actions/checkout action to v6 by @renovate[bot] in [#49](https://github.com/nicholas-fedor/todo/pull/49)
- Update commitlint monorepo to v20 by @renovate[bot] in [#51](https://github.com/nicholas-fedor/todo/pull/51)
- Update docker/login-action action to v4 by @renovate[bot] in [#53](https://github.com/nicholas-fedor/todo/pull/53)
- Update sigstore/cosign-installer action to v4 by @renovate[bot] in [#54](https://github.com/nicholas-fedor/todo/pull/54)
- Update module golang.org/x/net to v0.53.0 by @renovate[bot] in [#46](https://github.com/nicholas-fedor/todo/pull/46)
- Update actions/setup-go digest to 40f1582 by @renovate[bot] in [#34](https://github.com/nicholas-fedor/todo/pull/34)
- Update module github.com/templui/templui to v1.10.0 by @renovate[bot] in [#43](https://github.com/nicholas-fedor/todo/pull/43)
- Update module github.com/gofiber/fiber/v3 to v3.2.0 by @renovate[bot] in [#42](https://github.com/nicholas-fedor/todo/pull/42)
- Update module github.com/gofiber/utils/v2 to v2.0.4 by @renovate[bot] in [#39](https://github.com/nicholas-fedor/todo/pull/39)
- Update dependency typescript to v6 by @renovate[bot] in [#52](https://github.com/nicholas-fedor/todo/pull/52)
- Lock file maintenance by @renovate[bot] in [#55](https://github.com/nicholas-fedor/todo/pull/55)
- Update module golang.org/x/text to v0.36.0 by @renovate[bot] in [#48](https://github.com/nicholas-fedor/todo/pull/48)
- Update module github.com/valyala/fasthttp to v1.70.0 by @renovate[bot] in [#44](https://github.com/nicholas-fedor/todo/pull/44)
- Update module github.com/mattn/go-isatty to v0.0.21 by @renovate[bot] in [#40](https://github.com/nicholas-fedor/todo/pull/40)
- Update module github.com/tinylib/msgp to v1.6.4 by @renovate[bot] in [#41](https://github.com/nicholas-fedor/todo/pull/41)
- Update module github.com/gofiber/schema to v1.7.1 by @renovate[bot] in [#38](https://github.com/nicholas-fedor/todo/pull/38)
- Update docker/login-action digest to c94ce9f by @renovate[bot] in [#35](https://github.com/nicholas-fedor/todo/pull/35)
- Update sigstore/cosign-installer digest to 398d4b0 by @renovate[bot] in [#37](https://github.com/nicholas-fedor/todo/pull/37)
- Update securego/gosec digest to 223e19b by @renovate[bot] in [#36](https://github.com/nicholas-fedor/todo/pull/36)
- Update actions/checkout digest to 34e1148 by @renovate[bot] in [#31](https://github.com/nicholas-fedor/todo/pull/31)
- Pin dependencies by @renovate[bot] in [#30](https://github.com/nicholas-fedor/todo/pull/30)
- Add gitattributes by @nicholas-fedor in [#24](https://github.com/nicholas-fedor/todo/pull/24)

### New Contributors

- @renovate[bot] made their first contribution in [#61](https://github.com/nicholas-fedor/todo/pull/61)

## [0.1.1] - 2026-04-25

### Added

- Add volume mountpoint by @nicholas-fedor in [#20](https://github.com/nicholas-fedor/todo/pull/20)

### Changed

- Enable cosign image digest signing by @nicholas-fedor in [#17](https://github.com/nicholas-fedor/todo/pull/17)

### Chores

- Update documentation and build info by @nicholas-fedor in [#22](https://github.com/nicholas-fedor/todo/pull/22)
- Fix broken comment by @nicholas-fedor in [#16](https://github.com/nicholas-fedor/todo/pull/16)

## [0.1.0] - 2026-04-25

### Added

- Add changelog pr automerge step by @nicholas-fedor in [#4](https://github.com/nicholas-fedor/todo/pull/4)

### Changed

- Update Dockerfile to support GoReleaser dockers_v2 by @nicholas-fedor in [#14](https://github.com/nicholas-fedor/todo/pull/14)
- Set build id to static value by @nicholas-fedor in [#10](https://github.com/nicholas-fedor/todo/pull/10)
- Enable verifiable builds by using package directory for main by @nicholas-fedor in [#8](https://github.com/nicholas-fedor/todo/pull/8)
- Reorganize project structure and add CI/CD infrastructure by @nicholas-fedor in [#1](https://github.com/nicholas-fedor/todo/pull/1)
- Initialize project structure by @nicholas-fedor
- Initial commit by @nicholas-fedor

### Fixed

- Fix git-cliff null reference error by @nicholas-fedor in [#2](https://github.com/nicholas-fedor/todo/pull/2)

### Removed

- Remove template substitution by @nicholas-fedor in [#12](https://github.com/nicholas-fedor/todo/pull/12)
- Remove conflicting gomod.mod flag by @nicholas-fedor in [#6](https://github.com/nicholas-fedor/todo/pull/6)

### New Contributors

- @nicholas-fedor made their first contribution in [#14](https://github.com/nicholas-fedor/todo/pull/14)
- @github-actions[bot] made their first contribution in [#13](https://github.com/nicholas-fedor/todo/pull/13)

## Compare Releases

- [unreleased](https://github.com/nicholas-fedor/todo/compare/v0.1.2...HEAD)
- [0.1.2](https://github.com/nicholas-fedor/todo/compare/v0.1.1...v0.1.2)
- [0.1.1](https://github.com/nicholas-fedor/todo/compare/v0.1.0...v0.1.1)

<!-- generated by git-cliff -->
