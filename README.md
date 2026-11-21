# Go Distroless Image Optimization

A hands-on container engineering project demonstrating Docker
multi-stage builds, static Go compilation, and minimal Distroless
runtime images.

## Objectives

- Develop a lightweight Go web application.
- Establish a baseline Docker image.
- Implement Docker multi-stage builds.
- Compile statically linked Go binaries.
- Use Google Distroless as the production runtime.
- Run containers as a non-root user.
- Measure image size and layer optimization.
- Explore container security and minimal runtime environments.

## Technology Stack

- Go
- Docker
- Docker Multi-stage Builds
- Google Distroless
- HTML5
- CSS3
- JavaScript

## Application

ContainerScope is a Go-powered dashboard for demonstrating
container image optimization and multi-stage Docker architecture.

## Architecture

The application will use a multi-stage Docker build:

1. Builder stage: Compile the Go application.
2. Runtime stage: Copy the compiled binary into a minimal
   Distroless image.
3. Execution: Run the application as a non-root user.

Frontend assets will be embedded into the Go executable using
the standard library embed package.

## Project Status

Currently under development.

- [x] Go development environment
- [x] GitHub repository initialization
- [x] Application configuration
- [x] Initial frontend structure
- [x] Embedded asset foundation
- [ ] Go HTTP server
- [ ] Dashboard and health endpoints
- [ ] Baseline Docker image
- [ ] Multi-stage Docker build
- [ ] Distroless production image
- [ ] Image size comparison
- [ ] Container security verification

## License

MIT
