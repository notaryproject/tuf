# Metadata examples

## Root
This file would be located on the tuf repository.
```
{
  "signatures": [
    {
      "keyid": "3451d8b97ad5f60f9fbd61f4644cccaa8b40750e3964ec3443ede4a7b63b116f",
      "sig": "2f8ed129c675f30df3f2b58e2dd9bf81ae54676b131c14641c0b2fa5501bf66f891b79579fe71e5f24380e74760e4054aa1082a26ccccc4452b79146ce4dc3d2384d5737033b14326ea98e5d1d8bc711f81fb3a3a4141fd8b4915d97023c853b9684ab9441b6c7b78bc7e8b89ada5148326b453cab74ea86c6b5a4c608f15efe70167d17446a22173a5e28587dd7cab5410a509e653043593a76ddb19126eaa02b48a9525f014d7de980def889195a44fd553a0feda773d31a4feb7bdf46d34dd90c62e51395110b77fdf60f744c8599d38c845f75ee3b8d6ac792cf2918cb6759606e6802f43b024c3be97293abfca69c0e173e6aab67e49cac965955f52dea811ba272a92ce65fec0be4380059bbfd8a0fe0b8bc676e4497070dcfef920053bdaa2621e15913a522e79df0983b4782c0576d811af5e62a67881b3c0f5281595c001d4d3c4278742e73614336885732a666b0ec61f6366683aedc77768bba4cbc476c7e3a47ab29ff64db4a23b23b48c5c7c71bf5fdaf6fd11ce8022096307a"
    },
    {
      "keyid": "3abe61be8817225882a9b34b64ea22a2ce1f8f6466dd06a30db84559ddd66d20",
      "sig": "31fef83a63e88fb0c471880a41f6586a21e0cdd8cca9db10ecf629d9d264e8fc00f9979b69c71edf77caf986a54d07b8d15b241385880946286489881ce16afcc20cee9af3d6e3bf7485b2a077314a25f3ce8d546aae05f52d5960c1bb31340545cdbfa160fa796a046a00d2c1883a99eb0a87c0226ffd9f73bdc1272a3c4ec495a7977791772f84e5c1cfd0b9b7e5ffe8dfdc600a54a8833c7e6b283c5419b12f0570d7121f3b1e4c6771ccde9b70a04588d162b78e6eed5e4fc14e0ae56088fafcc8335b676e89f108be88228a916c8362eda8b1118e3b090d6c2d4c9c3d9f9be2aaa4e45760027152a5a075c7c89046f2a9be8043c7e9a625a9b21b87dea39dcccd3280888f3caaed63ecdb3c98ca28e5bb086d1dcfcf326f868bdff38fd49c400b14c3d95197fb96bbb686e7b17675803ec68067913080ee3e4735b7a16ad551cd563ddd2e02d4a8060f4b4c859bb69f2dc35a613a69253f58618adebaa005dd662ed81e897630415f5a0c6a7237d27c15e21aaf9b95a30859e4526d0890"
    }
  ],
  "signed": {
    "_type": "root",
    "consistent_snapshot": false,
    "expires": "2022-05-19T22:05:27Z",
    "keys": {
      "3451d8b97ad5f60f9fbd61f4644cccaa8b40750e3964ec3443ede4a7b63b116f": {
        "keyid_hash_algorithms": [
          "sha256",
          "sha512"
        ],
        "keytype": "rsa",
        "keyval": {
          "public": "-----BEGIN PUBLIC KEY-----\nMIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA2aUzndeF7l8sit9yaW68\ntwKK7aVehBMzQwZGni+8TlnXPzkjaEO4+z3nNx8R1Sk7p1bznl70CLEZZ6r5MqqG\nOD6PvrjasKCPKN9m4/8nVOUtVxMyzJ5uSL/Xdrzfzm96SSuTf0bDxb872zxLR/Qh\nLk7rw+LZyodf/Is0Nd3uIHJ8CwZEuAlWn/6TrS5xLBm1Wkj1QvRLRJE1s1dzjUEH\neVFHJooJCkRC377XQSDNwPD3vPSL6zJCdMe7XskwpQnROC74qBMp1V+wFLWaR77S\nwizA9/i997zN8WRfcmKeD/Cxuj8KMTntyUxze5jN/q1jv60XddxS1drontx9QeZx\nhBeKgeZbyH/XTXnu+2ojyW6PFgtM6WNevYOBrY6sPndFYj7wOnbMhvTYXghsfoiE\nN5JhP3U18CkZjSchKxqduJYPgetdYSQF/uvvS4pTYN3o5kk5J2qUiDBpdmkhk1xx\n5nzACkcKCyEq86tdda7yhtOUL0mbR9fJUhXS21qZtVI9AgMBAAE=\n-----END PUBLIC KEY-----"
        },
        "scheme": "rsassa-pss-sha256"
      },
      "3abe61be8817225882a9b34b64ea22a2ce1f8f6466dd06a30db84559ddd66d20": {
        "keyid_hash_algorithms": [
          "sha256",
          "sha512"
        ],
        "keytype": "rsa",
        "keyval": {
          "public": "-----BEGIN PUBLIC KEY-----\nMIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA3Zr6T+fEPHYkeJazT97k\nHAazthHsszlDg0CezgmBEwSqA96vQBjb9+ryQW35o4z9AJDuiaEwnk6G/94C+GG/\nDwi6wyHbAcjYbm6TbHo1DzvUAfl6QeTVQWm70PE1M00jKbCo5TcZmmiMD3vOMhQP\nA6R0sQ2/ntKIXpbANdEFfFNA/I92aRgT9rpQKF+W3wUJ+/bqf1wTauUFg01evyeL\n8fbIa1ci2ez4QCebv45VFIGbe066+TQzeIoMs9Qlv1FBFZrR32BKDxJCqH8Jqz8W\np00MDCZ8xjVC1qSU71Oe+jbct4/qEQNGYJpGwVuZTbv9m5XEQUFNfJWAB5K1npHr\nYNCVNCSy0BrlQMKi0Wn/2UbRLg1vQXuJSz5FyXk3u9oUbyRyyjxmrFQfsHaTKpRu\nomVA98ev1tgKqhBA5hdqZIuWQxq32t++5kt8wIzos9QOz3ATsXea3K/SF3OY9ndU\nRlxw4OzfAVauWS/P1BZAtkJ+nqy2Wft61wJwo/gsCTgbAgMBAAE=\n-----END PUBLIC KEY-----"
        },
        "scheme": "rsassa-pss-sha256"
      },
      "527503249a89815eebd644de6d6f7139cae231db9107b892e5cc466bb69b0296": {
        "keyid_hash_algorithms": [
          "sha256",
          "sha512"
        ],
        "keytype": "ed25519",
        "keyval": {
          "public": "a923ee2d8fbc81eb76c1fa8142e80898b1d5d84aa2a24c8ae077964a9f208d54"
        },
        "scheme": "ed25519"
      },
      "cc608e9c20fbbba9e4b3fd39d83062721d1a9ed2960f362816e83f8fc4f15273": {
        "keyid_hash_algorithms": [
          "sha256",
          "sha512"
        ],
        "keytype": "ed25519",
        "keyval": {
          "public": "96e04b85fea35c23d25bfc0184d89cd9dc00903577558f49241f330f22a86de4"
        },
        "scheme": "ed25519"
      },
      "f387737ecee1851bcccf788b71826e43075976d08b0e179426174ef5165abc81": {
        "keyid_hash_algorithms": [
          "sha256",
          "sha512"
        ],
        "keytype": "ed25519",
        "keyval": {
          "public": "9c967b28d2dcd599bf63d37f5c7a8d84d2d67400f8d511f9baf9c35d5881e50f"
        },
        "scheme": "ed25519"
      }
    },
    "roles": {
      "root": {
        "keyids": [
          "3451d8b97ad5f60f9fbd61f4644cccaa8b40750e3964ec3443ede4a7b63b116f",
          "3abe61be8817225882a9b34b64ea22a2ce1f8f6466dd06a30db84559ddd66d20"
        ],
        "threshold": 2
      },
      "snapshot": {
        "keyids": [
          "f387737ecee1851bcccf788b71826e43075976d08b0e179426174ef5165abc81"
        ],
        "threshold": 1
      },
      "targets": {
        "keyids": [
          "527503249a89815eebd644de6d6f7139cae231db9107b892e5cc466bb69b0296"
        ],
        "threshold": 1
      },
      "timestamp": {
        "keyids": [
          "cc608e9c20fbbba9e4b3fd39d83062721d1a9ed2960f362816e83f8fc4f15273"
        ],
        "threshold": 1
      }
    },
    "spec_version": "1.0.0",
    "version": 1
  }
}
```

