using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Insights.v2015_04_01.TenantActivityLogs;

internal class Definition : ResourceDefinition
{
    public string Name => "TenantActivityLogs";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new ListOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(EventLevelConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(EventDataModel),
        typeof(HTTPRequestInfoModel),
        typeof(LocalizableStringModel),
        typeof(SenderAuthorizationModel),
    };
}
