# fastfood-operations

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=SOAT-46_fastfood-operations&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=SOAT-46_fastfood-operations)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=SOAT-46_fastfood-operations&metric=coverage)](https://sonarcloud.io/summary/new_code?id=SOAT-46_fastfood-operations)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=SOAT-46_fastfood-operations&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=SOAT-46_fastfood-operations)

## Documentation

For a complete understanding of the functionalities and how to interact with the API, use Postman.

### Environment Variables

This project requires certain environment variables to be set.
You can find a template for these variables in the .env.example file.
To create your own .env file, run the following command:

```shell
cp .env.example .env
```

Then, edit the .env file to include the appropriate values for your setup.

### Building and running your application

When you're ready, start your application by running:

```shell
make docker-up
```

Your application will be accessible at http://localhost:8080, and the Swagger
documentation can be found at http://localhost:8080/swagger/index.html.


### Generate database migrations

To generate database migrations, run the following command:

```shell
goose -dir ./fastfood-operations/db/migrations create create_XXX_table sql
```
