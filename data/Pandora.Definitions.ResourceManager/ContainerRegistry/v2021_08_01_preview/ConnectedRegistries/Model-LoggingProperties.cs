using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.ContainerRegistry.v2021_08_01_preview.ConnectedRegistries;


internal class LoggingPropertiesModel
{
    [JsonPropertyName("auditLogStatus")]
    public AuditLogStatusConstant? AuditLogStatus { get; set; }

    [JsonPropertyName("logLevel")]
    public LogLevelConstant? LogLevel { get; set; }
}