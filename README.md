# terraform-provider-keyring

![Travis (.org)](https://img.shields.io/travis/rremer/terraform-provider-keyring?label=linux%2Fosx) ![AppVeyor](https://img.shields.io/appveyor/ci/rremer/terraform-provider-keyring?label=windows)

A terraform provider for leveraging local keyrings on all operating systems.

## Installation

```sh
go get
go build
```

OS-specific installation examples below, but please reference the official [Terraform Plugin Discovery] documentation.

Linux / OSX:

```
cp terraform-provider-keyring ~/.terraform.d/plugins/
```

Windows:

```
copy terraform-provider-keyring %APPDATA%\terraform.d\plugins\
```

## Usage

### Hashicorp Vault

For an organization with many terraform projects sourcing similar secrets, you may want a 'bootstrap' project which developers run locally once, and then source that secret as an output elsewhere. Here's an example sourcing a secret from [Hashicorp Vault]. This would make one remote call to the Vault service, then cache the secret in the developer's local keyring. Other terraform projects would not need to make remote calls to Vault and instead reference the ```${data.keyring_secret.example.secret}```.

```hcl
data "vault_generic_secret" "example" {
  path = "secret/example"
}

resource "keyring_secret" "example" {
        name   = "example"
        secret = "${data.vault_generic_secret.example.data["auth_token"]}"
}

data "keyring_secret" "example" {
        name = "${keyring_secret.example.name}"
}
```

### Replacing file() and hard-coded secrets

Anywhere you reference terraform's ```file()``` method to fetch secrets like RSA private keys could be replaced with this provider.

Before:
```hcl
resource "null_resource" "example_sshable_instance" {
  connection {
   user        = "ubuntu"
   host        = "example.com"
   private_key = "${file("~/.ssh/id_rsa")}"
 }
}
```

After:
```hcl
data "keyring_secret" "ssh" {
        name = "example"
}

resource "null_resource" "example_sshable_instance" {
  connection {
   user        = "ubuntu"
   host        = "example.com"
   private_key = "${data.keyring_secret.ssh.secret}"
 }
}
```

Inserting secrets into your keyring is OS/distribution specific. Here's some common ones:

#### GNOME Keyring

Installation of secret-tool (or equivelant GUIs like Seahorse) varies, but the majority of Linux distrobutions implement [GNOME Keyring] for secrets storage. By default, most display managers will unlock a default login keyring for use.

Ubuntu:
```sh
sudo apt-get install -y libsecret-tools
secret-tool store --label='example' id example
```
... when prompted, paste in your private key.

#### OSX Keychain

OSX Keychain, leveraging the Login keychain (currently not configurable). Not including screenshots of that for brevity.

```sh
security add-generic-password -U -s terraform -a example -w <YOUR_PRIVATE_KEY>
```
... note that ```-s terraform``` defines the service label, which is a constant in the terraform provider for consistency with other OS implementations which don't support this construct. ```-a``` can be whatever you want and must match ``data.keyring_secret.example.name```.

#### Windows Credential Manager

Since XP, Windows has shipped with a CLI and GUI for Windows Credential Manager.

```sh
cmdkey /generic terraform /user example /pass <YOUR_PRIVATE_KEY>
```
... note that ```/generic terraform``` defines the domain, which is a constant in the terraform provider for consistency with other OS implementations which don't support this construct. ```/user``` can be whatever you want and must match ```data.keyring_secret.example.name```.


[Terraform Plugin Discovery]:https://www.terraform.io/docs/extend/how-terraform-works.html#discovery
[Hashicorp Vault]:https://www.terraform.io/docs/providers/vault/index.html
[GNOME Keyring]:https://wiki.gnome.org/Projects/GnomeKeyring
