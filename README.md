# Torizon Containers

This repository containers all container images used by or with 
[TorizonCore](https://www.toradex.com/operating-systems/torizon-core), 
the Easy-to-use Industrial Linux Software Platform.

## Steps after forking

In order for the pipeline to push images to your respective fork's image 
registry, you need to setup a `GIT_TOKEN` variable by

1. Going into the project's "Settings"
2. `Project Access Tokens`
3. Choosing a Token Name (can be any)
4. Selecting the following capabilities: `api`, `read_api`, `read_repository`,
`write_repository`, `read_registry`, `write_registry`
5. "Create project access token"
6. Copying the newly generate token
7. Going to "CI/CD" under the project's "Settings"
8. Expanding the "Variables" menu
9. "Add a new variable"
10. `GIT_TOKEN` as "Key"
11. The token generated copied in step 6 as "Value"
12. Check "Mask variable"
13. Add
