# Frontend Setup

## Option A: local Node.js
Initialize Next.js app in this folder:

npx create-next-app@latest . --typescript --tailwind --eslint --app --src-dir --import-alias "@/*"

Install required libraries:

npm install @tanstack/react-query @tanstack/react-table react-hook-form zod @hookform/resolvers
npx shadcn@latest init

Suggested extra packages:

npm install date-fns clsx lucide-react

Initial pages to build:
- /login
- /dashboard
- /orders
- /orders/new
- /products
- /customers

## Option B: Docker (no local Node needed)
From project root:

docker run --rm -it -v "${PWD}/frontend:/app" -w /app node:20-alpine sh -c "npx create-next-app@latest . --typescript --tailwind --eslint --app --src-dir --import-alias '@/*'"

Install libraries:

docker run --rm -it -v "${PWD}/frontend:/app" -w /app node:20-alpine sh -c "npm install @tanstack/react-query @tanstack/react-table react-hook-form zod @hookform/resolvers"
