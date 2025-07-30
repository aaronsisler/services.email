# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Legend

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [2.3.0] Adding in the email service for request body parsing

### Added

- Default email service added through dependency injection
- Updating test cases to handle this DI process

## [2.2.0] Creating and deploying the base email POST endpoint

### Added

- Adding in the bad request validation for email handler
- Adding in testing using Godog and feature files
- Adding GHA and Makefile steps for running test

## [2.1.0] Creating and deploying the base email POST endpoint

### Added

- Base email handler
- GitHub Action, Makefile, and CF template changes to deploy to API Gateway

## [1.0.0] Creating and deploying the base application

### Added

- Base health handler
- GitHub Action, Makefile, and CF template changes to deploy to API Gateway
