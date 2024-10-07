/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v2alpha1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v2alpha1

import (
	"io"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalDNSDomain writes a value of the 'DNS_domain' type to the given writer.
func MarshalDNSDomain(object *DNSDomain, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeDNSDomain(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeDNSDomain writes a value of the 'DNS_domain' type to the given stream.
func writeDNSDomain(object *DNSDomain, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if object.bitmap_&1 != 0 {
		stream.WriteString(DNSDomainLinkKind)
	} else {
		stream.WriteString(DNSDomainKind)
	}
	count++
	if object.bitmap_&2 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if object.bitmap_&4 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = object.bitmap_&8 != 0 && object.cluster != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster")
		writeClusterLink(object.cluster, stream)
		count++
	}
	present_ = object.bitmap_&16 != 0 && object.organization != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("organization")
		writeOrganizationLink(object.organization, stream)
		count++
	}
	present_ = object.bitmap_&32 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("reserved_at_timestamp")
		stream.WriteString((object.reservedAtTimestamp).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&64 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("user_defined")
		stream.WriteBool(object.userDefined)
	}
	stream.WriteObjectEnd()
}

// UnmarshalDNSDomain reads a value of the 'DNS_domain' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalDNSDomain(source interface{}) (object *DNSDomain, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readDNSDomain(iterator)
	err = iterator.Error
	return
}

// readDNSDomain reads a value of the 'DNS_domain' type from the given iterator.
func readDNSDomain(iterator *jsoniter.Iterator) *DNSDomain {
	object := &DNSDomain{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == DNSDomainLinkKind {
				object.bitmap_ |= 1
			}
		case "id":
			object.id = iterator.ReadString()
			object.bitmap_ |= 2
		case "href":
			object.href = iterator.ReadString()
			object.bitmap_ |= 4
		case "cluster":
			value := readClusterLink(iterator)
			object.cluster = value
			object.bitmap_ |= 8
		case "organization":
			value := readOrganizationLink(iterator)
			object.organization = value
			object.bitmap_ |= 16
		case "reserved_at_timestamp":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.reservedAtTimestamp = value
			object.bitmap_ |= 32
		case "user_defined":
			value := iterator.ReadBool()
			object.userDefined = value
			object.bitmap_ |= 64
		default:
			iterator.ReadAny()
		}
	}
	return object
}
