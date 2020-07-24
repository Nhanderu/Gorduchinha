# Gorduchinha

The Brazilian football teams API and scraper!

## Structure

### `app`

This package contains the common code for all the applications

### `cmd/api`

Code for the API server. It depends on `app`.

To run it, execute `go run ./cmd/api`.

### `cmd/job-scraper`

Code for the scraper bot. It depends on `app`.

To run it, execute `go run ./cmd/job-scraper`.

## License

This project code is in the public domain. See the [LICENSE file][1].

[1]: https://github.com/Nhanderu/gorduchinha/blob/master/LICENSE
