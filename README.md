# go-ITS

## Description
Implementation of ETSI CDD (common data dictionary) and CAM protocol, custom CVM (Control Vehicle Message) protocol and encoding/decoding as a Golang library. 
This project was done as university work.

For more information, refer to documentation https://github.com/pprizbul/go-ITS/wiki/Technical-documentation

## Copyright
The ASN.1 files are taken from official ETSI documents. Please refer to them for more information.

- CAM: ETSI TS 103 900 - V2.1.1 (2023-11)
- CDD: ETSI TS 102 894-2 - V2.2.1 (2023-10)

ASN.1 files are available on https://forge.etsi.org/rep/ITS/asn1 in their respective folders.

Copyright notice:
```
Copyright 2019 ETSI

Redistribution and use in source and binary forms, with or without 
modification, are permitted provided that the following conditions are met:
1. Redistributions of source code must retain the above copyright notice, 
   this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice, 
   this list of conditions and the following disclaimer in the documentation 
   and/or other materials provided with the distribution.
3. Neither the name of the copyright holder nor the names of its contributors 
   may be used to endorse or promote products derived from this software without 
   specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND 
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED 
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. 
IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, 
INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, 
BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, 
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF 
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE 
OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED 
OF THE POSSIBILITY OF SUCH DAMAGE.

```


## Usage

Import the module into your Golang project with `go get github.com/pprizbul/go-ITS`

In code, you can import whatever package you need with `"github.com/pprizbul/go-ITS/<pkgName>"`

## Examples

Decoding message
```
n, clientAddr, err := conn.ReadFromUDP(buffer)
if err != nil {
    fmt.Println(err)
}

message, err := codec.Decode(buffer[:n])
if err != nil {
    fmt.Println(err)
}

switch message.(type){
case CAM.CAM:
    //Handle CAM
case ...
}
```

Encoding message
```
camMessage := CAM.CamPayload{} //Header is added dynamically

data, err := codec.Encode(camMessage)
if err != nil {
    fmt.Println(err)
}

_, err := conn.WriteToUDP(data, clientAddr)
if err != nil {
    fmt.Println(err)
}
```