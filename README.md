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

### Parking

- [Parking](parking/models/parking.go)
- [Parking statuses](parking/models/status.go)
- [Custom errors](parking/errors/errors.go)
- [Controller](parking/parkings.go)

### Place

- [Place](place/models/place.go)
- [Custom errors](place/errors/errors.go)
- [Controller](place/places.go)

## Build tool

The GoLang will be chosen for much of the project. Due to that, the **build tool** will be Go itself.

## Task runner

There are two main tasks manager written in GoLang [Task](https://github.com/go-task/task) and [Robo](https://github.com/tj/robo). As we can see on them repositories, *Task* is updated later and it has more contributors than Robo and the first has a lot of [documentation](https://taskfile.dev). So, we are going to choose **Tasks** as a task manager.

Tasks can be installed following [this steps](https://taskfile.dev/#/installation?id=install-script).

### Task file

All tasks are written in [taskfile.yml](./Taskfile.yml). Tasks are:

- syntax: for checking sources syntax

## Tests

GoLang has a main official package for automated testing and benchmarking called [Testing](https://golang.org/pkg/testing/) but it doesn't have assertions, so we have to choose a library to code faster. In this [repo](https://github.com/bmuschko/go-testing-frameworks) there is a comparison between a lot of GoLang testing libraries.

For this project, we are going to choose [Testify](https://github.com/stretchr/testify) for now. Later, maybe, we are going to choose [Goblin](https://github.com/franela/goblin) for BDD.
