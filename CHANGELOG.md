> **English** | [Tiếng Việt](CHANGELOG.vi.md)

# Changelog

All notable changes to the HiTechCloud SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.0] - 2026-11-16

### Added
- Bilingual Vietnamese/English standardization across the entire repository (README, docs, guides, SDK docs, examples, agents, skills)
- GitLab CI/CD pipeline (`.gitlab-ci.yml`) with tests and tag-based releases
- `workflow_dispatch` on the Release workflow to publish releases for existing tags
- Missing C# SDK README (`sdk/csharp/README.md` + `README.vi.md`)
- Bilingual code comments in examples

### Fixed
- Pinned `maven-resources-plugin` to `3.3.1` in `sdk/java/pom.xml` (the default lifecycle version `3.5.0` is not published on Maven Central) — unblocks the Java CI job
- Corrected development paths in README (`cd sdk/python`, `cd sdk/nodejs`)

## [1.2.0] - 2026-11-15

### Changed
- Moved repositories to the `hitechcloudvn` organization on GitHub and GitLab
- Normalized Go module path to `github.com/hitechcloudvn/hitechcloud-sdk/sdk/go`
- Pointed Julia, Objective-C, and Swift SDK references to the monorepo
- Standardized commit authorship to Nguyễn Thanh An (thanhan92f1)

## [1.1.0] - 2026-11-15

### Changed
- Synchronized latest User API sources
- Added build artifact tracking for release pipeline

## [1.0.0] - 2024-12-01

### Added
- Initial release with 7 language SDKs (Python, Node.js, PHP, Go, Ruby, Java, C#)
- 39 resource modules covering all User API endpoints
- OpenAPI 3.0 specification
- Comprehensive documentation
- CI/CD pipelines for all languages
- Examples for all supported languages
- Agent configurations and MCP server

[1.3.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/hitechcloudvn/hitechcloud-sdk/releases/tag/v1.0.0
