using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.ContainerRegistry.v2023_06_01_preview.CredentialSets;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum CredentialHealthStatusConstant
{
    [Description("Healthy")]
    Healthy,

    [Description("Unhealthy")]
    Unhealthy,
}
