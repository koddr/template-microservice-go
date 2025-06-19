# Microservice Go template

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![License][repo_license_img]][repo_license_url]

A lightweight Go microservice.

Example configuration:

```bash
# API server settings.
API_SERVER_PORT=8080

# API server auth settings.
API_SERVER_AUTH_USERNAME="user"
API_SERVER_AUTH_PASSWORD="password"

# API server DB settings.
API_SERVER_DATABASE_URL="postgres://postgres:password@localhost:5432/postgres"
```

## Endpoints

### `GET /api/v1/transactions`

Returns all transactions from the database.

### `GET /api/v1/transactions/filter`

Returns filtered by `created_at` field transactions from the database.

Required query params:

- `created_at_start` - string in format "2006-01-02"
- `created_at_end` - string in format "2006-01-02"

### `POST /api/v1/transaction`

Add a new transaction to the database.

## ⚠️ License

[`template-microservice-go`][repo_url] is free and open-source software licensed
under the [Apache 2.0 License][repo_license_url], created and supported by
[Vic Shóstak][author_url] with 🩵 for people and robots.

<!-- Go links -->

[go_report_url]: https://goreportcard.com/report/github.com/koddr/template-microservice-go
[go_dev_url]: https://pkg.go.dev/github.com/koddr/template-microservice-go
[go_version_img]: https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/koddr/template-microservice-go
[repo_license_url]: https://github.com/koddr/template-microservice-go/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none

<!-- Author links -->

[author_url]: https://github.com/koddr
