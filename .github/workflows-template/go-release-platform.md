## config-file

`go-release-platform.yml`

## usage in pipeline

```yaml
jobs:
  version:
    name: version
    uses: ./.github/workflows/version.yml

  go-build-check-main:
    name: go-build-check-main
    needs:
      - version
    if: ${{ ( github.event_name == 'push' && github.ref == 'refs/heads/main' ) || github.base_ref == 'main' }}
    uses: ./.github/workflows/go-release-platform.yml
    secrets: inherit
    with:
      version_name: latest
      go_build_id: ${{ needs.version.outputs.short_sha }}

  ### tag to release start
  go-release-platform:
    name: go-release-platform
    needs:
     - version
     # - golang-ci
    if: startsWith(github.ref, 'refs/tags/')
    uses: ./.github/workflows/go-release-platform.yml
    secrets: inherit
    with:
      version_name: ${{ needs.version.outputs.tag_name }}
      upload_artifact_name: go-release

  deploy-tag:
    needs:
      - version
      - go-release-platform
    name: deploy-tag
    uses: ./.github/workflows/deploy-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
    secrets: inherit
    with:
      prerelease: true
      tag_name: ${{ needs.version.outputs.tag_name }}
      tag_changes: ${{ needs.version.outputs.cc_changes }}
      download_artifact_name: go-release
  ### tag to release end
```