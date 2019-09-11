# infrastructure

## chain

a chain contains some blocks, has a config
to define the characteristics, such as `uid`,
`name`...

    -- chain
        -- block #0
        -- block #1
        -- ...   #N

### block

a block contains some logs that can never be changed

    -- block #N
        -- log #0
        -- log #1
        -- ... #N

### Crypto

- secret key

  256 bits

- public key

  pk = [sk]\*G

  format = hex

- elliptic curve

  P256

- signature

  ecdsa

### data struct

1. Msg

2. Tx

```js
{
    Signer:"", // pk
    Signature:"",
    Hash:"",
    Body:{
        Type:"",
        Action:"",
        TimeStamp:"",
        To:"", //path
        Encrypted: bool, //是否加密
        Content:{

        }
    }
}
```
