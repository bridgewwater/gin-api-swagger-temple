# this file must use as base Makefile
# include z-MakefileUtils/MakeDockerCompose.mk
# need var:
# ROOT_NAME is this project name
# ROOT_OWNER is this project owner
# INFO_BUILD_DOCKER_FROM_IMAGE for image running
# INFO_TEST_BUILD_DOCKER_PARENT_IMAGE for image local build
# INFO_BUILD_DOCKER_FILE for build docker image default Dockerfile
# INFO_TEST_BUILD_DOCKER_FILE for local build docker image file
# INFO_TEST_BUILD_DOCKER_PARENT_IMAGE for local build
# ROOT_PARENT_SWITCH_TAG is change parent image tag

# ENV_INFO_BUILD_DOCKER_TAG=latest
ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_PORT =${ROOT_DOCKER_CONTAINER_PORT}
ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_HOST =0.0.0.0
ENV_INFO_DOCKER_COMPOSE_FILE =${INFO_DOCKER_COMPOSE_DEFAULT_FILE}
ENV_INFO_BUILD_DOCKER_TAG =${ENV_DIST_VERSION}
ENV_INFO_DOCKER_REPOSITORY =${ROOT_NAME}
ENV_INFO_DOCKER_OWNER =${ROOT_OWNER}
ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME =${ROOT_NAME}
# if set ENV_INFO_PRIVATE_DOCKER_REGISTRY= will push to hub.docker.com
# private docker registry use harbor must create project name as ${ENV_INFO_DOCKER_OWNER}
#ENV_INFO_PRIVATE_DOCKER_REGISTRY=harbor.xxx.com/
ENV_INFO_PRIVATE_DOCKER_REGISTRY =${INFO_PRIVATE_DOCKER_REGISTRY}/
ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE ?=${ENV_INFO_DOCKER_OWNER}/${ENV_INFO_DOCKER_REPOSITORY}
ENV_INFO_BUILD_DOCKER_FROM_IMAGE =${INFO_BUILD_DOCKER_FROM_IMAGE}
ENV_INFO_BUILD_DOCKER_FILE =${INFO_BUILD_DOCKER_FILE}

ENV_INFO_TEST_BUILD_DOCKER_FILE ?=${INFO_TEST_BUILD_DOCKER_FILE}
ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE =${INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}:${ROOT_PARENT_SWITCH_TAG}
ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER =test-parent-${ENV_INFO_DOCKER_REPOSITORY}
ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME =test-${ENV_INFO_DOCKER_REPOSITORY}

## change this for ip-v4 get
#ROOT_LOCAL_IP_V4_LINUX = $$(ifconfig enp8s0 | grep inet | grep -v inet6 | cut -d ':' -f2 | cut -d ' ' -f1)
#ROOT_LOCAL_IP_V4_DARWIN = $$(ifconfig en0 | grep inet | grep -v inet6 | cut -d ' ' -f2)
#
#localIPLinux:
#	@echo "=> now run as docker with linux"
#	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_LINUX)"
#
#localIPDarwin:
#	@echo "=> now run as docker with darwin"
#	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_DARWIN)"

.PHONY: dockerEnv
dockerEnv:
	@echo "== docker env print start"
	@echo "ENV_INFO_DOCKER_REPOSITORY                     ${ENV_INFO_DOCKER_REPOSITORY}"
	@echo "ENV_INFO_DOCKER_OWNER                          ${ENV_INFO_DOCKER_OWNER}"
	@echo ""
	@echo "ENV_INFO_BUILD_DOCKER_FILE                     ${ENV_INFO_BUILD_DOCKER_FILE}"
	@echo "ENV_INFO_BUILD_DOCKER_FROM_IMAGE               ${ENV_INFO_BUILD_DOCKER_FROM_IMAGE}"
	@echo "ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE             ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}"
	@echo "ENV_INFO_BUILD_DOCKER_TAG                      ${ENV_INFO_BUILD_DOCKER_TAG}"
	@echo "ENV_INFO_PRIVATE_DOCKER_REGISTRY               ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}"
	@echo "REGISTRY tag as:                               ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"
	@echo ""
	@echo "ENV_INFO_TEST_BUILD_DOCKER_FILE                ${ENV_INFO_TEST_BUILD_DOCKER_FILE}"
	@echo "ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE        ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}"
	@echo "ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER    ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER}"
	@echo "ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME  ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}"
	@echo "ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME          ${ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME}"
	@echo ""
	@echo "== docker env print end"

