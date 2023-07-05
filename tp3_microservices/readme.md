# Installez docker-compose:

https://docs.docker.com/compose/install/

# Build l'image docker du proxy

```
docker build -t proxy .
```

# Lancez la stack avec

```
docker-compose up
```

# Testez et comparez les réponses des deux services:
- de l'api: http://localhost:8080/v1/pets
- du proxy: http://localhost/v1/pets