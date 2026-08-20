# Container images

- Dockerfile.controlplane builds a minimal non-root aetherd image.
- Dockerfile.node establishes non-root packaging for the Phase 2 daemon.
- Dockerfile.web builds the Next.js standalone workspace output.

The root docker-compose.yml is the canonical Phase 0 local stack.
