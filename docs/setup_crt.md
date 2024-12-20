# setup.crt

This file is the input for the first call to `FairPlaySAPExchange`. Its format
is proprietary and can be decoded using the schema provided in this document.
The use of the `.crt` extension is misleading, as the file is not PEM-encoded.

According to the HTTP response headers of the Akamai CDN, that file hasn't been
modified since `Tue, 30 Aug 2016 23:08:34 GMT`. The file embeds two DER-encoded
certificates. The first certificate is `Apple System Integration Certification
Authority` and the second certificate is `DRM Technologies A01`. The former
being an intermediate certificate signed by `Apple Root CA` and the latter being
a leaf certificate with `Digital Signature` and `Key Encipherment` key usage.

The key usage of the leaf certificate suggests that the key used by
`FairPlaySAPSign` is exchanged with Apple's servers using some standard process.

## .crt Schema

- `sign_sap_setup_cert.ksy`

```yaml
meta:
  id: sign_sap_setup_cert
  file-extension: crt
  endian: be
seq:
  - id: header
    type: header
  - id: certs
    type: cert
    repeat: expr
    repeat-expr: header.certs_len
types:
  header:
    seq:
      - id: version
        contents:
          - 0x01
        doc: yet to be verified
      - id: certs_len
        type: u1
        doc: yet to be verified
  cert:
    seq:
      - id: size
        type: u4
      - id: data
        size: size
        doc: DER-encoded Certificate
```

## .plist Schema

- `setupCert.plist`

```xml
<plist>
    <dict>
        <key>sign-sap-setup-cert</key>
        <data>${BASE64_ENCODED_SETUP_CRT}</data>
    </dict>
</plist>
```

## Download Locations

1. https://s.mzstatic.com/sap/setup.crt
2. https://s.mzstatic.com/sap/setupCert.plist
3. https://init.itunes.apple.com/WebObjects/MZInit.woa/wa/signSapSetupCert

## Possibly Useful

1. https://s.mzstatic.com/sap/fps-prod.crt
2. https://init.itunes.apple.com/WebObjects/MZInit.woa/wa/fpsCertificate
3. https://www.apple.com/appleca/AppleIncRootCertificate.cer
