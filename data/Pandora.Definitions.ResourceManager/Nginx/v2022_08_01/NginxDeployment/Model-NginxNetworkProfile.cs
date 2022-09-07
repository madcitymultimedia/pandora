using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Nginx.v2022_08_01.NginxDeployment;


internal class NginxNetworkProfileModel
{
    [JsonPropertyName("frontEndIPConfiguration")]
    public NginxFrontendIPConfigurationModel? FrontEndIPConfiguration { get; set; }

    [JsonPropertyName("networkInterfaceConfiguration")]
    public NginxNetworkInterfaceConfigurationModel? NetworkInterfaceConfiguration { get; set; }
}