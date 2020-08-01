# Gorduchinha

The Brazilian football teams API and scraper running @
[gorduchinha.herokuapp.com/api][1]!

## API

### `POST /v1/graphql`

Route for GraphQL queries. Check out the [schema][2].

### `GET /v1/teams`

Fetches all teams in the database.

### `GET /v1/teams/:abbr`

Fetches team by its abbreviation.

Current possible values:
- `sccp`
- `sep`
- `spfc`
- `sfc`
- `crf`
- `crvg`
- `ffc`
- `bfr`
- `cam`
- `cec`
- `gfbpa`
- `sci`

### `GET /v1/champs`

Fetches all championships in the database.

### `GET /v1/champs/:slug`

Fetches championship by its slug.

Current possible values:
- `national-league-1-div`
- `national-league-2-div`
- `national-cup`
- `world-cup`
- `intercontinental-cup`
- `south-american-cup-a`
- `south-american-cup-b`
- `south-american-supercup`
- `sp-state-cup`
- `rj-state-cup`
- `rs-state-cup`
- `mg-state-cup`

### `PUT /v1/trophies`

Executes scraper job that updates all trophies. It requires a secret key.

## Code structure

### `app/`

This package contains the common code for all the applications

### `cmd/api/`

Code for the API server. It depends on `app`.

To run it, execute `go run ./cmd/api`.

### `cmd/job-scraper/`

Code for the scraper bot. It depends on `app`.

To run it, execute `go run ./cmd/job-scraper`.

## License

This project code is in the public domain. See the [LICENSE file][3].

[1]: http://gorduchinha.herokuapp.com/api/
[2]: ./static/graphql/schema.gql
[3]: https://github.com/Nhanderu/gorduchinha/blob/master/LICENSE
