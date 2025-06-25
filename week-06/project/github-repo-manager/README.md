# GitHub Repo Manager - CLI Tool

This is a CLI tool to manage GitHub repositories.

## Features

- List user repositories
- Create new repositories
- Check repository stats
- Clone repositories

## Installation

```
go install github.com/your-username/github-repo-manager@latest
```

## Usage

```
# Initialize tool with your GitHub token
gh-manager auth --token YOUR_GITHUB_TOKEN

# List your repositories
gh-manager list

# Get info about a specific repository
gh-manager info username/repo-name

# Create a new repository
gh-manager create repo-name --description "Your description" --private

# Clone a repository
gh-manager clone username/repo-name
```

## Project Structure

- `cmd/` - Command implementations
- `api/` - GitHub API client
- `config/` - Configuration management
- `utils/` - Helper functions

## Dependencies

- github.com/spf13/cobra
- github.com/google/go-github/v53
- golang.org/x/oauth2
