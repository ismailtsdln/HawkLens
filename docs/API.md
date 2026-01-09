# API Documentation

HawkLens provides a modern REST + Streaming API for interacting with the platform.

## 🌐 Base URL
`http://localhost:8080/api/v1`

## 📡 Endpoints

### 1. Health Check
`GET /health`
Verifies if the API server is up and running.

**Response:**
```json
{
  "status": "up"
}
```

### 2. List Plugins
`GET /plugins`
Returns a list of all currently registered and active OSINT plugins.

**Response:**
```json
{
  "plugins": ["twitter", "youtube", "reddit", "instagram", "tiktok"]
}
```

### 3. Stream Scan (SSE)
`GET /scan-stream?query=[query]`
Starts a concurrent scan across all platforms and streams results back to the client using Server-Sent Events.

**Query Parameters:**
- `query` (required): The term to search for across social platforms.

**Event Stream Type:** `text/event-stream`

**Message Data Structure:**
```json
{
  "platform": "twitter",
  "data_type": "tweet",
  "data": {
    "text": "OSINT is powerful!",
    "id": "12345"
  }
}
```

### 4. List Results
`GET /results?platform=[platform]`
Retrieves previously saved results from the PostgreSQL persistence layer.

**Query Parameters:**
- `platform` (optional): Filter results by a specific platform.

---

## 🛠️ Usage Example (Javascript)

```javascript
const eventSource = new EventSource('/api/v1/scan-stream?query=cybersecurity');

eventSource.onmessage = (event) => {
    const result = JSON.parse(event.data);
    console.log(`New result from ${result.platform}:`, result.data);
};

eventSource.onerror = (err) => {
    console.error("Stream failed:", err);
    eventSource.close();
};
```
