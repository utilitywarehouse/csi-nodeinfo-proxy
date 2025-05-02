SHELL := /bin/bash

release:
	@sd "newTag: main" "newTag: $(VERSION)" deploy/trident/kustomization.yaml
	@git add deploy/trident/kustomization.yaml
	@git commit -m "Release $(VERSION)"
	@git tag -m "Release $(VERSION)" -a $(VERSION)
	@sd "newTag: $(VERSION)" "newTag: main" deploy/trident/kustomization.yaml
	@git add deploy/trident/kustomization.yaml
	@git commit -m "Clean up release $(VERSION)"
