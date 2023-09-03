# Go Dockertest Example

[Dockertest](https://github.com/ory/dockertest) is an amazing library that allows us to manage docker containers as part of a test suite. This means instead of having some docker compose stack or other setup (outside of tests) we can create a clean docker environment on every run.

This repo has a single entry point,