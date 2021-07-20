BASE_REGISTRY_IMAGE=weblogin
master=$BASE_REGISTRY_IMAGE/master:latest
(
    set -x
    docker build ../cmd/master/ -t $master
)
