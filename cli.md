# Proposed CLI

## High level CLI

`tuf-notary init <registry> [--repo=<repository>]`\
creates initial root and top-level targets metadata uploaded to `<registry>/<repository>`.
If not specified, the repository name will be `tuf-repo`

`tuf-notary delegate <from> <to> <keys>`\
Creates a new delegation in `<from>` with the rolename  `<to>` and public keys `<keys>`

`tuf-notary sign <artifact> <role> [--local]`\
Creates delegated tuf targets metadata for `<artifact>` using local keys from `<role>`.
This metadata will be uploaded alongside the artifact, and will automatically update snapshot and timestamp.
If local is set, the metadata will be generated but not uploaded.

`tuf-notary upload <role> <registry> <repository> <tuf-repo>`\
Uploads the delegated targets metadata for `<role>` to `<registry>/<repository>` (that may be signed locally on a disconnected machine), and
will automatically update top-level targets, snapshot, timestamp to `<registry>/<tuf-repo>`

`tuf-notary verify <artifact> <tuf-repo>`\
Downloads the delegated targets metadata that references `<artifact>`,
and the top-level tuf metadata from the `<tuf-repo>`,
then performs tuf verification.

## Low level CLI

`tuf-notary update-targets`\
will just update top-level targets metadata

`tuf-notary upload <role> <destination>`\
will just upload the given role metadata

`tuf-notary download <role> <location>`\
will download tuf metadata for `<role>` from `<location>`
