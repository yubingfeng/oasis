

package main 

import (
        "fmt"
	"github.com/rackspace/gophercloud/openstack"
)

func main(){

	authOpts, err := openstack.AuthOptionsFromEnv()
	if err != nil{
		fmt.Printf("Error in get auth options from env %s \n", err)
	}

	provider, err := openstack.AuthenticatedClient(authOpts) 
	if err != nil{
		fmt.Printf("Error in get provider::: %s \n", err)
		return
	}
        client  :=  openstack.NewIdentityV3(provider)
	fmt.Printf("the client is %s \n", client)
}
