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

- Use actionlint-action instead of manual download by @nicholas-fedor in [#133](https://github.com/nicholas-fedor/todo/pull/133)
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

- Update actions/setup-go action to v7 by @renovate[bot] in [#320](https://github.com/nicholas-fedor/todo/pull/320)
- Update module github.com/mattn/go-isatty to v0.0.23 by @renovate[bot] in [#321](https://github.com/nicholas-fedor/todo/pull/321)
- Update module github.com/gofiber/utils/v2 to v2.2.0 by @renovate[bot] in [#318](https://github.com/nicholas-fedor/todo/pull/318)
- Update securego/gosec action to v2.28.0 by @renovate[bot] in [#316](https://github.com/nicholas-fedor/todo/pull/316)
- Lock file maintenance by @renovate[bot] in [#287](https://github.com/nicholas-fedor/todo/pull/287)
- Update module github.com/gofiber/schema to v1.8.2 by @renovate[bot] in [#314](https://github.com/nicholas-fedor/todo/pull/314)
- Update module github.com/gofiber/utils/v2 to v2.1.2 by @renovate[bot] in [#312](https://github.com/nicholas-fedor/todo/pull/312)
- Update dependency typescript to v7 by @renovate[bot] in [#310](https://github.com/nicholas-fedor/todo/pull/310)
- Update dependency @commitlint/cli to v21.2.1 by @renovate[bot] in [#308](https://github.com/nicholas-fedor/todo/pull/308)
- Update module golang.org/x/net to v0.57.0 by @renovate[bot] in [#306](https://github.com/nicholas-fedor/todo/pull/306)
- Update module golang.org/x/text to v0.40.0 by @renovate[bot] in [#304](https://github.com/nicholas-fedor/todo/pull/304)
- Update module github.com/dgraph-io/badger/v4 to v4.9.4 by @renovate[bot] in [#303](https://github.com/nicholas-fedor/todo/pull/303)
- Update module golang.org/x/sys to v0.47.0 by @renovate[bot] in [#301](https://github.com/nicholas-fedor/todo/pull/301)
- Update github/codeql-action digest to 99df26d by @renovate[bot] in [#300](https://github.com/nicholas-fedor/todo/pull/300)
- Update module github.com/dgraph-io/ristretto/v2 to v2.4.2 by @renovate[bot] in [#298](https://github.com/nicholas-fedor/todo/pull/298)
- Update go module directive to v1.26.5 by @renovate[bot] in [#296](https://github.com/nicholas-fedor/todo/pull/296)
- Update module github.com/gofiber/schema to v1.8.1 by @renovate[bot] in [#295](https://github.com/nicholas-fedor/todo/pull/295)
- Update step-security/harden-runner action to v2.20.0 by @renovate[bot] in [#293](https://github.com/nicholas-fedor/todo/pull/293)
- Update module github.com/dgraph-io/ristretto/v2 to v2.4.1 by @renovate[bot] in [#291](https://github.com/nicholas-fedor/todo/pull/291)
- Update module github.com/dgraph-io/badger/v4 to v4.9.3 by @renovate[bot] in [#290](https://github.com/nicholas-fedor/todo/pull/290)
- Update module golang.org/x/text to v0.39.0 by @renovate[bot] in [#288](https://github.com/nicholas-fedor/todo/pull/288)
- Update nicholas-fedor/actionlint-action action to v1.0.18 by @renovate[bot] in [#285](https://github.com/nicholas-fedor/todo/pull/285)
- Update nicholas-fedor/actionlint-action action to v1.0.17 by @renovate[bot] in [#283](https://github.com/nicholas-fedor/todo/pull/283)
- Update docker/login-action digest to af1e73f by @renovate[bot] in [#281](https://github.com/nicholas-fedor/todo/pull/281)
- Update commitlint monorepo to v21.2.0 by @renovate[bot] in [#279](https://github.com/nicholas-fedor/todo/pull/279)
- Lock file maintenance by @renovate[bot] in [#261](https://github.com/nicholas-fedor/todo/pull/261)
- Update module github.com/gofiber/fiber/v3 to v3.4.0 by @renovate[bot] in [#276](https://github.com/nicholas-fedor/todo/pull/276)
- Update dependency tailwindcss to v4.3.2 by @renovate[bot] in [#274](https://github.com/nicholas-fedor/todo/pull/274)
- Update nicholas-fedor/actionlint-action action to v1.0.16 by @renovate[bot] in [#275](https://github.com/nicholas-fedor/todo/pull/275)
- Update github/codeql-action digest to 54f647b by @renovate[bot] in [#273](https://github.com/nicholas-fedor/todo/pull/273)
- Update docker/login-action digest to c99871d by @renovate[bot] in [#272](https://github.com/nicholas-fedor/todo/pull/272)
- Update module github.com/klauspost/compress to v1.19.0 by @renovate[bot] in [#270](https://github.com/nicholas-fedor/todo/pull/270)
- Update module github.com/klauspost/compress to v1.18.7 by @renovate[bot] in [#268](https://github.com/nicholas-fedor/todo/pull/268)
- Update module github.com/andybalholm/brotli to v1.2.2 by @renovate[bot] in [#266](https://github.com/nicholas-fedor/todo/pull/266)
- Update module github.com/valyala/fasthttp to v1.72.0 by @renovate[bot] in [#267](https://github.com/nicholas-fedor/todo/pull/267)
- Update goreleaser/goreleaser-action digest to f06c13b by @renovate[bot] in [#262](https://github.com/nicholas-fedor/todo/pull/262)
- Update module github.com/gofiber/utils/v2 to v2.1.1 by @renovate[bot] in [#263](https://github.com/nicholas-fedor/todo/pull/263)
- Update module github.com/templui/templui to v1.12.1 by @renovate[bot] in [#259](https://github.com/nicholas-fedor/todo/pull/259)
- Update nicholas-fedor/actionlint-action action to v1.0.15 by @renovate[bot] in [#257](https://github.com/nicholas-fedor/todo/pull/257)
- Update commitlint monorepo to v21.1.0 by @renovate[bot] in [#255](https://github.com/nicholas-fedor/todo/pull/255)
- Lock file maintenance by @renovate[bot] in [#245](https://github.com/nicholas-fedor/todo/pull/245)
- Update nicholas-fedor/actionlint-action action to v1.0.14 by @renovate[bot] in [#252](https://github.com/nicholas-fedor/todo/pull/252)
- Update actions/setup-go digest to 924ae3a by @renovate[bot] in [#250](https://github.com/nicholas-fedor/todo/pull/250)
- Update nicholas-fedor/actionlint-action action to v1.0.13 by @renovate[bot] in [#248](https://github.com/nicholas-fedor/todo/pull/248)
- Update nicholas-fedor/actionlint-action action to v1.0.12 by @renovate[bot] in [#246](https://github.com/nicholas-fedor/todo/pull/246)
- Update nicholas-fedor/actionlint-action action to v1.0.10 by @renovate[bot] in [#244](https://github.com/nicholas-fedor/todo/pull/244)
- Update actions/checkout action to v7 by @renovate[bot] in [#242](https://github.com/nicholas-fedor/todo/pull/242)
- Update alpine docker digest to 28bd5fe by @renovate[bot] in [#240](https://github.com/nicholas-fedor/todo/pull/240)
- Update alpine:3.24.1 docker digest to 28bd5fe by @renovate[bot] in [#241](https://github.com/nicholas-fedor/todo/pull/241)
- Update alpine docker tag to v3.24.1 by @renovate[bot] in [#238](https://github.com/nicholas-fedor/todo/pull/238)
- Update alpine docker digest to f5064d3 by @renovate[bot] in [#237](https://github.com/nicholas-fedor/todo/pull/237)
- Update tailwindcss monorepo to v4.3.1 by @renovate[bot] in [#235](https://github.com/nicholas-fedor/todo/pull/235)
- Update nicholas-fedor/actionlint-action action to v1.0.9 by @renovate[bot] in [#233](https://github.com/nicholas-fedor/todo/pull/233)
- Lock file maintenance by @renovate[bot] in [#231](https://github.com/nicholas-fedor/todo/pull/231)
- Update nicholas-fedor/actionlint-action action to v1.0.8 by @renovate[bot] in [#229](https://github.com/nicholas-fedor/todo/pull/229)
- Update nicholas-fedor/actionlint-action action to v1.0.7 by @renovate[bot] in [#227](https://github.com/nicholas-fedor/todo/pull/227)
- Update module golang.org/x/net to v0.56.0 by @renovate[bot] in [#225](https://github.com/nicholas-fedor/todo/pull/225)
- Update alpine:3.24.0 docker digest to a2d49ea by @renovate[bot] in [#223](https://github.com/nicholas-fedor/todo/pull/223)
- Update alpine docker digest to a2d49ea by @renovate[bot] in [#222](https://github.com/nicholas-fedor/todo/pull/222)
- Update alpine:3.24.0 docker digest to 8ddefa9 by @renovate[bot] in [#220](https://github.com/nicholas-fedor/todo/pull/220)
- Update alpine docker digest to fa1b3b8 by @renovate[bot] in [#219](https://github.com/nicholas-fedor/todo/pull/219)
- Update alpine docker digest to 4f4ba24 by @renovate[bot] in [#216](https://github.com/nicholas-fedor/todo/pull/216)
- Update alpine docker tag to v3.24.0 by @renovate[bot] in [#217](https://github.com/nicholas-fedor/todo/pull/217)
- Update module github.com/gofiber/schema to v1.8.0 by @renovate[bot] in [#214](https://github.com/nicholas-fedor/todo/pull/214)
- Update module github.com/dgraph-io/badger/v4 to v4.9.2 by @renovate[bot] in [#213](https://github.com/nicholas-fedor/todo/pull/213)
- Update nicholas-fedor/actionlint-action action to v1.0.6 by @renovate[bot] in [#211](https://github.com/nicholas-fedor/todo/pull/211)
- Lock file maintenance by @renovate[bot] in [#202](https://github.com/nicholas-fedor/todo/pull/202)
- Update module golang.org/x/crypto to v0.53.0 by @renovate[bot] in [#209](https://github.com/nicholas-fedor/todo/pull/209)
- Update module golang.org/x/text to v0.38.0 by @renovate[bot] in [#207](https://github.com/nicholas-fedor/todo/pull/207)
- Update module golang.org/x/sys to v0.46.0 by @renovate[bot] in [#205](https://github.com/nicholas-fedor/todo/pull/205)
- Update nicholas-fedor/actionlint-action action to v1.0.5 by @renovate[bot] in [#203](https://github.com/nicholas-fedor/todo/pull/203)
- Update nicholas-fedor/actionlint-action action to v1.0.4 by @renovate[bot] in [#200](https://github.com/nicholas-fedor/todo/pull/200)
- Update github/codeql-action digest to 8aad20d by @renovate[bot] in [#198](https://github.com/nicholas-fedor/todo/pull/198)
- Update actions/checkout digest to df4cb1c by @renovate[bot] in [#196](https://github.com/nicholas-fedor/todo/pull/196)
- Update go module directive to v1.26.4 by @renovate[bot] in [#195](https://github.com/nicholas-fedor/todo/pull/195)
- Update module github.com/gofiber/schema to v1.7.2 by @renovate[bot] in [#193](https://github.com/nicholas-fedor/todo/pull/193)
- Update actions/checkout action to v6.0.3 by @renovate[bot] in [#191](https://github.com/nicholas-fedor/todo/pull/191)
- Update github/codeql-action digest to 87557b9 by @renovate[bot] in [#190](https://github.com/nicholas-fedor/todo/pull/190)
- Update securego/gosec action to v2.27.1 by @renovate[bot] in [#188](https://github.com/nicholas-fedor/todo/pull/188)
- Update module github.com/templui/templui to v1.12.0 by @renovate[bot] in [#187](https://github.com/nicholas-fedor/todo/pull/187)
- Update module github.com/gofiber/utils/v2 to v2.1.0 by @renovate[bot] in [#185](https://github.com/nicholas-fedor/todo/pull/185)
- Update commitlint monorepo to v21.0.2 by @renovate[bot] in [#184](https://github.com/nicholas-fedor/todo/pull/184)
- Update module github.com/mattn/go-colorable to v0.1.15 by @renovate[bot] in [#182](https://github.com/nicholas-fedor/todo/pull/182)
- Update opentelemetry-go monorepo to v1.44.0 by @renovate[bot] in [#180](https://github.com/nicholas-fedor/todo/pull/180)
- Lock file maintenance by @renovate[bot] in [#178](https://github.com/nicholas-fedor/todo/pull/178)
- Update module github.com/templui/templui to v1.11.2 by @renovate[bot] in [#177](https://github.com/nicholas-fedor/todo/pull/177)
- Update github/codeql-action digest to 7211b7c by @renovate[bot] in [#176](https://github.com/nicholas-fedor/todo/pull/176)
- Update docker/login-action digest to 650006c by @renovate[bot] in [#175](https://github.com/nicholas-fedor/todo/pull/175)
- Update module github.com/gofiber/fiber/v3 to v3.3.0 by @renovate[bot] in [#173](https://github.com/nicholas-fedor/todo/pull/173)
- Update module golang.org/x/net to v0.55.0 by @renovate[bot] in [#171](https://github.com/nicholas-fedor/todo/pull/171)
- Update module golang.org/x/crypto to v0.52.0 by @renovate[bot] in [#169](https://github.com/nicholas-fedor/todo/pull/169)
- Update module golang.org/x/sys to v0.45.0 by @renovate[bot] in [#168](https://github.com/nicholas-fedor/todo/pull/168)
- Update step-security/harden-runner action to v2.19.4 by @renovate[bot] in [#166](https://github.com/nicholas-fedor/todo/pull/166)
- Update module github.com/gofiber/utils/v2 to v2.0.6 by @renovate[bot] in [#165](https://github.com/nicholas-fedor/todo/pull/165)
- Update goreleaser/goreleaser-action digest to 5daf1e9 by @renovate[bot] in [#163](https://github.com/nicholas-fedor/todo/pull/163)
- Lock file maintenance by @renovate[bot] in [#161](https://github.com/nicholas-fedor/todo/pull/161)
- Update commitlint monorepo to v21.0.1 by @renovate[bot] in [#159](https://github.com/nicholas-fedor/todo/pull/159)
- Update github/codeql-action digest to 9e0d7b8 by @renovate[bot] in [#158](https://github.com/nicholas-fedor/todo/pull/158)
- Update step-security/harden-runner action to v2.19.3 by @renovate[bot] in [#156](https://github.com/nicholas-fedor/todo/pull/156)
- Update step-security/harden-runner action to v2.19.2 by @renovate[bot] in [#154](https://github.com/nicholas-fedor/todo/pull/154)
- Update module github.com/templui/templui to v1.11.1 by @renovate[bot] in [#152](https://github.com/nicholas-fedor/todo/pull/152)
- Update dependency tailwindcss to v4.3.0 by @renovate[bot] in [#150](https://github.com/nicholas-fedor/todo/pull/150)
- Update commitlint monorepo to v21 by @renovate[bot] in [#146](https://github.com/nicholas-fedor/todo/pull/146)
- Update module github.com/templui/templui to v1.11.0 by @renovate[bot] in [#147](https://github.com/nicholas-fedor/todo/pull/147)
- Update module github.com/gofiber/utils/v2 to v2.0.5 by @renovate[bot] in [#145](https://github.com/nicholas-fedor/todo/pull/145)
- Update module github.com/templui/templui to v1.10.1 by @renovate[bot] in [#143](https://github.com/nicholas-fedor/todo/pull/143)
- Lock file maintenance by @renovate[bot] in [#141](https://github.com/nicholas-fedor/todo/pull/141)
- Update module github.com/a-h/templ to v0.3.1020 by @renovate[bot] in [#139](https://github.com/nicholas-fedor/todo/pull/139)
- Update nicholas-fedor/actionlint-action action to v1.0.3 by @renovate[bot] in [#138](https://github.com/nicholas-fedor/todo/pull/138)
- Update nicholas-fedor/actionlint-action action to v1.0.2 by @renovate[bot] in [#137](https://github.com/nicholas-fedor/todo/pull/137)
- Update nicholas-fedor/actionlint-action digest to 25b48e2 by @renovate[bot] in [#134](https://github.com/nicholas-fedor/todo/pull/134)
- Update module golang.org/x/net to v0.54.0 by @renovate[bot] in [#130](https://github.com/nicholas-fedor/todo/pull/130)
- Update module golang.org/x/crypto to v0.51.0 by @renovate[bot] in [#129](https://github.com/nicholas-fedor/todo/pull/129)
- Update module golang.org/x/text to v0.37.0 by @renovate[bot] in [#127](https://github.com/nicholas-fedor/todo/pull/127)
- Update module golang.org/x/sys to v0.44.0 by @renovate[bot] in [#125](https://github.com/nicholas-fedor/todo/pull/125)
- Update go module directive to v1.26.3 by @renovate[bot] in [#123](https://github.com/nicholas-fedor/todo/pull/123)
- Update github/codeql-action digest to 68bde55 by @renovate[bot] in [#122](https://github.com/nicholas-fedor/todo/pull/122)
- Update sigstore/cosign-installer action to v4.1.2 by @renovate[bot] in [#120](https://github.com/nicholas-fedor/todo/pull/120)
- Update module github.com/valyala/fasthttp to v1.71.0 by @renovate[bot] in [#118](https://github.com/nicholas-fedor/todo/pull/118)
- Lock file maintenance by @renovate[bot] in [#116](https://github.com/nicholas-fedor/todo/pull/116)
- Update commitlint monorepo to v20.5.3 by @renovate[bot] in [#114](https://github.com/nicholas-fedor/todo/pull/114)
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
