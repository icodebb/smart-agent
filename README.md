# Smart Agent

A smart agent to monitor kubernetes system and provide HTTP access.  

Forked from http-echo, there are informations in Iceman's Cookbook Kubernetes chapter about http-echo, also in Iceman's repo, there's cluster/go/http-echo which is the original version.  

Check summer2019 and cheers2019 in Iceman's Cookbook Kubernetes chapter as well.  

## Files Tree

## Build

The suggestion is to checkout code into $HOME/go/src/github.com/icodebb/smart-agent.

## Docker Images

[hub.docker.com/repository/docker/imiceman/http-echo](https://hub.docker.com/repository/docker/imiceman/http-echo)  
[hub.docker.com/repository/docker/imiceman/summer2019](https://hub.docker.com/repository/docker/imiceman/summer2019)  
[hub.docker.com/repository/docker/imiceman/cheers2019](https://hub.docker.com/repository/docker/imiceman/cheers2019)  

## Deployment

### Run from command line

E.g. on SuperMicro Linux:  

```bash
cd /home/ice/go/src/github.com/icodebb/smart-agent
go run src/main.go
```

Then open a Web browser, input http://10.207.27.21:9999/, it should show  
Hello, Friend.  

input http://10.207.27.21:9999/?name=Iceman, then it should show:  
Hello, Iceman.  

### Kubernetes

## Reference

[docker/doodle](https://github.com/docker/doodle) - summer2019 and cheers2019  
