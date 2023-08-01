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
			if row.Condition != nil {
				fmt.Println("Children: ", row.Condition.Children)
			}
		}
	}
	fmt.Println("-------------")

	fmt.Println("executing GetNetworkAccessPolicySetsbyID...")
	policy_set, _, err := Client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID("13727e31-93d1-4fc6-9a7e-80f6998b1916")

	if err != nil {
		fmt.Println(err)
	}
	if policy_set != nil && policy_set.Response != nil {

		fmt.Println("-------------")
		fmt.Println("ID: ", policy_set.Response.ID)
		fmt.Println("Name: ", policy_set.Response.Name)
		fmt.Println("Description: ", policy_set.Response.Description)
		fmt.Println("Condition: ", policy_set.Response.Condition)
		if policy_set.Response.Condition != nil {
			fmt.Println("Children: ", policy_set.Response.Condition.Children)
		}
	}
	fmt.Println("-------------")

	falseVal := false

	policy_set2 := &isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet{
		Name:        "Test_9983",
		Description: "Sample policy_set using ciscoise-go-sdk",
		ServiceName: "Default Network Access",
		State:       "enabled",
		IsProxy:     &falseVal,
		Condition: &isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
			ConditionType: "ConditionAndBlock",
			IsNegate:      &falseVal,
			Children: &[]isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
				{
					ConditionType: "ConditionReference",
					ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
				},
				{
					ConditionType: "ConditionReference",
					ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
				},
				{
					ConditionType: "ConditionAndBlock",
					IsNegate:      &falseVal,
					Children: &[]isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
						{
							ConditionType: "ConditionReference",
							ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
						},
						{
							ConditionType: "ConditionReference",
							ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
						},
						{
							ConditionType: "ConditionAndBlock",
							IsNegate:      &falseVal,
							Children: &[]isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
								{
									ConditionType: "ConditionReference",
									ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
								},
								{
									ConditionType: "ConditionReference",
									ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
								},
							},
						},
					},
				},
			},
		},
	}

	policy_set1 := &isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySet{
		Name:        "Test_9988_1",
		Description: "Sample policy_set using ciscoise-go-sdk",
		ServiceName: "Default Network Access",
		State:       "enabled",
		IsProxy:     &falseVal,
		Condition: &isegosdk.RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition{
			IsNegate:      &falseVal,
			ConditionType: "ConditionReference",
			ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
		},
	}

	policy_set3 := &isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID{
		Name:        "Test_9988_1",
		Description: "Sample policy_set using ciscoise-go-sdk",
		ServiceName: "Default Network Access",
		State:       "enabled",
		IsProxy:     &falseVal,
		Condition: &isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
			IsNegate:      &falseVal,
			ConditionType: "ConditionReference",
			ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
		},
	}

	fmt.Println(policy_set3)

	policy_set4 := &isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByID{
		Name:        "Test_9988_2",
		Description: "Sample policy_set using ciscoise-go-sdk",
		ServiceName: "Default Network Access",
		State:       "enabled",
		IsProxy:     &falseVal,
		Condition: &isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
			ConditionType: "ConditionAndBlock",
			IsNegate:      &falseVal,
			Children: &[]isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
				{
					ConditionType: "ConditionReference",
					ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
				},
				{
					ConditionType: "ConditionReference",
					ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
				},
				{
					ConditionType: "ConditionAndBlock",
					IsNegate:      &falseVal,
					Children: &[]isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
						{
							ConditionType: "ConditionReference",
							ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
						},
						{
							ConditionType: "ConditionReference",
							ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
						},
						{
							ConditionType: "ConditionAndBlock",
							IsNegate:      &falseVal,
							Children: &[]isegosdk.RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition{
								{
									ConditionType: "ConditionReference",
									ID:            "49ec9c17-6987-43ef-9478-a42111c59b81",
								},
								{
									ConditionType: "ConditionReference",
									ID:            "1e9f02d0-13e7-4a3b-8fbc-fb38330d74b5",
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(policy_set1)
	fmt.Println(policy_set2)
	fmt.Println("executing CreatePolicySet...")

	//policy_set_, _, err := Client.NetworkAccessPolicySet.CreateNetworkAccessPolicySet(policy_set2)
	//if err != nil {
	//		fmt.Println(err)
	//	}
	//	if policy_set_ != nil {
	//		fmt.Println(policy_set_)
	//	}

	//policy_set1_, _, err := Client.NetworkAccessPolicySet.CreateNetworkAccessPolicySet(policy_set1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if policy_set1_ != nil {
	//		fmt.Println(policy_set1_.Response.Condition)
	//	}

	fmt.Println("executing UpdateNetworkAccessPolicySetByID...")
	policy_setByID, _, err := Client.NetworkAccessPolicySet.UpdateNetworkAccessPolicySetByID("13727e31-93d1-4fc6-9a7e-80f6998b1916", policy_set4)
	if err != nil {
		fmt.Println(err)
	}

	if policy_setByID != nil {
		fmt.Println(policy_setByID.Response.Condition)
	}

	//if profByID != nil && profByID.AuthorizationProfile != nil {
	//	newProfileID := profByID.AuthorizationProfile.ID
	//	fmt.Println("-------------")
	//	fmt.Println("ID: ", profByID.AuthorizationProfile.ID)
	//	fmt.Println("Name: ", profByID.AuthorizationProfile.Name)
	//	fmt.Println("Description: ", profByID.AuthorizationProfile.Description)
	//	fmt.Println("AdvancedAttributes: ", profByID.AuthorizationProfile.AdvancedAttributes)
	//	fmt.Println("AccessType: ", profByID.AuthorizationProfile.AccessType)
	//	fmt.Println("AuthzProfileType: ", profByID.AuthorizationProfile.AuthzProfileType)
	//	fmt.Println("VLAN: ", profByID.AuthorizationProfile.VLAN)
	//	fmt.Println("Reauth: ", profByID.AuthorizationProfile.Reauth)
	//	fmt.Println("WebRedirection: ", profByID.AuthorizationProfile.WebRedirection)
	//	fmt.Println("TrackMovement: ", profByID.AuthorizationProfile.TrackMovement)
	//	fmt.Println("AgentlessPosture: ", profByID.AuthorizationProfile.AgentlessPosture)
	//	fmt.Println("ServiceTemplate: ", profByID.AuthorizationProfile.ServiceTemplate)
	//	fmt.Println("EasywiredSessionCandidate: ", profByID.AuthorizationProfile.EasywiredSessionCandidate)
	//	fmt.Println("DaclName: ", profByID.AuthorizationProfile.DaclName)
	//	fmt.Println("VoiceDomainPermission: ", profByID.AuthorizationProfile.VoiceDomainPermission)
	//	fmt.Println("Neat: ", profByID.AuthorizationProfile.Neat)
	//	fmt.Println("WebAuth: ", profByID.AuthorizationProfile.WebAuth)
	//	fmt.Println("ProfileName: ", profByID.AuthorizationProfile.ProfileName)
	//	fmt.Println("Link: ", profByID.AuthorizationProfile.Link)
	//	fmt.Println("-------------")

	//	fmt.Println("executing DeleteAuthorizationProfileByID...")
	//	_, err = Client.AuthorizationProfile.DeleteAuthorizationProfileByID(newProfileID)

	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
}
