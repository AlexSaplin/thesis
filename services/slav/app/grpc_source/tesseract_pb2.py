# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tesseract.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='tesseract.proto',
  package='tesseract',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0ftesseract.proto\x12\ttesseract\" \n\x02KV\x12\x0b\n\x03Key\x18\x01 \x01(\t\x12\r\n\x05Value\x18\x02 \x01(\t\"\xb2\x01\n\x0c\x41pplyRequest\x12\x0c\n\x04Name\x18\x01 \x01(\t\x12\n\n\x02ID\x18\x02 \x01(\t\x12\x0b\n\x03\x44NS\x18\x03 \x01(\t\x12\r\n\x05Scale\x18\x04 \x01(\r\x12\x0b\n\x03\x43PU\x18\x05 \x01(\r\x12\x0b\n\x03RAM\x18\x06 \x01(\r\x12\x0b\n\x03GPU\x18\x07 \x01(\t\x12\x0c\n\x04Port\x18\x08 \x01(\r\x12\r\n\x05Image\x18\t \x01(\t\x12\x1a\n\x03\x45nv\x18\n \x03(\x0b\x32\r.tesseract.KV\x12\x0c\n\x04\x41uth\x18\x0b \x01(\t\"\x0f\n\rApplyResponse\"\x1e\n\x10GetStatusRequest\x12\n\n\x02ID\x18\x01 \x01(\t\"E\n\x11GetStatusResponse\x12!\n\x06Status\x18\x01 \x01(\x0e\x32\x11.tesseract.Status\x12\r\n\x05\x45rror\x18\x02 \x01(\t\"\x1b\n\rDeleteRequest\x12\n\n\x02ID\x18\x01 \x01(\t\"\x10\n\x0e\x44\x65leteResponse*;\n\x06Status\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07RUNNING\x10\x01\x12\x0c\n\x08UPDATING\x10\x02\x12\t\n\x05\x45RROR\x10\x03\x32\xce\x01\n\tTesseract\x12:\n\x05\x41pply\x12\x17.tesseract.ApplyRequest\x1a\x18.tesseract.ApplyResponse\x12\x46\n\tGetStatus\x12\x1b.tesseract.GetStatusRequest\x1a\x1c.tesseract.GetStatusResponse\x12=\n\x06\x44\x65lete\x12\x18.tesseract.DeleteRequest\x1a\x19.tesseract.DeleteResponseb\x06proto3'
)

_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='tesseract.Status',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RUNNING', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UPDATING', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ERROR', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=412,
  serialized_end=471,
)
_sym_db.RegisterEnumDescriptor(_STATUS)

Status = enum_type_wrapper.EnumTypeWrapper(_STATUS)
UNKNOWN = 0
RUNNING = 1
UPDATING = 2
ERROR = 3



_KV = _descriptor.Descriptor(
  name='KV',
  full_name='tesseract.KV',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Key', full_name='tesseract.KV.Key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Value', full_name='tesseract.KV.Value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=30,
  serialized_end=62,
)


_APPLYREQUEST = _descriptor.Descriptor(
  name='ApplyRequest',
  full_name='tesseract.ApplyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Name', full_name='tesseract.ApplyRequest.Name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ID', full_name='tesseract.ApplyRequest.ID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='DNS', full_name='tesseract.ApplyRequest.DNS', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Scale', full_name='tesseract.ApplyRequest.Scale', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CPU', full_name='tesseract.ApplyRequest.CPU', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='RAM', full_name='tesseract.ApplyRequest.RAM', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='GPU', full_name='tesseract.ApplyRequest.GPU', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Port', full_name='tesseract.ApplyRequest.Port', index=7,
      number=8, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Image', full_name='tesseract.ApplyRequest.Image', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Env', full_name='tesseract.ApplyRequest.Env', index=9,
      number=10, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Auth', full_name='tesseract.ApplyRequest.Auth', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=65,
  serialized_end=243,
)


_APPLYRESPONSE = _descriptor.Descriptor(
  name='ApplyResponse',
  full_name='tesseract.ApplyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=245,
  serialized_end=260,
)


_GETSTATUSREQUEST = _descriptor.Descriptor(
  name='GetStatusRequest',
  full_name='tesseract.GetStatusRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='tesseract.GetStatusRequest.ID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=262,
  serialized_end=292,
)


_GETSTATUSRESPONSE = _descriptor.Descriptor(
  name='GetStatusResponse',
  full_name='tesseract.GetStatusResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Status', full_name='tesseract.GetStatusResponse.Status', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Error', full_name='tesseract.GetStatusResponse.Error', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=294,
  serialized_end=363,
)


