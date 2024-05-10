import * as ise from "@pulumi/ise";

const username = process.env["ISE_USERNAME"] || "";
const password = process.env["ISE_PASSWORD"] || "";
if (username === "" || password === "") {
    throw new Error("Missing ISE username and/or password");
}

const url = process.env["ISE_URL"] || "";
if (url === "") {
    throw new Error("Missing ISE URL");
}

const provider = new ise.Provider("example-ise-provider", {
    username: username, //Not required, the provider will pick this via ISE_USERNAME env variable
    password: password, //Not required, the provider will pick this via ISE_PASSWORD env variable
    url: url, //Not required, the provider will pick this via ISE_URL env variable
    retries: 0 //Defaults to 3
});

const exampleEmployeeGroup = new ise.identitymanagement.UserIdentityGroup("example-employee-group", {
    name: "PulumiEmployeeGroup",
    description: "My pulumi employee identity group",
}, {
    provider: provider
});
const exampleGuestsGroup = new ise.identitymanagement.UserIdentityGroup("example-guests-group", {
    name: "PulumiGuestGroup",
    description: "My pulumi guest identity group",
}, {
    provider: provider
});

new ise.identitymanagement.InternalUser("example-internal-user-employee", {
    name: "PulumiEmployee",
    password: "Cisco123",
    changePassword: true,
    email: "aaa@cisco.com",
    accountNameAlias: "Pulumi Employee 1",
    enablePassword: "Cisco123",
    enabled: true,
    passwordNeverExpires: false,
    firstName: "John",
    lastName: "Doe",
    passwordIdStore: "Internal Users",
    description: "My first Pulumi user",
    identityGroups: exampleEmployeeGroup.id
}, {
    provider: provider
});

new ise.identitymanagement.InternalUser("example-internal-user-guest", {
    name: "PulumiGuest",
    password: "Cisco123",
    changePassword: true,
    email: "aaa@cisco.com",
    accountNameAlias: "Pulumi Guest 1",
    enablePassword: "Cisco123",
    enabled: true,
    passwordNeverExpires: false,
    firstName: "Jane",
    lastName: "Doe",
    passwordIdStore: "Internal Users",
    description: "My first Pulumi guest",
    identityGroups: exampleGuestsGroup.id
}, {
    provider: provider
});

const parentEndpointIdentityGroupId = ise.identitymanagement.getEndpointIdentityGroupOutput({
    name: "Profiled"
}, {
    provider: provider
})

const endpointIdentityGroup = new ise.identitymanagement.EndpointIdentityGroup("example-endpoint-identity-group", {
    name: "PulumiEndpointGroup",
    description: "My endpoint identity group for Pulumi",
    parentEndpointIdentityGroupId: parentEndpointIdentityGroupId.id
}, {
    provider: provider
})

new ise.identitymanagement.Endpoint("example-endpoint", {
    name: "PulumiEndpoint",
    description: "My pulumi endpoint",
    mac: "aa:aa:aa:aa:aa:aa",
    groupId: endpointIdentityGroup.id,
    profileId: "061f5530-19f6-11ee-88d5-0050568fa0ed",
    staticProfileAssignment: true,
    staticGroupAssignment: true,
}, {
    provider: provider
});

const exampleDeviceGroup = new ise.network.DeviceGroup("example-device-group", {
    name: "All Device Types#All Pulumi Device Types#Group1",
    description: "My pulumi network device group",
    rootGroup: "All Device Types",
}, {
    provider: provider
});

new ise.network.Device("example-network-device", {
    name: "PulumiDevice1",
    description: "My Pulumi device description",
    ips: [{
        ipaddress: "0.0.0.1",
        mask: "32",
    }],
    networkDeviceGroups: [exampleDeviceGroup.name]
}, {
    provider: provider
})

new ise.networkaccess.PolicySet("example-network-access-policy-set", {
    name: "PulumiNetworkAccessPolicySet1",
    description: "My Pulumi network access policy set description",
    isProxy: false,
    rank: 0,
    serviceName: "Default Network Access",
    state: "disabled",
    conditionType: "ConditionAttributes",
    conditionIsNegate: false,
    conditionAttributeName: "Location",
    conditionAttributeValue: "All Locations",
    conditionDictionaryName: "DEVICE",
    conditionOperator: "equals",
}, {
    provider: provider
})
