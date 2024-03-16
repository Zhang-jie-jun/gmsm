
GM SM2/3/4 library based on Golang

基于Go语言的国密SM2/SM3/SM4加密算法库

GMSM包含以下主要功能

    SM2: 国密椭圆曲线算法库
        . 支持Generate Key, Sign, Verify基础操作
        . 支持加密和不加密的pem文件格式(加密方法参见RFC5958, 具体实现参加代码)
        . 支持证书的生成，证书的读写(接口兼容rsa和ecdsa的证书)
        . 支持证书链的操作(接口兼容rsa和ecdsa)
        . 支持crypto.Signer接口

    SM3: 国密hash算法库
       . 支持基础的sm3Sum操作
       . 支持hash.Hash接口

    SM4: 国密分组密码算法库
        . 支持Generate Key, Encrypt, Decrypt基础操作
        . 提供Cipher.Block接口
        . 支持加密和不加密的pem文件格式(加密方法为pem block加密, 具体函数为x509.EncryptPEMBlock)
    AES: AES对称密码算法库
         . 支持Generate Key, Encrypt, Decrypt基础操作
         . 提供16位，32位加解密接口
         . 支持cbc/cbf等加密模式
         
性能测试


# 国密GM/T Go API使用说明

GO SDK=1.14.1

## 国密gmsm包安装

```bash
go get -u github.com/Zhang-jie-jun/gmsm
```

## SM3密码杂凑算法 - SM3 cryptographic hash algorithm

遵循的SM3标准号为： GM/T 0004-2012

导入包
```Go
import github.com/Zhang-jie-jun/gmsm/sm3
```

### 代码示例

```Go
    data := "test"
    h := sm3.New()
    h.Write([]byte(data))
    sum := h.Sum(nil)
    fmt.Printf("digest value is: %x\n",sum)
```
### 方法列表

####  New 
创建哈希计算实例
```Go
func New() hash.Hash 
```

#### Sum 
返回SM3哈希算法摘要值
```Go
func Sum() []byte 
```

## SM4分组密码算法 - SM4 block cipher algorithm

遵循的SM4标准号为:  GM/T 0002-2012

导入包
```Go
import github.com/Zhang-jie-jun/gmsm/sm4
```

### 代码示例

```Go
    import  "crypto/cipher"
    import  "github.com/Zhang-jie-jun/gmsm/sm4"

    func main(){
        // 128比特密钥
        key := []byte("1234567890abcdef")
        // 128比特iv
        iv := make([]byte, sm4.BlockSize)
        data := []byte("Tongji Fintech Research Institute")
        ciphertxt,err := sm4Encrypt(key,iv, data)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Printf("加密结果: %x\n", ciphertxt)
    }

    func sm4Encrypt(key, iv, plainText []byte) ([]byte, error) {
        block, err := sm4.NewCipher(key)
        if err != nil {
            return nil, err
        }
        blockSize := block.BlockSize()
        origData := pkcs5Padding(plainText, blockSize)
        blockMode := cipher.NewCBCEncrypter(block, iv)
        cryted := make([]byte, len(origData))
        blockMode.CryptBlocks(cryted, origData)
        return cryted, nil
    }

    func sm4Decrypt(key, iv, cipherText []byte) ([]byte, error) {
        block, err := sm4.NewCipher(key)
    	if err != nil {
        	return nil, err
    	}
    	blockMode := cipher.NewCBCDecrypter(block, iv)
    	origData := make([]byte, len(cipherText))
    	blockMode.CryptBlocks(origData, cipherText)
    	origData = pkcs5UnPadding(origData)
    	return origData, nil
    }
    // pkcs5填充
    func pkcs5Padding(src []byte, blockSize int) []byte {
        padding := blockSize - len(src)%blockSize
    	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    	return append(src, padtext...)
    }

    func pkcs5UnPadding(src []byte) []byte {
        length := len(src)
        if(length==0){
            return nil
        }
    	unpadding := int(src[length-1])
    	return src[:(length - unpadding)]
    }
```

### 方法列表

#### NewCipher
创建SM4密码分组算法模型，参数key长度只支持128比特。
```Go
func NewCipher(key []byte) (cipher.Block, error)
```

## SM2椭圆曲线公钥密码算法 - Public key cryptographic algorithm SM2 based on elliptic curves

遵循的SM2标准号为： GM/T 0003.1-2012、GM/T 0003.2-2012、GM/T 0003.3-2012、GM/T 0003.4-2012、GM/T 0003.5-2012、GM/T 0009-2012、GM/T 0010-2012

导入包
```Go
import github.com/Zhang-jie-jun/gmsm/sm2
```

### 代码示例

```Go
    priv, err := sm2.GenerateKey() // 生成密钥对
    if err != nil {
    	log.Fatal(err)
    }
    msg := []byte("Tongji Fintech Research Institute")
    pub := &priv.PublicKey
    ciphertxt, err := pub.Encrypt(msg)
    if err != nil {
    	log.Fatal(err)
    }
    fmt.Printf("加密结果:%x\n",ciphertxt)
    plaintxt,err :=  priv.Decrypt(ciphertxt)
    if err != nil {
    	log.Fatal(err)
    }
    if !bytes.Equal(msg,plaintxt){
        log.Fatal("原文不匹配")
    }

    r,s,err := sm2.Sign(priv, msg)
    if err != nil {
    	log.Fatal(err)
    }
    isok := sm2.Verify(pub,msg,r,s)
    fmt.Printf("Verified: %v\n", isok)
```

### 方法列表

#### GenerateKey
生成随机秘钥。
```Go
func GenerateKey() (*PrivateKey, error) 
```

#### Sign
用私钥签名数据，成功返回以两个大数表示的签名结果，否则返回错误。
```Go
func Sign(priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
```

#### Verify
用公钥验证数据签名, 验证成功返回True，否则返回False。
```Go
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool 
```

#### Encrypt
用公钥加密数据,成功返回密文错误，否则返回错误。
```Go
func Encrypt(pub *PublicKey, data []byte) ([]byte, error) 
```

#### Decrypt
用私钥解密数据，成功返回原始明文数据，否则返回错误。
```Go
func Decrypt(priv *PrivateKey, data []byte) ([]byte, error)
```