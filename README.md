# URL Shortening Service

This is a URL shortening service developed in Go. It allows users to create, manage, and track short URLs. The service provides endpoints to shorten URLs, retrieve statistics, and manage links (view, delete).

## Features
- **Shorten URL**: Convert a long URL into a short one.
- **List Links**: View a list of all short URLs created by the user.
- **Redirect**: Access the original long URL using the shortened URL.
- **Delete**: Delete a specific short URL.
- **Statistics**: View statistics for each short URL (e.g., number of clicks, last accessed time).
- **Expiration**: Shortened links have a time-to-live (TTL) of 30 days.
