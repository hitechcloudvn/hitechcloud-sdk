> **English** | [Tiếng Việt](CONTRIBUTING.vi.md)

# Contributing to HiTechCloud SDK

Thank you for your interest in contributing!

## Development Setup

1. Clone the repository
2. Choose your language SDK under `sdk/`
3. Install dependencies
4. Run tests

## Adding a New Language

1. Create a new directory under `sdk/` with the language name
2. Implement the client and all resource classes
3. Add tests
4. Update the main README.md (and README.vi.md)
5. Add CI/CD workflow in `.github/workflows/` and `.gitlab-ci.yml`

## Code Style

- Follow each language's conventions
- Include docstrings/JSDoc for all public methods
- Match the existing resource structure
- Documentation is bilingual (English and Vietnamese): keep the `README.md` / `README.vi.md` pair in sync

## Pull Requests

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass (CI runs on GitHub Actions and GitLab CI)
5. Submit a pull request

## Commit Messages

- Use concise, imperative subjects (e.g. `fix: ...`, `feat: ...`, `docs: ...`)
- End commit messages with co-author trailers where applicable:

```
Co-Authored-By: pgt-vn <nguyenthanhan92f1@gmail.com>
Co-Authored-By: agent-hitechcloud <agent@hitechcloud.vn>
```

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
