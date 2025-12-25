# How to Deploy on Railway

This project is a monorepo containing both a Go backend and a Vue frontend. To deploy specifically on Railway, you need to create two separate services from the same repository.

## Prerequisites

- A [Railway](https://railway.app/) account.
- Your project connected to GitHub.

## 1. Deploying the Backend (Go)

1.  **New Service**: Click "New Service" -> "GitHub Repo" -> Select this repository.
2.  **Settings**:
    - Go to the **Settings** tab of the new service.
    - Scroll down to **Root Directory** and set it to `/backend`.
    - Railway should automatically detect the `Dockerfile` in the `backend` folder.
3.  **Variables**:
    - Go to the **Variables** tab.
    - Add the following variables:
        - `PORT`: (Railway adds this automatically, do not change)
        - `JWT_SECRET`: Set a strong secret key (e.g., `my-super-secret-key-123`).
        - `DB_HOST`: Your database host (see Database section below).
        - `DB_PORT`: Your database port.
        - `DB_USER`: Your database user.
        - `DB_PASSWORD`: Your database password.
        - `DB_NAME`: Your database name.
        - `DB_SSLMODE`: Set to `require` or `disable` depending on your DB (Railway Postgres usually requires `disable` or no setting if internal, but `require` for public). *Recommendation: Use Railway's internal networking variables if possible.*
4.  **Database**:
    - If you don't have a database, create a **PostgreSQL** service in Railway.
    - Railway provides variables like `DATABASE_URL`. You might need to parse this or just use the individual variables (`PGHOST`, `PGUSER`, etc.) ensuring they map to what the backend expects (`DB_HOST`, etc.).
    - **Tip**: You can assume standard mapped variables if you link the services, or manually copy them.

## 2. Deploying the Frontend (Vue)

1.  **New Service**: Click "New Service" -> "GitHub Repo" -> Select this repository *again*.
2.  **Settings**:
    - Go to the **Settings** tab.
    - Set **Root Directory** to `/frontend`.
    - Railway should detect it as a Node/Vite app.
    - **Build Command**: `npm install && npm run build` (should be default).
    - **Start Command**: Not needed if serving headers only, but for a web service, you typically serve the `dist` folder.
    - *However*, for a pure static site (SPA), it is often better to use a static hosting service (like Vercel or Netlify) or configuring Railway to serve the static files using Nginx or a simple node server.
    - If using Railway's default Node image, it might try to "start" something. Since it's a SPA, you might need a simple server or use `npm run preview` (not recommended for prod).
    - **Recommended**: Create a file named `serve.json` or use `serve -s dist` as the Start Command if using a Node environment.
    - **Start Command**: `npx serve -s dist -l $PORT`
3.  **Variables**:
    - `VITE_API_URL`: The URL of your deployed Backend service (e.g., `https://backend-production.up.railway.app`). **Important**: Do not add a trailing slash.

## Troubleshooting

- **Node Version**: The frontend requires Node.js v20+. We have updated `package.json` to specify this, so Railway should respect it.
- **Build Errors**: Check the "Build Logs" in Railway. 
- **404 on API**: Ensure `VITE_API_URL` is correct and the Backend is running.
