package tufnotary

import (
	"context"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

func UploadTUFMetadata(registry string, repository string, name string, contents []byte, reference string) (ocispec.Descriptor, error) {
	ref := registry + "/" + repository + ":" + name
	fileName := repository + "/repository/" + name + ".json"

	mediaType := "application/vnd.cncf.notary.tuf+json"

	ctx := context.Background()

	// TODO: add reference once it's supported in oras-go: https://github.com/oras-project/oras-go/pull/35

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

func DownloadTUFMetadata(registry string, repository string, name string) error {
	ref := registry + "/" + repository + ":" + name

	mediaType := "application/vnd.cncf.notary.tuf+json"
	ctx := context.Background()

	reg, err := content.NewRegistry(content.RegistryOptions{PlainHTTP: true})
	if err != nil {
		return err
	}

	fileStore := content.NewFile("")
	defer fileStore.Close()
	allowedMediaTypes := []string{mediaType}
	_, err = oras.Copy(ctx, reg, ref, fileStore, "", oras.WithAllowedMediaTypes(allowedMediaTypes))
	return err
}
