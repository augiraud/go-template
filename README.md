# go-template

A ready-to-use Go project template. Run the setup script and start coding, with CI, linting, releases and quality checks already wired up.

## Getting started

1. Create a repository from this template (`Use this template` on GitHub), then clone it.
2. Install git, [Go](https://go.dev/doc/install) and [Docker](https://docs.docker.com/get-docker/), preferably with your package manager. Go downloads the right version for you if yours is older than the one in `go.mod`.
3. Run `./scripts/setup.sh`. The first time, it asks for the project name and your GitHub user or organization, then renames the project.
4. Write your code in `internal/` and your entry point in `cmd/go-template/`. Tests go next to the code, in `*_test.go` files.

## Commands

- `./scripts/run.sh [args]`: run the app
- `./scripts/run_docker.sh [args]`: run the app in Docker, building the image first
- `./scripts/docker_build.sh`: build the Docker image
- `./scripts/local_ci.sh`: run all the checks exactly like CI does (modules, format, lint, tests with the race detector, vulnerabilities, release config, Docker image). With `--fix`, it first formats the code and fixes what can be fixed automatically.
- `go get <module>`: add a dependency

A git hook checks formatting and lint before each commit.

## Tools

golangci-lint, gotestsum, govulncheck and goreleaser are pinned in `tools/go.mod`, apart from the app's dependencies. The scripts run them with `go tool -modfile=tools/go.mod <tool>`, so there is nothing else to install. Dependabot keeps them up to date.

## Lint rules

`.golangci.yml` is based on the [golden config](https://github.com/maratori/golangci-lint-config) by Marat Reimers, with a few rules relaxed (listed at the top of the file). For example, tests don't have to call `t.Parallel()`.

## CI

GitHub Actions runs `./scripts/local_ci.sh` on every push to `main`, every `v*` tag and every pull request.

### Releases

Pushing a tag such as `v1.2.0` publishes a GitHub release with binaries for Linux, macOS and Windows (amd64 and arm64), once the checks pass. The binaries are built by [goreleaser](https://goreleaser.com), configured in `.goreleaser.yml`.

```sh
git tag v1.2.0
git push origin v1.2.0
```

`go-template -version` prints the version: the tag in released binaries and Docker images, and the module version otherwise.

### Sonar

CI skips Sonar until you set it up:

1. On [SonarQube Cloud](https://sonarcloud.io/projects/create), import the repository.
2. In the SonarQube Cloud project, go to Administration > Analysis Method and turn off Automatic Analysis.
3. Create a token in My Account > Security.
4. In the GitHub repository, add the token as a secret named `SONAR_TOKEN`: Settings > Secrets and variables > Actions > New repository secret.

The setup script uses SonarQube Cloud's default names: your GitHub user or organization in lowercase for the organization, and `<owner>_<project>` for the project key. If yours differ, edit `sonar-project.properties`.

More details in [SonarQube Cloud's documentation](https://docs.sonarsource.com/sonarqube-cloud/analyzing-source-code/ci-based-analysis/github-actions-for-sonarcloud).

### Publishing the Docker image

Disabled by default. To enable it, uncomment the last block of the `ci` job in `.github/workflows/ci.yml`.

Then every push to `main` publishes `ghcr.io/<owner>/<repository>:latest`, and every tag such as `v1.2.0` publishes `ghcr.io/<owner>/<repository>:v1.2.0`. Both also publish an image tagged with the commit hash.

## Configuration

- `go.mod`: module path, Go version and dependencies
- `tools/go.mod`: pinned development tools
- `.golangci.yml`: format and lint rules
- `.goreleaser.yml`: release binaries
- `.github/workflows/ci.yml`: CI and releases
- `.github/dependabot.yml`: weekly dependency updates
- `Dockerfile`: Docker image
- `.githooks/pre-commit`: checks before each commit

## License

MIT, see [LICENSE](LICENSE).
