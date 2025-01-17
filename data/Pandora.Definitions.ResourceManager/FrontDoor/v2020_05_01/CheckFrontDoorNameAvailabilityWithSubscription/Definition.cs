using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.FrontDoor.v2020_05_01.CheckFrontDoorNameAvailabilityWithSubscription;

internal class Definition : ResourceDefinition
{
    public string Name => "CheckFrontDoorNameAvailabilityWithSubscription";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new FrontDoorNameAvailabilityWithSubscriptionCheckOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(AvailabilityConstant),
        typeof(ResourceTypeConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(CheckNameAvailabilityInputModel),
        typeof(CheckNameAvailabilityOutputModel),
    };
}
