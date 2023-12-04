#!/bin/bash

# Start Golang backend using fresh
fresh &

# Navigate to the frontend directory and start SvelteKit dev server
cd /app/frontend
npm run dev -- --host 0.0.0.0
