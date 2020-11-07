# Docker Container

### Rodar um Container
```bash
docker container run ubuntu
```

### Visualizar todos container activos
```bash
docker container ls
```

### Visualizar todos container activs e não activos
```bash
docker container ls -a
```

### Rodar um comando no Container
```bash
docker container run ubuntu /bin/bash
```
### Acessar ocontainer
```bash
docker container run -it ubuntu /bin/bash
```

### Executar um container auto remove
```bash
docker container run -it --rm ubuntu /bin/bash
```

### Acessar um container pela porta em modo detached (in background)
```bash
docker container run -d -p 8080:80 nginx
```

### Viasulizar log docontainer em execução
```bash
docker container logs nginx
```

### Adicionar um nome ao Container
```bash
docker container run --name nginx2 -d -p 8080:80 nginx
```

### Remover todos container activos e não activos
```
docker container rm $(docker container ls -qa) -f
```

### Listar images docker local
```bash
docker image ls
```

### Remover todas images locais
```bash
docker image rm $(docker image ls -qa) -f
```

### Construir uma image docker pelo Dockerfile
```bash
docker image build -t angolarti/image:v1 .
```