# How to move an artifact signed by tuf-notary

<img src="images/Notary-v2_movement.jpg">

The following is the process for moving an artifact that has been signed with TUF.
We assume that both the source and destination registries have existing TUF metadata.

1. Copy the artifact and any delegated targets metadata to the target repository on the destination registry.
1. Update the top-level targets metadata on the destination repository to indicate that the original uploader's signature should be trusted
1. (Optional) Add a signature to the artifact from the verifier or other entity and update the top-level targets metadata accordingly
1. (Automated) Update the snapshot and timestamp metadata on the destination registry.


All of these steps can be combined into a single CLI/API call with flags for adding signatures and indicating which keys should be trusted.
