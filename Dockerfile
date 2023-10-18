# Use an official PostgreSQL image from Docker Hub
FROM postgres:latest

# Set environment variables from .env file
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}
