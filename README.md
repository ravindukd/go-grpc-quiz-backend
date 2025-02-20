# grpc-test

## GitHub Workflow Action

This repository includes a GitHub workflow action to build and release the application when a pull request is made to the release branch.

### Workflow File

The workflow file is located at `.github/workflows/build-and-release.yml`.

### Workflow Steps

1. **Checkout code**: Checks out the repository code.
2. **Set up Go**: Sets up the Go environment.
3. **Build application**: Builds the Go application.
4. **Create release**: Creates a release in GitHub.

### Trigger

The workflow is triggered when a pull request is made to the `release` branch.
