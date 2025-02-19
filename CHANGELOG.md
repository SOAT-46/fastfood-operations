# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The Unreleased section on `CHANGELOG.md` gets a version number and date;
3. Open a Pull Request with the bump version changes targeting the `main` branch;

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

### Added

- added docker build and push in pipeline
- added `PUT /v1/orders/:id` endpoint to update an order
- added `GET /v1/orders/:id` endpoint to get an order by the target ID
- added `GET /v1/orders` endpoint to get all orders
- added `POST /v1/orders` endpoint to create a new order
- added Dockerfile to containerize the application
- added project structure

### Changed

- changed database from Postgres to MongoDB
