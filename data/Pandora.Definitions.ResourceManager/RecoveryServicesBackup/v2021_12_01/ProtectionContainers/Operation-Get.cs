using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.RecoveryServicesBackup.v2021_12_01.ProtectionContainers;

internal class GetOperation : Operations.GetOperation
{
    public override ResourceID? ResourceId() => new ProtectionContainerId();

    public override Type? ResponseObject() => typeof(ProtectionContainerResourceModel);


}
