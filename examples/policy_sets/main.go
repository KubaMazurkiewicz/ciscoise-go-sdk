package main

import (
	"fmt"
	"os"

	isegosdk "github.com/KubaMazurkiewicz/ciscoise-go-sdk/sdk"
)

func main() {
	Client, err := isegosdk.NewClientWithOptions("https://10.48.190.181",
		"admin", "Devnet.12345",
		"false", "false",
		"false", "false")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("executing GetNetworkAccessPolicySets...")
	policy_sets, _, err := Client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()

	if err != nil {
		fmt.Println(err)
	}

	if policy_sets != nil && policy_sets.Response != nil {
		for _, row := range *policy_sets.Response {
			fmt.Println("-------------")
			fmt.Println("ID: ", row.ID)
			fmt.Println("Name: ", row.Name)
			fmt.Println("Description: ", row.Description)
			fmt.Println("Condition: ", row.Condition)
		}
	}
	fmt.Println("-------------")
	falseVal := false

	prof := &isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet{
		Name:        "Test_9988",
		Description: "Sample policy_set using ciscoise-go-sdk",
		ServiceName: "Default Network Access",
		State:       "enabled",
		IsProxy:     false
	}

	reqProf := &isegosdk.RequestAuthorizationProfileCreateAuthorizationProfile{
		AuthorizationProfile: prof,
	}

	fmt.Println("executing CreateAuthorizationProfile...")
	_, err = Client.AuthorizationProfile.CreateAuthorizationProfile(reqProf)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("executing GetAuthorizationProfileByName...")
	profByID, _, err := Client.AuthorizationProfile.GetAuthorizationProfileByName("Internet_Only")
	if err != nil {
		fmt.Println(err)
	}
	if profByID != nil && profByID.AuthorizationProfile != nil {
		newProfileID := profByID.AuthorizationProfile.ID
		fmt.Println("-------------")
		fmt.Println("ID: ", profByID.AuthorizationProfile.ID)
		fmt.Println("Name: ", profByID.AuthorizationProfile.Name)
		fmt.Println("Description: ", profByID.AuthorizationProfile.Description)
		fmt.Println("AdvancedAttributes: ", profByID.AuthorizationProfile.AdvancedAttributes)
		fmt.Println("AccessType: ", profByID.AuthorizationProfile.AccessType)
		fmt.Println("AuthzProfileType: ", profByID.AuthorizationProfile.AuthzProfileType)
		fmt.Println("VLAN: ", profByID.AuthorizationProfile.VLAN)
		fmt.Println("Reauth: ", profByID.AuthorizationProfile.Reauth)
		fmt.Println("WebRedirection: ", profByID.AuthorizationProfile.WebRedirection)
		fmt.Println("TrackMovement: ", profByID.AuthorizationProfile.TrackMovement)
		fmt.Println("AgentlessPosture: ", profByID.AuthorizationProfile.AgentlessPosture)
		fmt.Println("ServiceTemplate: ", profByID.AuthorizationProfile.ServiceTemplate)
		fmt.Println("EasywiredSessionCandidate: ", profByID.AuthorizationProfile.EasywiredSessionCandidate)
		fmt.Println("DaclName: ", profByID.AuthorizationProfile.DaclName)
		fmt.Println("VoiceDomainPermission: ", profByID.AuthorizationProfile.VoiceDomainPermission)
		fmt.Println("Neat: ", profByID.AuthorizationProfile.Neat)
		fmt.Println("WebAuth: ", profByID.AuthorizationProfile.WebAuth)
		fmt.Println("ProfileName: ", profByID.AuthorizationProfile.ProfileName)
		fmt.Println("Link: ", profByID.AuthorizationProfile.Link)
		fmt.Println("-------------")

		fmt.Println("executing DeleteAuthorizationProfileByID...")
		_, err = Client.AuthorizationProfile.DeleteAuthorizationProfileByID(newProfileID)

		if err != nil {
			fmt.Println(err)
		}
	}
}
