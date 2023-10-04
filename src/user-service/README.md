# User Service

## How to use

#### Generating keys
The ECDSA private key for signing [JWT](https://jwt.io) tokens:

    ssh-keygen -t ecdsa -f /path/to/key -m pem 

#### Create application configuration

```yaml
database:
    host: localhost
    port: 5432
    username: postgres
    password: password
    dbname: postgres
jwt:
    signKey: /path/to/key
```

#### Run

    go run main.go -config=/path/to/config
