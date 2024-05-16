# Foobar

- Tutorials: [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)
- Forum: [discuss.hashicorp.com](https://discuss.hashicorp.com/c/terraform-providers/tf-google/)
- Documentation: https://www.terraform.io/docs/providers/google/index.html
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

The Terraform Google provider is a plugin that allows [Terraform](https://www.terraform.io) to manage resources on Google Cloud Platform. This provider is maintained by the [Terraform team at Google](https://cloudplatform.googleblog.com/2017/03/partnering-on-open-source-Google-and-HashiCorp-engineers-on-managing-GCP-infrastructure.html) and the Terraform team at [HashiCorp](https://www.hashicorp.com/)

This is the `google` provider, containing generally available features. To use preview features or features at a beta [launch stage](https://cloud.google.com/products#product-launch-stages), you may use the [`google-beta` provider](https://github.com/hashicorp/terraform-provider-google-beta). Refer to the [provider versions](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_versions) documentation for more information about how to use `google-beta`.

## Quick Starts

- [Getting Started with the Google Provider](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/getting_started)
- [Provider Documentation](https://registry.terraform.io/providers/hashicorp/google/latest/docs)

## Provider Usage

Please see [instructions](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference) on how to configure the Google Provider.

### Upgrading the provider

The Google provider doesn't upgrade automatically once you've started using it. After a new release you can run

```bash
terraform init -upgrade
```

to upgrade to the latest stable version of the Google provider. See the [Terraform website](https://www.terraform.io/docs/configuration/providers.html#provider-versions)
for more information on provider upgrades, and how to set version constraints on your provider.

## Developing the provider

This repository is generated by [magic-modules](https://github.com/GoogleCloudPlatform/magic-modules).
If you wish to work on the provider, you'll need to make changes in [magic-modules](https://github.com/GoogleCloudPlatform/magic-modules). Any changes made directly to this repository will likely be overwritten.

For guidance on how to contribute, see our [contribution documentation](https://googlecloudplatform.github.io/magic-modules/).
If you have other development questions we don't cover, please [file an issue](https://github.com/hashicorp/terraform-provider-google/issues/new/choose)!















