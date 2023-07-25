# TUF

This repository is **not** in active maintenance. Please see the Notary Project [README](https://github.com/notaryproject/.github/blob/main/README.md) file to learn about overall Notary Project.

TUF is a project to implement the full TUF specification in a registry native way. This may
require upstream TUF spec changes or extensions, as there are some differences between the
registry model and common usage to other TUF use cases. This project will use existing
registry extensions where available but may need its own document types in addition.

The initial version of notary TUF-based implementation ran as an additional service on a registry, so was not
available everywhere and did not create native registry artifacts. In turn this meant
that moving signatures between registries was not supported. The notary TUF-based implementation also made some
changes to the TUF security model, like defaulting to TOFU, which in retrospect were
not a good model in a world of ephemeral cloud native hosts. It did not get widespread
adoption due to these reasons and others. This project aims to build a version suitable
for widespread adoption that resolves these issues.
