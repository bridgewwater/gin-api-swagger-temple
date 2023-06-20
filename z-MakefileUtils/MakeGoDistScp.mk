# this file must use as base Makefile job must has variate
#
# must as some include MakeDistTools.mk MakeGoDist.mk
# include z-MakefileUtils/MakeGoDistScp.mk

ENV_SERVER_TEST_SSH_ALIAS=aliyun-ecs-test
ENV_SERVER_REPO_SSH_ALIAS=aliyun-ecs-release
ENV_SERVER_TEST_FOLDER=/home/work/Document/
ENV_SERVER_REPO_FOLDER=/home/ubuntu/$(ROOT_NAME)

distScpTestOSTar: distTestOSTar
	$(info => can send file:)
	$(info scp ${ENV_PATH_INFO_ROOT_DIST_OS}/${ENV_INFO_DIST_BIN_NAME}-${ENV_INFO_DIST_ENV_TEST_NAME}-${ENV_INFO_DIST_GO_OS}-${ENV_INFO_DIST_GO_ARCH}-${ENV_INFO_DIST_VERSION}${ENV_DIST_MARK}.tar.gz ${ENV_SERVER_TEST_SSH_ALIAS}:${ENV_SERVER_TEST_FOLDER})
	@echo "=> must check below config of set for release OS Scp"

distScpReleaseOSTar: distReleaseOSTar
	$(info => can send file:)
	$(info scp ${ENV_PATH_INFO_ROOT_DIST_OS}/${ENV_INFO_DIST_BIN_NAME}-${ENV_INFO_DIST_ENV_RELEASE_NAME}-${ENV_INFO_DIST_GO_OS}-${ENV_INFO_DIST_GO_ARCH}-${ENV_INFO_DIST_VERSION}${ENV_DIST_MARK}.tar.gz ${ENV_SERVER_REPO_SSH_ALIAS}:${ENV_SERVER_REPO_FOLDER})
	@echo "=> must check below config of set for release OS Scp"