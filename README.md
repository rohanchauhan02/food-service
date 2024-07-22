- Install `goose` for db migrations :

```shell
    brew install goose
```

- Migration Status

```shell
cd database/migrations
goose mysql "root:password@tcp(localhost:3307)/loan_service_development?parseTime=true" status
```

- Migration up

```shell
cd database/migrations
goose mysql "root:password@tcp(localhost:3307)/loan_service_development?parseTime=true" up
```

- Migration down

```shell
cd database/migrations
goose mysql "root:password@tcp(localhost:3307)/loan_service_development?parseTime=true" down
```

- Seeders Status

```shell
cd database/seeders
goose mysql "root:password@tcp(localhost:3307)/loan_service_development?parseTime=true" status
```

- Seeders up

```shell
cd database/seeders
goose mysql "root:password@tcp(localhost:3307)/loan_service_development?parseTime=true" up
```

- Seeders down

```shell
cd database/seeders
goose mysql "root:password@tcp(localhost:3307)/loan_service_development?parseTime=true" down
```
