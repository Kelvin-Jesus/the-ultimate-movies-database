FROM postgres:11-alpine

# Set the environment variables for the user and password
ENV POSTGRES_USER="user1234"
ENV POSTGRES_PASSWORD="1234user"

# Expose the Postgres port
EXPOSE 5432

# Add a label with the container name
LABEL name="movies-db"