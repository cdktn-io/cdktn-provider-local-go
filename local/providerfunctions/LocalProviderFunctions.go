// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-local-go/local/v14/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Provider-defined functions of the local provider.
type LocalProviderFunctions interface {
	// Given a path string, will return true if the directory exists.
	//
	// This function works only with directories. If used with a file, the function will return an error.
	//
	// This function behaves similar to the built-in [`fileexists`](https://developer.hashicorp.com/terraform/language/functions/fileexists) function, however, `direxists` will not replace filesystem paths including `~` with the current user's home directory path. This functionality can be achieved by using the built-in [`pathexpand`](https://developer.hashicorp.com/terraform/language/functions/pathexpand) function with `direxists`, see example below.
	Direxists(path *string) cdktn.IResolvable
}

// The jsii proxy struct for LocalProviderFunctions
type jsiiProxy_LocalProviderFunctions struct {
	_ byte // padding
}

func NewLocalProviderFunctions(providerLocalName *string) LocalProviderFunctions {
	_init_.Initialize()

	if err := validateNewLocalProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_LocalProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-local.providerFunctions.LocalProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewLocalProviderFunctions_Override(l LocalProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-local.providerFunctions.LocalProviderFunctions",
		[]interface{}{providerLocalName},
		l,
	)
}

func (l *jsiiProxy_LocalProviderFunctions) Direxists(path *string) cdktn.IResolvable {
	if err := l.validateDirexistsParameters(path); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"direxists",
		[]interface{}{path},
		&returns,
	)

	return returns
}