.PHONY: dockerAllPull
dockerAllPull:
	docker pull ${ENV_INFO_BUILD_DOCKER_FROM_IMAGE}
	docker pull ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}

.PHONY: dockerCleanImages
dockerCleanImages:
	(while :; do echo 'y'; sleep 3; done) | docker image prune

.PHONY: dockerCleanPruneAll
dockerCleanPruneAll:
	(while :; do echo 'y'; sleep 3; done) | docker container prune
	(while :; do echo 'y'; sleep 3; done) | docker image prune

.PHONY: dockerRunContainerParentBuild
dockerRunContainerParentBuild:
	@echo "run rm container image: ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}"
	$(info docker run -d --rm --name ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER} ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE})
	docker run -d --rm --name ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER} ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE} tail -f /dev/null
	@echo ""
	@echo "-> run rm container name: ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER}"
	@echo "-> into container use: docker exec -it ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER} bash"

.PHONY: dockerRmContainerParentBuild
dockerRmContainerParentBuild:
	-docker rm -f ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_CONTAINER}

.PHONY: dockerPruneContainerParentBuild
dockerPruneContainerParentBuild: dockerRmContainerParentBuild
	-docker rmi -f ${ENV_INFO_TEST_BUILD_DOCKER_PARENT_IMAGE}

.PHONY: dockerComposeUp
dockerComposeUp: export ENV_WEB_HOST=0.0.0.0
dockerComposeUp: export ENV_WEB_PORT=${ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_PORT}
dockerComposeUp:
	@echo "-> docker compose up as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -p ${ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME} -f ${ENV_INFO_DOCKER_COMPOSE_FILE} up -d --build --remove-orphans --force-recreate

.PHONY: dockerComposePs
dockerComposePs: export ENV_WEB_HOST=0.0.0.0
dockerComposePs: export ENV_WEB_PORT=${ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_PORT}
dockerComposePs:
	@echo "-> docker compose ps as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -p ${ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME} -f ${ENV_INFO_DOCKER_COMPOSE_FILE} ps

.PHONY: dockerComposeLogs
dockerComposeLogs: export ENV_WEB_HOST=0.0.0.0
dockerComposeLogs: export ENV_WEB_PORT=${ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_PORT}
dockerComposeLogs:
	docker-compose -p ${ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME} -f ${ENV_INFO_DOCKER_COMPOSE_FILE} logs

.PHONY: dockerComposeFollowLogs
dockerComposeFollowLogs: export ENV_WEB_HOST=0.0.0.0
dockerComposeFollowLogs: export ENV_WEB_PORT=${ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_PORT}
dockerComposeFollowLogs:
	docker-compose -p ${ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME} -f ${ENV_INFO_DOCKER_COMPOSE_FILE} logs -f

.PHONY: dockerComposeRestart
dockerComposeRestart: export ENV_WEB_HOST=0.0.0.0
dockerComposeRestart: export ENV_WEB_PORT=${ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_PORT}
dockerComposeRestart:
	@echo "-> docker compose restart as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -p ${ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME} -f ${ENV_INFO_DOCKER_COMPOSE_FILE} restart

.PHONY: dockerComposeDown
dockerComposeDown: export ENV_WEB_HOST=0.0.0.0
dockerComposeDown: export ENV_WEB_PORT=${ENV_INFO_DOCKER_CONTAINER_PUBLISH_HOST_PORT}
dockerComposeDown:
	@echo "-> docker compose down as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -p ${ENV_LOCAL_DOCKER_COMPOSE_PROJECT_NAME} -f ${ENV_INFO_DOCKER_COMPOSE_FILE} down --remove-orphans --rmi local

.PHONY: dockerTestBuildLatest
dockerTestBuildLatest:
	docker build --rm=true \
	--build-arg ENV_BUILD_DIST_CODE_MARK=${ENV_DIST_CODE_MARK} \
	--tag ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG} --file ${ENV_INFO_TEST_BUILD_DOCKER_FILE} .

.PHONY: dockerTestRunLatest
dockerTestRunLatest:
	docker image inspect --format='{{ .Created}}' ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	$(warning you can change test docker run args at here for dev)
	-docker run --rm --name ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} \
	-e RUN_MODE=dev \
	${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	$(info for inner check can use like this)
	$(info docker run -it -d --entrypoint /bin/sh --name ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG})
	-docker inspect --format='{{ .State.Status}}' ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}

.PHONY: dockerTestLogLatest
dockerTestLogLatest:
	-docker logs ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}

