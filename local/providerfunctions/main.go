// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-local.providerFunctions.LocalProviderFunctions",
		reflect.TypeOf((*LocalProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "direxists", GoMethod: "Direxists"},
		},
		func() interface{} {
			return &jsiiProxy_LocalProviderFunctions{}
		},
	)
}
