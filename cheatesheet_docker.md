Fabriquer une image Docker:
``` bash
docker build -t <image_name> .
```

Lister les images locales:
``` bash
docker images
```

Supprimer une image:
``` bash
docker rmi <image_name>
```

Executer un conteneur:
``` bash
docker run --name <container_name> <image_name>
```

Executer un conteneur et expose le port:
``` bash
docker run -p <port_hote>:<container_port> <image_name>
```

Executer bash dans un conteneur:
``` bash
docker run -it <image_name> /bin/bash
```

Liste les conteneurs actifs:
``` bash
docker ps
```