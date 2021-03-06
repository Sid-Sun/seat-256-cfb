## seat-256-cfb

### seat-256-cfb is a CLI program which implements the [SeaTurtle Block Cipher](https://github.com/sid-sun/seaturtle) in CFB (cipher feedback) mode with 256-Bit length keys using SHA3-256 with files.

## Usage:

```
To encrypt: seat-256-cfb (--encrypt / -e) <input file> <passphrase file> <output file (optional)>

To decrypt: seat-256-cfb (--decrypt / -d) <encrypted input> <passphrase file> <output file (optional)>

To get version number: seat-256-cfb (--version / -v)

To get help: seat-256-cfb (--help / -h)
```

## Installation:

### Compressed Compiled Binaries and Debian package: 

> [amd64](https://github.com/Sid-Sun/seat-256-cfb/releases/latest/download/binaries_and_debian_package.tar)


### Use YAPPA (Yet Another PPA) :

```bash
curl -s --compressed "https://sid-sun.github.io/yappa/KEY.gpg" | sudo apt-key add -
curl -s --compressed "https://sid-sun.github.io/yappa/yappa.list" | sudo tee /etc/apt/sources.list.d/yappa.list
sudo apt update
sudo apt install seat-256-cfb
```

## Versioning system:

The Versioning system follows a Trickle-down approach (i.e. the version part after the updated part is to be set to 0s)

The version number consists of three parts:

1. Major 

    Major version is to be updated when using the SAME input and key, the output generated differs (ex: bug fixes)

2. Minor

    Minor version is to be updated when features are added or change are made to the core system which don't affect how it behaves with the same inputs (ex: performance improvements)

3. Infant

    Infant version is to be changed when the change doesn't affect the core system (ex: UX updates)


Updating on major and minor version changes is highly recommended.

### Cheers!