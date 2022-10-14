using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.VMware.v2021_12_01.PlacementPolicies;

internal class VirtualMachineId : ResourceID
{
    public string? CommonAlias => null;

    public string ID => "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/virtualMachines/{virtualMachineId}";

    public List<ResourceIDSegment> Segments => new List<ResourceIDSegment>
    {
        ResourceIDSegment.Static("staticSubscriptions", "subscriptions"),
        ResourceIDSegment.SubscriptionId("subscriptionId"),
        ResourceIDSegment.Static("staticResourceGroups", "resourceGroups"),
        ResourceIDSegment.ResourceGroup("resourceGroupName"),
        ResourceIDSegment.Static("staticProviders", "providers"),
        ResourceIDSegment.ResourceProvider("staticMicrosoftAVS", "Microsoft.AVS"),
        ResourceIDSegment.Static("staticPrivateClouds", "privateClouds"),
        ResourceIDSegment.UserSpecified("privateCloudName"),
        ResourceIDSegment.Static("staticClusters", "clusters"),
        ResourceIDSegment.UserSpecified("clusterName"),
        ResourceIDSegment.Static("staticVirtualMachines", "virtualMachines"),
        ResourceIDSegment.UserSpecified("virtualMachineId"),
    };
}