Fabriquer une image Docker:
docker build -t <image_name> .

Lister les images locales:
docker images

Supprimer une image:
docker rmi <image_name>

Executer un conteneur:
docker run --name <container_name> <image_name>

Executer un conteneur et expose le port:
docker run -p <port_hote>:<container_port> <image_name>

Executer bash dans un conteneur:
docker run --it <image_name> /bin/bash

Liste les conteneurs actifs:
docker ps