# Frontend-Backend Connection Documentation

This document describes how the frontend connects to the backend, including the authentication system, permission system, API endpoints, and navigation structure.

---

## Table of Contents

1. [API Configuration](#1-api-configuration)
2. [Authentication System](#2-authentication-system)
3. [Permission System](#3-permission-system)
4. [Sidebar Navigation](#4-sidebar-navigation)
5. [State Management](#5-state-management)
6. [API Endpoints](#6-api-endpoints)
7. [Request Flow Examples](#7-request-flow-examples)
8. [Key Files Reference](#8-key-files-reference)
9. [All Frontend Pages](#9-all-frontend-pages)
10. [Route to Permission Mapping](#10-route-to-permission-mapping-complete)
11. [API Files Summary](#11-api-files-summary)

---

## 1. API Configuration

### Base URL

The API base URL is configured in `/src/api/client.js`:

```javascript
const API_BASE_URL = import.meta.env.VITE_API_URL || '/api'
```

- **Default**: `/api` (for Docker nginx proxy)
- **Custom**: Set `VITE_API_URL` environment variable for local development

### Default Headers

All requests include:

```javascript
{
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}
```

### Authenticated Requests

Protected endpoints include the Authorization header:

```
Authorization: Bearer {accessToken}
```

---

## 2. Authentication System

### Token Storage (localStorage)

| Key | Description |
|-----|-------------|
| `accessToken` | JWT access token for API authentication |
| `refreshToken` | Token used to refresh expired access tokens |
| `isAuthenticated` | Flag set to `'true'` when logged in |
| `user` | JSON object with `{ id, username }` |
| `loginTime` | Timestamp of last login |
| `userPages` | JSON array of user's authorized pages |

### Authentication Endpoints

```
POST   /auth/login      - User login (no auth required)
POST   /auth/refresh    - Refresh access token (no auth required)
POST   /auth/logout     - User logout (no auth required)
GET    /auth/pages      - Get user's authorized pages (requires auth)
```

### Login Request

```javascript
POST /api/auth/login
Body: { username: string, password: string }
```

### Login Response

```json
{
  "success": true,
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs...",
    "userId": 1,
    "username": "admin"
  }
}
```

### Token Refresh Flow

1. API call returns 401 Unauthorized
2. System checks for valid refresh token
3. Calls `POST /auth/refresh` with `{ refreshToken }`
4. Receives new `accessToken` and `refreshToken`
5. Retries the failed request with new token
6. If refresh fails, clears session and redirects to `/login`

### Key Auth Files

| File | Purpose |
|------|---------|
| `/src/api/auth.js` | Auth API endpoint functions |
| `/src/services/authService.js` | Auth business logic (login, logout, session) |
| `/src/api/client.js` | HTTP client with token management |
| `/src/pages/login.vue` | Login page component |

---

## 3. Permission System

### Permission Types

```javascript
const PERMISSIONS = {
  READ: 'read',
  CREATE: 'create',
  UPDATE: 'update',
  DELETE: 'delete',
  UPDATE_PASSWORD: 'updatePassword',
  UPDATE_STATUS: 'updateStatus'
}
```

### User Pages Structure

Each authorized page has:

```javascript
{
  id: number,
  name: string,           // Display name (Arabic)
  route: string,          // API route (e.g., '/projects')
  icon: string,           // Icon name
  permissions: string[]   // ['read', 'create', 'update', 'delete']
}
```

### Route to Permission Mapping

Frontend routes map to backend permission routes:

| Frontend Route | Permission Route |
|----------------|------------------|
| `/project-management` | `/projects` |
| `/task-management` | `/projects` |
| `/engineers` | `/users` |
| `/expense-types` | `/expenseTypes` |
| `/human-resources` | `/humanResources` |
| `/team-management` | `/teamMembers` |
| `/working-day-details` | `/workdays` |
| `/work-days` | `/workdays` |
| `/labor-details` | `/workdayLabor` |
| `/equipment-details` | `/workdayEquipment` |
| `/materials-expenses-details` | `/workdayMaterials` |

### Using Permissions in Components

```javascript
import { usePermissions } from '@/composables/usePermissions'

const { canRead, canCreate, canUpdate, canDelete, hasPermission } = usePermissions()

// Check specific permissions
if (canCreate.value) {
  // Show create button
}

if (hasPermission('updatePassword')) {
  // Show change password option
}
```

### Route Protection

Routes are protected via Vue Router navigation guard:

```javascript
router.beforeEach(async (to, from, next) => {
  // Public pages bypass auth check
  if (to.path === '/login' || to.path === '/unauthorized') {
    return next()
  }

  // Check authentication
  if (!isAuthenticated) {
    return next('/login')
  }

  // Check page permission
  if (!authStore.canAccessPage(to.path)) {
    return next('/unauthorized')
  }

  next()
})
```

### Key Permission Files

| File | Purpose |
|------|---------|
| `/src/stores/auth.js` | Auth store with permission methods |
| `/src/composables/usePermissions.js` | Permission composable for components |
| `/src/router/index.js` | Route guards |

---

## 4. Sidebar Navigation

### Structure

The sidebar is built in `/src/App.vue` and displays:

1. Logo and title section
2. Dynamic menu from `authStore.pages`
3. Active state highlighting
4. Logout button
5. User info card at bottom

### Menu Item Filtering

API-only routes are hidden from the sidebar:

```javascript
const apiOnlyRoutes = [
  '/workdays',
  '/workdayLabor',
  '/workdayEquipment',
  '/workdayMaterials',
  '/rolePages',
  '/userRoles',
  '/workSubcategories'
]
```

### Sidebar Pages (from API)

Pages are fetched via `GET /auth/pages` after login. Example response:

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "لوحة التحكم",
      "route": "/dashboard",
      "icon": "mdi-view-dashboard",
      "permissions": ["read"]
    },
    {
      "id": 2,
      "name": "المشاريع",
      "route": "/projects",
      "icon": "mdi-folder-multiple",
      "permissions": ["read", "create", "update", "delete"]
    }
    // ... more pages
  ]
}
```

### Navigation Properties

- RTL layout (Arabic)
- Sidebar width: 320px
- Mobile: Temporary drawer
- Background: Blue gradient `linear-gradient(135deg, #1976d2 0%, #1565c0 100%)`

---

## 5. State Management

### Framework

**Pinia** (Vue 3 Composition API store)

### Auth Store (`/src/stores/auth.js`)

#### State

```javascript
const user = ref(null)          // { id, username }
const pages = ref([])           // User's authorized pages
const isLoading = ref(false)    // Loading state
```

#### Computed Properties

```javascript
isAuthenticated    // Check if user is logged in
isMobile           // Check if mobile device
allowedRoutes      // Array of accessible routes
permissionsMap     // Map of route -> permissions
```

#### Actions

```javascript
fetchPages()                         // Fetch pages from API
loadFromStorage()                    // Load from localStorage
setUser(userData)                    // Set current user
clearAuth()                          // Clear session
hasPermission(route, permission)     // Check permission
canAccessPage(route)                 // Check access
getPagePermissions(route)            // Get page permissions
```

### Persistence

On app initialization:

1. Load `user` from `localStorage.user`
2. Load `pages` from `localStorage.userPages`
3. If authenticated, refresh pages from API

---

## 6. API Endpoints

### Authentication

```
POST   /auth/login              - Login
POST   /auth/refresh            - Refresh token
POST   /auth/logout             - Logout
GET    /auth/pages              - Get authorized pages
```

### Dashboard

```
GET    /dashboard/stats         - Dashboard statistics
GET    /dashboard/financial     - Financial stats
GET    /dashboard/project-progress - Project progress
GET    /dashboard/task-summary  - Task summary
GET    /dashboard/activities    - Recent activities
```

### Projects

```
GET    /projects                - List projects (paginated)
GET    /projects/:id            - Get single project
GET    /projects/stats          - Project statistics
GET    /projects/dropdown       - Projects dropdown (id, name)
GET    /projects/:id/workdays   - Project workdays
GET    /projects/:id/team-members - Project team
GET    /projects/:id/expenses   - Project expenses
POST   /projects                - Create project
PUT    /projects/:id            - Update project
```

### Users

```
GET    /users                   - List users (paginated)
GET    /users/:id               - Get single user
GET    /users/:id/roles         - User roles
GET    /users/:id/teamMembers   - User team members
GET    /users/dropdown          - Users dropdown (id, fullName, jobTitle)
POST   /users                   - Create user
PUT    /users/:id               - Update user
PUT    /users/:id/password      - Update password
PATCH  /users/:id/status        - Update status
DELETE /users/:id               - Delete user
```

### Roles

```
GET    /roles                   - List roles
GET    /roles/:id               - Get single role
POST   /roles                   - Create role
PUT    /roles/:id               - Update role
DELETE /roles/:id               - Delete role
POST   /userRoles               - Assign role to user
DELETE /userRoles/:id           - Remove role from user
```

### Pages (Admin)

```
GET    /pages                   - List all pages
GET    /pages/:id               - Get single page
POST   /pages                   - Create page
PUT    /pages/:id               - Update page
DELETE /pages/:id               - Delete page
```

### Role-Page Permissions

```
GET    /rolePages               - List assignments
GET    /rolePages/:id           - Get single assignment
GET    /rolePages?roleId=X      - Get pages for role
POST   /rolePages               - Create assignment
PUT    /rolePages/:id           - Update permissions
PUT    /rolePages/role/:roleId  - Bulk update for role
DELETE /rolePages/:id           - Delete assignment
```

### Workdays

```
GET    /workdays                - List workdays (paginated)
GET    /workdays/:id            - Get single workday
GET    /projects/:id/workdays   - Workdays by project
POST   /workdays                - Create workday
PUT    /workdays/:id            - Update workday
PATCH  /workdays/:id/complete   - Mark complete
PATCH  /workdays/:id/uncomplete - Mark incomplete
DELETE /workdays/:id            - Delete workday
```

### Workday Details

```
GET    /workdays/:id/materials  - Get materials
GET    /workdays/:id/labor      - Get labor
GET    /workdays/:id/equipment  - Get equipment
POST   /workdayMaterials        - Create material
PUT    /workdayMaterials/:id    - Update material
DELETE /workdayMaterials/:id    - Delete material
POST   /workdayLabor            - Create labor
PUT    /workdayLabor/:id        - Update labor
DELETE /workdayLabor/:id        - Delete labor
POST   /workdayEquipment        - Create equipment
PUT    /workdayEquipment/:id    - Update equipment
DELETE /workdayEquipment/:id    - Delete equipment
```

### Categories

```
GET    /workCategories          - List categories (paginated)
GET    /workCategories/:id      - Get single category
GET    /workCategories/stats    - Category statistics
POST   /workCategories          - Create category
PUT    /workCategories/:id      - Update category
DELETE /workCategories/:id      - Delete category
GET    /workSubcategories       - List subcategories (paginated)
GET    /workSubcategories/:id   - Get single subcategory
POST   /workSubcategories       - Create subcategory
PUT    /workSubcategories/:id   - Update subcategory
DELETE /workSubcategories/:id   - Delete subcategory
```

### Team Members

```
GET    /teamMembers             - List team members (paginated)
GET    /teamMembers/:id         - Get single member
GET    /teamMembers/stats       - Team statistics
POST   /teamMembers             - Add to project
DELETE /teamMembers/:id         - Remove from project
```

### Expenses

```
GET    /expenses                - List expenses (paginated)
GET    /expenses/:id            - Get single expense
GET    /expenses/stats          - Expense statistics
GET    /projects/:id/expenses   - Expenses by project
POST   /expenses                - Create expense
PUT    /expenses/:id            - Update expense
DELETE /expenses/:id            - Delete expense
```

### Income

```
GET    /income                  - List income (paginated)
GET    /income/:id              - Get single income
GET    /income/stats            - Income statistics
POST   /income                  - Create income
PUT    /income/:id              - Update income
DELETE /income/:id              - Delete income
```

### Debtors

```
GET    /debtors                 - List debtors (paginated)
GET    /debtors/:id             - Get single debtor
GET    /debtors/stats           - Debtor statistics
GET    /debtors/activeWithRemaining - Active with remaining debt
POST   /debtors                 - Create debtor
PUT    /debtors/:id             - Update debtor
DELETE /debtors/:id             - Delete debtor
```

### Pagination

Default pagination config:

```javascript
DEFAULT_PAGE = 1
DEFAULT_LIMIT = 20
```

Query parameters: `?page=1&limit=20`

---

## 7. Request Flow Examples

### Login Flow

```
1. User enters credentials on /login
2. handleLogin({ username, password })
3. POST /api/auth/login
4. Response: { accessToken, refreshToken, userId, username }
5. Store tokens in localStorage
6. Set user in authStore
7. Fetch pages: GET /api/auth/pages
8. Store pages in authStore and localStorage
9. Redirect to / (home)
```

### Authenticated API Request Flow

```
1. Component calls API function (e.g., getProjects())
2. client.js adds Authorization header
3. Request sent to /api/projects
4. If 401: trigger token refresh
5. If refresh succeeds: retry with new token
6. If refresh fails: clear session, redirect to /login
7. Return response data to component
```

### Permission Check Flow

```
1. User navigates to /projects
2. Router guard checks authStore.canAccessPage('/projects')
3. Look up '/projects' in permissionsMap
4. If found and has 'read' permission: allow
5. If not found: redirect to /unauthorized
```

---

## 8. Key Files Reference

### API Layer

| File | Purpose |
|------|---------|
| `/src/api/client.js` | HTTP client, token management |
| `/src/api/auth.js` | Auth endpoints |
| `/src/api/projects.js` | Projects endpoints |
| `/src/api/users.js` | Users endpoints |
| `/src/api/roles.js` | Roles endpoints |
| `/src/api/pages.js` | Pages endpoints |
| `/src/api/rolePages.js` | Role-page permissions |
| `/src/api/workdays.js` | Workdays endpoints |
| `/src/api/materials.js` | Materials, labor, equipment |
| `/src/api/teamMembers.js` | Team members |
| `/src/api/expenses.js` | Expenses |
| `/src/api/income.js` | Income |
| `/src/api/debtors.js` | Debtors |
| `/src/api/categories.js` | Categories/subcategories |
| `/src/api/dashboard.js` | Dashboard stats |

### State & Auth

| File | Purpose |
|------|---------|
| `/src/stores/auth.js` | Pinia auth store |
| `/src/services/authService.js` | Auth business logic |
| `/src/composables/usePermissions.js` | Permission composable |

### Routing & Layout

| File | Purpose |
|------|---------|
| `/src/router/index.js` | Routes and guards |
| `/src/App.vue` | Main layout with sidebar |
| `/src/pages/login.vue` | Login page |

### Constants

| File | Purpose |
|------|---------|
| `/src/constants/pagination.js` | Pagination defaults |

---

## API Response Format

### Success Response

```json
{
  "success": true,
  "data": { ... } // or [...]
}
```

### Error Response

```json
{
  "success": false,
  "error": "Error message"
}
```

### Paginated Response

```json
{
  "success": true,
  "data": [...],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 100,
    "totalPages": 5
  }
}
```

---

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `VITE_API_URL` | API base URL | `/api` |

---

## 9. All Frontend Pages

### Pages Connected to Backend API

These pages fetch and send data to the backend:

| Page File | Route | API Used | Description |
|-----------|-------|----------|-------------|
| `/src/pages/index.vue` | `/` | Dashboard API | Home/Dashboard page |
| `/src/pages/login.vue` | `/login` | Auth API | Login page |
| `/src/pages/projects.vue` | `/projects` | Projects API | Projects list |
| `/src/pages/project-management.vue` | `/project-management` | Projects API | Project management |
| `/src/pages/users.vue` | `/users` | Users API | Users list |
| `/src/pages/engineers.vue` | `/engineers` | Users API | Engineers (alias for users) |
| `/src/pages/roles.vue` | `/roles` | Roles API | Role management |
| `/src/pages/pages.vue` | `/pages` | Pages API | Page management (admin) |
| `/src/pages/teamMembers.vue` | `/teamMembers` | TeamMembers API | Team members list |
| `/src/pages/team-members.vue` | `/team-members` | TeamMembers API | Team members (alternate) |
| `/src/pages/expenses.vue` | `/expenses` | Expenses API | Expenses list |
| `/src/pages/income.vue` | `/income` | Income API | Income list |
| `/src/pages/debtors.vue` | `/debtors` | Debtors API | Debtors list |
| `/src/pages/work-days.vue` | `/work-days` | Workdays API | Work days list |
| `/src/pages/work-day-details.vue` | `/work-day-details` | Workdays API | Work day details |
| `/src/pages/labor-details.vue` | `/labor-details` | Materials API (labor) | Labor entries |
| `/src/pages/equipment-details.vue` | `/equipment-details` | Materials API (equipment) | Equipment entries |
| `/src/pages/materials-expenses-details.vue` | `/materials-expenses-details` | Materials API | Materials entries |

### Placeholder Pages (No Backend Connection)

These pages use local mock data and do NOT connect to the backend. They need API integration:

| Page File | Route | Description |
|-----------|-------|-------------|
| `/src/pages/sales.vue` | `/sales` | Sales management (mock data) |
| `/src/pages/purchases.vue` | `/purchases` | Purchases management (mock data) |
| `/src/pages/inventory.vue` | `/inventory` | Inventory management (mock data) |
| `/src/pages/human-resources.vue` | `/human-resources` | Human resources (mock data) |
| `/src/pages/categories.vue` | `/categories` | Work categories (mock data) |
| `/src/pages/expense-types.vue` | `/expense-types` | Expense types (mock data) |
| `/src/pages/project-expenses.vue` | `/project-expenses` | Project expenses view (mock data) |
| `/src/pages/project-details.vue` | `/project-details` | Project details view (mock data) |
| `/src/pages/task-management.vue` | `/task-management` | Task management (mock data) |
| `/src/pages/team-management.vue` | `/team-management` | Team management (mock data) |
| `/src/pages/working-day-details.vue` | `/working-day-details` | Working day details (mock data) |

**Note**: These pages have full UI but use hardcoded/mock data. To integrate with backend, add API imports and replace mock data with API calls.

### System/Utility Pages

| Page File | Route | Description |
|-----------|-------|-------------|
| `/src/pages/unauthorized.vue` | `/unauthorized` | Unauthorized access page |
| `/src/pages/example-design-system.vue` | `/example-design-system` | Design system demo |
| `/src/pages/test-dialog.vue` | `/test-dialog` | Dialog testing |
| `/src/pages/test-icons-final.vue` | `/test-icons-final` | Icon testing |

---

## 10. Route to Permission Mapping (Complete)

```javascript
const routeToPermissionMap = {
  '/project-management': '/projects',
  '/task-management': '/projects',
  '/engineers': '/users',
  '/expense-types': '/expenseTypes',
  '/human-resources': '/humanResources',
  '/team-management': '/teamMembers',
  '/working-day-details': '/workdays',
  '/work-days': '/workdays',
  '/work-day-details': '/workdays',
  '/labor-details': '/workdayLabor',
  '/equipment-details': '/workdayEquipment',
  '/materials-expenses-details': '/workdayMaterials'
}
```

**Note**: Routes not in this map use their path directly for permission checks.

---

## 11. API Files Summary

| API File | Endpoints | Used By |
|----------|-----------|---------|
| `auth.js` | `/auth/*` | Login, logout, refresh, pages |
| `projects.js` | `/projects/*` | Projects, project-management |
| `users.js` | `/users/*` | Users, engineers |
| `roles.js` | `/roles/*`, `/userRoles/*` | Roles management |
| `pages.js` | `/pages/*` | Admin page management |
| `rolePages.js` | `/rolePages/*` | Role-page permissions |
| `workdays.js` | `/workdays/*`, `/workSubcategories/*` | Work days |
| `materials.js` | `/workdayMaterials/*`, `/workdayLabor/*`, `/workdayEquipment/*` | Work day details |
| `teamMembers.js` | `/teamMembers/*` | Team members |
| `expenses.js` | `/expenses/*`, `/projects/:id/expenses` | Expenses |
| `income.js` | `/income/*` | Income |
| `debtors.js` | `/debtors/*` | Debtors |
| `categories.js` | `/workCategories/*`, `/workSubcategories/*` | Categories |
| `dashboard.js` | `/dashboard/*` | Dashboard stats |

---

## Notes

- **Language**: Arabic (RTL layout)
- **Framework**: Vue 3 + Composition API
- **State Management**: Pinia
- **HTTP Client**: Custom fetch-based client
- **Auth**: JWT Bearer tokens
- **Token Refresh**: Automatic with retry
- **Placeholder Pages**: sales, purchases, inventory, human-resources need backend API integration
