# Powershell
```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host naren4b/dev sh

```
# metric Simulator container 
```
docker run --rm -p 8080:8080 --name=metrics naren4b/metric-simulator --ac=2021 --mc=10

```
# For running both metric simulator and prometheus container

```
docker-compose up
```