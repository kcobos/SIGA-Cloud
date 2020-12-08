# SIGA-Cloud

SIGA (Sistema Integral de Gestión de Aparcamientos (Integral Parking Management System)) pretends to be a solution for parkings problem at 21's century. Specifically, SIGA is designed for reserved parkings in public area, like authorities, ambulances or, its main aim, disabled parkings. It tries to get any kind of information about these parkings (location, parking type, occupation status), process the information and to serve it as easy as possible to all citizens and users. The original documentation about this project is on <https://github.com/kcobos/SIGA>.

This solution has two separated parts: the cloud, the main processing one, and the sensors (IoT) which are disposed in parkings attached to this project and provide the actual occupation status information to the solution. Due to that, this solution has to be a cloud system to be able to scale whatever part when more parkings are added.

The main information in this system is the parking occupation and location and users or vehicles which can park in these reserved parkings. Processing this information we have to get an historical occupation for each parking, automatic incident notifications, incident historical..., even to predict occupation status of parkings when a user uses the system to get a parking near a place.

## Cloud architecture

SIGA cloud architecture is an hybridization between microservices architecture, multilayer architecture and event-driven architecture. More information [here](docs/architecture.md).

## Road map

All milestones of this project must be a product. Due to that, we have the next planning:

1. Basic list: System has to list all city parking places to show them in a list or in a map. User stories:
   1. [Set up parking sensor](https://github.com/kcobos/SIGA-Cloud/issues/2)
   2. [Set up a parking place](https://github.com/kcobos/SIGA-Cloud/issues/3)
   3. [Show parkings](https://github.com/kcobos/SIGA-Cloud/issues/16)
   4. [Show places](https://github.com/kcobos/SIGA-Cloud/issues/6)
2. Parking status: All parking places in the system must be updated if the parking lot get occupied or freed. User stories:
   1. [Update parking](https://github.com/kcobos/SIGA-Cloud/issues/7)
   2. [Show places](https://github.com/kcobos/SIGA-Cloud/issues/6)
3. Reports: System has to log parking statuses and show reports. User stories:
   1. [Get parking report](https://github.com/kcobos/SIGA-Cloud/issues/9)
4. Navigation: Users can navigate to places and get notify if the place become occupied. User stories:
   1. [Find place](https://github.com/kcobos/SIGA-Cloud/issues/15)
   2. [Place navigation](https://github.com/kcobos/SIGA-Cloud/issues/8)

This project follows Kanban methodology using [GitHub project](https://github.com/kcobos/SIGA-Cloud/projects/2). The tasks on the top of "To do" and "In progress" lists have to do before than others.

## Classes

- [Custom errors](internal/errors/errors.go)
- [Parking](internal/models/parking.go)
- [Parking statuses](internal/models/status.go)
- [Parking controller](internal/controllers/parkings.go)
- [Place](internal/models/place.go)
- [Place controller](internal/controllers/places.go)

## Build tool

The GoLang will be chosen for much of the project. Due to that, the **build tool** will be Go itself.

## Task runner

There are few tasks manager written in GoLang like [Task](https://github.com/go-task/task), [Robo](https://github.com/tj/robo), [go-task](https://github.com/leandroveronezi/go-task), [realize](https://github.com/oxequa/realize). As we can see on them repositories, *Task* has been updated later and it has more contributors than the others and it has a lot of [documentation](https://taskfile.dev). So, we are going to choose **Tasks** as a task manager because it seems simpler and we only write tasks in a yaml file.

Tasks can be installed following [these steps](https://taskfile.dev/#/installation?id=install-script).

To run any task, run:

```bash
task task-name
```

### Task file

All tasks are written in [taskfile.yml](./Taskfile.yml). Tasks are:

- syntax: for checking sources syntax.
- test: for testing all project. See [tests](#test-files).
- installdeps: Install project dependencies. See [dependencies](go.mod).

## Tests

GoLang has a main official package for automated testing and benchmarking called [Testing](https://golang.org/pkg/testing/) but it doesn't have assertions, so we have to choose a library to code faster. In this [repo](https://github.com/bmuschko/go-testing-frameworks) there is a comparison between a lot of GoLang testing libraries.

For this project, [Goblin](https://github.com/franela/goblin) to test user stories like BDD. For bug and  enhancement issues we are going to choose [Testify](https://github.com/stretchr/testify) because maybe these issues don't have an user story. We have chosen that two due to their community and simplicity and we don't need a web interface for testing for now like [GoCenvey](https://github.com/smartystreets/goconvey).

We choose for reference the user story in the description each test.

To test this project, run:

```bash
task test
```

### Test files

- [Parking test](internal/models/parking_test.go)
- [Parking controller test](internal/controllers/parkings_test.go)
- [Place test](internal/models/place_test.go)
- [Place controller test](internal/controllers/places_test.go)

## Dockerize for testing

First of all, we must say we have chosen GoLang as programming language and our task runner is also written in GoLang too, so we don't need more than Go. There is no code which needs special libraries for now.

We have looked for Linux built Go images and we have found:

- **golang**: official image without any user. This image is built to Linux main platforms (x86, x86-64, ARM, ARM64). It has these tags:
  - normal ([debian buster](https://github.com/docker-library/golang/blob/a7f393378d8566caf777ad2e6b9dc9d014875a88/1.15/buster/Dockerfile)). Uncompressed size: 839 MB
  - [alpine](https://github.com/docker-library/golang/blob/a7f393378d8566caf777ad2e6b9dc9d014875a88/1.15/alpine3.12/Dockerfile). Uncompressed size: 300 MB
- **neroinc/fedora-golang**: non official image without any user. It has an example program.
  - [normal](https://github.com/NeroINC/docker/blob/master/fedora-golang/Dockerfile). Uncompressed size: 877 MB
- **jcajka/fedora-golang**: non official image without any user.
  - [normal](https://github.com/jcajka/fedora-golang/blob/master/f25/Dockerfile). Uncompressed size: 516 MB

So, the best option to build an image for that project is from *golang:alpine* image due it is official, it has the latest language version and it is the smallest. It is even used in [docker examples](https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds). [Read more image comparison here](https://github.com/kcobos/Ejercicios-CC/blob/master/Tema3/Comparacion_imagenes.md).

Due this image is for testing, we cannot optimize the image because we cannot do a [multi-stage](https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds) build.
