# Beholder

Cron job to collect and save Steam sale prices based on a given user wishlist to a database

## Building and Running

```
podman build -t beholder ./
```

```
podman run --env-files .env beholder
```