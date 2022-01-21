package tufnotary

import (
	"context"
	"encoding/json"

	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	artifactspec "github.com/oras-project/artifacts-spec/specs-go/v1"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

func UploadTUFMetadataWithReference(registry string, repository string, name string, contents []byte, reference ocispec.Descriptor) (ocispec.Descriptor, error) {
	ref := registry + "/" + repository + ":" + name
	fileName := repository + "/staged/" + name + ".json"

	mediaType := "application/vnd.cncf.notary.tuf+json"

	ctx := context.Background()

	memoryStore := content.NewMemory()
	desc, err := memoryStore.Add(fileName, mediaType, contents)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	config, configDesc, err := content.GenerateConfig(nil)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	var descs []artifactspec.Descriptor
	descs = append(descs, artifactspec.Descriptor{
		MediaType:   desc.MediaType,
		Digest:      desc.Digest,
		Size:        desc.Size,
		URLs:        desc.URLs,
		Annotations: desc.Annotations,
	})

	manifest := artifactspec.Manifest{
		MediaType:    "application/vnd.cncf.oras.artifact.manifest.v1+json",
		ArtifactType: mediaType,
		Blobs:        descs,
		Annotations:  nil,
		Subject: artifactspec.Descriptor{
			MediaType:   reference.MediaType,
			Digest:      reference.Digest,
			Size:        reference.Size,
			URLs:        reference.URLs,
			Annotations: reference.Annotations,
		},
	}

	manifestBytes, err := json.Marshal(manifest)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	manifestDescriptor := ocispec.Descriptor{
		MediaType: ocispec.MediaTypeImageManifest,
		Digest:    digest.FromBytes(manifestBytes),
		Size:      int64(len(manifestBytes)),
	}

	memoryStore.Set(configDesc, config)
	err = memoryStore.StoreManifest(ref, manifestDescriptor, manifestBytes)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	reg, err := content.NewRegistry(content.RegistryOptions{PlainHTTP: true})
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	desc, err = oras.Copy(ctx, memoryStore, ref, reg, "")
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	return desc, nil
}

func UploadTUFMetadata(registry string, repository string, name string, contents []byte) (ocispec.Descriptor, error) {
	ref := registry + "/" + repository + ":" + name
	fileName := repository + "/staged/" + name + ".json"

	mediaType := "application/vnd.cncf.notary.tuf+json"

	ctx := context.Background()

	memoryStore := content.NewMemory()
	desc, err := memoryStore.Add(fileName, mediaType, contents)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	manifest, manifestDesc, config, configDesc, err := content.GenerateManifestAndConfig(nil, nil, desc)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	memoryStore.Set(configDesc, config)
	err = memoryStore.StoreManifest(ref, manifestDesc, manifest)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	reg, err := content.NewRegistry(content.RegistryOptions{PlainHTTP: true})
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	desc, err = oras.Copy(ctx, memoryStore, ref, reg, "")
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	return desc, nil
}
