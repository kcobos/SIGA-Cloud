# SIGA-Cloud

SIGA (Sistema Integral de Gestión de Aparcamientos (Integral Parking Management System)) pretends to be a solution for parkings problem at 21's century. Specifically, SIGA is designed for reserved parkings in public area, like authorities, ambulances or, its main aim, disabled parkings. It tries to get any kind of information about these parkings (location, parking type, occupation status), process the information and to serve it as easy as possible to all citizens and users. The original documentation about this project is on <https://github.com/kcobos/SIGA>.

This solution has two separated parts: the cloud, the main processing one, and the sensors (IoT) which are disposed in parkings attached to this project and provide the actual occupation status information to the solution. Due to that, this solution has to be a cloud system to be able to scale whatever part when more parkings are added.

The main information in this system is the parking occupation and location and users or vehicles which can park in these reserved parkings. Processing this information we have to get an historical occupation for each parking, automatic incident notifications, incident historical..., even to predict occupation status of parkings when a user uses the system to get a parking near a place.

## Cloud architecture

SIGA cloud architecture is an hybridization between microservices architecture, multilayer architecture and event-driven architecture. More information [here](docs/architecture.md).
