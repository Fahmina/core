// source: proto/api/common/v1/common.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');
goog.object.extend(proto, google_protobuf_struct_pb);
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
goog.object.extend(proto, google_protobuf_duration_pb);
var google_api_annotations_pb = require('../../../../google/api/annotations_pb.js');
goog.object.extend(proto, google_api_annotations_pb);
var google_api_httpbody_pb = require('../../../../google/api/httpbody_pb.js');
goog.object.extend(proto, google_api_httpbody_pb);
goog.exportSymbol('proto.proto.api.common.v1.Pose', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.api.common.v1.Pose = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.api.common.v1.Pose, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.api.common.v1.Pose.displayName = 'proto.proto.api.common.v1.Pose';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.api.common.v1.Pose.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.api.common.v1.Pose.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.api.common.v1.Pose} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.api.common.v1.Pose.toObject = function(includeInstance, msg) {
  var f, obj = {
    x: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0),
    y: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
    z: jspb.Message.getFloatingPointFieldWithDefault(msg, 3, 0.0),
    oX: jspb.Message.getFloatingPointFieldWithDefault(msg, 4, 0.0),
    oY: jspb.Message.getFloatingPointFieldWithDefault(msg, 5, 0.0),
    oZ: jspb.Message.getFloatingPointFieldWithDefault(msg, 6, 0.0),
    theta: jspb.Message.getFloatingPointFieldWithDefault(msg, 7, 0.0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.api.common.v1.Pose}
 */
proto.proto.api.common.v1.Pose.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.api.common.v1.Pose;
  return proto.proto.api.common.v1.Pose.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.api.common.v1.Pose} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.api.common.v1.Pose}
 */
proto.proto.api.common.v1.Pose.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setX(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setY(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setZ(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setOX(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setOY(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setOZ(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setTheta(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.api.common.v1.Pose.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.api.common.v1.Pose.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.api.common.v1.Pose} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.api.common.v1.Pose.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getX();
  if (f !== 0.0) {
    writer.writeDouble(
      1,
      f
    );
  }
  f = message.getY();
  if (f !== 0.0) {
    writer.writeDouble(
      2,
      f
    );
  }
  f = message.getZ();
  if (f !== 0.0) {
    writer.writeDouble(
      3,
      f
    );
  }
  f = message.getOX();
  if (f !== 0.0) {
    writer.writeDouble(
      4,
      f
    );
  }
  f = message.getOY();
  if (f !== 0.0) {
    writer.writeDouble(
      5,
      f
    );
  }
  f = message.getOZ();
  if (f !== 0.0) {
    writer.writeDouble(
      6,
      f
    );
  }
  f = message.getTheta();
  if (f !== 0.0) {
    writer.writeDouble(
      7,
      f
    );
  }
};


/**
 * optional double x = 1;
 * @return {number}
 */
proto.proto.api.common.v1.Pose.prototype.getX = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.api.common.v1.Pose} returns this
 */
proto.proto.api.common.v1.Pose.prototype.setX = function(value) {
  return jspb.Message.setProto3FloatField(this, 1, value);
};


/**
 * optional double y = 2;
 * @return {number}
 */
proto.proto.api.common.v1.Pose.prototype.getY = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.api.common.v1.Pose} returns this
 */
proto.proto.api.common.v1.Pose.prototype.setY = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional double z = 3;
 * @return {number}
 */
proto.proto.api.common.v1.Pose.prototype.getZ = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.api.common.v1.Pose} returns this
 */
proto.proto.api.common.v1.Pose.prototype.setZ = function(value) {
  return jspb.Message.setProto3FloatField(this, 3, value);
};


/**
 * optional double o_x = 4;
 * @return {number}
 */
proto.proto.api.common.v1.Pose.prototype.getOX = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 4, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.api.common.v1.Pose} returns this
 */
proto.proto.api.common.v1.Pose.prototype.setOX = function(value) {
  return jspb.Message.setProto3FloatField(this, 4, value);
};


/**
 * optional double o_y = 5;
 * @return {number}
 */
proto.proto.api.common.v1.Pose.prototype.getOY = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 5, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.api.common.v1.Pose} returns this
 */
proto.proto.api.common.v1.Pose.prototype.setOY = function(value) {
  return jspb.Message.setProto3FloatField(this, 5, value);
};


/**
 * optional double o_z = 6;
 * @return {number}
 */
proto.proto.api.common.v1.Pose.prototype.getOZ = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 6, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.api.common.v1.Pose} returns this
 */
proto.proto.api.common.v1.Pose.prototype.setOZ = function(value) {
  return jspb.Message.setProto3FloatField(this, 6, value);
};


/**
 * optional double theta = 7;
 * @return {number}
 */
proto.proto.api.common.v1.Pose.prototype.getTheta = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 7, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.proto.api.common.v1.Pose} returns this
 */
proto.proto.api.common.v1.Pose.prototype.setTheta = function(value) {
  return jspb.Message.setProto3FloatField(this, 7, value);
};


goog.object.extend(exports, proto.proto.api.common.v1);