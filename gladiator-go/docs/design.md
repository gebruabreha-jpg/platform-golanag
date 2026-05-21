# Design

## Use cases

### Proxy Mode

When the `--connect-udp` command-line option is specified, the client will operate in `proxy mode`. In proxy mode, the client initiates the creation and activation of the `OuterLayer` to facilitate UDP-based proxying functionality.

When the `--pvd-server` command-line option is specified (mutual exclusive to `--connect-udp`), the client need to get the information of MASQUE proxy service from PVD server.

#### QUIC-aware forwarding

When the `--quic-forwarding` option is specified, the client will negotiate with `MASQUE` proxy to use the [QUIC-Aware Proxying][QUIC-Aware].
After the negotiation success, the `OuterLayer` forward the **short header** QUIC packets between the `InnerLayer` and the `MASQUE` proxy.

##### Negotiation

- Connection Request

Add header "proxy-quic-forwarding = ?1"

- Connection Response

Check the existing of `capsule-protocol` and its value.
Check the existing of `proxy-quic-forwarding` and its value.

- Register Client Connection ID & Register Target Connection ID

After the `InnerLayer` completes the QUIC handshake, the client needs to register Client Connection ID and Target Connection ID to the `MASQUE` proxy.

>To initiate QUIC-aware proxying, the client sends a `REGISTER_CLIENT_CID` capsule containing the initial Client Connection ID that the client has advertised to the target as well as a Virtual Connection ID that the proxy MUST use when sending forwarded mode packets.

After the client receive a `ACK_CLIENT_CID` capsule that contains the same Client Connection ID that was requested, the Client Connection ID registration success.

>Once the client has learned the target server's Connection ID, such as in the response to a QUIC Initial packet, it can send a `REGISTER_TARGET_CID` capsule containing the Target Connection ID to request the ability to forward packets. The client MUST wait for an `ACK_TARGET_CID` capsule that contains the echoed connection ID and Virtual Target Connection ID before using forwarded mode.

##### Forwarding

- Sending direction

The `OuterLayer` needs to replace the `target connection id` with the `virtual target connection id`.

- Receiving direction

The `OuterLayer` needs to replace the `virtual client connection id` with the real `client connection id`.

#### Multiple Target Servers

The client establishes separate Http3/QUIC connections for each unique target server (host + port) identified in the request URLs. Multiple instances of `InnerLayer` exist, each managing a unique QUIC connection to a specific target server, handling the user's HTTP requests to the target server and the HTTP responses from the target server.

The `OuterLayer` has one QUIC connection to the `Masque Proxy Server` (specified by option `--connect-udp`). In this connection, it opens multiple streams for tunneling the traffic to/from the multiple target servers (host + port).

The `OuterLayer` maintains a mapping relationship between its QUIC stream IDs and the instances of `InnerLayer`. So that in the ingress direction, it knows how to dispatch the data to the instances of the `InnerLayer`.

### Non-proxy mode

If the `--connect-udp` command-line option is **not** specified, the client will operate in `non-proxy mode`. The `OuterLayer` won't be created.

## Class view

![class view](https://plantuml.lmera.ericsson.se/svg/XLDRJuCm67tlhsXuwa1UrFXaD3fTjcHcc6ZsieJKKYOHMoo5P77-TyjbfUsOlZHybtE-2-SZEIvBE5qkLeNCE8kmAAc0tpy0a3Ooox6yhzAqSG88WGwyP9zKhB4axtqeC3s8rd7EUU6VFTQKoJbcIKw5FndmiyG2QnDjcCPkM9gevvWYOvAA51CApLZ713cBfMg5Ln9D3Wv1ST8-EzGGT2rCEKkphJ7i4ow_AZhDCyMfo7u0VYkbm4J2VcQ1MLbm8PTurvOBR0yo2SnX5unHeK63TH8GrObk8z36oVHL9Gt-mHe1z8ZNXqZd7xvFqZzN6L7BOfGq6cdM7DWkPE_1nCwbB2viC0mSvn_mDHSVNPTRAkDU6SXSRNINvvnWoTiiY3kBnbbHNwf4stFLGLsz-s4l1W3qOlmmXDwvPro07TEA1S3PJulqFUj4WAMjDiMH4t3KrUKucnfFSdwGKuitxTnwXaZ6cZdC1knWCmTlIdG_0lGt4uRSkXPWtU7rNZ4wFHqd31MT_QdOFsChl75J8z5mk21W74Uult6t_Wi0)

## Sequence view

### Main Procedures

![sequence view](https://plantuml.lmera.ericsson.se/svg/pLNBRjim4BphAnQ-r8-yqDvqA134RPgsGPpOQIy115gY4pKooPAK4_zzTabAfcmd1OeUUX5epN0vd9sLSocibXNqbgmK_K4Fc05Q9da3_9JwaMlZa2U_JXJJjinBnOI5tbbLNYbRIZ5Xsd3jY7bLPb9PI18iL5Ky9zn0xvAPtpN77LnOqp1ftIrvMwZlR1rgrONgeXT2SBt1Io6w-5LjOzpMoA-wbCcgy28SGwsAP0l_COOfrwjG0WcPSJ-xD7yTc6ZvlFP4MdKvCwRlkOshIOEsaCbm-B5nAcqMHufSpF1NVn2JoAGqB8wCnHCG3--94bgJXiKHAuhibvptfqJzAkU3933srPCYsrCIZK_fqHwNt6jPAq3pm71ZiAKGyz4b60u1gshEBDg29dgN4u8y8_YOmfVPR0BddsR9vcMW1EpzoGbSVKkumEQ4asUkbUYbhLLyicznvt198O2AKh27EHgbOTbUBY0RJWhpkgmAq1nxR2nMnvaG6EL3z6GutDcGWgddozB2WxMgsHMnd9rFkaHnrxSNP_1PioN7U4f2tk7YEPlNYmNgrNW_BPAjNZZHhXreILOmor0rlMzg6va_Oj88HgGf4XWkcUWiQli0_LVkX2rslBPaIpSAvvSZttV5KKFbHcMS9eISmWoRKQmpAzs28lHcBzxOYP-INSElwpA7WkUouA3OkfAik1DUS_TjwyQcHYE4Zjy8zpZc4HD6Sbtw8ituiUd4hAHlqMxCY9Pwi0rO8y57UOgLOOmNtTus6crsHeqcg2sy9sZqJGcSnfsThFNFz5fPUJ2rA7vv8Eq3VWy6S2SmWtuJnh2JGVX3msA5-oQ8UXSg0iCBS-6wmioTSBg6m_7fxFJyvlGIJLquaVlA-9iiue7v1-Forfpig3k-7K-YwP2KTvUqOw9wr8N5j_U6vxnSSTzXQROdCDhK7RXNplqFPet-n3hjUdsqV8nTYgEPsglHHnQVVpFCvFSB)

[Back](../README.md)

[QUIC-Aware]: https://datatracker.ietf.org/doc/draft-ietf-masque-quic-proxy/
