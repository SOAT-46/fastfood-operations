# fastfood-operations

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
