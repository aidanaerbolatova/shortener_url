# URL Shortening Service

This is a URL shortening service developed in Go. It allows users to create, manage, and track short URLs. The service provides endpoints to shorten URLs, retrieve statistics, and manage links (view, delete).

## Features
- **Shorten URL**: Convert a long URL into a short one.
- **List Links**: View a list of all short URLs created by the user.
- **Redirect**: Access the original long URL using the shortened URL.
- **Delete**: Delete a specific short URL.
- **Statistics**: View statistics for each short URL (e.g., number of clicks, last accessed time).
- **Expiration**: Shortened links have a time-to-live (TTL) of 30 days.

## Endpoints

### 1. `POST: /shortener`

**Request Body**:
```json
{
  "url": "http://example.com"
}
```

**Response**:
```json
{
  "shortened_url": "http://localhost:8080/shortenedID"
}
```

- Accepts a long URL and returns a shortened URL with a unique identifier.
- The system ensures that the same long URL is not shortened multiple times.
- Shortened links expire after 30 days.

### 2. `GET: /shortener`

**Response**:
```json
{
  "links": [
    {
      "shortened_url": "http://localhost:8080/shortenedID",
      "original_url": "http://example.com",
      "created_at": "2024-11-15T12:34:56Z",
      "expires_at": "2024-12-15T12:34:56Z"
    }
  ]
}
```

- Retrieves a list of all short URLs created by the user.

### 3. `GET: /{link}`

**Response**:
- Redirects the user to the original long URL.

- Example: `GET http://localhost:8080/{shortenedID}` will redirect to `http://example.com`.

### 4. `DELETE: /{link}`

**Response**:
```json
{
  "message": "Link successfully deleted."
}
```

- Deletes the shortened URL and removes it from the system.

### 5. `GET: /stats/{link}`

**Response**:
```json
{
  "link": "http://localhost:8080/{shortenedID}",
  "clicks": 42,
  "last_accessed": "2024-11-15T14:12:34Z"
}
```

- Provides statistics for a shortened link, including the number of clicks and the last time the link was accessed.

## Validation and Error Handling

- The service validates the URL format before shortening. Invalid URLs return a `400 Bad Request` error.
- If a non-existent shortened URL is accessed, a `404 Not Found` error is returned.
- If there are database errors or other unexpected issues, the service returns appropriate HTTP error codes and messages.