.PHONY: dockerTestRmLatest
dockerTestRmLatest:
	-docker rm -f ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME}

.PHONY: dockerTestRmiLatest
dockerTestRmiLatest:
	-docker rmi -f ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}

.PHONY: dockerTestBuildCheck
dockerTestBuildCheck: dockerTestRmLatest dockerTestRmiLatest dockerTestBuildLatest
	$(info -> finish build check ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG})

.PHONY: dockerTestRestartLatest
dockerTestRestartLatest: dockerTestBuildCheck dockerTestRunLatest
	@echo "restart ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

.PHONY: dockerTestStopLatest
dockerTestStopLatest: dockerComposeDown dockerTestRmLatest dockerTestRmiLatest
	@echo "stop and remove ${ENV_INFO_TEST_TAG_BUILD_DOCKER_CONTAINER_NAME} ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

.PHONY: dockerTestPruneLatest
dockerTestPruneLatest: dockerTestStopLatest
	@echo "prune and remove ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

.PHONY: dockerRmiBuild
dockerRmiBuild:
	-docker rmi -f ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	-docker rmi -f ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}

.PHONY: dockerBuild
dockerBuild:
	docker build --rm=true \
	--build-arg ENV_BUILD_DIST_CODE_MARK=${ENV_DIST_CODE_MARK} \
	--tag ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG} --file ${ENV_INFO_BUILD_DOCKER_FILE} .

.PHONY: dockerTag
dockerTag:
	docker tag ${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG} ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}

.PHONY: dockerBeforePush
dockerBeforePush: dockerRmiBuild dockerBuild dockerTag
	@echo "===== then now can push to ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}"

.PHONY: dockerPushBuild
dockerPushBuild: dockerBeforePush
	docker push ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}
	@echo "=> push ${ENV_INFO_PRIVATE_DOCKER_REGISTRY}${ENV_INFO_BUILD_DOCKER_SOURCE_IMAGE}:${ENV_INFO_BUILD_DOCKER_TAG}"

.PHONY: helpDocker
helpDocker:
	@echo "=== this make file can include z-MakefileUtils/MakeDockerCompose.mk then use"
	@echo "- must has file: [ ${ENV_INFO_BUILD_DOCKER_FILE} ${ENV_INFO_TEST_BUILD_DOCKER_FILE} ${ENV_INFO_DOCKER_COMPOSE_FILE} ]"
	@echo "- then change tag as:                       INFO_BUILD_DOCKER_TAG"
	@echo "- then change repository as:                INFO_REPOSITORY"
	@echo "- then change owner as:                     INFO_OWNER"
	@echo "- then change private docker repository as: INFO_PRIVATE_DOCKER_REGISTRY"
	@echo "- then change build parent image as:        INFO_TEST_BUILD_PARENT_IMAGE"
	@echo "- then change build image as:               INFO_BUILD_FROM_IMAGE"
	@echo ""
	@echo "#- check run docker by task"
	@echo "$$ make dockerEnv"
	@echo ""
	@echo "# - first use can pull images"
	@echo "$$ make dockerAllPull"
	@echo ""
	@echo "# - then use to show how to build docker parent image"
	@echo "$$ make dockerRunContainerParentBuild"
	@echo "# - and prune resource at parent image"
	@echo "$$ make dockerPruneContainerParentBuild"
	@echo ""
	@echo "# - test run container use ./${ENV_INFO_TEST_BUILD_DOCKER_FILE}"
	@echo "$$ make dockerTestBuildCheck"
	@echo ""
	@echo "# - run container use ./${ENV_INFO_TEST_BUILD_DOCKER_FILE}"
	@echo "$$ make dockerTestRunLatest"
	@echo ""
	@echo "# - then can run as docker-compose build image and up"
	@echo "$$ make dockerComposeUp"
	@echo "# - then see log as docker-compose"
	@echo "$$ make dockerComposeFollowLogs"
	@echo "# - down as docker-compose will auto remove local image"
	@echo "$$ make dockerComposeDown"
	@echo ""
	@echo "# - prune test container and image"
	@echo "$$ make dockerTestPruneLatest"
	@echo ""
	@echo "- build and tag as use ./${ENV_INFO_BUILD_DOCKER_FILE}"
	@echo "$$ make dockerBeforePush"
	@echo "- prune build"
	@echo "$$ make dockerRmiBuild"
	@echo "- then final push use"
	@echo "$$ make dockerPushBuild"
	@echo ""