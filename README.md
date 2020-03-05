# Smart Agent

A smart agent to monitor kubernetes system and provide HTTP access.  

Forked from http-echo, there are informations in Iceman's Cookbook Kubernetes chapter about http-echo, also in Iceman's repo, there's cluster/go/http-echo which is the original version.  

Check summer2019 and cheers2019 in Iceman's Cookbook Kubernetes chapter as well.  

## Files Tree

## Build

### Command Line  

The suggestion is to checkout code into $HOME/go/src/github.com/icodebb/smart-agent.

e.g. on SuperMicro Debian Linux:  

```bash
cd /home/ice/go/src/github.com/icodebb/smart-agent/cmd
build -o smart-agent
```

### Build Docker Image

e.g. on SuperMicro Debian Linux:

```bash
cd /home/ice/go/src/github.com/icodebb/smart-agent
docker build -t smart-agent .
docker images
docker tag smart-agent imiceman/smart-agent:1.0.0
docker login
docker push imiceman/smart-agent:1.0.0
```

## Docker Images

### smart-agent

[hub.docker.com/repository/docker/imiceman/smart-agent](https://hub.docker.com/repository/docker/imiceman/smart-agent)  

### Others

[hub.docker.com/repository/docker/imiceman/http-echo](https://hub.docker.com/repository/docker/imiceman/http-echo)  
[hub.docker.com/repository/docker/imiceman/summer2019](https://hub.docker.com/repository/docker/imiceman/summer2019)  
[hub.docker.com/repository/docker/imiceman/cheers2019](https://hub.docker.com/repository/docker/imiceman/cheers2019)  

## Deployment

### Run from command line

E.g. on SuperMicro Linux:  

```bash
cd /home/ice/go/src/github.com/icodebb/smart-agent/cmd
go run main.go
# or
go build -o smart-agent
./smart-agent
```

Then open a Web browser, input http://10.207.27.21:9999/, it should show  
Hello, Friend.  

input http://10.207.27.21:9999/?name=Iceman, then it should show:  
Hello, Iceman.  

### Kubernetes

## Reference

[docker/doodle](https://github.com/docker/doodle) - summer2019 and cheers2019  
