# SIGA-Cloud

SIGA (Sistema Integral de Gestión de Aparcamientos (Integral Parking Management System)) pretends to be a solution for parkings problem at 21's century. Specifically, SIGA is designed for reserved parkings in public area, like authorities, ambulances or, its main aim, disabled parkings. It tries to get any kind of information about these parkings (location, parking type, occupation status), process the information and to serve it as easy as possible to all citizens and users. The original documentation about this project is on <https://github.com/kcobos/SIGA>.

This solution has two separated parts: the cloud, the main processing one, and the sensors (IoT) which are disposed in parkings attached to this project and provide the actual occupation status information to the solution. Due to that, this solution has to be a cloud system to be able to scale whatever part when more parkings are added.

The main information in this system is the parking occupation and location and users or vehicles which can park in these reserved parkings. Processing this information we have to get an historical occupation for each parking, automatic incident notifications, incident historical..., even to predict occupation status of parkings when a user uses the system to get a parking near a place.

## Cloud architecture

SIGA cloud architecture is an hybridization between microservices architecture, multilayer architecture and event-driven architecture. More information [here](docs/architecture.md).

## Road map

All milestones of this project must be a product. Due to that, we have the next planning:

1. Basic administration: set up the project for a basic administration functionalities.
2. Basic functionalities: add functionalities to user can use the system.
3. Advance administration: add admins and operators functionalities.
4. Advance functionalities: add functionalities like look for a place.

This project follows Kanban methodology using [GitHub project](https://github.com/kcobos/SIGA-Cloud/projects/2). The tasks on the top of "To do" and "In progress" lists have to do before than others.

## Classes

### Parking

- [Parking](parking/models/parking.go)
- [Parking statuses](parking/models/status.go)
- [Custom errors](parking/errors/errors.go)
- [Controller](parking/parkings.go)

## Build tool

The GoLang will be chosen for much of the project. Due to that, the **build tool** will be Go itself.

## Task runner

There are two main tasks manager written in GoLang [Task](https://github.com/go-task/task) and [Robo](https://github.com/tj/robo). As we can see on them repositories, *Task* is updated later and it has more contributors than Robo and the first has a lot of [documentation](https://taskfile.dev). So, we are going to choose **Tasks** as a task manager.

Tasks can be installed following [this steps](https://taskfile.dev/#/installation?id=install-script).

### Task file

All tasks are written in [taskfile.yml](./Taskfile.yml). Tasks are:

- syntax: for checking sources syntax
