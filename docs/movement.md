# How to move an artifact signed by tuf-notary

There are a few scenarios for moving artifacts covered in this document:
* Mirroring the artifact and TUF metadata
* Copying the image without the original signature (or from a registry that doesn't use TUF)
* Copying the image and the original signature
* Copying the image, original signature, and adding an additional verifier's signature

In all but the first scenario we assume that the destination registry already has a top-level TUF repository set up.

All of these steps for each scenario can be combined into a single CLI/API call with flags indicating whether the original signature is copied, and if a new signature is to be added.

## Mirroring the artifact and TUF metadata
Copy the artifact, signature, and all top-level TUF metadata to the new registry. The mirror will want to ensure that copies are frequent enough that the timestamp and snapshot remain valid.

## Copying the image without the original signature

1. Copy the artifact to the target repository on the destination registry.
1. Add a TUF signature to the artifact from the verifier or other entity and update the top-level targets metadata accordingly
1. (Automated) Update the snapshot and timestamp metadata on the destination registry.

## Copying the image and the original signature

1. Copy the artifact and any delegated targets metadata to the target repository on the destination registry.
1. Update the top-level targets metadata on the destination repository to indicate that the original uploader's signature should be trusted
1. (Automated) Update the snapshot and timestamp metadata on the destination registry.

## Copying the image, original signature, and adding an additional verifier's signature

<img src="images/Notary-v2_movement.jpg">

1. Copy the artifact and any delegated targets metadata to the target repository on the destination registry.
1. Update the top-level targets metadata on the destination repository to indicate that the original uploader's signature should be trusted
1. Add a signature to the artifact from the verifier or other entity and update the top-level targets metadata accordingly
1. (Automated) Update the snapshot and timestamp metadata on the destination registry.
