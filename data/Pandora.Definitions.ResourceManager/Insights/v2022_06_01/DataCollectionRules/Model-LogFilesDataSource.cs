using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Insights.v2022_06_01.DataCollectionRules;


internal class LogFilesDataSourceModel
{
    [JsonPropertyName("filePatterns")]
    [Required]
    public List<string> FilePatterns { get; set; }

    [JsonPropertyName("format")]
    [Required]
    public KnownLogFilesDataSourceFormatConstant Format { get; set; }

    [JsonPropertyName("name")]
    public string? Name { get; set; }

    [JsonPropertyName("settings")]
    public LogFileSettingsModel? Settings { get; set; }

    [JsonPropertyName("streams")]
    [Required]
    public List<string> Streams { get; set; }
}