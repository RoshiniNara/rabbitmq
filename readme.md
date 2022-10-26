To test this code, Rabbit MQ instance should be running on the localhost on the port 5672

Install RabbitMQ via docker. 
    docker pull rabbitmq
    docker pull rabbitmq:3-management 
Run the docker instance from the download image 
    docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672  rabbitmq:3-management

Connection information 
    default rabbitmq username is guest 
    default rabbitmq password is guest 
