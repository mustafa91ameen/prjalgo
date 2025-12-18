# Authentication API

## Login

```
POST /api/v1/auth/login
```

**Auth Required:** No

### Request Body

```json
{
  "username": "string (required)",
  "password": "string (required)"
}
```

### Response

```json
{
  "data": {
    "accessToken": "jwt_token_here",
    "refreshToken": "refresh_token_here",
    "userId": 1,
    "username": "john_doe"
  }
}
```

### Errors

| Status | Message |
|--------|---------|
| 401 | invalid username or password |
| 401 | user account is inactive |

---

## Refresh Token

```
POST /api/v1/auth/refresh
```

**Auth Required:** No

### Request Body

```json
{
  "refreshToken": "string (required)"
}
```

### Response

```json
{
  "data": {
    "accessToken": "new_jwt_token",
    "refreshToken": "new_refresh_token",
    "userId": 1,
    "username": "john_doe"
  }
}
```

### Errors

| Status | Message |
|--------|---------|
| 401 | invalid or expired refresh token |

---

## Logout

```
POST /api/v1/auth/logout
```

**Auth Required:** No

### Request Body

```json
{
  "refreshToken": "string (required)"
}
```

### Response

```json
{
  "data": null
}
```

---

## Get My Pages

Get accessible pages for the authenticated user.

```
GET /api/v1/auth/pages
```

**Auth Required:** Yes

### Response

```json
{
  "data": [
    {
      "id": 1,
      "name": "Dashboard",
      "route": "/dashboard",
      "icon": "dashboard",
      "permissions": ["read", "write"]
    }
  ]
}
```
