# terraform-provider-keyring

[![Build Status](https://travis-ci.org/rremer/terraform-provider-keyring.svg?branch=master)](https://travis-ci.org/rremer/terraform-provider-keyring)

A terraform provider for leveraging GPG or system keyrings.

## Building

```sh
go get
go build
```


## Adding secrets

### Ubuntu

```sh
sudo apt-get install -y libsecret-tools
secret-tool store --label='terraformtest' id some-uuid
```

...enter your secret and then open up Seahors ("Passwords and Keys") and search the Login keychain for 'terraform'.
