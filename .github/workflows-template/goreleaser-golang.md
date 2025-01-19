## config-file

`goreleaser-golang.yml`

## usage in pipeline

```yaml
jobs:
  version:
    name: version
    uses: ./.github/workflows/version.yml

  go-goreleaser-dry-run: # use .golangci.yaml to check build
    name: go-goreleaser-dry-run
    needs:
      - version
    uses: ./.github/workflows/goreleaser-golang.yml
    if: ${{ github.ref_type != 'tag' }}
    secrets: inherit
    with:
      version_name: ${{ needs.version.outputs.cc_date }}
      goreleaser-build-timeout-minutes: 30
      dry_run: true # must set dry_run=true
      # upload_artifact_name: go-release

  ### tag to release start
  go-goreleaser-by-tag:
    name: go-goreleaser-by-tag
    needs:
      - version
    uses: ./.github/workflows/goreleaser-golang.yml
    if: startsWith(github.ref, 'refs/tags/')
    secrets: inherit
    with:
      version_name: ${{ needs.version.outputs.tag_name }}
      goreleaser-build-timeout-minutes: 30
      upload_artifact_name: go-release # use https://github.com/actions/upload-artifact to upload artifact
      # dry_run: false

  deploy-tag:
    needs:
      - version
      - go-goreleaser-by-tag
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