## top-level targets
This file would be located on the tuf repository.
```
{
  "signatures": [
    {
      "keyid": "527503249a89815eebd644de6d6f7139cae231db9107b892e5cc466bb69b0296",
      "sig": "3a0a955439e5c66797d702c2e14700d2b96038db951178e881b5f7fbebe325dccd330670261c4c46a4759be4b40998f2773644d7d34e70c144b2adec227cf700"
    }
  ],
  "signed": {
    "_type": "targets",
    "delegations": {
      "keys": {
        "ef5df3e0136f83569a4310192192328d7bb0137aaafe2d9ab2c706d8f6e84997": {
          "keyid_hash_algorithms": [
            "sha256",
            "sha512"
          ],
          "keytype": "ed25519",
          "keyval": {
            "public": "3ad7ee4e5e88e007f46dd15fac0ef7d7971e47ed347f511607e445f8fdadd92b"
          },
          "scheme": "ed25519"
        }
      },
      "roles": [
        {
          "keyids": [
            "ef5df3e0136f83569a4310192192328d7bb0137aaafe2d9ab2c706d8f6e84997"
          ],
          "name": "wabbit_networks",
          "paths": [
            "wabbit_networks/*"
          ],
          "terminating": false,
          "threshold": 1
        },
      ]
    },
    "expires": "2021-08-18T23:44:17Z",
    "spec_version": "1.0.0",
    "targets": {},
    "version": 3
  }
}
```

## wabbit networks targets
This file would be located alongside the target file in the wabbit networks repository
```
{
  "signatures": [
    {
      "keyid": "ef5df3e0136f83569a4310192192328d7bb0137aaafe2d9ab2c706d8f6e84997",
      "sig": "a80fb7d49269d8369e3d6f36b3262f1fb7bfe9d50bc168a42ea478b117e8923198a13d6e14fb7f9b58c265e3c138b0fbd22ca0c7149f2841a4e85342f3a9340f"
    }
  ],
  "signed": {
    "_type": "targets",
    "delegations": {
      "keys": {},
      "roles": []
    },
    "expires": "2021-08-18T23:44:18Z",
    "spec_version": "1.0.0",
    "targets": {
      "sha256:b2eb39acda1496e1727e365ad56877689d6aa8bbf59974414a5369c79f84632e": {
        "custom": {"mediatype": "application/vnd.oci.image.manifest.v1+json"},
        "hashes": {
          "sha256": "b2eb39acda1496e1727e365ad56877689d6aa8bbf59974414a5369c79f84632e"
        },
        "length": 30
      }
    },
    "version": 1
  }
}
```
