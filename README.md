# go-htmx-tailwind-example

Example CRUD app written in Go + HTMX + Tailwind CSS

This project implements a pure dynamic web app with SPA-like features but without heavy complex Javascript or Go frameworks to keep up with.  Just HTML/CSS + Go ❤️

![screenshot](./screenshot.jpeg)


## Deploy

Run the following command to deploy the app to AWS App Runner
```
make apprunner
```


## Develop

```
 Choose a make command to run

  init           initialize project (make init module=github.com/user/project)
  vet            vet code
  test           run unit tests
  build          build a binary
  docker-build   build project into a docker container image
  docker-run     run project in a container
  start          build and run local project
  css            build tailwindcss
  css-watch      watch build tailwindcss
  image          build code into a container image and push it to ECR
  apprunner      create App Runner service using Terraform
  destroy        tear down infrastructure
```
