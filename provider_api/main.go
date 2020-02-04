package main

import (
        "github.com/hashicorp/terraform/plugin"
        "github.com/hashicorp/terraform/terraform"
	"github.com/rajshreepardeshi/terraform_restapi/provider_api/restapi"
)




func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return restapi.Provider()
                },
        })
}

