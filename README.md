# Space Trouble

## Development

Build the binary

```sh
make
```

Start postgres

```sh
docker-compose up -Vd
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1:5432/space_trouble\?sslmode=disable up
```

Run the web server

```sh
./bin/server
```

## Usage

For example, you get information about bookings

```sh
http http://localhost:8080/api/v1/bookings
```

```json

```

## License

Project released under the terms of the MIT [license](./LICENSE).
