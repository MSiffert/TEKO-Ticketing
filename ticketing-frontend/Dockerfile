# Stage 1: Build the React application
FROM node:18 AS build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock) files
COPY package*.json ./
COPY yarn.lock* ./

# Install dependencies
RUN yarn install --silent

# Copy the rest of your app's source code from your host to your image filesystem.
COPY . .

# Build the app
RUN yarn build

# Stage 2: Serve the app with Nginx
FROM nginx:alpine

# Copy the built assets from the build stage to the appropriate folder for Nginx
COPY --from=build /app/build /usr/share/nginx/html

# Expose port 80 to the outside once the container has launched
EXPOSE 80

# Configure Nginx and start it
CMD ["nginx", "-g", "daemon off;"]
