# Architecture

The system overview is a microservices architecture where all microservices interact with each other but are separated. Each microservice controls a part of the overall system or performs a specific task and its data is unique.

Going deeper into each microservice, each of them has multiple layers, that is, they have a presentation layer and business logic. The third layer (database) is outside the microservice to scale the architecture.

Microservices that are connected to a database or to a database balancer have a responsibility to maintain the integrity of their data or send data to other microservices that require that part of the data. This last part could cause a bottleneck in the system.

To fix this, the calls between microservices could be asynchronous, so we could add a queue in the system. This adds complexity to the system, but prepares it to scale more easily.

In addition to microservices, the system needs a configuration method to change all microservices as easily as possible. In addition, the system needs a central registry to improve it and check for problems quickly.

The system needs an API gateway to be able to change the cloud without changing the applications that connect to this cloud and to balance the load if one service, or more, scales.

## Tools

### Configuration tool

All cloud architecture needs a configuration tool. All these tools are basically k-v (key-value) and we can find tools like Etcd or Consul. There is not much difference in scalability or availability between the two, nor in the maintenance of both tools in their respective repositories, but Etcd seems to us to be more widespread. For this reason, in this project Etcd will be used.

### Logging

