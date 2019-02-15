package yandex

import "github.com/hashicorp/terraform/helper/schema"

func resourceYandexIAMServiceAccountIAMBinding() *schema.Resource {
	return resourceIamBindingWithImport(IamServiceAccountSchema, newServiceAccountIamUpdater, serviceAccountIDParseFunc)
}
