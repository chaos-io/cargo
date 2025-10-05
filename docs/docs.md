# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [cargo/cargo.proto](#cargo_cargo-proto)
    - [Cargo](#cargo-Cargo)
    - [CarrierMovement](#cargo-CarrierMovement)
    - [Customer](#cargo-Customer)
    - [DeliveryHistory](#cargo-DeliveryHistory)
    - [DeliverySpecification](#cargo-DeliverySpecification)
    - [HandlingEvent](#cargo-HandlingEvent)
    - [Location](#cargo-Location)
  
- [Scalar Value Types](#scalar-value-types)



<a name="cargo_cargo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cargo/cargo.proto



<a name="cargo-Cargo"></a>

### Cargo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tracking_id | [string](#string) |  |  |
| customer_id | [string](#string) |  |  |
| Shipper | [Customer](#cargo-Customer) |  |  |
| Consignee | [Customer](#cargo-Customer) |  |  |
| Spec | [DeliverySpecification](#cargo-DeliverySpecification) |  |  |






<a name="cargo-CarrierMovement"></a>

### CarrierMovement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule_id | [string](#string) |  |  |
| from | [Location](#cargo-Location) |  |  |
| to | [Location](#cargo-Location) |  |  |






<a name="cargo-Customer"></a>

### Customer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="cargo-DeliveryHistory"></a>

### DeliveryHistory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| cargo_id | [string](#string) |  |  |
| schedule_id | [string](#string) |  |  |
| CompleteTime | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="cargo-DeliverySpecification"></a>

### DeliverySpecification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| From | [Location](#cargo-Location) |  |  |
| To | [Location](#cargo-Location) |  |  |
| ETA | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="cargo-HandlingEvent"></a>

### HandlingEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cargo | [Cargo](#cargo-Cargo) |  |  |
| event_type | [string](#string) |  |  |
| CompleteTime | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| movement | [CarrierMovement](#cargo-CarrierMovement) |  |  |






<a name="cargo-Location"></a>

### Location



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [string](#string) |  |  |
| city | [string](#string) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

