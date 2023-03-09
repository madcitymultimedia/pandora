using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Logz.v2020_10_01.TagRules;

internal class SubAccountTagRulesGetOperation : Operations.GetOperation
{
    public override ResourceID? ResourceId() => new AccountTagRuleId();

    public override Type? ResponseObject() => typeof(MonitoringTagRulesModel);


}