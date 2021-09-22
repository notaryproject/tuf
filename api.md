# Proposed API

## High level API

`tuf-notary init`\
creates initial root and top-level targets metadata

`tuf-notary delegate <from> <to> <keys>`\
Creates a new delegation in <from>

`tuf-notary sign <artifact> <role>`\
Creates delegated tuf targets metadata for <artifact> using local keys from <role>

`tuf-notary upload <role> <registry> <repository> <tuf-repo>`\
Uploads the delegated targets metadata for <role> to registry/repository, and
will automatically update top-level targets, root, snapshot, timestamp to registry/tuf-repo

`tuf-notary verify <artifact> <tuf-repo>`\
Downloads the delegated targets metadata that references artifact,
and the top-level tuf metadata from the tuf-repo,
then perform tuf verification.

## Low level API

`tuf-notary update-targets`\
will just update top-level targets metadata

`tuf-notary upload <role> <destination>`\
will just upload the given role metadata

`tuf-notary download <role> <location>`\
will download tuf metadata for role from location
