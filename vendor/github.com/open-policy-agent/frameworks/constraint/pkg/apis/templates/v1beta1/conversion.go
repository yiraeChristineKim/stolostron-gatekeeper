/*

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

package v1beta1

import (
	apisTemplates "github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	coreTemplates "github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/conversion"
)

func Convert_v1beta1_Validation_To_templates_Validation(in *Validation, out *coreTemplates.Validation, s conversion.Scope) error { //nolint:golint
	inSchema := in.OpenAPIV3Schema
	// to preserve legacy behavior, allow users to provide arbitrary parameters, regardless of whether the user specified them
	if inSchema == nil {
		inSchema = &apiextensionsv1beta1.JSONSchemaProps{}
	}

	inSchemaCopy := inSchema.DeepCopy()
	if err := apisTemplates.AddPreserveUnknownFields(inSchemaCopy); err != nil {
		return err
	}

	out.OpenAPIV3Schema = new(apiextensions.JSONSchemaProps)
	if err := apiextensionsv1beta1.Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(inSchemaCopy, out.OpenAPIV3Schema, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_CRDSpec_To_templates_CRDSpec(in *CRDSpec, out *templates.CRDSpec, s conversion.Scope) error { //nolint:golint
	if err := Convert_v1beta1_Names_To_templates_Names(&in.Names, &out.Names, s); err != nil {
		return err
	}
	validation := in.Validation
	if validation == nil {
		validation = &Validation{}
	}
	{
		in, out := &validation, &out.Validation
		*out = new(templates.Validation)
		if err := Convert_v1beta1_Validation_To_templates_Validation(*in, *out, s); err != nil {
			return err
		}
	}
	return nil
}