_DELETEREQUEST = _descriptor.Descriptor(
  name='DeleteRequest',
  full_name='tesseract.DeleteRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='tesseract.DeleteRequest.ID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=365,
  serialized_end=392,
)


_DELETERESPONSE = _descriptor.Descriptor(
  name='DeleteResponse',
  full_name='tesseract.DeleteResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=394,
  serialized_end=410,
)

_APPLYREQUEST.fields_by_name['Env'].message_type = _KV
_GETSTATUSRESPONSE.fields_by_name['Status'].enum_type = _STATUS
DESCRIPTOR.message_types_by_name['KV'] = _KV
DESCRIPTOR.message_types_by_name['ApplyRequest'] = _APPLYREQUEST
DESCRIPTOR.message_types_by_name['ApplyResponse'] = _APPLYRESPONSE
DESCRIPTOR.message_types_by_name['GetStatusRequest'] = _GETSTATUSREQUEST
DESCRIPTOR.message_types_by_name['GetStatusResponse'] = _GETSTATUSRESPONSE
DESCRIPTOR.message_types_by_name['DeleteRequest'] = _DELETEREQUEST
DESCRIPTOR.message_types_by_name['DeleteResponse'] = _DELETERESPONSE
DESCRIPTOR.enum_types_by_name['Status'] = _STATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

KV = _reflection.GeneratedProtocolMessageType('KV', (_message.Message,), {
  'DESCRIPTOR' : _KV,
  '__module__' : 'tesseract_pb2'
  # @@protoc_insertion_point(class_scope:tesseract.KV)
  })
_sym_db.RegisterMessage(KV)

ApplyRequest = _reflection.GeneratedProtocolMessageType('ApplyRequest', (_message.Message,), {
  'DESCRIPTOR' : _APPLYREQUEST,
  '__module__' : 'tesseract_pb2'
  # @@protoc_insertion_point(class_scope:tesseract.ApplyRequest)
  })
_sym_db.RegisterMessage(ApplyRequest)

ApplyResponse = _reflection.GeneratedProtocolMessageType('ApplyResponse', (_message.Message,), {
  'DESCRIPTOR' : _APPLYRESPONSE,
  '__module__' : 'tesseract_pb2'
  # @@protoc_insertion_point(class_scope:tesseract.ApplyResponse)
  })
_sym_db.RegisterMessage(ApplyResponse)

GetStatusRequest = _reflection.GeneratedProtocolMessageType('GetStatusRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETSTATUSREQUEST,
  '__module__' : 'tesseract_pb2'
  # @@protoc_insertion_point(class_scope:tesseract.GetStatusRequest)
  })
_sym_db.RegisterMessage(GetStatusRequest)

GetStatusResponse = _reflection.GeneratedProtocolMessageType('GetStatusResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETSTATUSRESPONSE,
  '__module__' : 'tesseract_pb2'
  # @@protoc_insertion_point(class_scope:tesseract.GetStatusResponse)
  })
_sym_db.RegisterMessage(GetStatusResponse)

DeleteRequest = _reflection.GeneratedProtocolMessageType('DeleteRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEREQUEST,
  '__module__' : 'tesseract_pb2'
  # @@protoc_insertion_point(class_scope:tesseract.DeleteRequest)
  })
_sym_db.RegisterMessage(DeleteRequest)

DeleteResponse = _reflection.GeneratedProtocolMessageType('DeleteResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETERESPONSE,
  '__module__' : 'tesseract_pb2'
  # @@protoc_insertion_point(class_scope:tesseract.DeleteResponse)
  })
_sym_db.RegisterMessage(DeleteResponse)



_TESSERACT = _descriptor.ServiceDescriptor(
  name='Tesseract',
  full_name='tesseract.Tesseract',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=474,
  serialized_end=680,
  methods=[
  _descriptor.MethodDescriptor(
    name='Apply',
    full_name='tesseract.Tesseract.Apply',
    index=0,
    containing_service=None,
    input_type=_APPLYREQUEST,
    output_type=_APPLYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetStatus',
    full_name='tesseract.Tesseract.GetStatus',
    index=1,
    containing_service=None,
    input_type=_GETSTATUSREQUEST,
    output_type=_GETSTATUSRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Delete',
    full_name='tesseract.Tesseract.Delete',
    index=2,
    containing_service=None,
    input_type=_DELETEREQUEST,
    output_type=_DELETERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TESSERACT)

DESCRIPTOR.services_by_name['Tesseract'] = _TESSERACT

# @@protoc_insertion_point(module_scope)
